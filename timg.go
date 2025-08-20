package main

import "llgo-test/volatile"

// Constants for TIMG0: Timer Group 0
const (
	// T0CONFIG
	// Position of ALARM_EN field.
	TIMG_T0CONFIG_ALARM_EN_Pos = 0xa
	// Bit mask of ALARM_EN field.
	TIMG_T0CONFIG_ALARM_EN_Msk = 0x400
	// Bit ALARM_EN.
	TIMG_T0CONFIG_ALARM_EN = 0x400
	// Position of LEVEL_INT_EN field.
	TIMG_T0CONFIG_LEVEL_INT_EN_Pos = 0xb
	// Bit mask of LEVEL_INT_EN field.
	TIMG_T0CONFIG_LEVEL_INT_EN_Msk = 0x800
	// Bit LEVEL_INT_EN.
	TIMG_T0CONFIG_LEVEL_INT_EN = 0x800
	// Position of EDGE_INT_EN field.
	TIMG_T0CONFIG_EDGE_INT_EN_Pos = 0xc
	// Bit mask of EDGE_INT_EN field.
	TIMG_T0CONFIG_EDGE_INT_EN_Msk = 0x1000
	// Bit EDGE_INT_EN.
	TIMG_T0CONFIG_EDGE_INT_EN = 0x1000
	// Position of DIVIDER field.
	TIMG_T0CONFIG_DIVIDER_Pos = 0xd
	// Bit mask of DIVIDER field.
	TIMG_T0CONFIG_DIVIDER_Msk = 0x1fffe000
	// Position of AUTORELOAD field.
	TIMG_T0CONFIG_AUTORELOAD_Pos = 0x1d
	// Bit mask of AUTORELOAD field.
	TIMG_T0CONFIG_AUTORELOAD_Msk = 0x20000000
	// Bit AUTORELOAD.
	TIMG_T0CONFIG_AUTORELOAD = 0x20000000
	// Position of INCREASE field.
	TIMG_T0CONFIG_INCREASE_Pos = 0x1e
	// Bit mask of INCREASE field.
	TIMG_T0CONFIG_INCREASE_Msk = 0x40000000
	// Bit INCREASE.
	TIMG_T0CONFIG_INCREASE = 0x40000000
	// Position of EN field.
	TIMG_T0CONFIG_EN_Pos = 0x1f
	// Bit mask of EN field.
	TIMG_T0CONFIG_EN_Msk = 0x80000000
	// Bit EN.
	TIMG_T0CONFIG_EN = 0x80000000

	// T0LO
	// Position of LO field.
	TIMG_T0LO_LO_Pos = 0x0
	// Bit mask of LO field.
	TIMG_T0LO_LO_Msk = 0xffffffff

	// T0HI
	// Position of HI field.
	TIMG_T0HI_HI_Pos = 0x0
	// Bit mask of HI field.
	TIMG_T0HI_HI_Msk = 0xffffffff

	// T0UPDATE
	// Position of UPDATE field.
	TIMG_T0UPDATE_UPDATE_Pos = 0x0
	// Bit mask of UPDATE field.
	TIMG_T0UPDATE_UPDATE_Msk = 0xffffffff

	// T0ALARMLO
	// Position of ALARM_LO field.
	TIMG_T0ALARMLO_ALARM_LO_Pos = 0x0
	// Bit mask of ALARM_LO field.
	TIMG_T0ALARMLO_ALARM_LO_Msk = 0xffffffff

	// T0ALARMHI
	// Position of ALARM_HI field.
	TIMG_T0ALARMHI_ALARM_HI_Pos = 0x0
	// Bit mask of ALARM_HI field.
	TIMG_T0ALARMHI_ALARM_HI_Msk = 0xffffffff

	// T0LOADLO
	// Position of LOAD_LO field.
	TIMG_T0LOADLO_LOAD_LO_Pos = 0x0
	// Bit mask of LOAD_LO field.
	TIMG_T0LOADLO_LOAD_LO_Msk = 0xffffffff

	// T0LOADHI
	// Position of LOAD_HI field.
	TIMG_T0LOADHI_LOAD_HI_Pos = 0x0
	// Bit mask of LOAD_HI field.
	TIMG_T0LOADHI_LOAD_HI_Msk = 0xffffffff

	// T0LOAD
	// Position of LOAD field.
	TIMG_T0LOAD_LOAD_Pos = 0x0
	// Bit mask of LOAD field.
	TIMG_T0LOAD_LOAD_Msk = 0xffffffff

	// T1CONFIG
	// Position of ALARM_EN field.
	TIMG_T1CONFIG_ALARM_EN_Pos = 0xa
	// Bit mask of ALARM_EN field.
	TIMG_T1CONFIG_ALARM_EN_Msk = 0x400
	// Bit ALARM_EN.
	TIMG_T1CONFIG_ALARM_EN = 0x400
	// Position of LEVEL_INT_EN field.
	TIMG_T1CONFIG_LEVEL_INT_EN_Pos = 0xb
	// Bit mask of LEVEL_INT_EN field.
	TIMG_T1CONFIG_LEVEL_INT_EN_Msk = 0x800
	// Bit LEVEL_INT_EN.
	TIMG_T1CONFIG_LEVEL_INT_EN = 0x800
	// Position of EDGE_INT_EN field.
	TIMG_T1CONFIG_EDGE_INT_EN_Pos = 0xc
	// Bit mask of EDGE_INT_EN field.
	TIMG_T1CONFIG_EDGE_INT_EN_Msk = 0x1000
	// Bit EDGE_INT_EN.
	TIMG_T1CONFIG_EDGE_INT_EN = 0x1000
	// Position of DIVIDER field.
	TIMG_T1CONFIG_DIVIDER_Pos = 0xd
	// Bit mask of DIVIDER field.
	TIMG_T1CONFIG_DIVIDER_Msk = 0x1fffe000
	// Position of AUTORELOAD field.
	TIMG_T1CONFIG_AUTORELOAD_Pos = 0x1d
	// Bit mask of AUTORELOAD field.
	TIMG_T1CONFIG_AUTORELOAD_Msk = 0x20000000
	// Bit AUTORELOAD.
	TIMG_T1CONFIG_AUTORELOAD = 0x20000000
	// Position of INCREASE field.
	TIMG_T1CONFIG_INCREASE_Pos = 0x1e
	// Bit mask of INCREASE field.
	TIMG_T1CONFIG_INCREASE_Msk = 0x40000000
	// Bit INCREASE.
	TIMG_T1CONFIG_INCREASE = 0x40000000
	// Position of EN field.
	TIMG_T1CONFIG_EN_Pos = 0x1f
	// Bit mask of EN field.
	TIMG_T1CONFIG_EN_Msk = 0x80000000
	// Bit EN.
	TIMG_T1CONFIG_EN = 0x80000000

	// T1LO
	// Position of LO field.
	TIMG_T1LO_LO_Pos = 0x0
	// Bit mask of LO field.
	TIMG_T1LO_LO_Msk = 0xffffffff

	// T1HI
	// Position of HI field.
	TIMG_T1HI_HI_Pos = 0x0
	// Bit mask of HI field.
	TIMG_T1HI_HI_Msk = 0xffffffff

	// T1UPDATE
	// Position of UPDATE field.
	TIMG_T1UPDATE_UPDATE_Pos = 0x0
	// Bit mask of UPDATE field.
	TIMG_T1UPDATE_UPDATE_Msk = 0xffffffff

	// T1ALARMLO
	// Position of ALARM_LO field.
	TIMG_T1ALARMLO_ALARM_LO_Pos = 0x0
	// Bit mask of ALARM_LO field.
	TIMG_T1ALARMLO_ALARM_LO_Msk = 0xffffffff

	// T1ALARMHI
	// Position of ALARM_HI field.
	TIMG_T1ALARMHI_ALARM_HI_Pos = 0x0
	// Bit mask of ALARM_HI field.
	TIMG_T1ALARMHI_ALARM_HI_Msk = 0xffffffff

	// T1LOADLO
	// Position of LOAD_LO field.
	TIMG_T1LOADLO_LOAD_LO_Pos = 0x0
	// Bit mask of LOAD_LO field.
	TIMG_T1LOADLO_LOAD_LO_Msk = 0xffffffff

	// T1LOADHI
	// Position of LOAD_HI field.
	TIMG_T1LOADHI_LOAD_HI_Pos = 0x0
	// Bit mask of LOAD_HI field.
	TIMG_T1LOADHI_LOAD_HI_Msk = 0xffffffff

	// T1LOAD
	// Position of LOAD field.
	TIMG_T1LOAD_LOAD_Pos = 0x0
	// Bit mask of LOAD field.
	TIMG_T1LOAD_LOAD_Msk = 0xffffffff

	// WDTCONFIG0
	// Position of WDT_FLASHBOOT_MOD_EN field.
	TIMG_WDTCONFIG0_WDT_FLASHBOOT_MOD_EN_Pos = 0xe
	// Bit mask of WDT_FLASHBOOT_MOD_EN field.
	TIMG_WDTCONFIG0_WDT_FLASHBOOT_MOD_EN_Msk = 0x4000
	// Bit WDT_FLASHBOOT_MOD_EN.
	TIMG_WDTCONFIG0_WDT_FLASHBOOT_MOD_EN = 0x4000
	// Position of WDT_SYS_RESET_LENGTH field.
	TIMG_WDTCONFIG0_WDT_SYS_RESET_LENGTH_Pos = 0xf
	// Bit mask of WDT_SYS_RESET_LENGTH field.
	TIMG_WDTCONFIG0_WDT_SYS_RESET_LENGTH_Msk = 0x38000
	// 100ns
	TIMG_WDTCONFIG0_WDT_SYS_RESET_LENGTH_NS100 = 0x0
	// 200ns
	TIMG_WDTCONFIG0_WDT_SYS_RESET_LENGTH_NS200 = 0x1
	// 300ns
	TIMG_WDTCONFIG0_WDT_SYS_RESET_LENGTH_NS300 = 0x2
	// 400ns
	TIMG_WDTCONFIG0_WDT_SYS_RESET_LENGTH_NS400 = 0x3
	// 500ns
	TIMG_WDTCONFIG0_WDT_SYS_RESET_LENGTH_NS500 = 0x4
	// 800ns
	TIMG_WDTCONFIG0_WDT_SYS_RESET_LENGTH_NS800 = 0x5
	// 1.6us
	TIMG_WDTCONFIG0_WDT_SYS_RESET_LENGTH_NS1600 = 0x6
	// 3.2us
	TIMG_WDTCONFIG0_WDT_SYS_RESET_LENGTH_NS3200 = 0x7
	// Position of WDT_CPU_RESET_LENGTH field.
	TIMG_WDTCONFIG0_WDT_CPU_RESET_LENGTH_Pos = 0x12
	// Bit mask of WDT_CPU_RESET_LENGTH field.
	TIMG_WDTCONFIG0_WDT_CPU_RESET_LENGTH_Msk = 0x1c0000
	// 100ns
	TIMG_WDTCONFIG0_WDT_CPU_RESET_LENGTH_NS100 = 0x0
	// 200ns
	TIMG_WDTCONFIG0_WDT_CPU_RESET_LENGTH_NS200 = 0x1
	// 300ns
	TIMG_WDTCONFIG0_WDT_CPU_RESET_LENGTH_NS300 = 0x2
	// 400ns
	TIMG_WDTCONFIG0_WDT_CPU_RESET_LENGTH_NS400 = 0x3
	// 500ns
	TIMG_WDTCONFIG0_WDT_CPU_RESET_LENGTH_NS500 = 0x4
	// 800ns
	TIMG_WDTCONFIG0_WDT_CPU_RESET_LENGTH_NS800 = 0x5
	// 1.6us
	TIMG_WDTCONFIG0_WDT_CPU_RESET_LENGTH_NS1600 = 0x6
	// 3.2us
	TIMG_WDTCONFIG0_WDT_CPU_RESET_LENGTH_NS3200 = 0x7
	// Position of WDT_LEVEL_INT_EN field.
	TIMG_WDTCONFIG0_WDT_LEVEL_INT_EN_Pos = 0x15
	// Bit mask of WDT_LEVEL_INT_EN field.
	TIMG_WDTCONFIG0_WDT_LEVEL_INT_EN_Msk = 0x200000
	// Bit WDT_LEVEL_INT_EN.
	TIMG_WDTCONFIG0_WDT_LEVEL_INT_EN = 0x200000
	// Position of WDT_EDGE_INT_EN field.
	TIMG_WDTCONFIG0_WDT_EDGE_INT_EN_Pos = 0x16
	// Bit mask of WDT_EDGE_INT_EN field.
	TIMG_WDTCONFIG0_WDT_EDGE_INT_EN_Msk = 0x400000
	// Bit WDT_EDGE_INT_EN.
	TIMG_WDTCONFIG0_WDT_EDGE_INT_EN = 0x400000
	// Position of WDT_STG3 field.
	TIMG_WDTCONFIG0_WDT_STG3_Pos = 0x17
	// Bit mask of WDT_STG3 field.
	TIMG_WDTCONFIG0_WDT_STG3_Msk = 0x1800000
	// Off
	TIMG_WDTCONFIG0_WDT_STG3_OFF = 0x0
	// Interrupt
	TIMG_WDTCONFIG0_WDT_STG3_INTERRUPT = 0x1
	// Reset CPU
	TIMG_WDTCONFIG0_WDT_STG3_RESET = 0x2
	// Reset system
	TIMG_WDTCONFIG0_WDT_STG3_RESET_SYS = 0x3
	// Position of WDT_STG2 field.
	TIMG_WDTCONFIG0_WDT_STG2_Pos = 0x19
	// Bit mask of WDT_STG2 field.
	TIMG_WDTCONFIG0_WDT_STG2_Msk = 0x6000000
	// Off
	TIMG_WDTCONFIG0_WDT_STG2_OFF = 0x0
	// Interrupt
	TIMG_WDTCONFIG0_WDT_STG2_INTERRUPT = 0x1
	// Reset CPU
	TIMG_WDTCONFIG0_WDT_STG2_RESET = 0x2
	// Reset system
	TIMG_WDTCONFIG0_WDT_STG2_RESET_SYS = 0x3
	// Position of WDT_STG1 field.
	TIMG_WDTCONFIG0_WDT_STG1_Pos = 0x1b
	// Bit mask of WDT_STG1 field.
	TIMG_WDTCONFIG0_WDT_STG1_Msk = 0x18000000
	// Off
	TIMG_WDTCONFIG0_WDT_STG1_OFF = 0x0
	// Interrupt
	TIMG_WDTCONFIG0_WDT_STG1_INTERRUPT = 0x1
	// Reset CPU
	TIMG_WDTCONFIG0_WDT_STG1_RESET = 0x2
	// Reset system
	TIMG_WDTCONFIG0_WDT_STG1_RESET_SYS = 0x3
	// Position of WDT_STG0 field.
	TIMG_WDTCONFIG0_WDT_STG0_Pos = 0x1d
	// Bit mask of WDT_STG0 field.
	TIMG_WDTCONFIG0_WDT_STG0_Msk = 0x60000000
	// Off
	TIMG_WDTCONFIG0_WDT_STG0_OFF = 0x0
	// Interrupt
	TIMG_WDTCONFIG0_WDT_STG0_INTERRUPT = 0x1
	// Reset CPU
	TIMG_WDTCONFIG0_WDT_STG0_RESET = 0x2
	// Reset system
	TIMG_WDTCONFIG0_WDT_STG0_RESET_SYS = 0x3
	// Position of WDT_EN field.
	TIMG_WDTCONFIG0_WDT_EN_Pos = 0x1f
	// Bit mask of WDT_EN field.
	TIMG_WDTCONFIG0_WDT_EN_Msk = 0x80000000
	// Bit WDT_EN.
	TIMG_WDTCONFIG0_WDT_EN = 0x80000000

	// WDTCONFIG1
	// Position of WDT_CLK_PRESCALE field.
	TIMG_WDTCONFIG1_WDT_CLK_PRESCALE_Pos = 0x10
	// Bit mask of WDT_CLK_PRESCALE field.
	TIMG_WDTCONFIG1_WDT_CLK_PRESCALE_Msk = 0xffff0000

	// WDTCONFIG2
	// Position of WDT_STG0_HOLD field.
	TIMG_WDTCONFIG2_WDT_STG0_HOLD_Pos = 0x0
	// Bit mask of WDT_STG0_HOLD field.
	TIMG_WDTCONFIG2_WDT_STG0_HOLD_Msk = 0xffffffff

	// WDTCONFIG3
	// Position of WDT_STG1_HOLD field.
	TIMG_WDTCONFIG3_WDT_STG1_HOLD_Pos = 0x0
	// Bit mask of WDT_STG1_HOLD field.
	TIMG_WDTCONFIG3_WDT_STG1_HOLD_Msk = 0xffffffff

	// WDTCONFIG4
	// Position of WDT_STG2_HOLD field.
	TIMG_WDTCONFIG4_WDT_STG2_HOLD_Pos = 0x0
	// Bit mask of WDT_STG2_HOLD field.
	TIMG_WDTCONFIG4_WDT_STG2_HOLD_Msk = 0xffffffff

	// WDTCONFIG5
	// Position of WDT_STG3_HOLD field.
	TIMG_WDTCONFIG5_WDT_STG3_HOLD_Pos = 0x0
	// Bit mask of WDT_STG3_HOLD field.
	TIMG_WDTCONFIG5_WDT_STG3_HOLD_Msk = 0xffffffff

	// WDTFEED
	// Position of WDT_FEED field.
	TIMG_WDTFEED_WDT_FEED_Pos = 0x0
	// Bit mask of WDT_FEED field.
	TIMG_WDTFEED_WDT_FEED_Msk = 0xffffffff

	// WDTWPROTECT
	// Position of WDT_WKEY field.
	TIMG_WDTWPROTECT_WDT_WKEY_Pos = 0x0
	// Bit mask of WDT_WKEY field.
	TIMG_WDTWPROTECT_WDT_WKEY_Msk = 0xffffffff

	// RTCCALICFG
	// Position of RTC_CALI_START_CYCLING field.
	TIMG_RTCCALICFG_RTC_CALI_START_CYCLING_Pos = 0xc
	// Bit mask of RTC_CALI_START_CYCLING field.
	TIMG_RTCCALICFG_RTC_CALI_START_CYCLING_Msk = 0x1000
	// Bit RTC_CALI_START_CYCLING.
	TIMG_RTCCALICFG_RTC_CALI_START_CYCLING = 0x1000
	// Position of RTC_CALI_CLK_SEL field.
	TIMG_RTCCALICFG_RTC_CALI_CLK_SEL_Pos = 0xd
	// Bit mask of RTC_CALI_CLK_SEL field.
	TIMG_RTCCALICFG_RTC_CALI_CLK_SEL_Msk = 0x6000
	// Position of RTC_CALI_RDY field.
	TIMG_RTCCALICFG_RTC_CALI_RDY_Pos = 0xf
	// Bit mask of RTC_CALI_RDY field.
	TIMG_RTCCALICFG_RTC_CALI_RDY_Msk = 0x8000
	// Bit RTC_CALI_RDY.
	TIMG_RTCCALICFG_RTC_CALI_RDY = 0x8000
	// Position of RTC_CALI_MAX field.
	TIMG_RTCCALICFG_RTC_CALI_MAX_Pos = 0x10
	// Bit mask of RTC_CALI_MAX field.
	TIMG_RTCCALICFG_RTC_CALI_MAX_Msk = 0x7fff0000
	// Position of RTC_CALI_START field.
	TIMG_RTCCALICFG_RTC_CALI_START_Pos = 0x1f
	// Bit mask of RTC_CALI_START field.
	TIMG_RTCCALICFG_RTC_CALI_START_Msk = 0x80000000
	// Bit RTC_CALI_START.
	TIMG_RTCCALICFG_RTC_CALI_START = 0x80000000

	// RTCCALICFG1
	// Position of RTC_CALI_VALUE field.
	TIMG_RTCCALICFG1_RTC_CALI_VALUE_Pos = 0x7
	// Bit mask of RTC_CALI_VALUE field.
	TIMG_RTCCALICFG1_RTC_CALI_VALUE_Msk = 0xffffff80

	// LACTCONFIG
	// Position of LACT_RTC_ONLY field.
	TIMG_LACTCONFIG_LACT_RTC_ONLY_Pos = 0x7
	// Bit mask of LACT_RTC_ONLY field.
	TIMG_LACTCONFIG_LACT_RTC_ONLY_Msk = 0x80
	// Bit LACT_RTC_ONLY.
	TIMG_LACTCONFIG_LACT_RTC_ONLY = 0x80
	// Position of LACT_CPST_EN field.
	TIMG_LACTCONFIG_LACT_CPST_EN_Pos = 0x8
	// Bit mask of LACT_CPST_EN field.
	TIMG_LACTCONFIG_LACT_CPST_EN_Msk = 0x100
	// Bit LACT_CPST_EN.
	TIMG_LACTCONFIG_LACT_CPST_EN = 0x100
	// Position of LACT_LAC_EN field.
	TIMG_LACTCONFIG_LACT_LAC_EN_Pos = 0x9
	// Bit mask of LACT_LAC_EN field.
	TIMG_LACTCONFIG_LACT_LAC_EN_Msk = 0x200
	// Bit LACT_LAC_EN.
	TIMG_LACTCONFIG_LACT_LAC_EN = 0x200
	// Position of LACT_ALARM_EN field.
	TIMG_LACTCONFIG_LACT_ALARM_EN_Pos = 0xa
	// Bit mask of LACT_ALARM_EN field.
	TIMG_LACTCONFIG_LACT_ALARM_EN_Msk = 0x400
	// Bit LACT_ALARM_EN.
	TIMG_LACTCONFIG_LACT_ALARM_EN = 0x400
	// Position of LACT_LEVEL_INT_EN field.
	TIMG_LACTCONFIG_LACT_LEVEL_INT_EN_Pos = 0xb
	// Bit mask of LACT_LEVEL_INT_EN field.
	TIMG_LACTCONFIG_LACT_LEVEL_INT_EN_Msk = 0x800
	// Bit LACT_LEVEL_INT_EN.
	TIMG_LACTCONFIG_LACT_LEVEL_INT_EN = 0x800
	// Position of LACT_EDGE_INT_EN field.
	TIMG_LACTCONFIG_LACT_EDGE_INT_EN_Pos = 0xc
	// Bit mask of LACT_EDGE_INT_EN field.
	TIMG_LACTCONFIG_LACT_EDGE_INT_EN_Msk = 0x1000
	// Bit LACT_EDGE_INT_EN.
	TIMG_LACTCONFIG_LACT_EDGE_INT_EN = 0x1000
	// Position of LACT_DIVIDER field.
	TIMG_LACTCONFIG_LACT_DIVIDER_Pos = 0xd
	// Bit mask of LACT_DIVIDER field.
	TIMG_LACTCONFIG_LACT_DIVIDER_Msk = 0x1fffe000
	// Position of LACT_AUTORELOAD field.
	TIMG_LACTCONFIG_LACT_AUTORELOAD_Pos = 0x1d
	// Bit mask of LACT_AUTORELOAD field.
	TIMG_LACTCONFIG_LACT_AUTORELOAD_Msk = 0x20000000
	// Bit LACT_AUTORELOAD.
	TIMG_LACTCONFIG_LACT_AUTORELOAD = 0x20000000
	// Position of LACT_INCREASE field.
	TIMG_LACTCONFIG_LACT_INCREASE_Pos = 0x1e
	// Bit mask of LACT_INCREASE field.
	TIMG_LACTCONFIG_LACT_INCREASE_Msk = 0x40000000
	// Bit LACT_INCREASE.
	TIMG_LACTCONFIG_LACT_INCREASE = 0x40000000
	// Position of LACT_EN field.
	TIMG_LACTCONFIG_LACT_EN_Pos = 0x1f
	// Bit mask of LACT_EN field.
	TIMG_LACTCONFIG_LACT_EN_Msk = 0x80000000
	// Bit LACT_EN.
	TIMG_LACTCONFIG_LACT_EN = 0x80000000

	// LACTRTC
	// Position of LACT_RTC_STEP_LEN field.
	TIMG_LACTRTC_LACT_RTC_STEP_LEN_Pos = 0x6
	// Bit mask of LACT_RTC_STEP_LEN field.
	TIMG_LACTRTC_LACT_RTC_STEP_LEN_Msk = 0xffffffc0

	// LACTLO
	// Position of LACT_LO field.
	TIMG_LACTLO_LACT_LO_Pos = 0x0
	// Bit mask of LACT_LO field.
	TIMG_LACTLO_LACT_LO_Msk = 0xffffffff

	// LACTHI
	// Position of LACT_HI field.
	TIMG_LACTHI_LACT_HI_Pos = 0x0
	// Bit mask of LACT_HI field.
	TIMG_LACTHI_LACT_HI_Msk = 0xffffffff

	// LACTUPDATE
	// Position of LACT_UPDATE field.
	TIMG_LACTUPDATE_LACT_UPDATE_Pos = 0x0
	// Bit mask of LACT_UPDATE field.
	TIMG_LACTUPDATE_LACT_UPDATE_Msk = 0xffffffff

	// LACTALARMLO
	// Position of LACT_ALARM_LO field.
	TIMG_LACTALARMLO_LACT_ALARM_LO_Pos = 0x0
	// Bit mask of LACT_ALARM_LO field.
	TIMG_LACTALARMLO_LACT_ALARM_LO_Msk = 0xffffffff

	// LACTALARMHI
	// Position of LACT_ALARM_HI field.
	TIMG_LACTALARMHI_LACT_ALARM_HI_Pos = 0x0
	// Bit mask of LACT_ALARM_HI field.
	TIMG_LACTALARMHI_LACT_ALARM_HI_Msk = 0xffffffff

	// LACTLOADLO
	// Position of LACT_LOAD_LO field.
	TIMG_LACTLOADLO_LACT_LOAD_LO_Pos = 0x0
	// Bit mask of LACT_LOAD_LO field.
	TIMG_LACTLOADLO_LACT_LOAD_LO_Msk = 0xffffffff

	// LACTLOADHI
	// Position of LACT_LOAD_HI field.
	TIMG_LACTLOADHI_LACT_LOAD_HI_Pos = 0x0
	// Bit mask of LACT_LOAD_HI field.
	TIMG_LACTLOADHI_LACT_LOAD_HI_Msk = 0xffffffff

	// LACTLOAD
	// Position of LACT_LOAD field.
	TIMG_LACTLOAD_LACT_LOAD_Pos = 0x0
	// Bit mask of LACT_LOAD field.
	TIMG_LACTLOAD_LACT_LOAD_Msk = 0xffffffff

	// INT_ENA_TIMERS
	// Position of T0_INT_ENA field.
	TIMG_INT_ENA_TIMERS_T0_INT_ENA_Pos = 0x0
	// Bit mask of T0_INT_ENA field.
	TIMG_INT_ENA_TIMERS_T0_INT_ENA_Msk = 0x1
	// Bit T0_INT_ENA.
	TIMG_INT_ENA_TIMERS_T0_INT_ENA = 0x1
	// Position of T1_INT_ENA field.
	TIMG_INT_ENA_TIMERS_T1_INT_ENA_Pos = 0x1
	// Bit mask of T1_INT_ENA field.
	TIMG_INT_ENA_TIMERS_T1_INT_ENA_Msk = 0x2
	// Bit T1_INT_ENA.
	TIMG_INT_ENA_TIMERS_T1_INT_ENA = 0x2
	// Position of WDT_INT_ENA field.
	TIMG_INT_ENA_TIMERS_WDT_INT_ENA_Pos = 0x2
	// Bit mask of WDT_INT_ENA field.
	TIMG_INT_ENA_TIMERS_WDT_INT_ENA_Msk = 0x4
	// Bit WDT_INT_ENA.
	TIMG_INT_ENA_TIMERS_WDT_INT_ENA = 0x4
	// Position of LACT_INT_ENA field.
	TIMG_INT_ENA_TIMERS_LACT_INT_ENA_Pos = 0x3
	// Bit mask of LACT_INT_ENA field.
	TIMG_INT_ENA_TIMERS_LACT_INT_ENA_Msk = 0x8
	// Bit LACT_INT_ENA.
	TIMG_INT_ENA_TIMERS_LACT_INT_ENA = 0x8

	// INT_RAW_TIMERS
	// Position of T0_INT_RAW field.
	TIMG_INT_RAW_TIMERS_T0_INT_RAW_Pos = 0x0
	// Bit mask of T0_INT_RAW field.
	TIMG_INT_RAW_TIMERS_T0_INT_RAW_Msk = 0x1
	// Bit T0_INT_RAW.
	TIMG_INT_RAW_TIMERS_T0_INT_RAW = 0x1
	// Position of T1_INT_RAW field.
	TIMG_INT_RAW_TIMERS_T1_INT_RAW_Pos = 0x1
	// Bit mask of T1_INT_RAW field.
	TIMG_INT_RAW_TIMERS_T1_INT_RAW_Msk = 0x2
	// Bit T1_INT_RAW.
	TIMG_INT_RAW_TIMERS_T1_INT_RAW = 0x2
	// Position of WDT_INT_RAW field.
	TIMG_INT_RAW_TIMERS_WDT_INT_RAW_Pos = 0x2
	// Bit mask of WDT_INT_RAW field.
	TIMG_INT_RAW_TIMERS_WDT_INT_RAW_Msk = 0x4
	// Bit WDT_INT_RAW.
	TIMG_INT_RAW_TIMERS_WDT_INT_RAW = 0x4
	// Position of LACT_INT_RAW field.
	TIMG_INT_RAW_TIMERS_LACT_INT_RAW_Pos = 0x3
	// Bit mask of LACT_INT_RAW field.
	TIMG_INT_RAW_TIMERS_LACT_INT_RAW_Msk = 0x8
	// Bit LACT_INT_RAW.
	TIMG_INT_RAW_TIMERS_LACT_INT_RAW = 0x8

	// INT_ST_TIMERS
	// Position of T0_INT_ST field.
	TIMG_INT_ST_TIMERS_T0_INT_ST_Pos = 0x0
	// Bit mask of T0_INT_ST field.
	TIMG_INT_ST_TIMERS_T0_INT_ST_Msk = 0x1
	// Bit T0_INT_ST.
	TIMG_INT_ST_TIMERS_T0_INT_ST = 0x1
	// Position of T1_INT_ST field.
	TIMG_INT_ST_TIMERS_T1_INT_ST_Pos = 0x1
	// Bit mask of T1_INT_ST field.
	TIMG_INT_ST_TIMERS_T1_INT_ST_Msk = 0x2
	// Bit T1_INT_ST.
	TIMG_INT_ST_TIMERS_T1_INT_ST = 0x2
	// Position of WDT_INT_ST field.
	TIMG_INT_ST_TIMERS_WDT_INT_ST_Pos = 0x2
	// Bit mask of WDT_INT_ST field.
	TIMG_INT_ST_TIMERS_WDT_INT_ST_Msk = 0x4
	// Bit WDT_INT_ST.
	TIMG_INT_ST_TIMERS_WDT_INT_ST = 0x4
	// Position of LACT_INT_ST field.
	TIMG_INT_ST_TIMERS_LACT_INT_ST_Pos = 0x3
	// Bit mask of LACT_INT_ST field.
	TIMG_INT_ST_TIMERS_LACT_INT_ST_Msk = 0x8
	// Bit LACT_INT_ST.
	TIMG_INT_ST_TIMERS_LACT_INT_ST = 0x8

	// INT_CLR_TIMERS
	// Position of T0_INT_CLR field.
	TIMG_INT_CLR_TIMERS_T0_INT_CLR_Pos = 0x0
	// Bit mask of T0_INT_CLR field.
	TIMG_INT_CLR_TIMERS_T0_INT_CLR_Msk = 0x1
	// Bit T0_INT_CLR.
	TIMG_INT_CLR_TIMERS_T0_INT_CLR = 0x1
	// Position of T1_INT_CLR field.
	TIMG_INT_CLR_TIMERS_T1_INT_CLR_Pos = 0x1
	// Bit mask of T1_INT_CLR field.
	TIMG_INT_CLR_TIMERS_T1_INT_CLR_Msk = 0x2
	// Bit T1_INT_CLR.
	TIMG_INT_CLR_TIMERS_T1_INT_CLR = 0x2
	// Position of WDT_INT_CLR field.
	TIMG_INT_CLR_TIMERS_WDT_INT_CLR_Pos = 0x2
	// Bit mask of WDT_INT_CLR field.
	TIMG_INT_CLR_TIMERS_WDT_INT_CLR_Msk = 0x4
	// Bit WDT_INT_CLR.
	TIMG_INT_CLR_TIMERS_WDT_INT_CLR = 0x4
	// Position of LACT_INT_CLR field.
	TIMG_INT_CLR_TIMERS_LACT_INT_CLR_Pos = 0x3
	// Bit mask of LACT_INT_CLR field.
	TIMG_INT_CLR_TIMERS_LACT_INT_CLR_Msk = 0x8
	// Bit LACT_INT_CLR.
	TIMG_INT_CLR_TIMERS_LACT_INT_CLR = 0x8

	// NTIMERS_DATE
	// Position of NTIMERS_DATE field.
	TIMG_NTIMERS_DATE_NTIMERS_DATE_Pos = 0x0
	// Bit mask of NTIMERS_DATE field.
	TIMG_NTIMERS_DATE_NTIMERS_DATE_Msk = 0xfffffff

	// TIMGCLK
	// Position of CLK_EN field.
	TIMG_TIMGCLK_CLK_EN_Pos = 0x1f
	// Bit mask of CLK_EN field.
	TIMG_TIMGCLK_CLK_EN_Msk = 0x80000000
	// Bit CLK_EN.
	TIMG_TIMGCLK_CLK_EN = 0x80000000
)

// Timer Group 0
type TIMG_Type struct {
	T0CONFIG       volatile.Register32 // 0x0
	T0LO           volatile.Register32 // 0x4
	T0HI           volatile.Register32 // 0x8
	T0UPDATE       volatile.Register32 // 0xC
	T0ALARMLO      volatile.Register32 // 0x10
	T0ALARMHI      volatile.Register32 // 0x14
	T0LOADLO       volatile.Register32 // 0x18
	T0LOADHI       volatile.Register32 // 0x1C
	T0LOAD         volatile.Register32 // 0x20
	T1CONFIG       volatile.Register32 // 0x24
	T1LO           volatile.Register32 // 0x28
	T1HI           volatile.Register32 // 0x2C
	T1UPDATE       volatile.Register32 // 0x30
	T1ALARMLO      volatile.Register32 // 0x34
	T1ALARMHI      volatile.Register32 // 0x38
	T1LOADLO       volatile.Register32 // 0x3C
	T1LOADHI       volatile.Register32 // 0x40
	T1LOAD         volatile.Register32 // 0x44
	WDTCONFIG0     volatile.Register32 // 0x48
	WDTCONFIG1     volatile.Register32 // 0x4C
	WDTCONFIG2     volatile.Register32 // 0x50
	WDTCONFIG3     volatile.Register32 // 0x54
	WDTCONFIG4     volatile.Register32 // 0x58
	WDTCONFIG5     volatile.Register32 // 0x5C
	WDTFEED        volatile.Register32 // 0x60
	WDTWPROTECT    volatile.Register32 // 0x64
	RTCCALICFG     volatile.Register32 // 0x68
	RTCCALICFG1    volatile.Register32 // 0x6C
	LACTCONFIG     volatile.Register32 // 0x70
	LACTRTC        volatile.Register32 // 0x74
	LACTLO         volatile.Register32 // 0x78
	LACTHI         volatile.Register32 // 0x7C
	LACTUPDATE     volatile.Register32 // 0x80
	LACTALARMLO    volatile.Register32 // 0x84
	LACTALARMHI    volatile.Register32 // 0x88
	LACTLOADLO     volatile.Register32 // 0x8C
	LACTLOADHI     volatile.Register32 // 0x90
	LACTLOAD       volatile.Register32 // 0x94
	INT_ENA_TIMERS volatile.Register32 // 0x98
	INT_RAW_TIMERS volatile.Register32 // 0x9C
	INT_ST_TIMERS  volatile.Register32 // 0xA0
	INT_CLR_TIMERS volatile.Register32 // 0xA4
	_              [80]byte
	NTIMERS_DATE   volatile.Register32 // 0xF8
	TIMGCLK        volatile.Register32 // 0xFC
}

// TIMG.T0CONFIG
func (o *TIMG_Type) SetT0CONFIG_ALARM_EN(value uint32) {
	volatile.StoreUint32(&o.T0CONFIG.Reg, volatile.LoadUint32(&o.T0CONFIG.Reg)&^(0x400)|value<<10)
}
func (o *TIMG_Type) GetT0CONFIG_ALARM_EN() uint32 {
	return (volatile.LoadUint32(&o.T0CONFIG.Reg) & 0x400) >> 10
}
func (o *TIMG_Type) SetT0CONFIG_LEVEL_INT_EN(value uint32) {
	volatile.StoreUint32(&o.T0CONFIG.Reg, volatile.LoadUint32(&o.T0CONFIG.Reg)&^(0x800)|value<<11)
}
func (o *TIMG_Type) GetT0CONFIG_LEVEL_INT_EN() uint32 {
	return (volatile.LoadUint32(&o.T0CONFIG.Reg) & 0x800) >> 11
}
func (o *TIMG_Type) SetT0CONFIG_EDGE_INT_EN(value uint32) {
	volatile.StoreUint32(&o.T0CONFIG.Reg, volatile.LoadUint32(&o.T0CONFIG.Reg)&^(0x1000)|value<<12)
}
func (o *TIMG_Type) GetT0CONFIG_EDGE_INT_EN() uint32 {
	return (volatile.LoadUint32(&o.T0CONFIG.Reg) & 0x1000) >> 12
}
func (o *TIMG_Type) SetT0CONFIG_DIVIDER(value uint32) {
	volatile.StoreUint32(&o.T0CONFIG.Reg, volatile.LoadUint32(&o.T0CONFIG.Reg)&^(0x1fffe000)|value<<13)
}
func (o *TIMG_Type) GetT0CONFIG_DIVIDER() uint32 {
	return (volatile.LoadUint32(&o.T0CONFIG.Reg) & 0x1fffe000) >> 13
}
func (o *TIMG_Type) SetT0CONFIG_AUTORELOAD(value uint32) {
	volatile.StoreUint32(&o.T0CONFIG.Reg, volatile.LoadUint32(&o.T0CONFIG.Reg)&^(0x20000000)|value<<29)
}
func (o *TIMG_Type) GetT0CONFIG_AUTORELOAD() uint32 {
	return (volatile.LoadUint32(&o.T0CONFIG.Reg) & 0x20000000) >> 29
}
func (o *TIMG_Type) SetT0CONFIG_INCREASE(value uint32) {
	volatile.StoreUint32(&o.T0CONFIG.Reg, volatile.LoadUint32(&o.T0CONFIG.Reg)&^(0x40000000)|value<<30)
}
func (o *TIMG_Type) GetT0CONFIG_INCREASE() uint32 {
	return (volatile.LoadUint32(&o.T0CONFIG.Reg) & 0x40000000) >> 30
}
func (o *TIMG_Type) SetT0CONFIG_EN(value uint32) {
	volatile.StoreUint32(&o.T0CONFIG.Reg, volatile.LoadUint32(&o.T0CONFIG.Reg)&^(0x80000000)|value<<31)
}
func (o *TIMG_Type) GetT0CONFIG_EN() uint32 {
	return (volatile.LoadUint32(&o.T0CONFIG.Reg) & 0x80000000) >> 31
}

// TIMG.T0LO
func (o *TIMG_Type) SetT0LO(value uint32) {
	volatile.StoreUint32(&o.T0LO.Reg, value)
}
func (o *TIMG_Type) GetT0LO() uint32 {
	return volatile.LoadUint32(&o.T0LO.Reg)
}

// TIMG.T0HI
func (o *TIMG_Type) SetT0HI(value uint32) {
	volatile.StoreUint32(&o.T0HI.Reg, value)
}
func (o *TIMG_Type) GetT0HI() uint32 {
	return volatile.LoadUint32(&o.T0HI.Reg)
}

// TIMG.T0UPDATE
func (o *TIMG_Type) SetT0UPDATE(value uint32) {
	volatile.StoreUint32(&o.T0UPDATE.Reg, value)
}
func (o *TIMG_Type) GetT0UPDATE() uint32 {
	return volatile.LoadUint32(&o.T0UPDATE.Reg)
}

// TIMG.T0ALARMLO
func (o *TIMG_Type) SetT0ALARMLO(value uint32) {
	volatile.StoreUint32(&o.T0ALARMLO.Reg, value)
}
func (o *TIMG_Type) GetT0ALARMLO() uint32 {
	return volatile.LoadUint32(&o.T0ALARMLO.Reg)
}

// TIMG.T0ALARMHI
func (o *TIMG_Type) SetT0ALARMHI(value uint32) {
	volatile.StoreUint32(&o.T0ALARMHI.Reg, value)
}
func (o *TIMG_Type) GetT0ALARMHI() uint32 {
	return volatile.LoadUint32(&o.T0ALARMHI.Reg)
}

// TIMG.T0LOADLO
func (o *TIMG_Type) SetT0LOADLO(value uint32) {
	volatile.StoreUint32(&o.T0LOADLO.Reg, value)
}
func (o *TIMG_Type) GetT0LOADLO() uint32 {
	return volatile.LoadUint32(&o.T0LOADLO.Reg)
}

// TIMG.T0LOADHI
func (o *TIMG_Type) SetT0LOADHI(value uint32) {
	volatile.StoreUint32(&o.T0LOADHI.Reg, value)
}
func (o *TIMG_Type) GetT0LOADHI() uint32 {
	return volatile.LoadUint32(&o.T0LOADHI.Reg)
}

// TIMG.T0LOAD
func (o *TIMG_Type) SetT0LOAD(value uint32) {
	volatile.StoreUint32(&o.T0LOAD.Reg, value)
}
func (o *TIMG_Type) GetT0LOAD() uint32 {
	return volatile.LoadUint32(&o.T0LOAD.Reg)
}

// TIMG.T1CONFIG
func (o *TIMG_Type) SetT1CONFIG_ALARM_EN(value uint32) {
	volatile.StoreUint32(&o.T1CONFIG.Reg, volatile.LoadUint32(&o.T1CONFIG.Reg)&^(0x400)|value<<10)
}
func (o *TIMG_Type) GetT1CONFIG_ALARM_EN() uint32 {
	return (volatile.LoadUint32(&o.T1CONFIG.Reg) & 0x400) >> 10
}
func (o *TIMG_Type) SetT1CONFIG_LEVEL_INT_EN(value uint32) {
	volatile.StoreUint32(&o.T1CONFIG.Reg, volatile.LoadUint32(&o.T1CONFIG.Reg)&^(0x800)|value<<11)
}
func (o *TIMG_Type) GetT1CONFIG_LEVEL_INT_EN() uint32 {
	return (volatile.LoadUint32(&o.T1CONFIG.Reg) & 0x800) >> 11
}
func (o *TIMG_Type) SetT1CONFIG_EDGE_INT_EN(value uint32) {
	volatile.StoreUint32(&o.T1CONFIG.Reg, volatile.LoadUint32(&o.T1CONFIG.Reg)&^(0x1000)|value<<12)
}
func (o *TIMG_Type) GetT1CONFIG_EDGE_INT_EN() uint32 {
	return (volatile.LoadUint32(&o.T1CONFIG.Reg) & 0x1000) >> 12
}
func (o *TIMG_Type) SetT1CONFIG_DIVIDER(value uint32) {
	volatile.StoreUint32(&o.T1CONFIG.Reg, volatile.LoadUint32(&o.T1CONFIG.Reg)&^(0x1fffe000)|value<<13)
}
func (o *TIMG_Type) GetT1CONFIG_DIVIDER() uint32 {
	return (volatile.LoadUint32(&o.T1CONFIG.Reg) & 0x1fffe000) >> 13
}
func (o *TIMG_Type) SetT1CONFIG_AUTORELOAD(value uint32) {
	volatile.StoreUint32(&o.T1CONFIG.Reg, volatile.LoadUint32(&o.T1CONFIG.Reg)&^(0x20000000)|value<<29)
}
func (o *TIMG_Type) GetT1CONFIG_AUTORELOAD() uint32 {
	return (volatile.LoadUint32(&o.T1CONFIG.Reg) & 0x20000000) >> 29
}
func (o *TIMG_Type) SetT1CONFIG_INCREASE(value uint32) {
	volatile.StoreUint32(&o.T1CONFIG.Reg, volatile.LoadUint32(&o.T1CONFIG.Reg)&^(0x40000000)|value<<30)
}
func (o *TIMG_Type) GetT1CONFIG_INCREASE() uint32 {
	return (volatile.LoadUint32(&o.T1CONFIG.Reg) & 0x40000000) >> 30
}
func (o *TIMG_Type) SetT1CONFIG_EN(value uint32) {
	volatile.StoreUint32(&o.T1CONFIG.Reg, volatile.LoadUint32(&o.T1CONFIG.Reg)&^(0x80000000)|value<<31)
}
func (o *TIMG_Type) GetT1CONFIG_EN() uint32 {
	return (volatile.LoadUint32(&o.T1CONFIG.Reg) & 0x80000000) >> 31
}

// TIMG.T1LO
func (o *TIMG_Type) SetT1LO(value uint32) {
	volatile.StoreUint32(&o.T1LO.Reg, value)
}
func (o *TIMG_Type) GetT1LO() uint32 {
	return volatile.LoadUint32(&o.T1LO.Reg)
}

// TIMG.T1HI
func (o *TIMG_Type) SetT1HI(value uint32) {
	volatile.StoreUint32(&o.T1HI.Reg, value)
}
func (o *TIMG_Type) GetT1HI() uint32 {
	return volatile.LoadUint32(&o.T1HI.Reg)
}

// TIMG.T1UPDATE
func (o *TIMG_Type) SetT1UPDATE(value uint32) {
	volatile.StoreUint32(&o.T1UPDATE.Reg, value)
}
func (o *TIMG_Type) GetT1UPDATE() uint32 {
	return volatile.LoadUint32(&o.T1UPDATE.Reg)
}

// TIMG.T1ALARMLO
func (o *TIMG_Type) SetT1ALARMLO(value uint32) {
	volatile.StoreUint32(&o.T1ALARMLO.Reg, value)
}
func (o *TIMG_Type) GetT1ALARMLO() uint32 {
	return volatile.LoadUint32(&o.T1ALARMLO.Reg)
}

// TIMG.T1ALARMHI
func (o *TIMG_Type) SetT1ALARMHI(value uint32) {
	volatile.StoreUint32(&o.T1ALARMHI.Reg, value)
}
func (o *TIMG_Type) GetT1ALARMHI() uint32 {
	return volatile.LoadUint32(&o.T1ALARMHI.Reg)
}

// TIMG.T1LOADLO
func (o *TIMG_Type) SetT1LOADLO(value uint32) {
	volatile.StoreUint32(&o.T1LOADLO.Reg, value)
}
func (o *TIMG_Type) GetT1LOADLO() uint32 {
	return volatile.LoadUint32(&o.T1LOADLO.Reg)
}

// TIMG.T1LOADHI
func (o *TIMG_Type) SetT1LOADHI(value uint32) {
	volatile.StoreUint32(&o.T1LOADHI.Reg, value)
}
func (o *TIMG_Type) GetT1LOADHI() uint32 {
	return volatile.LoadUint32(&o.T1LOADHI.Reg)
}

// TIMG.T1LOAD
func (o *TIMG_Type) SetT1LOAD(value uint32) {
	volatile.StoreUint32(&o.T1LOAD.Reg, value)
}
func (o *TIMG_Type) GetT1LOAD() uint32 {
	return volatile.LoadUint32(&o.T1LOAD.Reg)
}

// TIMG.WDTCONFIG0
func (o *TIMG_Type) SetWDTCONFIG0_WDT_FLASHBOOT_MOD_EN(value uint32) {
	volatile.StoreUint32(&o.WDTCONFIG0.Reg, volatile.LoadUint32(&o.WDTCONFIG0.Reg)&^(0x4000)|value<<14)
}
func (o *TIMG_Type) GetWDTCONFIG0_WDT_FLASHBOOT_MOD_EN() uint32 {
	return (volatile.LoadUint32(&o.WDTCONFIG0.Reg) & 0x4000) >> 14
}
func (o *TIMG_Type) SetWDTCONFIG0_WDT_SYS_RESET_LENGTH(value uint32) {
	volatile.StoreUint32(&o.WDTCONFIG0.Reg, volatile.LoadUint32(&o.WDTCONFIG0.Reg)&^(0x38000)|value<<15)
}
func (o *TIMG_Type) GetWDTCONFIG0_WDT_SYS_RESET_LENGTH() uint32 {
	return (volatile.LoadUint32(&o.WDTCONFIG0.Reg) & 0x38000) >> 15
}
func (o *TIMG_Type) SetWDTCONFIG0_WDT_CPU_RESET_LENGTH(value uint32) {
	volatile.StoreUint32(&o.WDTCONFIG0.Reg, volatile.LoadUint32(&o.WDTCONFIG0.Reg)&^(0x1c0000)|value<<18)
}
func (o *TIMG_Type) GetWDTCONFIG0_WDT_CPU_RESET_LENGTH() uint32 {
	return (volatile.LoadUint32(&o.WDTCONFIG0.Reg) & 0x1c0000) >> 18
}
func (o *TIMG_Type) SetWDTCONFIG0_WDT_LEVEL_INT_EN(value uint32) {
	volatile.StoreUint32(&o.WDTCONFIG0.Reg, volatile.LoadUint32(&o.WDTCONFIG0.Reg)&^(0x200000)|value<<21)
}
func (o *TIMG_Type) GetWDTCONFIG0_WDT_LEVEL_INT_EN() uint32 {
	return (volatile.LoadUint32(&o.WDTCONFIG0.Reg) & 0x200000) >> 21
}
func (o *TIMG_Type) SetWDTCONFIG0_WDT_EDGE_INT_EN(value uint32) {
	volatile.StoreUint32(&o.WDTCONFIG0.Reg, volatile.LoadUint32(&o.WDTCONFIG0.Reg)&^(0x400000)|value<<22)
}
func (o *TIMG_Type) GetWDTCONFIG0_WDT_EDGE_INT_EN() uint32 {
	return (volatile.LoadUint32(&o.WDTCONFIG0.Reg) & 0x400000) >> 22
}
func (o *TIMG_Type) SetWDTCONFIG0_WDT_STG3(value uint32) {
	volatile.StoreUint32(&o.WDTCONFIG0.Reg, volatile.LoadUint32(&o.WDTCONFIG0.Reg)&^(0x1800000)|value<<23)
}
func (o *TIMG_Type) GetWDTCONFIG0_WDT_STG3() uint32 {
	return (volatile.LoadUint32(&o.WDTCONFIG0.Reg) & 0x1800000) >> 23
}
func (o *TIMG_Type) SetWDTCONFIG0_WDT_STG2(value uint32) {
	volatile.StoreUint32(&o.WDTCONFIG0.Reg, volatile.LoadUint32(&o.WDTCONFIG0.Reg)&^(0x6000000)|value<<25)
}
func (o *TIMG_Type) GetWDTCONFIG0_WDT_STG2() uint32 {
	return (volatile.LoadUint32(&o.WDTCONFIG0.Reg) & 0x6000000) >> 25
}
func (o *TIMG_Type) SetWDTCONFIG0_WDT_STG1(value uint32) {
	volatile.StoreUint32(&o.WDTCONFIG0.Reg, volatile.LoadUint32(&o.WDTCONFIG0.Reg)&^(0x18000000)|value<<27)
}
func (o *TIMG_Type) GetWDTCONFIG0_WDT_STG1() uint32 {
	return (volatile.LoadUint32(&o.WDTCONFIG0.Reg) & 0x18000000) >> 27
}
func (o *TIMG_Type) SetWDTCONFIG0_WDT_STG0(value uint32) {
	volatile.StoreUint32(&o.WDTCONFIG0.Reg, volatile.LoadUint32(&o.WDTCONFIG0.Reg)&^(0x60000000)|value<<29)
}
func (o *TIMG_Type) GetWDTCONFIG0_WDT_STG0() uint32 {
	return (volatile.LoadUint32(&o.WDTCONFIG0.Reg) & 0x60000000) >> 29
}
func (o *TIMG_Type) SetWDTCONFIG0_WDT_EN(value uint32) {
	volatile.StoreUint32(&o.WDTCONFIG0.Reg, volatile.LoadUint32(&o.WDTCONFIG0.Reg)&^(0x80000000)|value<<31)
}
func (o *TIMG_Type) GetWDTCONFIG0_WDT_EN() uint32 {
	return (volatile.LoadUint32(&o.WDTCONFIG0.Reg) & 0x80000000) >> 31
}

// TIMG.WDTCONFIG1
func (o *TIMG_Type) SetWDTCONFIG1_WDT_CLK_PRESCALE(value uint32) {
	volatile.StoreUint32(&o.WDTCONFIG1.Reg, volatile.LoadUint32(&o.WDTCONFIG1.Reg)&^(0xffff0000)|value<<16)
}
func (o *TIMG_Type) GetWDTCONFIG1_WDT_CLK_PRESCALE() uint32 {
	return (volatile.LoadUint32(&o.WDTCONFIG1.Reg) & 0xffff0000) >> 16
}

// TIMG.WDTCONFIG2
func (o *TIMG_Type) SetWDTCONFIG2(value uint32) {
	volatile.StoreUint32(&o.WDTCONFIG2.Reg, value)
}
func (o *TIMG_Type) GetWDTCONFIG2() uint32 {
	return volatile.LoadUint32(&o.WDTCONFIG2.Reg)
}

// TIMG.WDTCONFIG3
func (o *TIMG_Type) SetWDTCONFIG3(value uint32) {
	volatile.StoreUint32(&o.WDTCONFIG3.Reg, value)
}
func (o *TIMG_Type) GetWDTCONFIG3() uint32 {
	return volatile.LoadUint32(&o.WDTCONFIG3.Reg)
}

// TIMG.WDTCONFIG4
func (o *TIMG_Type) SetWDTCONFIG4(value uint32) {
	volatile.StoreUint32(&o.WDTCONFIG4.Reg, value)
}
func (o *TIMG_Type) GetWDTCONFIG4() uint32 {
	return volatile.LoadUint32(&o.WDTCONFIG4.Reg)
}

// TIMG.WDTCONFIG5
func (o *TIMG_Type) SetWDTCONFIG5(value uint32) {
	volatile.StoreUint32(&o.WDTCONFIG5.Reg, value)
}
func (o *TIMG_Type) GetWDTCONFIG5() uint32 {
	return volatile.LoadUint32(&o.WDTCONFIG5.Reg)
}

// TIMG.WDTFEED
func (o *TIMG_Type) SetWDTFEED(value uint32) {
	volatile.StoreUint32(&o.WDTFEED.Reg, value)
}
func (o *TIMG_Type) GetWDTFEED() uint32 {
	return volatile.LoadUint32(&o.WDTFEED.Reg)
}

// TIMG.WDTWPROTECT
func (o *TIMG_Type) SetWDTWPROTECT(value uint32) {
	volatile.StoreUint32(&o.WDTWPROTECT.Reg, value)
}
func (o *TIMG_Type) GetWDTWPROTECT() uint32 {
	return volatile.LoadUint32(&o.WDTWPROTECT.Reg)
}

// TIMG.RTCCALICFG
func (o *TIMG_Type) SetRTCCALICFG_RTC_CALI_START_CYCLING(value uint32) {
	volatile.StoreUint32(&o.RTCCALICFG.Reg, volatile.LoadUint32(&o.RTCCALICFG.Reg)&^(0x1000)|value<<12)
}
func (o *TIMG_Type) GetRTCCALICFG_RTC_CALI_START_CYCLING() uint32 {
	return (volatile.LoadUint32(&o.RTCCALICFG.Reg) & 0x1000) >> 12
}
func (o *TIMG_Type) SetRTCCALICFG_RTC_CALI_CLK_SEL(value uint32) {
	volatile.StoreUint32(&o.RTCCALICFG.Reg, volatile.LoadUint32(&o.RTCCALICFG.Reg)&^(0x6000)|value<<13)
}
func (o *TIMG_Type) GetRTCCALICFG_RTC_CALI_CLK_SEL() uint32 {
	return (volatile.LoadUint32(&o.RTCCALICFG.Reg) & 0x6000) >> 13
}
func (o *TIMG_Type) SetRTCCALICFG_RTC_CALI_RDY(value uint32) {
	volatile.StoreUint32(&o.RTCCALICFG.Reg, volatile.LoadUint32(&o.RTCCALICFG.Reg)&^(0x8000)|value<<15)
}
func (o *TIMG_Type) GetRTCCALICFG_RTC_CALI_RDY() uint32 {
	return (volatile.LoadUint32(&o.RTCCALICFG.Reg) & 0x8000) >> 15
}
func (o *TIMG_Type) SetRTCCALICFG_RTC_CALI_MAX(value uint32) {
	volatile.StoreUint32(&o.RTCCALICFG.Reg, volatile.LoadUint32(&o.RTCCALICFG.Reg)&^(0x7fff0000)|value<<16)
}
func (o *TIMG_Type) GetRTCCALICFG_RTC_CALI_MAX() uint32 {
	return (volatile.LoadUint32(&o.RTCCALICFG.Reg) & 0x7fff0000) >> 16
}
func (o *TIMG_Type) SetRTCCALICFG_RTC_CALI_START(value uint32) {
	volatile.StoreUint32(&o.RTCCALICFG.Reg, volatile.LoadUint32(&o.RTCCALICFG.Reg)&^(0x80000000)|value<<31)
}
func (o *TIMG_Type) GetRTCCALICFG_RTC_CALI_START() uint32 {
	return (volatile.LoadUint32(&o.RTCCALICFG.Reg) & 0x80000000) >> 31
}

// TIMG.RTCCALICFG1
func (o *TIMG_Type) SetRTCCALICFG1_RTC_CALI_VALUE(value uint32) {
	volatile.StoreUint32(&o.RTCCALICFG1.Reg, volatile.LoadUint32(&o.RTCCALICFG1.Reg)&^(0xffffff80)|value<<7)
}
func (o *TIMG_Type) GetRTCCALICFG1_RTC_CALI_VALUE() uint32 {
	return (volatile.LoadUint32(&o.RTCCALICFG1.Reg) & 0xffffff80) >> 7
}

// TIMG.LACTCONFIG
func (o *TIMG_Type) SetLACTCONFIG_LACT_RTC_ONLY(value uint32) {
	volatile.StoreUint32(&o.LACTCONFIG.Reg, volatile.LoadUint32(&o.LACTCONFIG.Reg)&^(0x80)|value<<7)
}
func (o *TIMG_Type) GetLACTCONFIG_LACT_RTC_ONLY() uint32 {
	return (volatile.LoadUint32(&o.LACTCONFIG.Reg) & 0x80) >> 7
}
func (o *TIMG_Type) SetLACTCONFIG_LACT_CPST_EN(value uint32) {
	volatile.StoreUint32(&o.LACTCONFIG.Reg, volatile.LoadUint32(&o.LACTCONFIG.Reg)&^(0x100)|value<<8)
}
func (o *TIMG_Type) GetLACTCONFIG_LACT_CPST_EN() uint32 {
	return (volatile.LoadUint32(&o.LACTCONFIG.Reg) & 0x100) >> 8
}
func (o *TIMG_Type) SetLACTCONFIG_LACT_LAC_EN(value uint32) {
	volatile.StoreUint32(&o.LACTCONFIG.Reg, volatile.LoadUint32(&o.LACTCONFIG.Reg)&^(0x200)|value<<9)
}
func (o *TIMG_Type) GetLACTCONFIG_LACT_LAC_EN() uint32 {
	return (volatile.LoadUint32(&o.LACTCONFIG.Reg) & 0x200) >> 9
}
func (o *TIMG_Type) SetLACTCONFIG_LACT_ALARM_EN(value uint32) {
	volatile.StoreUint32(&o.LACTCONFIG.Reg, volatile.LoadUint32(&o.LACTCONFIG.Reg)&^(0x400)|value<<10)
}
func (o *TIMG_Type) GetLACTCONFIG_LACT_ALARM_EN() uint32 {
	return (volatile.LoadUint32(&o.LACTCONFIG.Reg) & 0x400) >> 10
}
func (o *TIMG_Type) SetLACTCONFIG_LACT_LEVEL_INT_EN(value uint32) {
	volatile.StoreUint32(&o.LACTCONFIG.Reg, volatile.LoadUint32(&o.LACTCONFIG.Reg)&^(0x800)|value<<11)
}
func (o *TIMG_Type) GetLACTCONFIG_LACT_LEVEL_INT_EN() uint32 {
	return (volatile.LoadUint32(&o.LACTCONFIG.Reg) & 0x800) >> 11
}
func (o *TIMG_Type) SetLACTCONFIG_LACT_EDGE_INT_EN(value uint32) {
	volatile.StoreUint32(&o.LACTCONFIG.Reg, volatile.LoadUint32(&o.LACTCONFIG.Reg)&^(0x1000)|value<<12)
}
func (o *TIMG_Type) GetLACTCONFIG_LACT_EDGE_INT_EN() uint32 {
	return (volatile.LoadUint32(&o.LACTCONFIG.Reg) & 0x1000) >> 12
}
func (o *TIMG_Type) SetLACTCONFIG_LACT_DIVIDER(value uint32) {
	volatile.StoreUint32(&o.LACTCONFIG.Reg, volatile.LoadUint32(&o.LACTCONFIG.Reg)&^(0x1fffe000)|value<<13)
}
func (o *TIMG_Type) GetLACTCONFIG_LACT_DIVIDER() uint32 {
	return (volatile.LoadUint32(&o.LACTCONFIG.Reg) & 0x1fffe000) >> 13
}
func (o *TIMG_Type) SetLACTCONFIG_LACT_AUTORELOAD(value uint32) {
	volatile.StoreUint32(&o.LACTCONFIG.Reg, volatile.LoadUint32(&o.LACTCONFIG.Reg)&^(0x20000000)|value<<29)
}
func (o *TIMG_Type) GetLACTCONFIG_LACT_AUTORELOAD() uint32 {
	return (volatile.LoadUint32(&o.LACTCONFIG.Reg) & 0x20000000) >> 29
}
func (o *TIMG_Type) SetLACTCONFIG_LACT_INCREASE(value uint32) {
	volatile.StoreUint32(&o.LACTCONFIG.Reg, volatile.LoadUint32(&o.LACTCONFIG.Reg)&^(0x40000000)|value<<30)
}
func (o *TIMG_Type) GetLACTCONFIG_LACT_INCREASE() uint32 {
	return (volatile.LoadUint32(&o.LACTCONFIG.Reg) & 0x40000000) >> 30
}
func (o *TIMG_Type) SetLACTCONFIG_LACT_EN(value uint32) {
	volatile.StoreUint32(&o.LACTCONFIG.Reg, volatile.LoadUint32(&o.LACTCONFIG.Reg)&^(0x80000000)|value<<31)
}
func (o *TIMG_Type) GetLACTCONFIG_LACT_EN() uint32 {
	return (volatile.LoadUint32(&o.LACTCONFIG.Reg) & 0x80000000) >> 31
}

// TIMG.LACTRTC
func (o *TIMG_Type) SetLACTRTC_LACT_RTC_STEP_LEN(value uint32) {
	volatile.StoreUint32(&o.LACTRTC.Reg, volatile.LoadUint32(&o.LACTRTC.Reg)&^(0xffffffc0)|value<<6)
}
func (o *TIMG_Type) GetLACTRTC_LACT_RTC_STEP_LEN() uint32 {
	return (volatile.LoadUint32(&o.LACTRTC.Reg) & 0xffffffc0) >> 6
}

// TIMG.LACTLO
func (o *TIMG_Type) SetLACTLO(value uint32) {
	volatile.StoreUint32(&o.LACTLO.Reg, value)
}
func (o *TIMG_Type) GetLACTLO() uint32 {
	return volatile.LoadUint32(&o.LACTLO.Reg)
}

// TIMG.LACTHI
func (o *TIMG_Type) SetLACTHI(value uint32) {
	volatile.StoreUint32(&o.LACTHI.Reg, value)
}
func (o *TIMG_Type) GetLACTHI() uint32 {
	return volatile.LoadUint32(&o.LACTHI.Reg)
}

// TIMG.LACTUPDATE
func (o *TIMG_Type) SetLACTUPDATE(value uint32) {
	volatile.StoreUint32(&o.LACTUPDATE.Reg, value)
}
func (o *TIMG_Type) GetLACTUPDATE() uint32 {
	return volatile.LoadUint32(&o.LACTUPDATE.Reg)
}

// TIMG.LACTALARMLO
func (o *TIMG_Type) SetLACTALARMLO(value uint32) {
	volatile.StoreUint32(&o.LACTALARMLO.Reg, value)
}
func (o *TIMG_Type) GetLACTALARMLO() uint32 {
	return volatile.LoadUint32(&o.LACTALARMLO.Reg)
}

// TIMG.LACTALARMHI
func (o *TIMG_Type) SetLACTALARMHI(value uint32) {
	volatile.StoreUint32(&o.LACTALARMHI.Reg, value)
}
func (o *TIMG_Type) GetLACTALARMHI() uint32 {
	return volatile.LoadUint32(&o.LACTALARMHI.Reg)
}

// TIMG.LACTLOADLO
func (o *TIMG_Type) SetLACTLOADLO(value uint32) {
	volatile.StoreUint32(&o.LACTLOADLO.Reg, value)
}
func (o *TIMG_Type) GetLACTLOADLO() uint32 {
	return volatile.LoadUint32(&o.LACTLOADLO.Reg)
}

// TIMG.LACTLOADHI
func (o *TIMG_Type) SetLACTLOADHI(value uint32) {
	volatile.StoreUint32(&o.LACTLOADHI.Reg, value)
}
func (o *TIMG_Type) GetLACTLOADHI() uint32 {
	return volatile.LoadUint32(&o.LACTLOADHI.Reg)
}

// TIMG.LACTLOAD
func (o *TIMG_Type) SetLACTLOAD(value uint32) {
	volatile.StoreUint32(&o.LACTLOAD.Reg, value)
}
func (o *TIMG_Type) GetLACTLOAD() uint32 {
	return volatile.LoadUint32(&o.LACTLOAD.Reg)
}

// TIMG.INT_ENA_TIMERS
func (o *TIMG_Type) SetINT_ENA_TIMERS_T0_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA_TIMERS.Reg, volatile.LoadUint32(&o.INT_ENA_TIMERS.Reg)&^(0x1)|value)
}
func (o *TIMG_Type) GetINT_ENA_TIMERS_T0_INT_ENA() uint32 {
	return volatile.LoadUint32(&o.INT_ENA_TIMERS.Reg) & 0x1
}
func (o *TIMG_Type) SetINT_ENA_TIMERS_T1_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA_TIMERS.Reg, volatile.LoadUint32(&o.INT_ENA_TIMERS.Reg)&^(0x2)|value<<1)
}
func (o *TIMG_Type) GetINT_ENA_TIMERS_T1_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA_TIMERS.Reg) & 0x2) >> 1
}
func (o *TIMG_Type) SetINT_ENA_TIMERS_WDT_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA_TIMERS.Reg, volatile.LoadUint32(&o.INT_ENA_TIMERS.Reg)&^(0x4)|value<<2)
}
func (o *TIMG_Type) GetINT_ENA_TIMERS_WDT_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA_TIMERS.Reg) & 0x4) >> 2
}
func (o *TIMG_Type) SetINT_ENA_TIMERS_LACT_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.INT_ENA_TIMERS.Reg, volatile.LoadUint32(&o.INT_ENA_TIMERS.Reg)&^(0x8)|value<<3)
}
func (o *TIMG_Type) GetINT_ENA_TIMERS_LACT_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.INT_ENA_TIMERS.Reg) & 0x8) >> 3
}

// TIMG.INT_RAW_TIMERS
func (o *TIMG_Type) SetINT_RAW_TIMERS_T0_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW_TIMERS.Reg, volatile.LoadUint32(&o.INT_RAW_TIMERS.Reg)&^(0x1)|value)
}
func (o *TIMG_Type) GetINT_RAW_TIMERS_T0_INT_RAW() uint32 {
	return volatile.LoadUint32(&o.INT_RAW_TIMERS.Reg) & 0x1
}
func (o *TIMG_Type) SetINT_RAW_TIMERS_T1_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW_TIMERS.Reg, volatile.LoadUint32(&o.INT_RAW_TIMERS.Reg)&^(0x2)|value<<1)
}
func (o *TIMG_Type) GetINT_RAW_TIMERS_T1_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW_TIMERS.Reg) & 0x2) >> 1
}
func (o *TIMG_Type) SetINT_RAW_TIMERS_WDT_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW_TIMERS.Reg, volatile.LoadUint32(&o.INT_RAW_TIMERS.Reg)&^(0x4)|value<<2)
}
func (o *TIMG_Type) GetINT_RAW_TIMERS_WDT_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW_TIMERS.Reg) & 0x4) >> 2
}
func (o *TIMG_Type) SetINT_RAW_TIMERS_LACT_INT_RAW(value uint32) {
	volatile.StoreUint32(&o.INT_RAW_TIMERS.Reg, volatile.LoadUint32(&o.INT_RAW_TIMERS.Reg)&^(0x8)|value<<3)
}
func (o *TIMG_Type) GetINT_RAW_TIMERS_LACT_INT_RAW() uint32 {
	return (volatile.LoadUint32(&o.INT_RAW_TIMERS.Reg) & 0x8) >> 3
}

// TIMG.INT_ST_TIMERS
func (o *TIMG_Type) SetINT_ST_TIMERS_T0_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST_TIMERS.Reg, volatile.LoadUint32(&o.INT_ST_TIMERS.Reg)&^(0x1)|value)
}
func (o *TIMG_Type) GetINT_ST_TIMERS_T0_INT_ST() uint32 {
	return volatile.LoadUint32(&o.INT_ST_TIMERS.Reg) & 0x1
}
func (o *TIMG_Type) SetINT_ST_TIMERS_T1_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST_TIMERS.Reg, volatile.LoadUint32(&o.INT_ST_TIMERS.Reg)&^(0x2)|value<<1)
}
func (o *TIMG_Type) GetINT_ST_TIMERS_T1_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST_TIMERS.Reg) & 0x2) >> 1
}
func (o *TIMG_Type) SetINT_ST_TIMERS_WDT_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST_TIMERS.Reg, volatile.LoadUint32(&o.INT_ST_TIMERS.Reg)&^(0x4)|value<<2)
}
func (o *TIMG_Type) GetINT_ST_TIMERS_WDT_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST_TIMERS.Reg) & 0x4) >> 2
}
func (o *TIMG_Type) SetINT_ST_TIMERS_LACT_INT_ST(value uint32) {
	volatile.StoreUint32(&o.INT_ST_TIMERS.Reg, volatile.LoadUint32(&o.INT_ST_TIMERS.Reg)&^(0x8)|value<<3)
}
func (o *TIMG_Type) GetINT_ST_TIMERS_LACT_INT_ST() uint32 {
	return (volatile.LoadUint32(&o.INT_ST_TIMERS.Reg) & 0x8) >> 3
}

// TIMG.INT_CLR_TIMERS
func (o *TIMG_Type) SetINT_CLR_TIMERS_T0_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR_TIMERS.Reg, volatile.LoadUint32(&o.INT_CLR_TIMERS.Reg)&^(0x1)|value)
}
func (o *TIMG_Type) GetINT_CLR_TIMERS_T0_INT_CLR() uint32 {
	return volatile.LoadUint32(&o.INT_CLR_TIMERS.Reg) & 0x1
}
func (o *TIMG_Type) SetINT_CLR_TIMERS_T1_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR_TIMERS.Reg, volatile.LoadUint32(&o.INT_CLR_TIMERS.Reg)&^(0x2)|value<<1)
}
func (o *TIMG_Type) GetINT_CLR_TIMERS_T1_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR_TIMERS.Reg) & 0x2) >> 1
}
func (o *TIMG_Type) SetINT_CLR_TIMERS_WDT_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR_TIMERS.Reg, volatile.LoadUint32(&o.INT_CLR_TIMERS.Reg)&^(0x4)|value<<2)
}
func (o *TIMG_Type) GetINT_CLR_TIMERS_WDT_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR_TIMERS.Reg) & 0x4) >> 2
}
func (o *TIMG_Type) SetINT_CLR_TIMERS_LACT_INT_CLR(value uint32) {
	volatile.StoreUint32(&o.INT_CLR_TIMERS.Reg, volatile.LoadUint32(&o.INT_CLR_TIMERS.Reg)&^(0x8)|value<<3)
}
func (o *TIMG_Type) GetINT_CLR_TIMERS_LACT_INT_CLR() uint32 {
	return (volatile.LoadUint32(&o.INT_CLR_TIMERS.Reg) & 0x8) >> 3
}

// TIMG.NTIMERS_DATE
func (o *TIMG_Type) SetNTIMERS_DATE(value uint32) {
	volatile.StoreUint32(&o.NTIMERS_DATE.Reg, volatile.LoadUint32(&o.NTIMERS_DATE.Reg)&^(0xfffffff)|value)
}
func (o *TIMG_Type) GetNTIMERS_DATE() uint32 {
	return volatile.LoadUint32(&o.NTIMERS_DATE.Reg) & 0xfffffff
}

// TIMG.TIMGCLK
func (o *TIMG_Type) SetTIMGCLK_CLK_EN(value uint32) {
	volatile.StoreUint32(&o.TIMGCLK.Reg, volatile.LoadUint32(&o.TIMGCLK.Reg)&^(0x80000000)|value<<31)
}
func (o *TIMG_Type) GetTIMGCLK_CLK_EN() uint32 {
	return (volatile.LoadUint32(&o.TIMGCLK.Reg) & 0x80000000) >> 31
}
