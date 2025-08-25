# ESP32 Boot Issues Diagnosis and Solutions

## Problem Analysis

Based on the boot logs provided, your ESP32 is experiencing:

1. **Multiple Reset Events**:
   - `rst:0x1 (POWERON_RESET)` - Initial power-on reset
   - `rst:0x10 (RTCWDT_RTC_RESET)` - RTC Watchdog Timer resets (repeated)

2. **Corrupted Serial Output**:
   - Garbled characters: `p��xf~�����fx~���������~���xf�枀���f�`fx`f���f��f���fx��������x�������x�������~x`��f�xf��f���`
   - This indicates UART communication issues

3. **Boot Loop**: The device keeps resetting instead of running your application

## Root Causes

### 1. UART Baud Rate Mismatch
Your code sets UART to 115200 baud, but the monitor might be using a different rate:
```go
uart.CLKDIV.Set(peripheralClock/115200 - 1)
```

### 2. Watchdog Timer Issues
The RTC watchdog is causing resets. Your code disables watchdogs but this might happen too late:
```go
getRTCCNTL().WDTWPROTECT.Set(0x50D83AA1)
getRTCCNTL().WDTCONFIG0.Set(0)
getTIMG().WDTCONFIG0.Set(0)
```

### 3. Clock Configuration Problems
The clock switching from XTAL to PLL might be unstable:
```go
getRTCCNTL().CLK_CONF.Set((1 << RTC_CNTL_CLK_CONF_SOC_CLK_SEL_Pos) | ...)
```

### 4. Power Supply Issues
- Insufficient power supply current
- Voltage drops during boot
- Poor power supply filtering

## Solutions

### 1. Fix UART Configuration

Update your `uart.go` with better initialization:

```go
func initUART() {
    uart := getUART()

    // Reset UART first
    uart.CONF0.Set(0x80000000) // Reset bit
    uart.CONF0.Set(0)          // Clear reset

    // Set baud rate with proper calculation
    // Use 40MHz as base clock for more stable timing
    const baseClock = 40000000
    const baudRate = 115200
    uart.CLKDIV.Set(baseClock/baudRate - 1)

    // Configure: 8N1 (8 data bits, no parity, 1 stop bit)
    uart.CONF0.Set(0x800001c)

    // Reset and clear FIFOs
    uart.CONF0.SetBits(0x60000)   // Set RXFIFO_RST and TXFIFO_RST
    uart.CONF0.ClearBits(0x60000) // Clear reset bits

    // Enable UART
    uart.CONF0.SetBits(0x2000000) // CLK_EN bit
}
```

### 2. Improve Watchdog Handling

Add this to the very beginning of `main()`:

```go
func main() {
    // Disable watchdogs IMMEDIATELY
    disableWatchdogs()

    // Initialize UART early for debugging
    initUART()
    c.Printf(c.Str("ESP32 Starting...\r\n"))

    // Rest of your code...
}

func disableWatchdogs() {
    rtc := getRTCCNTL()
    timg := getTIMG()

    // Disable RTC watchdog
    rtc.WDTWPROTECT.Set(0x50D83AA1) // Unlock
    rtc.WDTCONFIG0.Set(0)           // Disable
    rtc.WDTWPROTECT.Set(0)          // Lock

    // Disable Timer Group watchdog
    timg.WDTWPROTECT.Set(0x50D83AA1) // Unlock
    timg.WDTCONFIG0.Set(0)           // Disable
    timg.WDTWPROTECT.Set(0)          // Lock
}
```

### 3. Stable Clock Configuration

Replace your clock configuration with:

```go
func configureClocks() {
    rtc := getRTCCNTL()

    // Keep XTAL clock stable first
