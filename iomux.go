package main

import "llgo-test/volatile"

// Constants for IO_MUX: Input/Output Multiplexer
const (
	// PIN_CTRL
	// Position of CLK1 field.
	IO_MUX_PIN_CTRL_CLK1_Pos = 0x0
	// Bit mask of CLK1 field.
	IO_MUX_PIN_CTRL_CLK1_Msk = 0xf
	// Position of CLK2 field.
	IO_MUX_PIN_CTRL_CLK2_Pos = 0x4
	// Bit mask of CLK2 field.
	IO_MUX_PIN_CTRL_CLK2_Msk = 0xf0
	// Position of CLK3 field.
	IO_MUX_PIN_CTRL_CLK3_Pos = 0x8
	// Bit mask of CLK3 field.
	IO_MUX_PIN_CTRL_CLK3_Msk = 0xf00

	// GPIO36
	// Position of MCU_OE field.
	IO_MUX_GPIO36_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO36_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO36_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO36_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO36_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO36_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO36_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO36_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO36_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO36_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO36_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO36_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO36_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO36_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO36_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO36_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO36_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO36_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO36_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO36_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO36_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO36_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO36_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO36_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO36_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO36_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO36_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO36_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO36_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO36_MCU_SEL_Msk = 0x7000

	// GPIO37
	// Position of MCU_OE field.
	IO_MUX_GPIO37_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO37_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO37_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO37_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO37_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO37_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO37_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO37_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO37_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO37_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO37_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO37_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO37_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO37_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO37_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO37_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO37_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO37_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO37_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO37_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO37_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO37_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO37_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO37_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO37_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO37_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO37_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO37_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO37_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO37_MCU_SEL_Msk = 0x7000

	// GPIO38
	// Position of MCU_OE field.
	IO_MUX_GPIO38_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO38_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO38_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO38_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO38_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO38_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO38_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO38_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO38_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO38_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO38_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO38_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO38_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO38_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO38_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO38_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO38_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO38_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO38_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO38_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO38_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO38_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO38_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO38_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO38_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO38_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO38_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO38_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO38_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO38_MCU_SEL_Msk = 0x7000

	// GPIO39
	// Position of MCU_OE field.
	IO_MUX_GPIO39_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO39_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO39_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO39_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO39_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO39_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO39_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO39_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO39_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO39_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO39_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO39_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO39_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO39_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO39_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO39_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO39_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO39_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO39_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO39_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO39_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO39_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO39_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO39_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO39_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO39_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO39_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO39_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO39_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO39_MCU_SEL_Msk = 0x7000

	// GPIO34
	// Position of MCU_OE field.
	IO_MUX_GPIO34_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO34_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO34_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO34_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO34_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO34_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO34_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO34_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO34_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO34_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO34_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO34_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO34_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO34_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO34_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO34_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO34_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO34_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO34_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO34_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO34_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO34_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO34_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO34_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO34_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO34_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO34_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO34_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO34_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO34_MCU_SEL_Msk = 0x7000

	// GPIO35
	// Position of MCU_OE field.
	IO_MUX_GPIO35_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO35_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO35_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO35_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO35_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO35_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO35_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO35_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO35_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO35_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO35_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO35_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO35_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO35_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO35_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO35_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO35_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO35_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO35_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO35_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO35_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO35_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO35_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO35_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO35_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO35_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO35_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO35_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO35_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO35_MCU_SEL_Msk = 0x7000

	// GPIO32
	// Position of MCU_OE field.
	IO_MUX_GPIO32_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO32_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO32_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO32_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO32_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO32_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO32_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO32_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO32_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO32_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO32_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO32_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO32_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO32_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO32_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO32_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO32_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO32_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO32_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO32_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO32_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO32_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO32_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO32_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO32_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO32_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO32_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO32_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO32_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO32_MCU_SEL_Msk = 0x7000

	// GPIO33
	// Position of MCU_OE field.
	IO_MUX_GPIO33_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO33_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO33_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO33_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO33_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO33_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO33_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO33_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO33_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO33_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO33_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO33_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO33_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO33_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO33_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO33_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO33_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO33_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO33_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO33_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO33_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO33_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO33_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO33_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO33_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO33_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO33_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO33_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO33_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO33_MCU_SEL_Msk = 0x7000

	// GPIO25
	// Position of MCU_OE field.
	IO_MUX_GPIO25_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO25_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO25_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO25_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO25_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO25_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO25_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO25_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO25_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO25_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO25_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO25_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO25_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO25_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO25_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO25_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO25_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO25_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO25_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO25_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO25_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO25_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO25_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO25_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO25_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO25_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO25_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO25_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO25_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO25_MCU_SEL_Msk = 0x7000

	// GPIO26
	// Position of MCU_OE field.
	IO_MUX_GPIO26_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO26_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO26_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO26_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO26_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO26_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO26_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO26_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO26_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO26_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO26_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO26_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO26_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO26_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO26_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO26_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO26_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO26_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO26_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO26_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO26_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO26_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO26_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO26_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO26_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO26_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO26_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO26_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO26_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO26_MCU_SEL_Msk = 0x7000

	// GPIO27
	// Position of MCU_OE field.
	IO_MUX_GPIO27_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO27_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO27_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO27_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO27_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO27_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO27_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO27_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO27_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO27_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO27_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO27_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO27_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO27_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO27_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO27_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO27_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO27_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO27_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO27_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO27_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO27_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO27_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO27_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO27_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO27_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO27_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO27_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO27_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO27_MCU_SEL_Msk = 0x7000

	// GPIO14
	// Position of MCU_OE field.
	IO_MUX_GPIO14_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO14_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO14_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO14_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO14_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO14_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO14_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO14_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO14_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO14_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO14_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO14_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO14_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO14_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO14_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO14_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO14_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO14_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO14_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO14_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO14_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO14_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO14_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO14_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO14_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO14_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO14_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO14_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO14_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO14_MCU_SEL_Msk = 0x7000

	// GPIO12
	// Position of MCU_OE field.
	IO_MUX_GPIO12_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO12_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO12_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO12_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO12_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO12_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO12_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO12_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO12_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO12_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO12_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO12_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO12_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO12_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO12_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO12_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO12_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO12_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO12_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO12_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO12_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO12_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO12_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO12_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO12_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO12_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO12_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO12_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO12_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO12_MCU_SEL_Msk = 0x7000

	// GPIO13
	// Position of MCU_OE field.
	IO_MUX_GPIO13_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO13_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO13_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO13_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO13_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO13_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO13_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO13_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO13_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO13_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO13_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO13_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO13_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO13_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO13_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO13_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO13_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO13_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO13_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO13_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO13_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO13_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO13_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO13_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO13_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO13_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO13_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO13_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO13_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO13_MCU_SEL_Msk = 0x7000

	// GPIO15
	// Position of MCU_OE field.
	IO_MUX_GPIO15_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO15_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO15_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO15_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO15_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO15_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO15_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO15_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO15_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO15_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO15_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO15_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO15_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO15_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO15_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO15_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO15_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO15_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO15_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO15_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO15_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO15_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO15_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO15_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO15_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO15_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO15_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO15_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO15_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO15_MCU_SEL_Msk = 0x7000

	// GPIO2
	// Position of MCU_OE field.
	IO_MUX_GPIO2_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO2_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO2_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO2_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO2_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO2_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO2_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO2_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO2_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO2_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO2_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO2_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO2_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO2_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO2_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO2_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO2_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO2_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO2_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO2_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO2_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO2_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO2_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO2_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO2_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO2_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO2_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO2_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO2_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO2_MCU_SEL_Msk = 0x7000

	// GPIO0
	// Position of MCU_OE field.
	IO_MUX_GPIO0_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO0_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO0_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO0_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO0_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO0_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO0_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO0_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO0_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO0_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO0_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO0_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO0_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO0_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO0_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO0_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO0_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO0_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO0_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO0_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO0_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO0_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO0_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO0_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO0_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO0_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO0_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO0_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO0_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO0_MCU_SEL_Msk = 0x7000

	// GPIO4
	// Position of MCU_OE field.
	IO_MUX_GPIO4_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO4_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO4_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO4_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO4_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO4_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO4_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO4_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO4_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO4_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO4_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO4_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO4_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO4_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO4_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO4_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO4_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO4_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO4_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO4_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO4_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO4_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO4_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO4_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO4_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO4_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO4_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO4_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO4_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO4_MCU_SEL_Msk = 0x7000

	// GPIO16
	// Position of MCU_OE field.
	IO_MUX_GPIO16_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO16_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO16_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO16_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO16_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO16_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO16_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO16_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO16_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO16_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO16_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO16_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO16_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO16_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO16_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO16_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO16_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO16_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO16_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO16_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO16_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO16_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO16_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO16_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO16_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO16_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO16_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO16_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO16_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO16_MCU_SEL_Msk = 0x7000

	// GPIO17
	// Position of MCU_OE field.
	IO_MUX_GPIO17_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO17_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO17_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO17_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO17_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO17_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO17_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO17_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO17_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO17_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO17_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO17_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO17_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO17_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO17_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO17_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO17_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO17_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO17_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO17_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO17_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO17_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO17_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO17_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO17_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO17_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO17_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO17_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO17_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO17_MCU_SEL_Msk = 0x7000

	// GPIO9
	// Position of MCU_OE field.
	IO_MUX_GPIO9_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO9_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO9_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO9_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO9_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO9_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO9_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO9_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO9_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO9_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO9_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO9_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO9_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO9_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO9_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO9_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO9_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO9_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO9_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO9_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO9_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO9_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO9_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO9_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO9_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO9_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO9_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO9_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO9_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO9_MCU_SEL_Msk = 0x7000

	// GPIO10
	// Position of MCU_OE field.
	IO_MUX_GPIO10_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO10_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO10_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO10_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO10_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO10_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO10_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO10_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO10_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO10_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO10_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO10_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO10_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO10_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO10_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO10_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO10_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO10_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO10_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO10_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO10_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO10_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO10_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO10_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO10_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO10_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO10_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO10_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO10_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO10_MCU_SEL_Msk = 0x7000

	// GPIO11
	// Position of MCU_OE field.
	IO_MUX_GPIO11_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO11_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO11_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO11_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO11_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO11_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO11_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO11_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO11_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO11_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO11_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO11_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO11_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO11_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO11_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO11_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO11_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO11_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO11_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO11_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO11_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO11_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO11_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO11_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO11_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO11_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO11_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO11_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO11_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO11_MCU_SEL_Msk = 0x7000

	// GPIO6
	// Position of MCU_OE field.
	IO_MUX_GPIO6_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO6_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO6_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO6_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO6_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO6_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO6_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO6_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO6_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO6_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO6_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO6_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO6_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO6_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO6_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO6_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO6_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO6_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO6_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO6_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO6_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO6_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO6_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO6_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO6_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO6_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO6_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO6_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO6_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO6_MCU_SEL_Msk = 0x7000

	// GPIO7
	// Position of MCU_OE field.
	IO_MUX_GPIO7_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO7_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO7_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO7_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO7_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO7_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO7_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO7_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO7_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO7_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO7_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO7_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO7_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO7_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO7_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO7_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO7_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO7_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO7_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO7_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO7_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO7_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO7_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO7_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO7_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO7_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO7_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO7_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO7_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO7_MCU_SEL_Msk = 0x7000

	// GPIO8
	// Position of MCU_OE field.
	IO_MUX_GPIO8_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO8_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO8_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO8_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO8_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO8_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO8_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO8_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO8_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO8_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO8_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO8_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO8_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO8_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO8_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO8_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO8_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO8_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO8_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO8_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO8_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO8_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO8_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO8_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO8_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO8_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO8_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO8_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO8_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO8_MCU_SEL_Msk = 0x7000

	// GPIO5
	// Position of MCU_OE field.
	IO_MUX_GPIO5_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO5_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO5_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO5_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO5_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO5_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO5_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO5_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO5_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO5_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO5_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO5_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO5_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO5_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO5_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO5_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO5_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO5_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO5_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO5_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO5_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO5_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO5_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO5_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO5_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO5_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO5_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO5_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO5_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO5_MCU_SEL_Msk = 0x7000

	// GPIO18
	// Position of MCU_OE field.
	IO_MUX_GPIO18_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO18_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO18_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO18_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO18_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO18_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO18_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO18_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO18_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO18_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO18_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO18_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO18_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO18_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO18_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO18_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO18_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO18_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO18_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO18_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO18_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO18_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO18_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO18_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO18_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO18_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO18_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO18_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO18_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO18_MCU_SEL_Msk = 0x7000

	// GPIO19
	// Position of MCU_OE field.
	IO_MUX_GPIO19_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO19_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO19_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO19_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO19_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO19_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO19_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO19_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO19_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO19_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO19_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO19_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO19_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO19_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO19_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO19_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO19_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO19_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO19_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO19_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO19_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO19_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO19_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO19_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO19_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO19_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO19_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO19_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO19_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO19_MCU_SEL_Msk = 0x7000

	// GPIO20
	// Position of MCU_OE field.
	IO_MUX_GPIO20_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO20_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO20_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO20_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO20_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO20_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO20_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO20_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO20_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO20_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO20_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO20_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO20_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO20_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO20_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO20_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO20_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO20_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO20_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO20_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO20_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO20_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO20_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO20_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO20_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO20_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO20_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO20_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO20_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO20_MCU_SEL_Msk = 0x7000

	// GPIO21
	// Position of MCU_OE field.
	IO_MUX_GPIO21_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO21_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO21_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO21_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO21_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO21_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO21_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO21_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO21_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO21_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO21_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO21_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO21_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO21_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO21_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO21_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO21_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO21_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO21_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO21_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO21_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO21_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO21_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO21_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO21_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO21_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO21_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO21_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO21_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO21_MCU_SEL_Msk = 0x7000

	// GPIO22
	// Position of MCU_OE field.
	IO_MUX_GPIO22_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO22_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO22_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO22_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO22_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO22_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO22_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO22_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO22_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO22_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO22_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO22_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO22_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO22_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO22_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO22_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO22_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO22_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO22_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO22_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO22_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO22_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO22_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO22_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO22_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO22_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO22_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO22_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO22_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO22_MCU_SEL_Msk = 0x7000

	// GPIO3
	// Position of MCU_OE field.
	IO_MUX_GPIO3_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO3_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO3_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO3_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO3_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO3_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO3_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO3_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO3_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO3_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO3_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO3_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO3_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO3_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO3_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO3_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO3_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO3_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO3_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO3_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO3_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO3_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO3_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO3_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO3_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO3_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO3_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO3_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO3_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO3_MCU_SEL_Msk = 0x7000

	// GPIO1
	// Position of MCU_OE field.
	IO_MUX_GPIO1_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO1_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO1_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO1_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO1_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO1_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO1_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO1_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO1_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO1_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO1_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO1_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO1_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO1_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO1_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO1_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO1_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO1_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO1_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO1_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO1_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO1_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO1_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO1_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO1_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO1_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO1_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO1_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO1_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO1_MCU_SEL_Msk = 0x7000

	// GPIO23
	// Position of MCU_OE field.
	IO_MUX_GPIO23_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO23_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO23_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO23_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO23_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO23_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO23_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO23_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO23_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO23_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO23_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO23_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO23_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO23_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO23_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO23_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO23_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO23_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO23_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO23_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO23_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO23_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO23_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO23_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO23_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO23_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO23_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO23_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO23_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO23_MCU_SEL_Msk = 0x7000

	// GPIO24
	// Position of MCU_OE field.
	IO_MUX_GPIO24_MCU_OE_Pos = 0x0
	// Bit mask of MCU_OE field.
	IO_MUX_GPIO24_MCU_OE_Msk = 0x1
	// Bit MCU_OE.
	IO_MUX_GPIO24_MCU_OE = 0x1
	// Position of SLP_SEL field.
	IO_MUX_GPIO24_SLP_SEL_Pos = 0x1
	// Bit mask of SLP_SEL field.
	IO_MUX_GPIO24_SLP_SEL_Msk = 0x2
	// Bit SLP_SEL.
	IO_MUX_GPIO24_SLP_SEL = 0x2
	// Position of MCU_WPD field.
	IO_MUX_GPIO24_MCU_WPD_Pos = 0x2
	// Bit mask of MCU_WPD field.
	IO_MUX_GPIO24_MCU_WPD_Msk = 0x4
	// Bit MCU_WPD.
	IO_MUX_GPIO24_MCU_WPD = 0x4
	// Position of MCU_WPU field.
	IO_MUX_GPIO24_MCU_WPU_Pos = 0x3
	// Bit mask of MCU_WPU field.
	IO_MUX_GPIO24_MCU_WPU_Msk = 0x8
	// Bit MCU_WPU.
	IO_MUX_GPIO24_MCU_WPU = 0x8
	// Position of MCU_IE field.
	IO_MUX_GPIO24_MCU_IE_Pos = 0x4
	// Bit mask of MCU_IE field.
	IO_MUX_GPIO24_MCU_IE_Msk = 0x10
	// Bit MCU_IE.
	IO_MUX_GPIO24_MCU_IE = 0x10
	// Position of MCU_DRV field.
	IO_MUX_GPIO24_MCU_DRV_Pos = 0x5
	// Bit mask of MCU_DRV field.
	IO_MUX_GPIO24_MCU_DRV_Msk = 0x60
	// Position of FUN_WPD field.
	IO_MUX_GPIO24_FUN_WPD_Pos = 0x7
	// Bit mask of FUN_WPD field.
	IO_MUX_GPIO24_FUN_WPD_Msk = 0x80
	// Bit FUN_WPD.
	IO_MUX_GPIO24_FUN_WPD = 0x80
	// Position of FUN_WPU field.
	IO_MUX_GPIO24_FUN_WPU_Pos = 0x8
	// Bit mask of FUN_WPU field.
	IO_MUX_GPIO24_FUN_WPU_Msk = 0x100
	// Bit FUN_WPU.
	IO_MUX_GPIO24_FUN_WPU = 0x100
	// Position of FUN_IE field.
	IO_MUX_GPIO24_FUN_IE_Pos = 0x9
	// Bit mask of FUN_IE field.
	IO_MUX_GPIO24_FUN_IE_Msk = 0x200
	// Bit FUN_IE.
	IO_MUX_GPIO24_FUN_IE = 0x200
	// Position of FUN_DRV field.
	IO_MUX_GPIO24_FUN_DRV_Pos = 0xa
	// Bit mask of FUN_DRV field.
	IO_MUX_GPIO24_FUN_DRV_Msk = 0xc00
	// Position of MCU_SEL field.
	IO_MUX_GPIO24_MCU_SEL_Pos = 0xc
	// Bit mask of MCU_SEL field.
	IO_MUX_GPIO24_MCU_SEL_Msk = 0x7000
)

// Input/Output Multiplexer
type IO_MUX_Type struct {
	PIN_CTRL volatile.Register32 // 0x0
	GPIO36   volatile.Register32 // 0x4
	GPIO37   volatile.Register32 // 0x8
	GPIO38   volatile.Register32 // 0xC
	GPIO39   volatile.Register32 // 0x10
	GPIO34   volatile.Register32 // 0x14
	GPIO35   volatile.Register32 // 0x18
	GPIO32   volatile.Register32 // 0x1C
	GPIO33   volatile.Register32 // 0x20
	GPIO25   volatile.Register32 // 0x24
	GPIO26   volatile.Register32 // 0x28
	GPIO27   volatile.Register32 // 0x2C
	GPIO14   volatile.Register32 // 0x30
	GPIO12   volatile.Register32 // 0x34
	GPIO13   volatile.Register32 // 0x38
	GPIO15   volatile.Register32 // 0x3C
	GPIO2    volatile.Register32 // 0x40
	GPIO0    volatile.Register32 // 0x44
	GPIO4    volatile.Register32 // 0x48
	GPIO16   volatile.Register32 // 0x4C
	GPIO17   volatile.Register32 // 0x50
	GPIO9    volatile.Register32 // 0x54
	GPIO10   volatile.Register32 // 0x58
	GPIO11   volatile.Register32 // 0x5C
	GPIO6    volatile.Register32 // 0x60
	GPIO7    volatile.Register32 // 0x64
	GPIO8    volatile.Register32 // 0x68
	GPIO5    volatile.Register32 // 0x6C
	GPIO18   volatile.Register32 // 0x70
	GPIO19   volatile.Register32 // 0x74
	GPIO20   volatile.Register32 // 0x78
	GPIO21   volatile.Register32 // 0x7C
	GPIO22   volatile.Register32 // 0x80
	GPIO3    volatile.Register32 // 0x84
	GPIO1    volatile.Register32 // 0x88
	GPIO23   volatile.Register32 // 0x8C
	GPIO24   volatile.Register32 // 0x90
}

// IO_MUX.PIN_CTRL
func (o *IO_MUX_Type) SetPIN_CTRL_CLK1(value uint32) {
	volatile.StoreUint32(&o.PIN_CTRL.Reg, volatile.LoadUint32(&o.PIN_CTRL.Reg)&^(0xf)|value)
}
func (o *IO_MUX_Type) GetPIN_CTRL_CLK1() uint32 {
	return volatile.LoadUint32(&o.PIN_CTRL.Reg) & 0xf
}
func (o *IO_MUX_Type) SetPIN_CTRL_CLK2(value uint32) {
	volatile.StoreUint32(&o.PIN_CTRL.Reg, volatile.LoadUint32(&o.PIN_CTRL.Reg)&^(0xf0)|value<<4)
}
func (o *IO_MUX_Type) GetPIN_CTRL_CLK2() uint32 {
	return (volatile.LoadUint32(&o.PIN_CTRL.Reg) & 0xf0) >> 4
}
func (o *IO_MUX_Type) SetPIN_CTRL_CLK3(value uint32) {
	volatile.StoreUint32(&o.PIN_CTRL.Reg, volatile.LoadUint32(&o.PIN_CTRL.Reg)&^(0xf00)|value<<8)
}
func (o *IO_MUX_Type) GetPIN_CTRL_CLK3() uint32 {
	return (volatile.LoadUint32(&o.PIN_CTRL.Reg) & 0xf00) >> 8
}

// IO_MUX.GPIO36
func (o *IO_MUX_Type) SetGPIO36_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO36.Reg, volatile.LoadUint32(&o.GPIO36.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO36_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO36.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO36_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO36.Reg, volatile.LoadUint32(&o.GPIO36.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO36_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO36.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO36_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO36.Reg, volatile.LoadUint32(&o.GPIO36.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO36_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO36.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO36_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO36.Reg, volatile.LoadUint32(&o.GPIO36.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO36_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO36.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO36_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO36.Reg, volatile.LoadUint32(&o.GPIO36.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO36_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO36.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO36_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO36.Reg, volatile.LoadUint32(&o.GPIO36.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO36_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO36.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO36_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO36.Reg, volatile.LoadUint32(&o.GPIO36.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO36_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO36.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO36_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO36.Reg, volatile.LoadUint32(&o.GPIO36.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO36_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO36.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO36_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO36.Reg, volatile.LoadUint32(&o.GPIO36.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO36_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO36.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO36_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO36.Reg, volatile.LoadUint32(&o.GPIO36.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO36_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO36.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO36_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO36.Reg, volatile.LoadUint32(&o.GPIO36.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO36_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO36.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO37
func (o *IO_MUX_Type) SetGPIO37_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO37.Reg, volatile.LoadUint32(&o.GPIO37.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO37_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO37.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO37_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO37.Reg, volatile.LoadUint32(&o.GPIO37.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO37_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO37.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO37_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO37.Reg, volatile.LoadUint32(&o.GPIO37.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO37_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO37.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO37_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO37.Reg, volatile.LoadUint32(&o.GPIO37.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO37_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO37.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO37_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO37.Reg, volatile.LoadUint32(&o.GPIO37.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO37_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO37.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO37_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO37.Reg, volatile.LoadUint32(&o.GPIO37.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO37_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO37.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO37_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO37.Reg, volatile.LoadUint32(&o.GPIO37.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO37_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO37.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO37_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO37.Reg, volatile.LoadUint32(&o.GPIO37.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO37_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO37.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO37_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO37.Reg, volatile.LoadUint32(&o.GPIO37.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO37_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO37.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO37_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO37.Reg, volatile.LoadUint32(&o.GPIO37.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO37_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO37.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO37_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO37.Reg, volatile.LoadUint32(&o.GPIO37.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO37_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO37.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO38
func (o *IO_MUX_Type) SetGPIO38_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO38.Reg, volatile.LoadUint32(&o.GPIO38.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO38_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO38.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO38_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO38.Reg, volatile.LoadUint32(&o.GPIO38.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO38_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO38.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO38_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO38.Reg, volatile.LoadUint32(&o.GPIO38.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO38_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO38.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO38_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO38.Reg, volatile.LoadUint32(&o.GPIO38.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO38_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO38.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO38_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO38.Reg, volatile.LoadUint32(&o.GPIO38.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO38_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO38.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO38_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO38.Reg, volatile.LoadUint32(&o.GPIO38.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO38_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO38.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO38_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO38.Reg, volatile.LoadUint32(&o.GPIO38.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO38_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO38.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO38_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO38.Reg, volatile.LoadUint32(&o.GPIO38.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO38_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO38.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO38_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO38.Reg, volatile.LoadUint32(&o.GPIO38.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO38_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO38.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO38_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO38.Reg, volatile.LoadUint32(&o.GPIO38.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO38_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO38.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO38_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO38.Reg, volatile.LoadUint32(&o.GPIO38.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO38_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO38.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO39
func (o *IO_MUX_Type) SetGPIO39_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO39.Reg, volatile.LoadUint32(&o.GPIO39.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO39_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO39.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO39_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO39.Reg, volatile.LoadUint32(&o.GPIO39.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO39_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO39.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO39_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO39.Reg, volatile.LoadUint32(&o.GPIO39.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO39_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO39.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO39_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO39.Reg, volatile.LoadUint32(&o.GPIO39.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO39_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO39.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO39_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO39.Reg, volatile.LoadUint32(&o.GPIO39.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO39_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO39.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO39_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO39.Reg, volatile.LoadUint32(&o.GPIO39.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO39_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO39.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO39_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO39.Reg, volatile.LoadUint32(&o.GPIO39.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO39_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO39.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO39_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO39.Reg, volatile.LoadUint32(&o.GPIO39.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO39_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO39.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO39_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO39.Reg, volatile.LoadUint32(&o.GPIO39.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO39_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO39.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO39_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO39.Reg, volatile.LoadUint32(&o.GPIO39.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO39_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO39.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO39_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO39.Reg, volatile.LoadUint32(&o.GPIO39.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO39_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO39.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO34
func (o *IO_MUX_Type) SetGPIO34_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO34.Reg, volatile.LoadUint32(&o.GPIO34.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO34_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO34.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO34_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO34.Reg, volatile.LoadUint32(&o.GPIO34.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO34_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO34.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO34_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO34.Reg, volatile.LoadUint32(&o.GPIO34.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO34_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO34.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO34_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO34.Reg, volatile.LoadUint32(&o.GPIO34.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO34_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO34.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO34_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO34.Reg, volatile.LoadUint32(&o.GPIO34.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO34_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO34.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO34_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO34.Reg, volatile.LoadUint32(&o.GPIO34.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO34_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO34.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO34_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO34.Reg, volatile.LoadUint32(&o.GPIO34.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO34_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO34.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO34_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO34.Reg, volatile.LoadUint32(&o.GPIO34.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO34_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO34.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO34_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO34.Reg, volatile.LoadUint32(&o.GPIO34.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO34_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO34.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO34_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO34.Reg, volatile.LoadUint32(&o.GPIO34.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO34_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO34.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO34_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO34.Reg, volatile.LoadUint32(&o.GPIO34.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO34_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO34.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO35
func (o *IO_MUX_Type) SetGPIO35_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO35.Reg, volatile.LoadUint32(&o.GPIO35.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO35_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO35.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO35_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO35.Reg, volatile.LoadUint32(&o.GPIO35.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO35_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO35.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO35_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO35.Reg, volatile.LoadUint32(&o.GPIO35.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO35_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO35.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO35_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO35.Reg, volatile.LoadUint32(&o.GPIO35.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO35_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO35.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO35_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO35.Reg, volatile.LoadUint32(&o.GPIO35.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO35_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO35.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO35_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO35.Reg, volatile.LoadUint32(&o.GPIO35.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO35_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO35.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO35_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO35.Reg, volatile.LoadUint32(&o.GPIO35.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO35_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO35.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO35_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO35.Reg, volatile.LoadUint32(&o.GPIO35.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO35_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO35.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO35_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO35.Reg, volatile.LoadUint32(&o.GPIO35.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO35_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO35.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO35_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO35.Reg, volatile.LoadUint32(&o.GPIO35.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO35_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO35.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO35_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO35.Reg, volatile.LoadUint32(&o.GPIO35.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO35_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO35.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO32
func (o *IO_MUX_Type) SetGPIO32_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO32.Reg, volatile.LoadUint32(&o.GPIO32.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO32_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO32.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO32_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO32.Reg, volatile.LoadUint32(&o.GPIO32.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO32_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO32.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO32_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO32.Reg, volatile.LoadUint32(&o.GPIO32.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO32_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO32.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO32_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO32.Reg, volatile.LoadUint32(&o.GPIO32.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO32_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO32.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO32_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO32.Reg, volatile.LoadUint32(&o.GPIO32.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO32_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO32.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO32_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO32.Reg, volatile.LoadUint32(&o.GPIO32.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO32_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO32.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO32_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO32.Reg, volatile.LoadUint32(&o.GPIO32.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO32_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO32.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO32_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO32.Reg, volatile.LoadUint32(&o.GPIO32.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO32_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO32.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO32_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO32.Reg, volatile.LoadUint32(&o.GPIO32.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO32_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO32.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO32_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO32.Reg, volatile.LoadUint32(&o.GPIO32.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO32_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO32.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO32_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO32.Reg, volatile.LoadUint32(&o.GPIO32.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO32_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO32.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO33
func (o *IO_MUX_Type) SetGPIO33_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO33.Reg, volatile.LoadUint32(&o.GPIO33.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO33_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO33.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO33_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO33.Reg, volatile.LoadUint32(&o.GPIO33.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO33_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO33.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO33_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO33.Reg, volatile.LoadUint32(&o.GPIO33.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO33_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO33.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO33_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO33.Reg, volatile.LoadUint32(&o.GPIO33.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO33_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO33.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO33_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO33.Reg, volatile.LoadUint32(&o.GPIO33.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO33_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO33.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO33_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO33.Reg, volatile.LoadUint32(&o.GPIO33.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO33_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO33.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO33_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO33.Reg, volatile.LoadUint32(&o.GPIO33.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO33_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO33.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO33_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO33.Reg, volatile.LoadUint32(&o.GPIO33.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO33_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO33.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO33_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO33.Reg, volatile.LoadUint32(&o.GPIO33.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO33_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO33.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO33_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO33.Reg, volatile.LoadUint32(&o.GPIO33.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO33_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO33.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO33_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO33.Reg, volatile.LoadUint32(&o.GPIO33.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO33_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO33.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO25
func (o *IO_MUX_Type) SetGPIO25_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO25.Reg, volatile.LoadUint32(&o.GPIO25.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO25_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO25.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO25_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO25.Reg, volatile.LoadUint32(&o.GPIO25.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO25_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO25.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO25_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO25.Reg, volatile.LoadUint32(&o.GPIO25.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO25_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO25.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO25_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO25.Reg, volatile.LoadUint32(&o.GPIO25.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO25_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO25.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO25_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO25.Reg, volatile.LoadUint32(&o.GPIO25.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO25_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO25.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO25_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO25.Reg, volatile.LoadUint32(&o.GPIO25.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO25_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO25.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO25_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO25.Reg, volatile.LoadUint32(&o.GPIO25.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO25_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO25.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO25_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO25.Reg, volatile.LoadUint32(&o.GPIO25.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO25_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO25.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO25_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO25.Reg, volatile.LoadUint32(&o.GPIO25.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO25_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO25.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO25_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO25.Reg, volatile.LoadUint32(&o.GPIO25.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO25_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO25.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO25_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO25.Reg, volatile.LoadUint32(&o.GPIO25.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO25_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO25.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO26
func (o *IO_MUX_Type) SetGPIO26_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO26.Reg, volatile.LoadUint32(&o.GPIO26.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO26_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO26.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO26_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO26.Reg, volatile.LoadUint32(&o.GPIO26.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO26_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO26.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO26_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO26.Reg, volatile.LoadUint32(&o.GPIO26.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO26_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO26.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO26_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO26.Reg, volatile.LoadUint32(&o.GPIO26.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO26_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO26.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO26_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO26.Reg, volatile.LoadUint32(&o.GPIO26.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO26_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO26.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO26_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO26.Reg, volatile.LoadUint32(&o.GPIO26.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO26_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO26.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO26_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO26.Reg, volatile.LoadUint32(&o.GPIO26.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO26_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO26.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO26_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO26.Reg, volatile.LoadUint32(&o.GPIO26.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO26_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO26.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO26_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO26.Reg, volatile.LoadUint32(&o.GPIO26.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO26_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO26.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO26_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO26.Reg, volatile.LoadUint32(&o.GPIO26.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO26_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO26.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO26_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO26.Reg, volatile.LoadUint32(&o.GPIO26.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO26_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO26.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO27
func (o *IO_MUX_Type) SetGPIO27_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO27.Reg, volatile.LoadUint32(&o.GPIO27.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO27_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO27.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO27_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO27.Reg, volatile.LoadUint32(&o.GPIO27.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO27_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO27.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO27_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO27.Reg, volatile.LoadUint32(&o.GPIO27.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO27_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO27.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO27_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO27.Reg, volatile.LoadUint32(&o.GPIO27.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO27_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO27.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO27_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO27.Reg, volatile.LoadUint32(&o.GPIO27.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO27_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO27.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO27_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO27.Reg, volatile.LoadUint32(&o.GPIO27.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO27_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO27.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO27_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO27.Reg, volatile.LoadUint32(&o.GPIO27.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO27_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO27.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO27_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO27.Reg, volatile.LoadUint32(&o.GPIO27.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO27_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO27.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO27_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO27.Reg, volatile.LoadUint32(&o.GPIO27.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO27_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO27.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO27_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO27.Reg, volatile.LoadUint32(&o.GPIO27.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO27_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO27.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO27_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO27.Reg, volatile.LoadUint32(&o.GPIO27.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO27_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO27.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO14
func (o *IO_MUX_Type) SetGPIO14_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO14.Reg, volatile.LoadUint32(&o.GPIO14.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO14_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO14.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO14_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO14.Reg, volatile.LoadUint32(&o.GPIO14.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO14_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO14.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO14_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO14.Reg, volatile.LoadUint32(&o.GPIO14.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO14_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO14.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO14_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO14.Reg, volatile.LoadUint32(&o.GPIO14.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO14_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO14.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO14_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO14.Reg, volatile.LoadUint32(&o.GPIO14.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO14_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO14.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO14_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO14.Reg, volatile.LoadUint32(&o.GPIO14.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO14_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO14.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO14_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO14.Reg, volatile.LoadUint32(&o.GPIO14.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO14_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO14.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO14_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO14.Reg, volatile.LoadUint32(&o.GPIO14.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO14_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO14.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO14_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO14.Reg, volatile.LoadUint32(&o.GPIO14.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO14_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO14.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO14_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO14.Reg, volatile.LoadUint32(&o.GPIO14.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO14_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO14.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO14_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO14.Reg, volatile.LoadUint32(&o.GPIO14.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO14_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO14.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO12
func (o *IO_MUX_Type) SetGPIO12_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO12.Reg, volatile.LoadUint32(&o.GPIO12.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO12_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO12.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO12_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO12.Reg, volatile.LoadUint32(&o.GPIO12.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO12_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO12.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO12_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO12.Reg, volatile.LoadUint32(&o.GPIO12.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO12_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO12.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO12_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO12.Reg, volatile.LoadUint32(&o.GPIO12.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO12_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO12.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO12_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO12.Reg, volatile.LoadUint32(&o.GPIO12.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO12_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO12.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO12_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO12.Reg, volatile.LoadUint32(&o.GPIO12.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO12_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO12.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO12_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO12.Reg, volatile.LoadUint32(&o.GPIO12.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO12_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO12.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO12_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO12.Reg, volatile.LoadUint32(&o.GPIO12.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO12_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO12.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO12_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO12.Reg, volatile.LoadUint32(&o.GPIO12.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO12_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO12.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO12_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO12.Reg, volatile.LoadUint32(&o.GPIO12.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO12_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO12.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO12_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO12.Reg, volatile.LoadUint32(&o.GPIO12.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO12_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO12.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO13
func (o *IO_MUX_Type) SetGPIO13_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO13.Reg, volatile.LoadUint32(&o.GPIO13.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO13_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO13.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO13_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO13.Reg, volatile.LoadUint32(&o.GPIO13.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO13_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO13.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO13_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO13.Reg, volatile.LoadUint32(&o.GPIO13.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO13_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO13.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO13_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO13.Reg, volatile.LoadUint32(&o.GPIO13.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO13_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO13.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO13_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO13.Reg, volatile.LoadUint32(&o.GPIO13.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO13_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO13.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO13_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO13.Reg, volatile.LoadUint32(&o.GPIO13.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO13_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO13.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO13_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO13.Reg, volatile.LoadUint32(&o.GPIO13.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO13_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO13.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO13_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO13.Reg, volatile.LoadUint32(&o.GPIO13.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO13_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO13.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO13_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO13.Reg, volatile.LoadUint32(&o.GPIO13.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO13_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO13.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO13_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO13.Reg, volatile.LoadUint32(&o.GPIO13.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO13_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO13.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO13_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO13.Reg, volatile.LoadUint32(&o.GPIO13.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO13_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO13.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO15
func (o *IO_MUX_Type) SetGPIO15_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO15.Reg, volatile.LoadUint32(&o.GPIO15.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO15_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO15.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO15_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO15.Reg, volatile.LoadUint32(&o.GPIO15.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO15_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO15.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO15_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO15.Reg, volatile.LoadUint32(&o.GPIO15.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO15_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO15.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO15_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO15.Reg, volatile.LoadUint32(&o.GPIO15.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO15_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO15.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO15_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO15.Reg, volatile.LoadUint32(&o.GPIO15.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO15_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO15.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO15_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO15.Reg, volatile.LoadUint32(&o.GPIO15.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO15_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO15.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO15_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO15.Reg, volatile.LoadUint32(&o.GPIO15.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO15_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO15.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO15_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO15.Reg, volatile.LoadUint32(&o.GPIO15.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO15_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO15.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO15_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO15.Reg, volatile.LoadUint32(&o.GPIO15.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO15_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO15.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO15_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO15.Reg, volatile.LoadUint32(&o.GPIO15.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO15_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO15.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO15_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO15.Reg, volatile.LoadUint32(&o.GPIO15.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO15_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO15.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO2
func (o *IO_MUX_Type) SetGPIO2_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO2.Reg, volatile.LoadUint32(&o.GPIO2.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO2_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO2.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO2_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO2.Reg, volatile.LoadUint32(&o.GPIO2.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO2_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO2.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO2_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO2.Reg, volatile.LoadUint32(&o.GPIO2.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO2_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO2.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO2_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO2.Reg, volatile.LoadUint32(&o.GPIO2.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO2_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO2.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO2_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO2.Reg, volatile.LoadUint32(&o.GPIO2.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO2_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO2.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO2_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO2.Reg, volatile.LoadUint32(&o.GPIO2.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO2_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO2.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO2_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO2.Reg, volatile.LoadUint32(&o.GPIO2.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO2_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO2.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO2_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO2.Reg, volatile.LoadUint32(&o.GPIO2.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO2_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO2.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO2_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO2.Reg, volatile.LoadUint32(&o.GPIO2.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO2_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO2.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO2_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO2.Reg, volatile.LoadUint32(&o.GPIO2.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO2_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO2.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO2_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO2.Reg, volatile.LoadUint32(&o.GPIO2.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO2_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO2.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO0
func (o *IO_MUX_Type) SetGPIO0_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO0.Reg, volatile.LoadUint32(&o.GPIO0.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO0_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO0.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO0_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO0.Reg, volatile.LoadUint32(&o.GPIO0.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO0_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO0.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO0_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO0.Reg, volatile.LoadUint32(&o.GPIO0.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO0_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO0.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO0_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO0.Reg, volatile.LoadUint32(&o.GPIO0.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO0_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO0.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO0_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO0.Reg, volatile.LoadUint32(&o.GPIO0.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO0_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO0.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO0_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO0.Reg, volatile.LoadUint32(&o.GPIO0.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO0_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO0.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO0_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO0.Reg, volatile.LoadUint32(&o.GPIO0.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO0_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO0.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO0_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO0.Reg, volatile.LoadUint32(&o.GPIO0.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO0_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO0.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO0_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO0.Reg, volatile.LoadUint32(&o.GPIO0.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO0_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO0.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO0_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO0.Reg, volatile.LoadUint32(&o.GPIO0.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO0_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO0.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO0_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO0.Reg, volatile.LoadUint32(&o.GPIO0.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO0_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO0.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO4
func (o *IO_MUX_Type) SetGPIO4_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO4.Reg, volatile.LoadUint32(&o.GPIO4.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO4_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO4.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO4_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO4.Reg, volatile.LoadUint32(&o.GPIO4.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO4_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO4.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO4_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO4.Reg, volatile.LoadUint32(&o.GPIO4.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO4_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO4.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO4_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO4.Reg, volatile.LoadUint32(&o.GPIO4.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO4_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO4.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO4_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO4.Reg, volatile.LoadUint32(&o.GPIO4.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO4_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO4.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO4_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO4.Reg, volatile.LoadUint32(&o.GPIO4.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO4_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO4.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO4_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO4.Reg, volatile.LoadUint32(&o.GPIO4.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO4_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO4.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO4_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO4.Reg, volatile.LoadUint32(&o.GPIO4.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO4_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO4.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO4_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO4.Reg, volatile.LoadUint32(&o.GPIO4.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO4_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO4.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO4_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO4.Reg, volatile.LoadUint32(&o.GPIO4.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO4_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO4.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO4_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO4.Reg, volatile.LoadUint32(&o.GPIO4.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO4_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO4.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO16
func (o *IO_MUX_Type) SetGPIO16_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO16.Reg, volatile.LoadUint32(&o.GPIO16.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO16_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO16.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO16_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO16.Reg, volatile.LoadUint32(&o.GPIO16.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO16_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO16.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO16_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO16.Reg, volatile.LoadUint32(&o.GPIO16.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO16_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO16.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO16_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO16.Reg, volatile.LoadUint32(&o.GPIO16.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO16_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO16.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO16_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO16.Reg, volatile.LoadUint32(&o.GPIO16.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO16_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO16.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO16_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO16.Reg, volatile.LoadUint32(&o.GPIO16.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO16_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO16.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO16_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO16.Reg, volatile.LoadUint32(&o.GPIO16.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO16_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO16.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO16_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO16.Reg, volatile.LoadUint32(&o.GPIO16.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO16_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO16.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO16_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO16.Reg, volatile.LoadUint32(&o.GPIO16.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO16_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO16.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO16_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO16.Reg, volatile.LoadUint32(&o.GPIO16.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO16_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO16.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO16_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO16.Reg, volatile.LoadUint32(&o.GPIO16.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO16_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO16.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO17
func (o *IO_MUX_Type) SetGPIO17_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO17.Reg, volatile.LoadUint32(&o.GPIO17.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO17_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO17.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO17_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO17.Reg, volatile.LoadUint32(&o.GPIO17.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO17_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO17.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO17_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO17.Reg, volatile.LoadUint32(&o.GPIO17.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO17_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO17.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO17_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO17.Reg, volatile.LoadUint32(&o.GPIO17.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO17_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO17.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO17_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO17.Reg, volatile.LoadUint32(&o.GPIO17.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO17_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO17.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO17_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO17.Reg, volatile.LoadUint32(&o.GPIO17.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO17_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO17.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO17_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO17.Reg, volatile.LoadUint32(&o.GPIO17.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO17_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO17.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO17_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO17.Reg, volatile.LoadUint32(&o.GPIO17.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO17_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO17.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO17_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO17.Reg, volatile.LoadUint32(&o.GPIO17.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO17_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO17.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO17_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO17.Reg, volatile.LoadUint32(&o.GPIO17.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO17_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO17.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO17_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO17.Reg, volatile.LoadUint32(&o.GPIO17.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO17_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO17.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO9
func (o *IO_MUX_Type) SetGPIO9_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO9.Reg, volatile.LoadUint32(&o.GPIO9.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO9_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO9.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO9_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO9.Reg, volatile.LoadUint32(&o.GPIO9.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO9_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO9.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO9_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO9.Reg, volatile.LoadUint32(&o.GPIO9.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO9_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO9.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO9_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO9.Reg, volatile.LoadUint32(&o.GPIO9.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO9_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO9.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO9_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO9.Reg, volatile.LoadUint32(&o.GPIO9.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO9_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO9.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO9_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO9.Reg, volatile.LoadUint32(&o.GPIO9.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO9_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO9.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO9_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO9.Reg, volatile.LoadUint32(&o.GPIO9.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO9_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO9.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO9_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO9.Reg, volatile.LoadUint32(&o.GPIO9.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO9_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO9.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO9_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO9.Reg, volatile.LoadUint32(&o.GPIO9.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO9_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO9.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO9_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO9.Reg, volatile.LoadUint32(&o.GPIO9.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO9_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO9.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO9_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO9.Reg, volatile.LoadUint32(&o.GPIO9.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO9_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO9.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO10
func (o *IO_MUX_Type) SetGPIO10_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO10.Reg, volatile.LoadUint32(&o.GPIO10.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO10_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO10.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO10_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO10.Reg, volatile.LoadUint32(&o.GPIO10.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO10_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO10.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO10_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO10.Reg, volatile.LoadUint32(&o.GPIO10.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO10_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO10.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO10_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO10.Reg, volatile.LoadUint32(&o.GPIO10.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO10_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO10.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO10_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO10.Reg, volatile.LoadUint32(&o.GPIO10.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO10_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO10.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO10_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO10.Reg, volatile.LoadUint32(&o.GPIO10.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO10_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO10.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO10_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO10.Reg, volatile.LoadUint32(&o.GPIO10.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO10_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO10.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO10_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO10.Reg, volatile.LoadUint32(&o.GPIO10.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO10_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO10.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO10_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO10.Reg, volatile.LoadUint32(&o.GPIO10.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO10_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO10.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO10_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO10.Reg, volatile.LoadUint32(&o.GPIO10.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO10_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO10.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO10_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO10.Reg, volatile.LoadUint32(&o.GPIO10.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO10_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO10.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO11
func (o *IO_MUX_Type) SetGPIO11_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO11.Reg, volatile.LoadUint32(&o.GPIO11.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO11_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO11.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO11_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO11.Reg, volatile.LoadUint32(&o.GPIO11.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO11_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO11.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO11_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO11.Reg, volatile.LoadUint32(&o.GPIO11.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO11_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO11.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO11_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO11.Reg, volatile.LoadUint32(&o.GPIO11.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO11_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO11.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO11_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO11.Reg, volatile.LoadUint32(&o.GPIO11.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO11_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO11.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO11_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO11.Reg, volatile.LoadUint32(&o.GPIO11.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO11_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO11.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO11_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO11.Reg, volatile.LoadUint32(&o.GPIO11.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO11_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO11.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO11_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO11.Reg, volatile.LoadUint32(&o.GPIO11.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO11_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO11.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO11_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO11.Reg, volatile.LoadUint32(&o.GPIO11.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO11_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO11.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO11_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO11.Reg, volatile.LoadUint32(&o.GPIO11.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO11_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO11.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO11_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO11.Reg, volatile.LoadUint32(&o.GPIO11.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO11_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO11.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO6
func (o *IO_MUX_Type) SetGPIO6_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO6.Reg, volatile.LoadUint32(&o.GPIO6.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO6_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO6.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO6_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO6.Reg, volatile.LoadUint32(&o.GPIO6.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO6_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO6.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO6_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO6.Reg, volatile.LoadUint32(&o.GPIO6.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO6_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO6.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO6_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO6.Reg, volatile.LoadUint32(&o.GPIO6.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO6_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO6.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO6_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO6.Reg, volatile.LoadUint32(&o.GPIO6.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO6_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO6.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO6_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO6.Reg, volatile.LoadUint32(&o.GPIO6.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO6_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO6.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO6_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO6.Reg, volatile.LoadUint32(&o.GPIO6.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO6_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO6.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO6_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO6.Reg, volatile.LoadUint32(&o.GPIO6.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO6_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO6.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO6_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO6.Reg, volatile.LoadUint32(&o.GPIO6.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO6_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO6.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO6_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO6.Reg, volatile.LoadUint32(&o.GPIO6.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO6_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO6.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO6_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO6.Reg, volatile.LoadUint32(&o.GPIO6.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO6_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO6.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO7
func (o *IO_MUX_Type) SetGPIO7_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO7.Reg, volatile.LoadUint32(&o.GPIO7.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO7_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO7.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO7_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO7.Reg, volatile.LoadUint32(&o.GPIO7.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO7_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO7.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO7_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO7.Reg, volatile.LoadUint32(&o.GPIO7.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO7_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO7.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO7_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO7.Reg, volatile.LoadUint32(&o.GPIO7.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO7_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO7.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO7_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO7.Reg, volatile.LoadUint32(&o.GPIO7.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO7_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO7.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO7_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO7.Reg, volatile.LoadUint32(&o.GPIO7.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO7_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO7.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO7_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO7.Reg, volatile.LoadUint32(&o.GPIO7.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO7_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO7.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO7_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO7.Reg, volatile.LoadUint32(&o.GPIO7.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO7_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO7.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO7_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO7.Reg, volatile.LoadUint32(&o.GPIO7.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO7_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO7.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO7_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO7.Reg, volatile.LoadUint32(&o.GPIO7.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO7_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO7.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO7_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO7.Reg, volatile.LoadUint32(&o.GPIO7.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO7_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO7.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO8
func (o *IO_MUX_Type) SetGPIO8_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO8.Reg, volatile.LoadUint32(&o.GPIO8.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO8_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO8.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO8_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO8.Reg, volatile.LoadUint32(&o.GPIO8.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO8_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO8.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO8_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO8.Reg, volatile.LoadUint32(&o.GPIO8.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO8_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO8.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO8_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO8.Reg, volatile.LoadUint32(&o.GPIO8.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO8_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO8.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO8_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO8.Reg, volatile.LoadUint32(&o.GPIO8.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO8_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO8.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO8_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO8.Reg, volatile.LoadUint32(&o.GPIO8.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO8_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO8.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO8_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO8.Reg, volatile.LoadUint32(&o.GPIO8.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO8_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO8.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO8_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO8.Reg, volatile.LoadUint32(&o.GPIO8.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO8_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO8.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO8_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO8.Reg, volatile.LoadUint32(&o.GPIO8.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO8_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO8.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO8_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO8.Reg, volatile.LoadUint32(&o.GPIO8.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO8_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO8.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO8_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO8.Reg, volatile.LoadUint32(&o.GPIO8.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO8_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO8.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO5
func (o *IO_MUX_Type) SetGPIO5_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO5.Reg, volatile.LoadUint32(&o.GPIO5.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO5_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO5.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO5_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO5.Reg, volatile.LoadUint32(&o.GPIO5.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO5_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO5.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO5_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO5.Reg, volatile.LoadUint32(&o.GPIO5.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO5_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO5.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO5_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO5.Reg, volatile.LoadUint32(&o.GPIO5.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO5_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO5.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO5_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO5.Reg, volatile.LoadUint32(&o.GPIO5.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO5_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO5.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO5_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO5.Reg, volatile.LoadUint32(&o.GPIO5.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO5_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO5.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO5_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO5.Reg, volatile.LoadUint32(&o.GPIO5.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO5_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO5.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO5_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO5.Reg, volatile.LoadUint32(&o.GPIO5.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO5_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO5.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO5_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO5.Reg, volatile.LoadUint32(&o.GPIO5.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO5_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO5.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO5_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO5.Reg, volatile.LoadUint32(&o.GPIO5.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO5_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO5.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO5_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO5.Reg, volatile.LoadUint32(&o.GPIO5.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO5_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO5.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO18
func (o *IO_MUX_Type) SetGPIO18_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO18.Reg, volatile.LoadUint32(&o.GPIO18.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO18_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO18.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO18_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO18.Reg, volatile.LoadUint32(&o.GPIO18.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO18_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO18.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO18_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO18.Reg, volatile.LoadUint32(&o.GPIO18.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO18_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO18.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO18_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO18.Reg, volatile.LoadUint32(&o.GPIO18.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO18_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO18.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO18_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO18.Reg, volatile.LoadUint32(&o.GPIO18.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO18_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO18.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO18_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO18.Reg, volatile.LoadUint32(&o.GPIO18.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO18_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO18.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO18_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO18.Reg, volatile.LoadUint32(&o.GPIO18.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO18_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO18.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO18_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO18.Reg, volatile.LoadUint32(&o.GPIO18.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO18_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO18.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO18_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO18.Reg, volatile.LoadUint32(&o.GPIO18.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO18_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO18.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO18_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO18.Reg, volatile.LoadUint32(&o.GPIO18.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO18_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO18.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO18_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO18.Reg, volatile.LoadUint32(&o.GPIO18.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO18_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO18.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO19
func (o *IO_MUX_Type) SetGPIO19_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO19.Reg, volatile.LoadUint32(&o.GPIO19.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO19_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO19.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO19_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO19.Reg, volatile.LoadUint32(&o.GPIO19.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO19_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO19.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO19_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO19.Reg, volatile.LoadUint32(&o.GPIO19.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO19_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO19.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO19_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO19.Reg, volatile.LoadUint32(&o.GPIO19.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO19_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO19.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO19_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO19.Reg, volatile.LoadUint32(&o.GPIO19.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO19_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO19.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO19_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO19.Reg, volatile.LoadUint32(&o.GPIO19.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO19_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO19.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO19_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO19.Reg, volatile.LoadUint32(&o.GPIO19.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO19_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO19.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO19_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO19.Reg, volatile.LoadUint32(&o.GPIO19.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO19_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO19.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO19_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO19.Reg, volatile.LoadUint32(&o.GPIO19.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO19_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO19.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO19_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO19.Reg, volatile.LoadUint32(&o.GPIO19.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO19_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO19.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO19_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO19.Reg, volatile.LoadUint32(&o.GPIO19.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO19_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO19.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO20
func (o *IO_MUX_Type) SetGPIO20_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO20.Reg, volatile.LoadUint32(&o.GPIO20.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO20_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO20.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO20_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO20.Reg, volatile.LoadUint32(&o.GPIO20.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO20_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO20.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO20_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO20.Reg, volatile.LoadUint32(&o.GPIO20.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO20_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO20.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO20_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO20.Reg, volatile.LoadUint32(&o.GPIO20.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO20_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO20.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO20_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO20.Reg, volatile.LoadUint32(&o.GPIO20.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO20_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO20.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO20_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO20.Reg, volatile.LoadUint32(&o.GPIO20.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO20_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO20.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO20_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO20.Reg, volatile.LoadUint32(&o.GPIO20.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO20_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO20.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO20_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO20.Reg, volatile.LoadUint32(&o.GPIO20.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO20_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO20.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO20_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO20.Reg, volatile.LoadUint32(&o.GPIO20.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO20_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO20.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO20_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO20.Reg, volatile.LoadUint32(&o.GPIO20.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO20_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO20.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO20_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO20.Reg, volatile.LoadUint32(&o.GPIO20.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO20_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO20.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO21
func (o *IO_MUX_Type) SetGPIO21_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO21.Reg, volatile.LoadUint32(&o.GPIO21.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO21_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO21.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO21_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO21.Reg, volatile.LoadUint32(&o.GPIO21.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO21_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO21.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO21_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO21.Reg, volatile.LoadUint32(&o.GPIO21.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO21_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO21.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO21_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO21.Reg, volatile.LoadUint32(&o.GPIO21.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO21_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO21.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO21_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO21.Reg, volatile.LoadUint32(&o.GPIO21.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO21_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO21.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO21_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO21.Reg, volatile.LoadUint32(&o.GPIO21.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO21_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO21.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO21_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO21.Reg, volatile.LoadUint32(&o.GPIO21.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO21_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO21.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO21_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO21.Reg, volatile.LoadUint32(&o.GPIO21.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO21_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO21.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO21_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO21.Reg, volatile.LoadUint32(&o.GPIO21.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO21_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO21.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO21_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO21.Reg, volatile.LoadUint32(&o.GPIO21.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO21_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO21.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO21_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO21.Reg, volatile.LoadUint32(&o.GPIO21.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO21_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO21.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO22
func (o *IO_MUX_Type) SetGPIO22_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO22.Reg, volatile.LoadUint32(&o.GPIO22.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO22_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO22.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO22_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO22.Reg, volatile.LoadUint32(&o.GPIO22.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO22_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO22.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO22_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO22.Reg, volatile.LoadUint32(&o.GPIO22.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO22_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO22.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO22_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO22.Reg, volatile.LoadUint32(&o.GPIO22.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO22_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO22.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO22_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO22.Reg, volatile.LoadUint32(&o.GPIO22.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO22_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO22.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO22_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO22.Reg, volatile.LoadUint32(&o.GPIO22.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO22_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO22.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO22_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO22.Reg, volatile.LoadUint32(&o.GPIO22.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO22_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO22.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO22_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO22.Reg, volatile.LoadUint32(&o.GPIO22.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO22_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO22.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO22_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO22.Reg, volatile.LoadUint32(&o.GPIO22.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO22_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO22.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO22_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO22.Reg, volatile.LoadUint32(&o.GPIO22.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO22_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO22.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO22_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO22.Reg, volatile.LoadUint32(&o.GPIO22.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO22_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO22.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO3
func (o *IO_MUX_Type) SetGPIO3_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO3.Reg, volatile.LoadUint32(&o.GPIO3.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO3_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO3.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO3_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO3.Reg, volatile.LoadUint32(&o.GPIO3.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO3_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO3.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO3_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO3.Reg, volatile.LoadUint32(&o.GPIO3.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO3_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO3.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO3_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO3.Reg, volatile.LoadUint32(&o.GPIO3.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO3_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO3.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO3_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO3.Reg, volatile.LoadUint32(&o.GPIO3.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO3_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO3.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO3_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO3.Reg, volatile.LoadUint32(&o.GPIO3.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO3_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO3.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO3_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO3.Reg, volatile.LoadUint32(&o.GPIO3.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO3_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO3.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO3_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO3.Reg, volatile.LoadUint32(&o.GPIO3.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO3_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO3.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO3_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO3.Reg, volatile.LoadUint32(&o.GPIO3.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO3_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO3.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO3_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO3.Reg, volatile.LoadUint32(&o.GPIO3.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO3_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO3.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO3_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO3.Reg, volatile.LoadUint32(&o.GPIO3.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO3_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO3.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO1
func (o *IO_MUX_Type) SetGPIO1_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO1.Reg, volatile.LoadUint32(&o.GPIO1.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO1_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO1.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO1_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO1.Reg, volatile.LoadUint32(&o.GPIO1.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO1_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO1.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO1_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO1.Reg, volatile.LoadUint32(&o.GPIO1.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO1_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO1.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO1_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO1.Reg, volatile.LoadUint32(&o.GPIO1.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO1_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO1.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO1_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO1.Reg, volatile.LoadUint32(&o.GPIO1.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO1_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO1.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO1_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO1.Reg, volatile.LoadUint32(&o.GPIO1.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO1_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO1.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO1_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO1.Reg, volatile.LoadUint32(&o.GPIO1.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO1_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO1.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO1_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO1.Reg, volatile.LoadUint32(&o.GPIO1.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO1_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO1.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO1_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO1.Reg, volatile.LoadUint32(&o.GPIO1.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO1_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO1.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO1_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO1.Reg, volatile.LoadUint32(&o.GPIO1.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO1_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO1.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO1_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO1.Reg, volatile.LoadUint32(&o.GPIO1.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO1_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO1.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO23
func (o *IO_MUX_Type) SetGPIO23_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO23.Reg, volatile.LoadUint32(&o.GPIO23.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO23_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO23.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO23_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO23.Reg, volatile.LoadUint32(&o.GPIO23.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO23_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO23.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO23_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO23.Reg, volatile.LoadUint32(&o.GPIO23.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO23_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO23.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO23_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO23.Reg, volatile.LoadUint32(&o.GPIO23.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO23_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO23.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO23_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO23.Reg, volatile.LoadUint32(&o.GPIO23.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO23_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO23.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO23_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO23.Reg, volatile.LoadUint32(&o.GPIO23.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO23_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO23.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO23_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO23.Reg, volatile.LoadUint32(&o.GPIO23.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO23_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO23.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO23_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO23.Reg, volatile.LoadUint32(&o.GPIO23.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO23_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO23.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO23_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO23.Reg, volatile.LoadUint32(&o.GPIO23.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO23_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO23.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO23_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO23.Reg, volatile.LoadUint32(&o.GPIO23.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO23_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO23.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO23_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO23.Reg, volatile.LoadUint32(&o.GPIO23.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO23_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO23.Reg) & 0x7000) >> 12
}

// IO_MUX.GPIO24
func (o *IO_MUX_Type) SetGPIO24_MCU_OE(value uint32) {
	volatile.StoreUint32(&o.GPIO24.Reg, volatile.LoadUint32(&o.GPIO24.Reg)&^(0x1)|value)
}
func (o *IO_MUX_Type) GetGPIO24_MCU_OE() uint32 {
	return volatile.LoadUint32(&o.GPIO24.Reg) & 0x1
}
func (o *IO_MUX_Type) SetGPIO24_SLP_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO24.Reg, volatile.LoadUint32(&o.GPIO24.Reg)&^(0x2)|value<<1)
}
func (o *IO_MUX_Type) GetGPIO24_SLP_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO24.Reg) & 0x2) >> 1
}
func (o *IO_MUX_Type) SetGPIO24_MCU_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO24.Reg, volatile.LoadUint32(&o.GPIO24.Reg)&^(0x4)|value<<2)
}
func (o *IO_MUX_Type) GetGPIO24_MCU_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO24.Reg) & 0x4) >> 2
}
func (o *IO_MUX_Type) SetGPIO24_MCU_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO24.Reg, volatile.LoadUint32(&o.GPIO24.Reg)&^(0x8)|value<<3)
}
func (o *IO_MUX_Type) GetGPIO24_MCU_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO24.Reg) & 0x8) >> 3
}
func (o *IO_MUX_Type) SetGPIO24_MCU_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO24.Reg, volatile.LoadUint32(&o.GPIO24.Reg)&^(0x10)|value<<4)
}
func (o *IO_MUX_Type) GetGPIO24_MCU_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO24.Reg) & 0x10) >> 4
}
func (o *IO_MUX_Type) SetGPIO24_MCU_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO24.Reg, volatile.LoadUint32(&o.GPIO24.Reg)&^(0x60)|value<<5)
}
func (o *IO_MUX_Type) GetGPIO24_MCU_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO24.Reg) & 0x60) >> 5
}
func (o *IO_MUX_Type) SetGPIO24_FUN_WPD(value uint32) {
	volatile.StoreUint32(&o.GPIO24.Reg, volatile.LoadUint32(&o.GPIO24.Reg)&^(0x80)|value<<7)
}
func (o *IO_MUX_Type) GetGPIO24_FUN_WPD() uint32 {
	return (volatile.LoadUint32(&o.GPIO24.Reg) & 0x80) >> 7
}
func (o *IO_MUX_Type) SetGPIO24_FUN_WPU(value uint32) {
	volatile.StoreUint32(&o.GPIO24.Reg, volatile.LoadUint32(&o.GPIO24.Reg)&^(0x100)|value<<8)
}
func (o *IO_MUX_Type) GetGPIO24_FUN_WPU() uint32 {
	return (volatile.LoadUint32(&o.GPIO24.Reg) & 0x100) >> 8
}
func (o *IO_MUX_Type) SetGPIO24_FUN_IE(value uint32) {
	volatile.StoreUint32(&o.GPIO24.Reg, volatile.LoadUint32(&o.GPIO24.Reg)&^(0x200)|value<<9)
}
func (o *IO_MUX_Type) GetGPIO24_FUN_IE() uint32 {
	return (volatile.LoadUint32(&o.GPIO24.Reg) & 0x200) >> 9
}
func (o *IO_MUX_Type) SetGPIO24_FUN_DRV(value uint32) {
	volatile.StoreUint32(&o.GPIO24.Reg, volatile.LoadUint32(&o.GPIO24.Reg)&^(0xc00)|value<<10)
}
func (o *IO_MUX_Type) GetGPIO24_FUN_DRV() uint32 {
	return (volatile.LoadUint32(&o.GPIO24.Reg) & 0xc00) >> 10
}
func (o *IO_MUX_Type) SetGPIO24_MCU_SEL(value uint32) {
	volatile.StoreUint32(&o.GPIO24.Reg, volatile.LoadUint32(&o.GPIO24.Reg)&^(0x7000)|value<<12)
}
func (o *IO_MUX_Type) GetGPIO24_MCU_SEL() uint32 {
	return (volatile.LoadUint32(&o.GPIO24.Reg) & 0x7000) >> 12
}
