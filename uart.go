package main

import (
	"llgo-test/volatile"
	"unsafe"
)

const peripheralClock = 80000000 // 80MHz

// 获取UART指针的函数
func getUART() *UART_Type {
	return (*UART_Type)(unsafe.Pointer(uintptr(0x3ff40000)))
}

func writeByte(b byte) {
	uart := getUART()
	for (uart.STATUS.Get()>>16)&0xff >= 128 {
		// Read UART_TXFIFO_CNT from the status register, which indicates how
		// many bytes there are in the transmit buffer. Wait until there are
		// less than 128 bytes in this buffer (the default buffer size).
	}
	// Write to the FIFO register (offset 0x0)
	uart.FIFO.Set(uint32(b))
}

// 初始化UART
func initUART() {
	uart := getUART()

	// 设置波特率为115200
	// 波特率 = 时钟频率 / (CLKDIV + 1)
	// CLKDIV = (时钟频率 / 波特率) - 1
	uart.CLKDIV.Set(peripheralClock/115200 - 1)

	// 配置UART参数：8位数据，无奇偶校验，1个停止位
	uart.CONF0.Set(0x800001c) // 8位数据位，无奇偶校验，1个停止位

	// 重置FIFO
	uart.CONF0.SetBits(0x60000)   // 设置RXFIFO_RST和TXFIFO_RST位
	uart.CONF0.ClearBits(0x60000) // 清除重置位
}

// 便捷函数 - 输出预定义消息（避免使用string和数组）
func printStartup() {
	writeByte('E')
	writeByte('S')
	writeByte('P')
	writeByte('3')
	writeByte('2')
	writeByte(' ')
	writeByte('s')
	writeByte('t')
	writeByte('a')
	writeByte('r')
	writeByte('t')
	writeByte('\r')
	writeByte('\n')
}

func printOK() {
	writeByte('O')
	writeByte('K')
	writeByte('\r')
	writeByte('\n')
}

func printError() {
	writeByte('E')
	writeByte('R')
	writeByte('R')
	writeByte('\r')
	writeByte('\n')
}

func printDone() {
	writeByte('D')
	writeByte('O')
	writeByte('N')
	writeByte('E')
	writeByte('\r')
	writeByte('\n')
}

func printHigh() {
	writeByte('H')
	writeByte('I')
	writeByte('G')
	writeByte('H')
	writeByte('\r')
	writeByte('\n')
}

func printLow() {
	writeByte('L')
	writeByte('O')
	writeByte('W')
	writeByte('\r')
	writeByte('\n')
}

// 将单个十六进制数字转换为ASCII字符
func hexDigitToChar(digit uint8) uint8 {
	if digit < 10 {
		return 0x30 + digit // '0' = 0x30
	}
	return 0x41 + (digit - 10) // 'A' = 0x41
}

// 将int32转换为16进制字符串并按byte输出到UART
func printInt32Hex(value int32) {
	// 转换为uint32以便进行位操作
	uvalue := uint32(value)

	// 输出"0x"前缀
	writeByte(0x30) // '0'
	writeByte(0x78) // 'x'

	// 从最高位开始，每4位转换为一个十六进制字符
	for i := uint32(0); i < 8; i++ {
		// 提取第i个十六进制位（从高位开始）
		digit := uint8((uvalue >> ((7 - i) * 4)) & 0xF)
		writeByte(hexDigitToChar(digit))
	}
}

// 将int16转换为16进制字符串并按byte输出到UART
func printInt16Hex(value int16) {
	// 转换为uint16以便进行位操作
	uvalue := uint16(value)

	// 输出"0x"前缀
	writeByte(0x30) // '0'
	writeByte(0x78) // 'x'

	// 从最高位开始，每4位转换为一个十六进制字符
	for i := uint32(0); i < 4; i++ {
		// 提取第i个十六进制位（从高位开始）
		digit := uint8((uvalue >> ((3 - i) * 4)) & 0xF)
		writeByte(hexDigitToChar(digit))
	}
}

// 将int8转换为16进制字符串并按byte输出到UART
func printInt8Hex(value int8) {
	// 转换为uint8以便进行位操作
	uvalue := uint8(value)

	// 输出"0x"前缀
	writeByte(0x30) // '0'
	writeByte(0x78) // 'x'

	// 输出高4位
	digit := (uvalue >> 4) & 0xF
	writeByte(hexDigitToChar(digit))

	// 输出低4位
	digit = uvalue & 0xF
	writeByte(hexDigitToChar(digit))
}

// 通用的int转16进制输出函数（默认为int32）
func printIntHex(value int) {
	printInt32Hex(int32(value))
}

// 将uint32按字节输出（大端序）
func printUint32BytesBE(value uint32) {
	writeByte(uint8((value >> 24) & 0xFF)) // 最高字节
	writeByte(uint8((value >> 16) & 0xFF))
	writeByte(uint8((value >> 8) & 0xFF))
	writeByte(uint8(value & 0xFF)) // 最低字节
}

// 将uint32按字节输出（小端序）
func printUint32BytesLE(value uint32) {
	writeByte(uint8(value & 0xFF)) // 最低字节
	writeByte(uint8((value >> 8) & 0xFF))
	writeByte(uint8((value >> 16) & 0xFF))
	writeByte(uint8((value >> 24) & 0xFF)) // 最高字节
}

// 将int32按字节输出（大端序）
func printInt32BytesBE(value int32) {
	printUint32BytesBE(uint32(value))
}

// 将int32按字节输出（小端序）
func printInt32BytesLE(value int32) {
	printUint32BytesLE(uint32(value))
}

// 将uint16按字节输出（大端序）
func printUint16BytesBE(value uint16) {
	writeByte(uint8((value >> 8) & 0xFF)) // 高字节
	writeByte(uint8(value & 0xFF))        // 低字节
}

// 将uint16按字节输出（小端序）
func printUint16BytesLE(value uint16) {
	writeByte(uint8(value & 0xFF))        // 低字节
	writeByte(uint8((value >> 8) & 0xFF)) // 高字节
}

// 将int16按字节输出（大端序）
func printInt16BytesBE(value int16) {
	printUint16BytesBE(uint16(value))
}

// 将int16按字节输出（小端序）
func printInt16BytesLE(value int16) {
	printUint16BytesLE(uint16(value))
}

// UART (Universal Asynchronous Receiver-Transmitter) Controller 0
type UART_Type struct {
	FIFO           volatile.Register32 // 0x0
	INT_RAW        volatile.Register32 // 0x4
	INT_ST         volatile.Register32 // 0x8
	INT_ENA        volatile.Register32 // 0xC
	INT_CLR        volatile.Register32 // 0x10
	CLKDIV         volatile.Register32 // 0x14
	AUTOBAUD       volatile.Register32 // 0x18
	STATUS         volatile.Register32 // 0x1C
	CONF0          volatile.Register32 // 0x20
	CONF1          volatile.Register32 // 0x24
	LOWPULSE       volatile.Register32 // 0x28
	HIGHPULSE      volatile.Register32 // 0x2C
	RXD_CNT        volatile.Register32 // 0x30
	FLOW_CONF      volatile.Register32 // 0x34
	SLEEP_CONF     volatile.Register32 // 0x38
	SWFC_CONF      volatile.Register32 // 0x3C
	IDLE_CONF      volatile.Register32 // 0x40
	RS485_CONF     volatile.Register32 // 0x44
	AT_CMD_PRECNT  volatile.Register32 // 0x48
	AT_CMD_POSTCNT volatile.Register32 // 0x4C
	AT_CMD_GAPTOUT volatile.Register32 // 0x50
	AT_CMD_CHAR    volatile.Register32 // 0x54
	MEM_CONF       volatile.Register32 // 0x58
	MEM_TX_STATUS  volatile.Register32 // 0x5C
	MEM_RX_STATUS  volatile.Register32 // 0x60
	MEM_CNT_STATUS volatile.Register32 // 0x64
	POSPULSE       volatile.Register32 // 0x68
	NEGPULSE       volatile.Register32 // 0x6C
	_              [8]byte
	DATE           volatile.Register32 // 0x78
	ID             volatile.Register32 // 0x7C
}

// UART.FIFO
func (o *UART_Type) SetFIFO_RXFIFO_RD_BYTE(value uint32) {
	volatile.StoreUint32(&o.FIFO.Reg, volatile.LoadUint32(&o.FIFO.Reg)&^(0xff)|value)
}
func (o *UART_Type) GetFIFO_RXFIFO_RD_BYTE() uint32 {
	return volatile.LoadUint32(&o.FIFO.Reg) & 0xff
}

// UART.INT_RAW
func (o *UART_Type) SetINT_RAW_RXFIFO_FULL_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x1)|value)
}
func (o *UART_Type) GetINT_RAW_RXFIFO_FULL_INT_RAW() uint32 {
	return volatile.LoadUint32(&o.INT_RAW.Reg) & 0x1
}
func (o *UART_Type) SetINT_RAW_TXFIFO_EMPTY_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x2)|value<<1)
}
func (o *UART_Type) GetINT_RAW_TXFIFO_EMPTY_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x2) >> 1
}
func (o *UART_Type) SetINT_RAW_PARITY_ERR_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x4)|value<<2)
}
func (o *UART_Type) GetINT_RAW_PARITY_ERR_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x4) >> 2
}
func (o *UART_Type) SetINT_RAW_FRM_ERR_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x8)|value<<3)
}
func (o *UART_Type) GetINT_RAW_FRM_ERR_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x8) >> 3
}
func (o *UART_Type) SetINT_RAW_RXFIFO_OVF_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x10)|value<<4)
}
func (o *UART_Type) GetINT_RAW_RXFIFO_OVF_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x10) >> 4
}
func (o *UART_Type) SetINT_RAW_DSR_CHG_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x20)|value<<5)
}
func (o *UART_Type) GetINT_RAW_DSR_CHG_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x20) >> 5
}
func (o *UART_Type) SetINT_RAW_CTS_CHG_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x40)|value<<6)
}
func (o *UART_Type) GetINT_RAW_CTS_CHG_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x40) >> 6
}
func (o *UART_Type) SetINT_RAW_BRK_DET_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x80)|value<<7)
}
func (o *UART_Type) GetINT_RAW_BRK_DET_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x80) >> 7
}
func (o *UART_Type) SetINT_RAW_RXFIFO_TOUT_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x100)|value<<8)
}
func (o *UART_Type) GetINT_RAW_RXFIFO_TOUT_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x100) >> 8
}
func (o *UART_Type) SetINT_RAW_SW_XON_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x200)|value<<9)
}
func (o *UART_Type) GetINT_RAW_SW_XON_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x200) >> 9
}
func (o *UART_Type) SetINT_RAW_SW_XOFF_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x400)|value<<10)
}
func (o *UART_Type) GetINT_RAW_SW_XOFF_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x400) >> 10
}
func (o *UART_Type) SetINT_RAW_GLITCH_DET_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x800)|value<<11)
}
func (o *UART_Type) GetINT_RAW_GLITCH_DET_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x800) >> 11
}
func (o *UART_Type) SetINT_RAW_TX_BRK_DONE_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x1000)|value<<12)
}
func (o *UART_Type) GetINT_RAW_TX_BRK_DONE_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x1000) >> 12
}
func (o *UART_Type) SetINT_RAW_TX_BRK_IDLE_DONE_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x2000)|value<<13)
}
func (o *UART_Type) GetINT_RAW_TX_BRK_IDLE_DONE_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x2000) >> 13
}
func (o *UART_Type) SetINT_RAW_TX_DONE_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x4000)|value<<14)
}
func (o *UART_Type) GetINT_RAW_TX_DONE_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x4000) >> 14
}
func (o *UART_Type) SetINT_RAW_RS485_PARITY_ERR_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x8000)|value<<15)
}
func (o *UART_Type) GetINT_RAW_RS485_PARITY_ERR_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x8000) >> 15
}
func (o *UART_Type) SetINT_RAW_RS485_FRM_ERR_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x10000)|value<<16)
}
func (o *UART_Type) GetINT_RAW_RS485_FRM_ERR_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x10000) >> 16
}
func (o *UART_Type) SetINT_RAW_RS485_CLASH_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x20000)|value<<17)
}
func (o *UART_Type) GetINT_RAW_RS485_CLASH_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x20000) >> 17
}
func (o *UART_Type) SetINT_RAW_AT_CMD_CHAR_DET_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW.Reg, volatile.LoadUint32(&o.INT_RAW.Reg)&^(0x40000)|value<<18)
}
func (o *UART_Type) GetINT_RAW_AT_CMD_CHAR_DET_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW.Reg) & 0x40000) >> 18
}

// UART.INT_ST
func (o *UART_Type) SetINT_ST_RXFIFO_FULL_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x1)|value)
}
func (o *UART_Type) GetINT_ST_RXFIFO_FULL_INT_ST() uint32 {
	return volatile.LoadUint32(&o.INT_ST.Reg) & 0x1
}
func (o *UART_Type) SetINT_ST_TXFIFO_EMPTY_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x2)|value<<1)
}
func (o *UART_Type) GetINT_ST_TXFIFO_EMPTY_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x2) >> 1
}
func (o *UART_Type) SetINT_ST_PARITY_ERR_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x4)|value<<2)
}
func (o *UART_Type) GetINT_ST_PARITY_ERR_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x4) >> 2
}
func (o *UART_Type) SetINT_ST_FRM_ERR_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x8)|value<<3)
}
func (o *UART_Type) GetINT_ST_FRM_ERR_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x8) >> 3
}
func (o *UART_Type) SetINT_ST_RXFIFO_OVF_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x10)|value<<4)
}
func (o *UART_Type) GetINT_ST_RXFIFO_OVF_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x10) >> 4
}
func (o *UART_Type) SetINT_ST_DSR_CHG_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x20)|value<<5)
}
func (o *UART_Type) GetINT_ST_DSR_CHG_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x20) >> 5
}
func (o *UART_Type) SetINT_ST_CTS_CHG_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x40)|value<<6)
}
func (o *UART_Type) GetINT_ST_CTS_CHG_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x40) >> 6
}
func (o *UART_Type) SetINT_ST_BRK_DET_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x80)|value<<7)
}
func (o *UART_Type) GetINT_ST_BRK_DET_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x80) >> 7
}
func (o *UART_Type) SetINT_ST_RXFIFO_TOUT_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x100)|value<<8)
}
func (o *UART_Type) GetINT_ST_RXFIFO_TOUT_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x100) >> 8
}
func (o *UART_Type) SetINT_ST_SW_XON_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x200)|value<<9)
}
func (o *UART_Type) GetINT_ST_SW_XON_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x200) >> 9
}
func (o *UART_Type) SetINT_ST_SW_XOFF_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x400)|value<<10)
}
func (o *UART_Type) GetINT_ST_SW_XOFF_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x400) >> 10
}
func (o *UART_Type) SetINT_ST_GLITCH_DET_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x800)|value<<11)
}
func (o *UART_Type) GetINT_ST_GLITCH_DET_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x800) >> 11
}
func (o *UART_Type) SetINT_ST_TX_BRK_DONE_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x1000)|value<<12)
}
func (o *UART_Type) GetINT_ST_TX_BRK_DONE_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x1000) >> 12
}
func (o *UART_Type) SetINT_ST_TX_BRK_IDLE_DONE_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x2000)|value<<13)
}
func (o *UART_Type) GetINT_ST_TX_BRK_IDLE_DONE_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x2000) >> 13
}
func (o *UART_Type) SetINT_ST_TX_DONE_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x4000)|value<<14)
}
func (o *UART_Type) GetINT_ST_TX_DONE_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x4000) >> 14
}
func (o *UART_Type) SetINT_ST_RS485_PARITY_ERR_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x8000)|value<<15)
}
func (o *UART_Type) GetINT_ST_RS485_PARITY_ERR_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x8000) >> 15
}
func (o *UART_Type) SetINT_ST_RS485_FRM_ERR_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x10000)|value<<16)
}
func (o *UART_Type) GetINT_ST_RS485_FRM_ERR_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x10000) >> 16
}
func (o *UART_Type) SetINT_ST_RS485_CLASH_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x20000)|value<<17)
}
func (o *UART_Type) GetINT_ST_RS485_CLASH_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x20000) >> 17
}
func (o *UART_Type) SetINT_ST_AT_CMD_CHAR_DET_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST.Reg, volatile.LoadUint32(&o.INT_ST.Reg)&^(0x40000)|value<<18)
}
func (o *UART_Type) GetINT_ST_AT_CMD_CHAR_DET_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST.Reg) & 0x40000) >> 18
}

// UART.INT_ENA
func (o *UART_Type) SetINT_ENA_RXFIFO_FULL_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x1)|value)
}
func (o *UART_Type) GetINT_ENA_RXFIFO_FULL_INT_ENA() uint32 {
	return volatile.LoadUint32(&o.INT_ENA.Reg) & 0x1
}
func (o *UART_Type) SetINT_ENA_TXFIFO_EMPTY_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x2)|value<<1)
}
func (o *UART_Type) GetINT_ENA_TXFIFO_EMPTY_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x2) >> 1
}
func (o *UART_Type) SetINT_ENA_PARITY_ERR_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x4)|value<<2)
}
func (o *UART_Type) GetINT_ENA_PARITY_ERR_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x4) >> 2
}
func (o *UART_Type) SetINT_ENA_FRM_ERR_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x8)|value<<3)
}
func (o *UART_Type) GetINT_ENA_FRM_ERR_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x8) >> 3
}
func (o *UART_Type) SetINT_ENA_RXFIFO_OVF_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x10)|value<<4)
}
func (o *UART_Type) GetINT_ENA_RXFIFO_OVF_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x10) >> 4
}
func (o *UART_Type) SetINT_ENA_DSR_CHG_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x20)|value<<5)
}
func (o *UART_Type) GetINT_ENA_DSR_CHG_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x20) >> 5
}
func (o *UART_Type) SetINT_ENA_CTS_CHG_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x40)|value<<6)
}
func (o *UART_Type) GetINT_ENA_CTS_CHG_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x40) >> 6
}
func (o *UART_Type) SetINT_ENA_BRK_DET_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x80)|value<<7)
}
func (o *UART_Type) GetINT_ENA_BRK_DET_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x80) >> 7
}
func (o *UART_Type) SetINT_ENA_RXFIFO_TOUT_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x100)|value<<8)
}
func (o *UART_Type) GetINT_ENA_RXFIFO_TOUT_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x100) >> 8
}
func (o *UART_Type) SetINT_ENA_SW_XON_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x200)|value<<9)
}
func (o *UART_Type) GetINT_ENA_SW_XON_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x200) >> 9
}
func (o *UART_Type) SetINT_ENA_SW_XOFF_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x400)|value<<10)
}
func (o *UART_Type) GetINT_ENA_SW_XOFF_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x400) >> 10
}
func (o *UART_Type) SetINT_ENA_GLITCH_DET_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x800)|value<<11)
}
func (o *UART_Type) GetINT_ENA_GLITCH_DET_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x800) >> 11
}
func (o *UART_Type) SetINT_ENA_TX_BRK_DONE_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x1000)|value<<12)
}
func (o *UART_Type) GetINT_ENA_TX_BRK_DONE_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x1000) >> 12
}
func (o *UART_Type) SetINT_ENA_TX_BRK_IDLE_DONE_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x2000)|value<<13)
}
func (o *UART_Type) GetINT_ENA_TX_BRK_IDLE_DONE_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x2000) >> 13
}
func (o *UART_Type) SetINT_ENA_TX_DONE_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x4000)|value<<14)
}
func (o *UART_Type) GetINT_ENA_TX_DONE_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x4000) >> 14
}
func (o *UART_Type) SetINT_ENA_RS485_PARITY_ERR_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x8000)|value<<15)
}
func (o *UART_Type) GetINT_ENA_RS485_PARITY_ERR_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x8000) >> 15
}
func (o *UART_Type) SetINT_ENA_RS485_FRM_ERR_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x10000)|value<<16)
}
func (o *UART_Type) GetINT_ENA_RS485_FRM_ERR_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x10000) >> 16
}
func (o *UART_Type) SetINT_ENA_RS485_CLASH_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x20000)|value<<17)
}
func (o *UART_Type) GetINT_ENA_RS485_CLASH_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x20000) >> 17
}
func (o *UART_Type) SetINT_ENA_AT_CMD_CHAR_DET_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA.Reg, volatile.LoadUint32(&o.INT_ENA.Reg)&^(0x40000)|value<<18)
}
func (o *UART_Type) GetINT_ENA_AT_CMD_CHAR_DET_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA.Reg) & 0x40000) >> 18
}

// UART.INT_CLR
func (o *UART_Type) SetINT_CLR_RXFIFO_FULL_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x1)|value)
}
func (o *UART_Type) GetINT_CLR_RXFIFO_FULL_INT_CLR() uint32 {
	return volatile.LoadUint32(&o.INT_CLR.Reg) & 0x1
}
func (o *UART_Type) SetINT_CLR_TXFIFO_EMPTY_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x2)|value<<1)
}
func (o *UART_Type) GetINT_CLR_TXFIFO_EMPTY_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x2) >> 1
}
func (o *UART_Type) SetINT_CLR_PARITY_ERR_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x4)|value<<2)
}
func (o *UART_Type) GetINT_CLR_PARITY_ERR_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x4) >> 2
}
func (o *UART_Type) SetINT_CLR_FRM_ERR_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x8)|value<<3)
}
func (o *UART_Type) GetINT_CLR_FRM_ERR_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x8) >> 3
}
func (o *UART_Type) SetINT_CLR_RXFIFO_OVF_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x10)|value<<4)
}
func (o *UART_Type) GetINT_CLR_RXFIFO_OVF_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x10) >> 4
}
func (o *UART_Type) SetINT_CLR_DSR_CHG_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x20)|value<<5)
}
func (o *UART_Type) GetINT_CLR_DSR_CHG_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x20) >> 5
}
func (o *UART_Type) SetINT_CLR_CTS_CHG_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x40)|value<<6)
}
func (o *UART_Type) GetINT_CLR_CTS_CHG_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x40) >> 6
}
func (o *UART_Type) SetINT_CLR_BRK_DET_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x80)|value<<7)
}
func (o *UART_Type) GetINT_CLR_BRK_DET_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x80) >> 7
}
func (o *UART_Type) SetINT_CLR_RXFIFO_TOUT_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x100)|value<<8)
}
func (o *UART_Type) GetINT_CLR_RXFIFO_TOUT_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x100) >> 8
}
func (o *UART_Type) SetINT_CLR_SW_XON_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x200)|value<<9)
}
func (o *UART_Type) GetINT_CLR_SW_XON_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x200) >> 9
}
func (o *UART_Type) SetINT_CLR_SW_XOFF_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x400)|value<<10)
}
func (o *UART_Type) GetINT_CLR_SW_XOFF_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x400) >> 10
}
func (o *UART_Type) SetINT_CLR_GLITCH_DET_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x800)|value<<11)
}
func (o *UART_Type) GetINT_CLR_GLITCH_DET_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x800) >> 11
}
func (o *UART_Type) SetINT_CLR_TX_BRK_DONE_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x1000)|value<<12)
}
func (o *UART_Type) GetINT_CLR_TX_BRK_DONE_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x1000) >> 12
}
func (o *UART_Type) SetINT_CLR_TX_BRK_IDLE_DONE_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x2000)|value<<13)
}
func (o *UART_Type) GetINT_CLR_TX_BRK_IDLE_DONE_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x2000) >> 13
}
func (o *UART_Type) SetINT_CLR_TX_DONE_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x4000)|value<<14)
}
func (o *UART_Type) GetINT_CLR_TX_DONE_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x4000) >> 14
}
func (o *UART_Type) SetINT_CLR_RS485_PARITY_ERR_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x8000)|value<<15)
}
func (o *UART_Type) GetINT_CLR_RS485_PARITY_ERR_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x8000) >> 15
}
func (o *UART_Type) SetINT_CLR_RS485_FRM_ERR_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x10000)|value<<16)
}
func (o *UART_Type) GetINT_CLR_RS485_FRM_ERR_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x10000) >> 16
}
func (o *UART_Type) SetINT_CLR_RS485_CLASH_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x20000)|value<<17)
}
func (o *UART_Type) GetINT_CLR_RS485_CLASH_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x20000) >> 17
}
func (o *UART_Type) SetINT_CLR_AT_CMD_CHAR_DET_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR.Reg, volatile.LoadUint32(&o.INT_CLR.Reg)&^(0x40000)|value<<18)
}
func (o *UART_Type) GetINT_CLR_AT_CMD_CHAR_DET_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR.Reg) & 0x40000) >> 18
}

// UART.CLKDIV
func (o *UART_Type) SetCLKDIV(value uint32) {
	volatile.StoreUint32(&o.CLKDIV.Reg, volatile.LoadUint32(&o.CLKDIV.Reg)&^(0xfffff)|value)
}
func (o *UART_Type) GetCLKDIV() uint32 {
	return volatile.LoadUint32(&o.CLKDIV.Reg) & 0xfffff
}
func (o *UART_Type) SetCLKDIV_FRAG(value uint32) {
	volatile.StoreUint32(&o.CLKDIV.Reg, volatile.LoadUint32(&o.CLKDIV.Reg)&^(0xf00000)|value<<20)
}
func (o *UART_Type) GetCLKDIV_FRAG() uint32 {
	return (volatile.LoadUint32(&o.CLKDIV.Reg) & 0xf00000) >> 20
}

// UART.AUTOBAUD
func (o *UART_Type) SetAUTOBAUD_EN(value uint32) {
	volatile.StoreUint32(&o.AUTOBAUD.Reg, volatile.LoadUint32(&o.AUTOBAUD.Reg)&^(0x1)|value)
}
func (o *UART_Type) GetAUTOBAUD_EN() uint32 {
	return volatile.LoadUint32(&o.AUTOBAUD.Reg) & 0x1
}
func (o *UART_Type) SetAUTOBAUD_GLITCH_FILT(value uint32) {
	volatile.StoreUint32(&o.AUTOBAUD.Reg, volatile.LoadUint32(&o.AUTOBAUD.Reg)&^(0xff00)|value<<8)
}
func (o *UART_Type) GetAUTOBAUD_GLITCH_FILT() uint32 {
	return (volatile.LoadUint32(&o.AUTOBAUD.Reg) & 0xff00) >> 8
}

// UART.STATUS
func (o *UART_Type) SetSTATUS_RXFIFO_CNT(value uint32) {
	volatile.StoreUint32(&o.STATUS.Reg, volatile.LoadUint32(&o.STATUS.Reg)&^(0xff)|value)
}
func (o *UART_Type) GetSTATUS_RXFIFO_CNT() uint32 {
	return volatile.LoadUint32(&o.STATUS.Reg) & 0xff
}
func (o *UART_Type) SetSTATUS_ST_URX_OUT(value uint32) {
	volatile.StoreUint32(&o.STATUS.Reg, volatile.LoadUint32(&o.STATUS.Reg)&^(0xf00)|value<<8)
}
func (o *UART_Type) GetSTATUS_ST_URX_OUT() uint32 {
	return (volatile.LoadUint32(&o.STATUS.Reg) & 0xf00) >> 8
}
func (o *UART_Type) SetSTATUS_DSRN(value uint32) {
	volatile.StoreUint32(&o.STATUS.Reg, volatile.LoadUint32(&o.STATUS.Reg)&^(0x2000)|value<<13)
}
func (o *UART_Type) GetSTATUS_DSRN() uint32 {
	return (volatile.LoadUint32(&o.STATUS.Reg) & 0x2000) >> 13
}
func (o *UART_Type) SetSTATUS_CTSN(value uint32) {
	volatile.StoreUint32(&o.STATUS.Reg, volatile.LoadUint32(&o.STATUS.Reg)&^(0x4000)|value<<14)
}
func (o *UART_Type) GetSTATUS_CTSN() uint32 {
	return (volatile.LoadUint32(&o.STATUS.Reg) & 0x4000) >> 14
}
func (o *UART_Type) SetSTATUS_RXD(value uint32) {
	volatile.StoreUint32(&o.STATUS.Reg, volatile.LoadUint32(&o.STATUS.Reg)&^(0x8000)|value<<15)
}
func (o *UART_Type) GetSTATUS_RXD() uint32 {
	return (volatile.LoadUint32(&o.STATUS.Reg) & 0x8000) >> 15
}
func (o *UART_Type) SetSTATUS_TXFIFO_CNT(value uint32) {
	volatile.StoreUint32(&o.STATUS.Reg, volatile.LoadUint32(&o.STATUS.Reg)&^(0xff0000)|value<<16)
}
func (o *UART_Type) GetSTATUS_TXFIFO_CNT() uint32 {
	return (volatile.LoadUint32(&o.STATUS.Reg) & 0xff0000) >> 16
}
func (o *UART_Type) SetSTATUS_ST_UTX_OUT(value uint32) {
	volatile.StoreUint32(&o.STATUS.Reg, volatile.LoadUint32(&o.STATUS.Reg)&^(0xf000000)|value<<24)
}
func (o *UART_Type) GetSTATUS_ST_UTX_OUT() uint32 {
	return (volatile.LoadUint32(&o.STATUS.Reg) & 0xf000000) >> 24
}
func (o *UART_Type) SetSTATUS_DTRN(value uint32) {
	volatile.StoreUint32(&o.STATUS.Reg, volatile.LoadUint32(&o.STATUS.Reg)&^(0x20000000)|value<<29)
}
func (o *UART_Type) GetSTATUS_DTRN() uint32 {
	return (volatile.LoadUint32(&o.STATUS.Reg) & 0x20000000) >> 29
}
func (o *UART_Type) SetSTATUS_RTSN(value uint32) {
	volatile.StoreUint32(&o.STATUS.Reg, volatile.LoadUint32(&o.STATUS.Reg)&^(0x40000000)|value<<30)
}
func (o *UART_Type) GetSTATUS_RTSN() uint32 {
	return (volatile.LoadUint32(&o.STATUS.Reg) & 0x40000000) >> 30
}
func (o *UART_Type) SetSTATUS_TXD(value uint32) {
	volatile.StoreUint32(&o.STATUS.Reg, volatile.LoadUint32(&o.STATUS.Reg)&^(0x80000000)|value<<31)
}
func (o *UART_Type) GetSTATUS_TXD() uint32 {
	return (volatile.LoadUint32(&o.STATUS.Reg) & 0x80000000) >> 31
}

// UART.CONF0
func (o *UART_Type) SetCONF0_PARITY(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x1)|value)
}
func (o *UART_Type) GetCONF0_PARITY() uint32 {
	return volatile.LoadUint32(&o.CONF0.Reg) & 0x1
}
func (o *UART_Type) SetCONF0_PARITY_EN(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x2)|value<<1)
}
func (o *UART_Type) GetCONF0_PARITY_EN() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x2) >> 1
}
func (o *UART_Type) SetCONF0_BIT_NUM(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0xc)|value<<2)
}
func (o *UART_Type) GetCONF0_BIT_NUM() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0xc) >> 2
}
func (o *UART_Type) SetCONF0_STOP_BIT_NUM(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x30)|value<<4)
}
func (o *UART_Type) GetCONF0_STOP_BIT_NUM() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x30) >> 4
}
func (o *UART_Type) SetCONF0_SW_RTS(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x40)|value<<6)
}
func (o *UART_Type) GetCONF0_SW_RTS() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x40) >> 6
}
func (o *UART_Type) SetCONF0_SW_DTR(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x80)|value<<7)
}
func (o *UART_Type) GetCONF0_SW_DTR() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x80) >> 7
}
func (o *UART_Type) SetCONF0_TXD_BRK(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x100)|value<<8)
}
func (o *UART_Type) GetCONF0_TXD_BRK() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x100) >> 8
}
func (o *UART_Type) SetCONF0_IRDA_DPLX(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x200)|value<<9)
}
func (o *UART_Type) GetCONF0_IRDA_DPLX() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x200) >> 9
}
func (o *UART_Type) SetCONF0_IRDA_TX_EN(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x400)|value<<10)
}
func (o *UART_Type) GetCONF0_IRDA_TX_EN() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x400) >> 10
}
func (o *UART_Type) SetCONF0_IRDA_WCTL(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x800)|value<<11)
}
func (o *UART_Type) GetCONF0_IRDA_WCTL() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x800) >> 11
}
func (o *UART_Type) SetCONF0_IRDA_TX_INV(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x1000)|value<<12)
}
func (o *UART_Type) GetCONF0_IRDA_TX_INV() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x1000) >> 12
}
func (o *UART_Type) SetCONF0_IRDA_RX_INV(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x2000)|value<<13)
}
func (o *UART_Type) GetCONF0_IRDA_RX_INV() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x2000) >> 13
}
func (o *UART_Type) SetCONF0_LOOPBACK(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x4000)|value<<14)
}
func (o *UART_Type) GetCONF0_LOOPBACK() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x4000) >> 14
}
func (o *UART_Type) SetCONF0_TX_FLOW_EN(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x8000)|value<<15)
}
func (o *UART_Type) GetCONF0_TX_FLOW_EN() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x8000) >> 15
}
func (o *UART_Type) SetCONF0_IRDA_EN(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x10000)|value<<16)
}
func (o *UART_Type) GetCONF0_IRDA_EN() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x10000) >> 16
}
func (o *UART_Type) SetCONF0_RXFIFO_RST(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x20000)|value<<17)
}
func (o *UART_Type) GetCONF0_RXFIFO_RST() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x20000) >> 17
}
func (o *UART_Type) SetCONF0_TXFIFO_RST(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x40000)|value<<18)
}
func (o *UART_Type) GetCONF0_TXFIFO_RST() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x40000) >> 18
}
func (o *UART_Type) SetCONF0_RXD_INV(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x80000)|value<<19)
}
func (o *UART_Type) GetCONF0_RXD_INV() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x80000) >> 19
}
func (o *UART_Type) SetCONF0_CTS_INV(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x100000)|value<<20)
}
func (o *UART_Type) GetCONF0_CTS_INV() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x100000) >> 20
}
func (o *UART_Type) SetCONF0_DSR_INV(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x200000)|value<<21)
}
func (o *UART_Type) GetCONF0_DSR_INV() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x200000) >> 21
}
func (o *UART_Type) SetCONF0_TXD_INV(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x400000)|value<<22)
}
func (o *UART_Type) GetCONF0_TXD_INV() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x400000) >> 22
}
func (o *UART_Type) SetCONF0_RTS_INV(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x800000)|value<<23)
}
func (o *UART_Type) GetCONF0_RTS_INV() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x800000) >> 23
}
func (o *UART_Type) SetCONF0_DTR_INV(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x1000000)|value<<24)
}
func (o *UART_Type) GetCONF0_DTR_INV() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x1000000) >> 24
}
func (o *UART_Type) SetCONF0_CLK_EN(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x2000000)|value<<25)
}
func (o *UART_Type) GetCONF0_CLK_EN() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x2000000) >> 25
}
func (o *UART_Type) SetCONF0_ERR_WR_MASK(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x4000000)|value<<26)
}
func (o *UART_Type) GetCONF0_ERR_WR_MASK() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x4000000) >> 26
}
func (o *UART_Type) SetCONF0_TICK_REF_ALWAYS_ON(value uint32) {
	volatile.StoreUint32(&o.CONF0.Reg, volatile.LoadUint32(&o.CONF0.Reg)&^(0x8000000)|value<<27)
}
func (o *UART_Type) GetCONF0_TICK_REF_ALWAYS_ON() uint32 {
	return (volatile.LoadUint32(&o.CONF0.Reg) & 0x8000000) >> 27
}

// UART.CONF1
func (o *UART_Type) SetCONF1_RXFIFO_FULL_THRHD(value uint32) {
	volatile.StoreUint32(&o.CONF1.Reg, volatile.LoadUint32(&o.CONF1.Reg)&^(0x7f)|value)
}
func (o *UART_Type) GetCONF1_RXFIFO_FULL_THRHD() uint32 {
	return volatile.LoadUint32(&o.CONF1.Reg) & 0x7f
}
func (o *UART_Type) SetCONF1_TXFIFO_EMPTY_THRHD(value uint32) {
	volatile.StoreUint32(&o.CONF1.Reg, volatile.LoadUint32(&o.CONF1.Reg)&^(0x7f00)|value<<8)
}
func (o *UART_Type) GetCONF1_TXFIFO_EMPTY_THRHD() uint32 {
	return (volatile.LoadUint32(&o.CONF1.Reg) & 0x7f00) >> 8
}
func (o *UART_Type) SetCONF1_RX_FLOW_THRHD(value uint32) {
	volatile.StoreUint32(&o.CONF1.Reg, volatile.LoadUint32(&o.CONF1.Reg)&^(0x7f0000)|value<<16)
}
func (o *UART_Type) GetCONF1_RX_FLOW_THRHD() uint32 {
	return (volatile.LoadUint32(&o.CONF1.Reg) & 0x7f0000) >> 16
}
func (o *UART_Type) SetCONF1_RX_FLOW_EN(value uint32) {
	volatile.StoreUint32(&o.CONF1.Reg, volatile.LoadUint32(&o.CONF1.Reg)&^(0x800000)|value<<23)
}
func (o *UART_Type) GetCONF1_RX_FLOW_EN() uint32 {
	return (volatile.LoadUint32(&o.CONF1.Reg) & 0x800000) >> 23
}
func (o *UART_Type) SetCONF1_RX_TOUT_THRHD(value uint32) {
	volatile.StoreUint32(&o.CONF1.Reg, volatile.LoadUint32(&o.CONF1.Reg)&^(0x7f000000)|value<<24)
}
func (o *UART_Type) GetCONF1_RX_TOUT_THRHD() uint32 {
	return (volatile.LoadUint32(&o.CONF1.Reg) & 0x7f000000) >> 24
}
func (o *UART_Type) SetCONF1_RX_TOUT_EN(value uint32) {
	volatile.StoreUint32(&o.CONF1.Reg, volatile.LoadUint32(&o.CONF1.Reg)&^(0x80000000)|value<<31)
}
func (o *UART_Type) GetCONF1_RX_TOUT_EN() uint32 {
	return (volatile.LoadUint32(&o.CONF1.Reg) & 0x80000000) >> 31
}

// UART.LOWPULSE
func (o *UART_Type) SetLOWPULSE_MIN_CNT(value uint32) {
	volatile.StoreUint32(&o.LOWPULSE.Reg, volatile.LoadUint32(&o.LOWPULSE.Reg)&^(0xfffff)|value)
}
func (o *UART_Type) GetLOWPULSE_MIN_CNT() uint32 {
	return volatile.LoadUint32(&o.LOWPULSE.Reg) & 0xfffff
}

// UART.HIGHPULSE
func (o *UART_Type) SetHIGHPULSE_MIN_CNT(value uint32) {
	volatile.StoreUint32(&o.HIGHPULSE.Reg, volatile.LoadUint32(&o.HIGHPULSE.Reg)&^(0xfffff)|value)
}
func (o *UART_Type) GetHIGHPULSE_MIN_CNT() uint32 {
	return volatile.LoadUint32(&o.HIGHPULSE.Reg) & 0xfffff
}

// UART.RXD_CNT
func (o *UART_Type) SetRXD_CNT_RXD_EDGE_CNT(value uint32) {
	volatile.StoreUint32(&o.RXD_CNT.Reg, volatile.LoadUint32(&o.RXD_CNT.Reg)&^(0x3ff)|value)
}
func (o *UART_Type) GetRXD_CNT_RXD_EDGE_CNT() uint32 {
	return volatile.LoadUint32(&o.RXD_CNT.Reg) & 0x3ff
}

// UART.FLOW_CONF
func (o *UART_Type) SetFLOW_CONF_SW_FLOW_CON_EN(value uint32) {
	volatile.StoreUint32(&o.FLOW_CONF.Reg, volatile.LoadUint32(&o.FLOW_CONF.Reg)&^(0x1)|value)
}
func (o *UART_Type) GetFLOW_CONF_SW_FLOW_CON_EN() uint32 {
	return volatile.LoadUint32(&o.FLOW_CONF.Reg) & 0x1
}
func (o *UART_Type) SetFLOW_CONF_XONOFF_DEL(value uint32) {
	volatile.StoreUint32(&o.FLOW_CONF.Reg, volatile.LoadUint32(&o.FLOW_CONF.Reg)&^(0x2)|value<<1)
}
func (o *UART_Type) GetFLOW_CONF_XONOFF_DEL() uint32 {
	return (volatile.LoadUint32(&o.FLOW_CONF.Reg) & 0x2) >> 1
}
func (o *UART_Type) SetFLOW_CONF_FORCE_XON(value uint32) {
	volatile.StoreUint32(&o.FLOW_CONF.Reg, volatile.LoadUint32(&o.FLOW_CONF.Reg)&^(0x4)|value<<2)
}
func (o *UART_Type) GetFLOW_CONF_FORCE_XON() uint32 {
	return (volatile.LoadUint32(&o.FLOW_CONF.Reg) & 0x4) >> 2
}
func (o *UART_Type) SetFLOW_CONF_FORCE_XOFF(value uint32) {
	volatile.StoreUint32(&o.FLOW_CONF.Reg, volatile.LoadUint32(&o.FLOW_CONF.Reg)&^(0x8)|value<<3)
}
func (o *UART_Type) GetFLOW_CONF_FORCE_XOFF() uint32 {
	return (volatile.LoadUint32(&o.FLOW_CONF.Reg) & 0x8) >> 3
}
func (o *UART_Type) SetFLOW_CONF_SEND_XON(value uint32) {
	volatile.StoreUint32(&o.FLOW_CONF.Reg, volatile.LoadUint32(&o.FLOW_CONF.Reg)&^(0x10)|value<<4)
}
func (o *UART_Type) GetFLOW_CONF_SEND_XON() uint32 {
	return (volatile.LoadUint32(&o.FLOW_CONF.Reg) & 0x10) >> 4
}
func (o *UART_Type) SetFLOW_CONF_SEND_XOFF(value uint32) {
	volatile.StoreUint32(&o.FLOW_CONF.Reg, volatile.LoadUint32(&o.FLOW_CONF.Reg)&^(0x20)|value<<5)
}
func (o *UART_Type) GetFLOW_CONF_SEND_XOFF() uint32 {
	return (volatile.LoadUint32(&o.FLOW_CONF.Reg) & 0x20) >> 5
}

// UART.SLEEP_CONF
func (o *UART_Type) SetSLEEP_CONF_ACTIVE_THRESHOLD(value uint32) {
	volatile.StoreUint32(&o.SLEEP_CONF.Reg, volatile.LoadUint32(&o.SLEEP_CONF.Reg)&^(0x3ff)|value)
}
func (o *UART_Type) GetSLEEP_CONF_ACTIVE_THRESHOLD() uint32 {
	return volatile.LoadUint32(&o.SLEEP_CONF.Reg) & 0x3ff
}

// UART.SWFC_CONF
func (o *UART_Type) SetSWFC_CONF_XON_THRESHOLD(value uint32) {
	volatile.StoreUint32(&o.SWFC_CONF.Reg, volatile.LoadUint32(&o.SWFC_CONF.Reg)&^(0xff)|value)
}
func (o *UART_Type) GetSWFC_CONF_XON_THRESHOLD() uint32 {
	return volatile.LoadUint32(&o.SWFC_CONF.Reg) & 0xff
}
func (o *UART_Type) SetSWFC_CONF_XOFF_THRESHOLD(value uint32) {
	volatile.StoreUint32(&o.SWFC_CONF.Reg, volatile.LoadUint32(&o.SWFC_CONF.Reg)&^(0xff00)|value<<8)
}
func (o *UART_Type) GetSWFC_CONF_XOFF_THRESHOLD() uint32 {
	return (volatile.LoadUint32(&o.SWFC_CONF.Reg) & 0xff00) >> 8
}
func (o *UART_Type) SetSWFC_CONF_XON_CHAR(value uint32) {
	volatile.StoreUint32(&o.SWFC_CONF.Reg, volatile.LoadUint32(&o.SWFC_CONF.Reg)&^(0xff0000)|value<<16)
}
func (o *UART_Type) GetSWFC_CONF_XON_CHAR() uint32 {
	return (volatile.LoadUint32(&o.SWFC_CONF.Reg) & 0xff0000) >> 16
}
func (o *UART_Type) SetSWFC_CONF_XOFF_CHAR(value uint32) {
	volatile.StoreUint32(&o.SWFC_CONF.Reg, volatile.LoadUint32(&o.SWFC_CONF.Reg)&^(0xff000000)|value<<24)
}
func (o *UART_Type) GetSWFC_CONF_XOFF_CHAR() uint32 {
	return (volatile.LoadUint32(&o.SWFC_CONF.Reg) & 0xff000000) >> 24
}

// UART.IDLE_CONF
func (o *UART_Type) SetIDLE_CONF_RX_IDLE_THRHD(value uint32) {
	volatile.StoreUint32(&o.IDLE_CONF.Reg, volatile.LoadUint32(&o.IDLE_CONF.Reg)&^(0x3ff)|value)
}
func (o *UART_Type) GetIDLE_CONF_RX_IDLE_THRHD() uint32 {
	return volatile.LoadUint32(&o.IDLE_CONF.Reg) & 0x3ff
}
func (o *UART_Type) SetIDLE_CONF_TX_IDLE_NUM(value uint32) {
	volatile.StoreUint32(&o.IDLE_CONF.Reg, volatile.LoadUint32(&o.IDLE_CONF.Reg)&^(0xffc00)|value<<10)
}
func (o *UART_Type) GetIDLE_CONF_TX_IDLE_NUM() uint32 {
	return (volatile.LoadUint32(&o.IDLE_CONF.Reg) & 0xffc00) >> 10
}
func (o *UART_Type) SetIDLE_CONF_TX_BRK_NUM(value uint32) {
	volatile.StoreUint32(&o.IDLE_CONF.Reg, volatile.LoadUint32(&o.IDLE_CONF.Reg)&^(0xff00000)|value<<20)
}
func (o *UART_Type) GetIDLE_CONF_TX_BRK_NUM() uint32 {
	return (volatile.LoadUint32(&o.IDLE_CONF.Reg) & 0xff00000) >> 20
}

// UART.RS485_CONF
func (o *UART_Type) SetRS485_CONF_RS485_EN(value uint32) {
	volatile.StoreUint32(&o.RS485_CONF.Reg, volatile.LoadUint32(&o.RS485_CONF.Reg)&^(0x1)|value)
}
func (o *UART_Type) GetRS485_CONF_RS485_EN() uint32 {
	return volatile.LoadUint32(&o.RS485_CONF.Reg) & 0x1
}
func (o *UART_Type) SetRS485_CONF_DL0_EN(value uint32) {
	volatile.StoreUint32(&o.RS485_CONF.Reg, volatile.LoadUint32(&o.RS485_CONF.Reg)&^(0x2)|value<<1)
}
func (o *UART_Type) GetRS485_CONF_DL0_EN() uint32 {
	return (volatile.LoadUint32(&o.RS485_CONF.Reg) & 0x2) >> 1
}
func (o *UART_Type) SetRS485_CONF_DL1_EN(value uint32) {
	volatile.StoreUint32(&o.RS485_CONF.Reg, volatile.LoadUint32(&o.RS485_CONF.Reg)&^(0x4)|value<<2)
}
func (o *UART_Type) GetRS485_CONF_DL1_EN() uint32 {
	return (volatile.LoadUint32(&o.RS485_CONF.Reg) & 0x4) >> 2
}
func (o *UART_Type) SetRS485_CONF_RS485TX_RX_EN(value uint32) {
	volatile.StoreUint32(&o.RS485_CONF.Reg, volatile.LoadUint32(&o.RS485_CONF.Reg)&^(0x8)|value<<3)
}
func (o *UART_Type) GetRS485_CONF_RS485TX_RX_EN() uint32 {
	return (volatile.LoadUint32(&o.RS485_CONF.Reg) & 0x8) >> 3
}
func (o *UART_Type) SetRS485_CONF_RS485RXBY_TX_EN(value uint32) {
	volatile.StoreUint32(&o.RS485_CONF.Reg, volatile.LoadUint32(&o.RS485_CONF.Reg)&^(0x10)|value<<4)
}
func (o *UART_Type) GetRS485_CONF_RS485RXBY_TX_EN() uint32 {
	return (volatile.LoadUint32(&o.RS485_CONF.Reg) & 0x10) >> 4
}
func (o *UART_Type) SetRS485_CONF_RS485_RX_DLY_NUM(value uint32) {
	volatile.StoreUint32(&o.RS485_CONF.Reg, volatile.LoadUint32(&o.RS485_CONF.Reg)&^(0x20)|value<<5)
}
func (o *UART_Type) GetRS485_CONF_RS485_RX_DLY_NUM() uint32 {
	return (volatile.LoadUint32(&o.RS485_CONF.Reg) & 0x20) >> 5
}
func (o *UART_Type) SetRS485_CONF_RS485_TX_DLY_NUM(value uint32) {
	volatile.StoreUint32(&o.RS485_CONF.Reg, volatile.LoadUint32(&o.RS485_CONF.Reg)&^(0x3c0)|value<<6)
}
func (o *UART_Type) GetRS485_CONF_RS485_TX_DLY_NUM() uint32 {
	return (volatile.LoadUint32(&o.RS485_CONF.Reg) & 0x3c0) >> 6
}

// UART.AT_CMD_PRECNT
func (o *UART_Type) SetAT_CMD_PRECNT_PRE_IDLE_NUM(value uint32) {
	volatile.StoreUint32(&o.AT_CMD_PRECNT.Reg, volatile.LoadUint32(&o.AT_CMD_PRECNT.Reg)&^(0xffffff)|value)
}
func (o *UART_Type) GetAT_CMD_PRECNT_PRE_IDLE_NUM() uint32 {
	return volatile.LoadUint32(&o.AT_CMD_PRECNT.Reg) & 0xffffff
}

// UART.AT_CMD_POSTCNT
func (o *UART_Type) SetAT_CMD_POSTCNT_POST_IDLE_NUM(value uint32) {
	volatile.StoreUint32(&o.AT_CMD_POSTCNT.Reg, volatile.LoadUint32(&o.AT_CMD_POSTCNT.Reg)&^(0xffffff)|value)
}
func (o *UART_Type) GetAT_CMD_POSTCNT_POST_IDLE_NUM() uint32 {
	return volatile.LoadUint32(&o.AT_CMD_POSTCNT.Reg) & 0xffffff
}

// UART.AT_CMD_GAPTOUT
func (o *UART_Type) SetAT_CMD_GAPTOUT_RX_GAP_TOUT(value uint32) {
	volatile.StoreUint32(&o.AT_CMD_GAPTOUT.Reg, volatile.LoadUint32(&o.AT_CMD_GAPTOUT.Reg)&^(0xffffff)|value)
}
func (o *UART_Type) GetAT_CMD_GAPTOUT_RX_GAP_TOUT() uint32 {
	return volatile.LoadUint32(&o.AT_CMD_GAPTOUT.Reg) & 0xffffff
}

// UART.AT_CMD_CHAR
func (o *UART_Type) SetAT_CMD_CHAR(value uint32) {
	volatile.StoreUint32(&o.AT_CMD_CHAR.Reg, volatile.LoadUint32(&o.AT_CMD_CHAR.Reg)&^(0xff)|value)
}
func (o *UART_Type) GetAT_CMD_CHAR() uint32 {
	return volatile.LoadUint32(&o.AT_CMD_CHAR.Reg) & 0xff
}
func (o *UART_Type) SetAT_CMD_CHAR_CHAR_NUM(value uint32) {
	volatile.StoreUint32(&o.AT_CMD_CHAR.Reg, volatile.LoadUint32(&o.AT_CMD_CHAR.Reg)&^(0xff00)|value<<8)
}
func (o *UART_Type) GetAT_CMD_CHAR_CHAR_NUM() uint32 {
	return (volatile.LoadUint32(&o.AT_CMD_CHAR.Reg) & 0xff00) >> 8
}

// UART.MEM_CONF
func (o *UART_Type) SetMEM_CONF_MEM_PD(value uint32) {
	volatile.StoreUint32(&o.MEM_CONF.Reg, volatile.LoadUint32(&o.MEM_CONF.Reg)&^(0x1)|value)
}
func (o *UART_Type) GetMEM_CONF_MEM_PD() uint32 {
	return volatile.LoadUint32(&o.MEM_CONF.Reg) & 0x1
}
func (o *UART_Type) SetMEM_CONF_RX_SIZE(value uint32) {
	volatile.StoreUint32(&o.MEM_CONF.Reg, volatile.LoadUint32(&o.MEM_CONF.Reg)&^(0x78)|value<<3)
}
func (o *UART_Type) GetMEM_CONF_RX_SIZE() uint32 {
	return (volatile.LoadUint32(&o.MEM_CONF.Reg) & 0x78) >> 3
}
func (o *UART_Type) SetMEM_CONF_TX_SIZE(value uint32) {
	volatile.StoreUint32(&o.MEM_CONF.Reg, volatile.LoadUint32(&o.MEM_CONF.Reg)&^(0x780)|value<<7)
}
func (o *UART_Type) GetMEM_CONF_TX_SIZE() uint32 {
	return (volatile.LoadUint32(&o.MEM_CONF.Reg) & 0x780) >> 7
}
func (o *UART_Type) SetMEM_CONF_RX_FLOW_THRHD_H3(value uint32) {
	volatile.StoreUint32(&o.MEM_CONF.Reg, volatile.LoadUint32(&o.MEM_CONF.Reg)&^(0x38000)|value<<15)
}
func (o *UART_Type) GetMEM_CONF_RX_FLOW_THRHD_H3() uint32 {
	return (volatile.LoadUint32(&o.MEM_CONF.Reg) & 0x38000) >> 15
}
func (o *UART_Type) SetMEM_CONF_RX_TOUT_THRHD_H3(value uint32) {
	volatile.StoreUint32(&o.MEM_CONF.Reg, volatile.LoadUint32(&o.MEM_CONF.Reg)&^(0x1c0000)|value<<18)
}
func (o *UART_Type) GetMEM_CONF_RX_TOUT_THRHD_H3() uint32 {
	return (volatile.LoadUint32(&o.MEM_CONF.Reg) & 0x1c0000) >> 18
}
func (o *UART_Type) SetMEM_CONF_XON_THRESHOLD_H2(value uint32) {
	volatile.StoreUint32(&o.MEM_CONF.Reg, volatile.LoadUint32(&o.MEM_CONF.Reg)&^(0x600000)|value<<21)
}
func (o *UART_Type) GetMEM_CONF_XON_THRESHOLD_H2() uint32 {
	return (volatile.LoadUint32(&o.MEM_CONF.Reg) & 0x600000) >> 21
}
func (o *UART_Type) SetMEM_CONF_XOFF_THRESHOLD_H2(value uint32) {
	volatile.StoreUint32(&o.MEM_CONF.Reg, volatile.LoadUint32(&o.MEM_CONF.Reg)&^(0x1800000)|value<<23)
}
func (o *UART_Type) GetMEM_CONF_XOFF_THRESHOLD_H2() uint32 {
	return (volatile.LoadUint32(&o.MEM_CONF.Reg) & 0x1800000) >> 23
}
func (o *UART_Type) SetMEM_CONF_RX_MEM_FULL_THRHD(value uint32) {
	volatile.StoreUint32(&o.MEM_CONF.Reg, volatile.LoadUint32(&o.MEM_CONF.Reg)&^(0xe000000)|value<<25)
}
func (o *UART_Type) GetMEM_CONF_RX_MEM_FULL_THRHD() uint32 {
	return (volatile.LoadUint32(&o.MEM_CONF.Reg) & 0xe000000) >> 25
}
func (o *UART_Type) SetMEM_CONF_TX_MEM_EMPTY_THRHD(value uint32) {
	volatile.StoreUint32(&o.MEM_CONF.Reg, volatile.LoadUint32(&o.MEM_CONF.Reg)&^(0x70000000)|value<<28)
}
func (o *UART_Type) GetMEM_CONF_TX_MEM_EMPTY_THRHD() uint32 {
	return (volatile.LoadUint32(&o.MEM_CONF.Reg) & 0x70000000) >> 28
}

// UART.MEM_TX_STATUS
func (o *UART_Type) SetMEM_TX_STATUS(value uint32) {
	volatile.StoreUint32(&o.MEM_TX_STATUS.Reg, volatile.LoadUint32(&o.MEM_TX_STATUS.Reg)&^(0xffffff)|value)
}
func (o *UART_Type) GetMEM_TX_STATUS() uint32 {
	return volatile.LoadUint32(&o.MEM_TX_STATUS.Reg) & 0xffffff
}

// UART.MEM_RX_STATUS
func (o *UART_Type) SetMEM_RX_STATUS(value uint32) {
	volatile.StoreUint32(&o.MEM_RX_STATUS.Reg, volatile.LoadUint32(&o.MEM_RX_STATUS.Reg)&^(0xffffff)|value)
}
func (o *UART_Type) GetMEM_RX_STATUS() uint32 {
	return volatile.LoadUint32(&o.MEM_RX_STATUS.Reg) & 0xffffff
}
func (o *UART_Type) SetMEM_RX_STATUS_MEM_RX_RD_ADDR(value uint32) {
	volatile.StoreUint32(&o.MEM_RX_STATUS.Reg, volatile.LoadUint32(&o.MEM_RX_STATUS.Reg)&^(0x1ffc)|value<<2)
}
func (o *UART_Type) GetMEM_RX_STATUS_MEM_RX_RD_ADDR() uint32 {
	return (volatile.LoadUint32(&o.MEM_RX_STATUS.Reg) & 0x1ffc) >> 2
}
func (o *UART_Type) SetMEM_RX_STATUS_MEM_RX_WR_ADDR(value uint32) {
	volatile.StoreUint32(&o.MEM_RX_STATUS.Reg, volatile.LoadUint32(&o.MEM_RX_STATUS.Reg)&^(0xffe000)|value<<13)
}
func (o *UART_Type) GetMEM_RX_STATUS_MEM_RX_WR_ADDR() uint32 {
	return (volatile.LoadUint32(&o.MEM_RX_STATUS.Reg) & 0xffe000) >> 13
}

// UART.MEM_CNT_STATUS
func (o *UART_Type) SetMEM_CNT_STATUS_RX_MEM_CNT(value uint32) {
	volatile.StoreUint32(&o.MEM_CNT_STATUS.Reg, volatile.LoadUint32(&o.MEM_CNT_STATUS.Reg)&^(0x7)|value)
}
func (o *UART_Type) GetMEM_CNT_STATUS_RX_MEM_CNT() uint32 {
	return volatile.LoadUint32(&o.MEM_CNT_STATUS.Reg) & 0x7
}
func (o *UART_Type) SetMEM_CNT_STATUS_TX_MEM_CNT(value uint32) {
	volatile.StoreUint32(&o.MEM_CNT_STATUS.Reg, volatile.LoadUint32(&o.MEM_CNT_STATUS.Reg)&^(0x38)|value<<3)
}
func (o *UART_Type) GetMEM_CNT_STATUS_TX_MEM_CNT() uint32 {
	return (volatile.LoadUint32(&o.MEM_CNT_STATUS.Reg) & 0x38) >> 3
}

// UART.POSPULSE
func (o *UART_Type) SetPOSPULSE_POSEDGE_MIN_CNT(value uint32) {
	volatile.StoreUint32(&o.POSPULSE.Reg, volatile.LoadUint32(&o.POSPULSE.Reg)&^(0xfffff)|value)
}
func (o *UART_Type) GetPOSPULSE_POSEDGE_MIN_CNT() uint32 {
	return volatile.LoadUint32(&o.POSPULSE.Reg) & 0xfffff
}

// UART.NEGPULSE
func (o *UART_Type) SetNEGPULSE_NEGEDGE_MIN_CNT(value uint32) {
	volatile.StoreUint32(&o.NEGPULSE.Reg, volatile.LoadUint32(&o.NEGPULSE.Reg)&^(0xfffff)|value)
}
func (o *UART_Type) GetNEGPULSE_NEGEDGE_MIN_CNT() uint32 {
	return volatile.LoadUint32(&o.NEGPULSE.Reg) & 0xfffff
}

// UART.DATE
func (o *UART_Type) SetDATE(value uint32) {
	volatile.StoreUint32(&o.DATE.Reg, value)
}
func (o *UART_Type) GetDATE() uint32 {
	return volatile.LoadUint32(&o.DATE.Reg)
}

// UART.ID
func (o *UART_Type) SetID(value uint32) {
	volatile.StoreUint32(&o.ID.Reg, value)
}
func (o *UART_Type) GetID() uint32 {
	return volatile.LoadUint32(&o.ID.Reg)
}
