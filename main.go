package main

import (
	"unsafe"

	"github.com/goplus/lib/c"

	_ "llgo-test/libc"
)

// 直接定义硬件地址常量，避免创建全局变量
const (
	GPIO_BASE     = 0x3ff44000
	IO_MUX_BASE   = 0x3ff49000
	RTC_CNTL_BASE = 0x3ff48000
	TIMG0_BASE    = 0x3ff5f000
)

func getTIMG() *TIMG_Type {
	return (*TIMG_Type)(unsafe.Pointer(uintptr(TIMG0_BASE)))
}

// 获取外设指针的函数
func getGPIO() *GPIO_Type {
	return (*GPIO_Type)(unsafe.Pointer(uintptr(GPIO_BASE)))
}

func getIOMUX() *IO_MUX_Type {
	return (*IO_MUX_Type)(unsafe.Pointer(uintptr(IO_MUX_BASE)))
}

func getRTCCNTL() *RTC_CNTL_Type {
	return (*RTC_CNTL_Type)(unsafe.Pointer(uintptr(RTC_CNTL_BASE)))
}

//go:linkname sbss _sbss
var sbss [0]byte

//go:linkname ebss _ebss
var ebss [0]byte

func enableButton() {
	const btnGPIO = 34
	var btnGPIOMux = &getIOMUX().GPIO34

	var muxConfig uint32 // The mux configuration.

	// Configure this pin as a GPIO pin.
	const function = 3 // function 3 is GPIO for every pin
	muxConfig |= (function - 1) << IO_MUX_GPIO0_MCU_SEL_Pos

	// Make this pin an input pin (always).
	muxConfig |= IO_MUX_GPIO0_FUN_IE

	// Set drive strength: 0 is lowest, 3 is highest.
	muxConfig |= 2 << IO_MUX_GPIO0_FUN_DRV_Pos
	btnGPIOMux.Set(muxConfig)

	// gpio input enable
	getGPIO().ENABLE1_W1TC.Set(1 << (btnGPIO - 32))

}

func getButtonStatus() bool {
	const btnGPIO = 34

	return getGPIO().IN1.Get()&(1<<(btnGPIO-32)) != 0
}

func clearbss() {

	sbssPtr := unsafe.Pointer(&sbss)
	ebssPtr := unsafe.Pointer(&ebss)
	for ptr := sbssPtr; ptr != ebssPtr; ptr = unsafe.Add(ptr, 4) {
		val := *(*uint32)(ptr)
		printIntHex(int(val))
	}

	size := uintptr(ebssPtr) - uintptr(sbssPtr)

	c.Memset(sbssPtr, 0, size)

	printOK()

	for ptr := sbssPtr; ptr != ebssPtr; ptr = unsafe.Add(ptr, 4) {
		val := *(*uint32)(ptr)
		printIntHex(int(val))
	}
}

func main() {

	// 初始化UART用于调试输出
	initUART()

	clearbss()
	// Disable the protection on the watchdog timer (needed when started from
	// the bootloader).
	getRTCCNTL().WDTWPROTECT.Set(0x50D83AA1)

	// Disable both watchdog timers that are enabled by default on startup.
	// Note that these watchdogs can be protected, but the ROM bootloader
	// doesn't seem to protect them.
	getRTCCNTL().WDTCONFIG0.Set(0)

	getTIMG().WDTCONFIG0.Set(0)

	// Switch SoC clock source to PLL (instead of the default which is XTAL).
	// This switches the CPU (and APB) clock from 40MHz to 80MHz.
	// Options:
	//   RTC_CNTL_CLK_CONF_SOC_CLK_SEL:       PLL (1)       (default XTAL)
	//   RTC_CNTL_CLK_CONF_CK8M_DIV_SEL:      2             (default)
	//   RTC_CNTL_CLK_CONF_DIG_CLK8M_D256_EN: Enable        (default)
	//   RTC_CNTL_CLK_CONF_CK8M_DIV:          divide by 256 (default)
	// The only real change made here is modifying RTC_CNTL_CLK_CONF_SOC_CLK_SEL,
	// but setting a fixed value produces smaller code.
	getRTCCNTL().CLK_CONF.Set((1 << RTC_CNTL_CLK_CONF_SOC_CLK_SEL_Pos) |
		(2 << RTC_CNTL_CLK_CONF_CK8M_DIV_SEL_Pos) |
		(1 << RTC_CNTL_CLK_CONF_DIG_CLK8M_D256_EN_Pos) |
		(1 << RTC_CNTL_CLK_CONF_CK8M_DIV_Pos))

	getTIMG().T0CONFIG.Set(TIMG_T0CONFIG_EN | TIMG_T0CONFIG_INCREASE | 2<<TIMG_T0CONFIG_DIVIDER_Pos)
	// esp.TIMG0.T0CONFIG.Set(1 << esp.TIMG_T0CONFIG_T0_DIVCNT_RST_Pos)
	// esp.TIMG0.T0CONFIG.Set(esp.TIMG_T0CONFIG_T0_EN)

	// Set the timer counter value to 0.
	getTIMG().T0LOADLO.Set(0)
	getTIMG().T0LOADHI.Set(0)
	getTIMG().T0LOAD.Set(0)

	enableButton()

	c.Printf(c.Str("111"))

	// // 主循环 - 可以添加按钮状态检测
	for {
		if !getButtonStatus() {
			printDone()
		}
	}
}
