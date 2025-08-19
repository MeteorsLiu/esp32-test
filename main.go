package main

import "unsafe"

// 直接定义硬件地址常量，避免创建全局变量
const (
	GPIO_BASE     = 0x3ff44000
	IO_MUX_BASE   = 0x3ff49000
	RTC_CNTL_BASE = 0x3ff48000
)

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

const (
	LLGoFiles   = "esp32.S"
	LLGoPackage = "link"
)

func enableButton() {
	const btnGPIO = 34
	var btnGPIOMux = &getIOMUX().GPIO35

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
	ptr := unsafe.Pointer(&_sbss)
	for ptr != unsafe.Pointer(&_ebss) {
		*(*uint32)(ptr) = 0
		ptr = unsafe.Add(ptr, 4)
	}
}

func main() {
	// 初始化UART用于调试输出
	initUART()
	printStartup()

	// Disable the protection on the watchdog timer (needed when started from
	// the bootloader).
	getRTCCNTL().WDTWPROTECT.Set(0x50D83AA1)
	printOK()

	// Disable both watchdog timers that are enabled by default on startup.
	// Note that these watchdogs can be protected, but the ROM bootloader
	// doesn't seem to protect them.
	getRTCCNTL().WDTCONFIG0.Set(0)
	printOK()

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
	printOK()

	clearbss()
	printOK()

	enableButton()
	printOK()

	printDone()

	// 主循环 - 可以添加按钮状态检测
	for {
		// 可以在这里添加按钮状态检测和其他功能
	}
}
