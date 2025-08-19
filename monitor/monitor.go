package main

import (
	"debug/dwarf"
	"debug/elf"
	"debug/macho"
	"debug/pe"
	"errors"
	"fmt"
	"go/token"
	"io"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"time"

	"github.com/mattn/go-tty"
	"go.bug.st/serial"
)

func main() {
	const timeout = time.Second * 3
	var exit func() // function to be called before exiting
	var serialConn io.ReadWriter

	var err error
	wait := 300
	port := "/dev/cu.usbserial-10"

	br := 115200

	wait = 300
	var p serial.Port
	for i := 0; i <= wait; i++ {
		p, err = serial.Open(port, &serial.Mode{BaudRate: br})
		if err != nil {
			if i < wait {
				time.Sleep(10 * time.Millisecond)
				continue
			}
			panic(err)
		}
		serialConn = p
		break
	}
	defer p.Close()

	tty, err := tty.Open()
	if err != nil {
		panic(err)
	}
	defer tty.Close()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	defer signal.Stop(sig)

	go func() {
		<-sig
		tty.Close()
		if exit != nil {
			exit()
		}
		os.Exit(0)
	}()

	fmt.Printf("Connected to %s. Press Ctrl-C to exit.\n", port)

	errCh := make(chan error, 1)

	go func() {
		buf := make([]byte, 100*1024)
		writer := newOutputWriter(os.Stdout, "../t.elf")
		for {
			n, err := serialConn.Read(buf)
			if err != nil {
				errCh <- fmt.Errorf("read error: %w", err)
				return
			}
			writer.Write(buf[:n])
		}
	}()

	go func() {
		for {
			r, err := tty.ReadRune()
			if err != nil {
				errCh <- err
				return
			}
			if r == 0 {
				continue
			}
			serialConn.Write([]byte(string(r)))
		}
	}()

	select {}
}

var addressMatch = regexp.MustCompile(`^panic: runtime error at 0x([0-9a-f]+): `)

// Extract the address from the "panic: runtime error at" message.
func extractPanicAddress(line []byte) uint64 {
	matches := addressMatch.FindSubmatch(line)
	if matches != nil {
		address, err := strconv.ParseUint(string(matches[1]), 16, 64)
		if err == nil {
			return address
		}
	}
	return 0
}

// Convert an address in the binary to a source address location.
func addressToLine(executable string, address uint64) (token.Position, error) {
	data, err := readDWARF(executable)
	if err != nil {
		return token.Position{}, err
	}
	r := data.Reader()

	for {
		e, err := r.Next()
		if err != nil {
			return token.Position{}, err
		}
		if e == nil {
			break
		}
		switch e.Tag {
		case dwarf.TagCompileUnit:
			r.SkipChildren()
			lr, err := data.LineReader(e)
			if err != nil {
				return token.Position{}, err
			}
			var lineEntry = dwarf.LineEntry{
				EndSequence: true,
			}
			for {
				// Read the next .debug_line entry.
				prevLineEntry := lineEntry
				err := lr.Next(&lineEntry)
				if err != nil {
					if err == io.EOF {
						break
					}
					return token.Position{}, err
				}

				if prevLineEntry.EndSequence && lineEntry.Address == 0 {
					// Tombstone value. This symbol has been removed, for
					// example by the --gc-sections linker flag. It is still
					// here in the debug information because the linker can't
					// just remove this reference.
					// Read until the next EndSequence so that this sequence is
					// skipped.
					// For more details, see (among others):
					// https://reviews.llvm.org/D84825
					for {
						err := lr.Next(&lineEntry)
						if err != nil {
							return token.Position{}, err
						}
						if lineEntry.EndSequence {
							break
						}
					}
				}

				if !prevLineEntry.EndSequence {
					// The chunk describes the code from prevLineEntry to
					// lineEntry.
					if prevLineEntry.Address <= address && lineEntry.Address > address {
						return token.Position{
							Filename: prevLineEntry.File.Name,
							Line:     prevLineEntry.Line,
							Column:   prevLineEntry.Column,
						}, nil
					}
				}
			}
		}
	}

	return token.Position{}, nil // location not found
}

// Read the DWARF debug information from a given file (in various formats).
func readDWARF(executable string) (*dwarf.Data, error) {
	f, err := os.Open(executable)
	if err != nil {
		return nil, err
	}
	if file, err := elf.NewFile(f); err == nil {
		return file.DWARF()
	} else if file, err := macho.NewFile(f); err == nil {
		return file.DWARF()
	} else if file, err := pe.NewFile(f); err == nil {
		return file.DWARF()
	} else {
		return nil, errors.New("unknown binary format")
	}
}

type outputWriter struct {
	out        io.Writer
	executable string
	line       []byte
}

// newOutputWriter returns an io.Writer that will intercept panic addresses and
// will try to insert a source location in the output if the source location can
// be found in the executable.
func newOutputWriter(out io.Writer, executable string) *outputWriter {
	return &outputWriter{
		out:        out,
		executable: executable,
	}
}

func (w *outputWriter) Write(p []byte) (n int, err error) {
	start := 0
	for i, c := range p {
		if c == '\n' {
			w.out.Write(p[start : i+1])
			start = i + 1
			address := extractPanicAddress(w.line)
			if address != 0 {
				loc, err := addressToLine(w.executable, address)
				fmt.Println(loc)

				if err == nil && loc.Filename != "" {
					fmt.Printf("[tinygo: panic at %s]\n", loc.String())
				}
			}
			w.line = w.line[:0]
		} else {
			w.line = append(w.line, c)
		}
	}
	w.out.Write(p[start:])
	n = len(p)
	return
}
