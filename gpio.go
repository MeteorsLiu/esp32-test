package main

import "llgo-test/volatile"

// General Purpose Input/Output
type GPIO_Type struct {
	BT_SELECT          volatile.Register32 // 0x0
	OUT                volatile.Register32 // 0x4
	OUT_W1TS           volatile.Register32 // 0x8
	OUT_W1TC           volatile.Register32 // 0xC
	OUT1               volatile.Register32 // 0x10
	OUT1_W1TS          volatile.Register32 // 0x14
	OUT1_W1TC          volatile.Register32 // 0x18
	SDIO_SELECT        volatile.Register32 // 0x1C
	ENABLE             volatile.Register32 // 0x20
	ENABLE_W1TS        volatile.Register32 // 0x24
	ENABLE_W1TC        volatile.Register32 // 0x28
	ENABLE1            volatile.Register32 // 0x2C
	ENABLE1_W1TS       volatile.Register32 // 0x30
	ENABLE1_W1TC       volatile.Register32 // 0x34
	STRAP              volatile.Register32 // 0x38
	IN                 volatile.Register32 // 0x3C
	IN1                volatile.Register32 // 0x40
	STATUS             volatile.Register32 // 0x44
	STATUS_W1TS        volatile.Register32 // 0x48
	STATUS_W1TC        volatile.Register32 // 0x4C
	STATUS1            volatile.Register32 // 0x50
	STATUS1_W1TS       volatile.Register32 // 0x54
	STATUS1_W1TC       volatile.Register32 // 0x58
	_                  [4]byte
	ACPU_INT           volatile.Register32 // 0x60
	ACPU_NMI_INT       volatile.Register32 // 0x64
	PCPU_INT           volatile.Register32 // 0x68
	PCPU_NMI_INT       volatile.Register32 // 0x6C
	CPUSDIO_INT        volatile.Register32 // 0x70
	ACPU_INT1          volatile.Register32 // 0x74
	ACPU_NMI_INT1      volatile.Register32 // 0x78
	PCPU_INT1          volatile.Register32 // 0x7C
	PCPU_NMI_INT1      volatile.Register32 // 0x80
	CPUSDIO_INT1       volatile.Register32 // 0x84
	PIN0               volatile.Register32 // 0x88
	PIN1               volatile.Register32 // 0x8C
	PIN2               volatile.Register32 // 0x90
	PIN3               volatile.Register32 // 0x94
	PIN4               volatile.Register32 // 0x98
	PIN5               volatile.Register32 // 0x9C
	PIN6               volatile.Register32 // 0xA0
	PIN7               volatile.Register32 // 0xA4
	PIN8               volatile.Register32 // 0xA8
	PIN9               volatile.Register32 // 0xAC
	PIN10              volatile.Register32 // 0xB0
	PIN11              volatile.Register32 // 0xB4
	PIN12              volatile.Register32 // 0xB8
	PIN13              volatile.Register32 // 0xBC
	PIN14              volatile.Register32 // 0xC0
	PIN15              volatile.Register32 // 0xC4
	PIN16              volatile.Register32 // 0xC8
	PIN17              volatile.Register32 // 0xCC
	PIN18              volatile.Register32 // 0xD0
	PIN19              volatile.Register32 // 0xD4
	PIN20              volatile.Register32 // 0xD8
	PIN21              volatile.Register32 // 0xDC
	PIN22              volatile.Register32 // 0xE0
	PIN23              volatile.Register32 // 0xE4
	PIN24              volatile.Register32 // 0xE8
	PIN25              volatile.Register32 // 0xEC
	PIN26              volatile.Register32 // 0xF0
	PIN27              volatile.Register32 // 0xF4
	PIN28              volatile.Register32 // 0xF8
	PIN29              volatile.Register32 // 0xFC
	PIN30              volatile.Register32 // 0x100
	PIN31              volatile.Register32 // 0x104
	PIN32              volatile.Register32 // 0x108
	PIN33              volatile.Register32 // 0x10C
	PIN34              volatile.Register32 // 0x110
	PIN35              volatile.Register32 // 0x114
	PIN36              volatile.Register32 // 0x118
	PIN37              volatile.Register32 // 0x11C
	PIN38              volatile.Register32 // 0x120
	PIN39              volatile.Register32 // 0x124
	CALI_CONF          volatile.Register32 // 0x128
	CALI_DATA          volatile.Register32 // 0x12C
	FUNC0_IN_SEL_CFG   volatile.Register32 // 0x130
	FUNC1_IN_SEL_CFG   volatile.Register32 // 0x134
	FUNC2_IN_SEL_CFG   volatile.Register32 // 0x138
	FUNC3_IN_SEL_CFG   volatile.Register32 // 0x13C
	FUNC4_IN_SEL_CFG   volatile.Register32 // 0x140
	FUNC5_IN_SEL_CFG   volatile.Register32 // 0x144
	FUNC6_IN_SEL_CFG   volatile.Register32 // 0x148
	FUNC7_IN_SEL_CFG   volatile.Register32 // 0x14C
	FUNC8_IN_SEL_CFG   volatile.Register32 // 0x150
	FUNC9_IN_SEL_CFG   volatile.Register32 // 0x154
	FUNC10_IN_SEL_CFG  volatile.Register32 // 0x158
	FUNC11_IN_SEL_CFG  volatile.Register32 // 0x15C
	FUNC12_IN_SEL_CFG  volatile.Register32 // 0x160
	FUNC13_IN_SEL_CFG  volatile.Register32 // 0x164
	FUNC14_IN_SEL_CFG  volatile.Register32 // 0x168
	FUNC15_IN_SEL_CFG  volatile.Register32 // 0x16C
	FUNC16_IN_SEL_CFG  volatile.Register32 // 0x170
	FUNC17_IN_SEL_CFG  volatile.Register32 // 0x174
	FUNC18_IN_SEL_CFG  volatile.Register32 // 0x178
	FUNC19_IN_SEL_CFG  volatile.Register32 // 0x17C
	FUNC20_IN_SEL_CFG  volatile.Register32 // 0x180
	FUNC21_IN_SEL_CFG  volatile.Register32 // 0x184
	FUNC22_IN_SEL_CFG  volatile.Register32 // 0x188
	FUNC23_IN_SEL_CFG  volatile.Register32 // 0x18C
	FUNC24_IN_SEL_CFG  volatile.Register32 // 0x190
	FUNC25_IN_SEL_CFG  volatile.Register32 // 0x194
	FUNC26_IN_SEL_CFG  volatile.Register32 // 0x198
	FUNC27_IN_SEL_CFG  volatile.Register32 // 0x19C
	FUNC28_IN_SEL_CFG  volatile.Register32 // 0x1A0
	FUNC29_IN_SEL_CFG  volatile.Register32 // 0x1A4
	FUNC30_IN_SEL_CFG  volatile.Register32 // 0x1A8
	FUNC31_IN_SEL_CFG  volatile.Register32 // 0x1AC
	FUNC32_IN_SEL_CFG  volatile.Register32 // 0x1B0
	FUNC33_IN_SEL_CFG  volatile.Register32 // 0x1B4
	FUNC34_IN_SEL_CFG  volatile.Register32 // 0x1B8
	FUNC35_IN_SEL_CFG  volatile.Register32 // 0x1BC
	FUNC36_IN_SEL_CFG  volatile.Register32 // 0x1C0
	FUNC37_IN_SEL_CFG  volatile.Register32 // 0x1C4
	FUNC38_IN_SEL_CFG  volatile.Register32 // 0x1C8
	FUNC39_IN_SEL_CFG  volatile.Register32 // 0x1CC
	FUNC40_IN_SEL_CFG  volatile.Register32 // 0x1D0
	FUNC41_IN_SEL_CFG  volatile.Register32 // 0x1D4
	FUNC42_IN_SEL_CFG  volatile.Register32 // 0x1D8
	FUNC43_IN_SEL_CFG  volatile.Register32 // 0x1DC
	FUNC44_IN_SEL_CFG  volatile.Register32 // 0x1E0
	FUNC45_IN_SEL_CFG  volatile.Register32 // 0x1E4
	FUNC46_IN_SEL_CFG  volatile.Register32 // 0x1E8
	FUNC47_IN_SEL_CFG  volatile.Register32 // 0x1EC
	FUNC48_IN_SEL_CFG  volatile.Register32 // 0x1F0
	FUNC49_IN_SEL_CFG  volatile.Register32 // 0x1F4
	FUNC50_IN_SEL_CFG  volatile.Register32 // 0x1F8
	FUNC51_IN_SEL_CFG  volatile.Register32 // 0x1FC
	FUNC52_IN_SEL_CFG  volatile.Register32 // 0x200
	FUNC53_IN_SEL_CFG  volatile.Register32 // 0x204
	FUNC54_IN_SEL_CFG  volatile.Register32 // 0x208
	FUNC55_IN_SEL_CFG  volatile.Register32 // 0x20C
	FUNC56_IN_SEL_CFG  volatile.Register32 // 0x210
	FUNC57_IN_SEL_CFG  volatile.Register32 // 0x214
	FUNC58_IN_SEL_CFG  volatile.Register32 // 0x218
	FUNC59_IN_SEL_CFG  volatile.Register32 // 0x21C
	FUNC60_IN_SEL_CFG  volatile.Register32 // 0x220
	FUNC61_IN_SEL_CFG  volatile.Register32 // 0x224
	FUNC62_IN_SEL_CFG  volatile.Register32 // 0x228
	FUNC63_IN_SEL_CFG  volatile.Register32 // 0x22C
	FUNC64_IN_SEL_CFG  volatile.Register32 // 0x230
	FUNC65_IN_SEL_CFG  volatile.Register32 // 0x234
	FUNC66_IN_SEL_CFG  volatile.Register32 // 0x238
	FUNC67_IN_SEL_CFG  volatile.Register32 // 0x23C
	FUNC68_IN_SEL_CFG  volatile.Register32 // 0x240
	FUNC69_IN_SEL_CFG  volatile.Register32 // 0x244
	FUNC70_IN_SEL_CFG  volatile.Register32 // 0x248
	FUNC71_IN_SEL_CFG  volatile.Register32 // 0x24C
	FUNC72_IN_SEL_CFG  volatile.Register32 // 0x250
	FUNC73_IN_SEL_CFG  volatile.Register32 // 0x254
	FUNC74_IN_SEL_CFG  volatile.Register32 // 0x258
	FUNC75_IN_SEL_CFG  volatile.Register32 // 0x25C
	FUNC76_IN_SEL_CFG  volatile.Register32 // 0x260
	FUNC77_IN_SEL_CFG  volatile.Register32 // 0x264
	FUNC78_IN_SEL_CFG  volatile.Register32 // 0x268
	FUNC79_IN_SEL_CFG  volatile.Register32 // 0x26C
	FUNC80_IN_SEL_CFG  volatile.Register32 // 0x270
	FUNC81_IN_SEL_CFG  volatile.Register32 // 0x274
	FUNC82_IN_SEL_CFG  volatile.Register32 // 0x278
	FUNC83_IN_SEL_CFG  volatile.Register32 // 0x27C
	FUNC84_IN_SEL_CFG  volatile.Register32 // 0x280
	FUNC85_IN_SEL_CFG  volatile.Register32 // 0x284
	FUNC86_IN_SEL_CFG  volatile.Register32 // 0x288
	FUNC87_IN_SEL_CFG  volatile.Register32 // 0x28C
	FUNC88_IN_SEL_CFG  volatile.Register32 // 0x290
	FUNC89_IN_SEL_CFG  volatile.Register32 // 0x294
	FUNC90_IN_SEL_CFG  volatile.Register32 // 0x298
	FUNC91_IN_SEL_CFG  volatile.Register32 // 0x29C
	FUNC92_IN_SEL_CFG  volatile.Register32 // 0x2A0
	FUNC93_IN_SEL_CFG  volatile.Register32 // 0x2A4
	FUNC94_IN_SEL_CFG  volatile.Register32 // 0x2A8
	FUNC95_IN_SEL_CFG  volatile.Register32 // 0x2AC
	FUNC96_IN_SEL_CFG  volatile.Register32 // 0x2B0
	FUNC97_IN_SEL_CFG  volatile.Register32 // 0x2B4
	FUNC98_IN_SEL_CFG  volatile.Register32 // 0x2B8
	FUNC99_IN_SEL_CFG  volatile.Register32 // 0x2BC
	FUNC100_IN_SEL_CFG volatile.Register32 // 0x2C0
	FUNC101_IN_SEL_CFG volatile.Register32 // 0x2C4
	FUNC102_IN_SEL_CFG volatile.Register32 // 0x2C8
	FUNC103_IN_SEL_CFG volatile.Register32 // 0x2CC
	FUNC104_IN_SEL_CFG volatile.Register32 // 0x2D0
	FUNC105_IN_SEL_CFG volatile.Register32 // 0x2D4
	FUNC106_IN_SEL_CFG volatile.Register32 // 0x2D8
	FUNC107_IN_SEL_CFG volatile.Register32 // 0x2DC
	FUNC108_IN_SEL_CFG volatile.Register32 // 0x2E0
	FUNC109_IN_SEL_CFG volatile.Register32 // 0x2E4
	FUNC110_IN_SEL_CFG volatile.Register32 // 0x2E8
	FUNC111_IN_SEL_CFG volatile.Register32 // 0x2EC
	FUNC112_IN_SEL_CFG volatile.Register32 // 0x2F0
	FUNC113_IN_SEL_CFG volatile.Register32 // 0x2F4
	FUNC114_IN_SEL_CFG volatile.Register32 // 0x2F8
	FUNC115_IN_SEL_CFG volatile.Register32 // 0x2FC
	FUNC116_IN_SEL_CFG volatile.Register32 // 0x300
	FUNC117_IN_SEL_CFG volatile.Register32 // 0x304
	FUNC118_IN_SEL_CFG volatile.Register32 // 0x308
	FUNC119_IN_SEL_CFG volatile.Register32 // 0x30C
	FUNC120_IN_SEL_CFG volatile.Register32 // 0x310
	FUNC121_IN_SEL_CFG volatile.Register32 // 0x314
	FUNC122_IN_SEL_CFG volatile.Register32 // 0x318
	FUNC123_IN_SEL_CFG volatile.Register32 // 0x31C
	FUNC124_IN_SEL_CFG volatile.Register32 // 0x320
	FUNC125_IN_SEL_CFG volatile.Register32 // 0x324
	FUNC126_IN_SEL_CFG volatile.Register32 // 0x328
	FUNC127_IN_SEL_CFG volatile.Register32 // 0x32C
	FUNC128_IN_SEL_CFG volatile.Register32 // 0x330
	FUNC129_IN_SEL_CFG volatile.Register32 // 0x334
	FUNC130_IN_SEL_CFG volatile.Register32 // 0x338
	FUNC131_IN_SEL_CFG volatile.Register32 // 0x33C
	FUNC132_IN_SEL_CFG volatile.Register32 // 0x340
	FUNC133_IN_SEL_CFG volatile.Register32 // 0x344
	FUNC134_IN_SEL_CFG volatile.Register32 // 0x348
	FUNC135_IN_SEL_CFG volatile.Register32 // 0x34C
	FUNC136_IN_SEL_CFG volatile.Register32 // 0x350
	FUNC137_IN_SEL_CFG volatile.Register32 // 0x354
	FUNC138_IN_SEL_CFG volatile.Register32 // 0x358
	FUNC139_IN_SEL_CFG volatile.Register32 // 0x35C
	FUNC140_IN_SEL_CFG volatile.Register32 // 0x360
	FUNC141_IN_SEL_CFG volatile.Register32 // 0x364
	FUNC142_IN_SEL_CFG volatile.Register32 // 0x368
	FUNC143_IN_SEL_CFG volatile.Register32 // 0x36C
	FUNC144_IN_SEL_CFG volatile.Register32 // 0x370
	FUNC145_IN_SEL_CFG volatile.Register32 // 0x374
	FUNC146_IN_SEL_CFG volatile.Register32 // 0x378
	FUNC147_IN_SEL_CFG volatile.Register32 // 0x37C
	FUNC148_IN_SEL_CFG volatile.Register32 // 0x380
	FUNC149_IN_SEL_CFG volatile.Register32 // 0x384
	FUNC150_IN_SEL_CFG volatile.Register32 // 0x388
	FUNC151_IN_SEL_CFG volatile.Register32 // 0x38C
	FUNC152_IN_SEL_CFG volatile.Register32 // 0x390
	FUNC153_IN_SEL_CFG volatile.Register32 // 0x394
	FUNC154_IN_SEL_CFG volatile.Register32 // 0x398
	FUNC155_IN_SEL_CFG volatile.Register32 // 0x39C
	FUNC156_IN_SEL_CFG volatile.Register32 // 0x3A0
	FUNC157_IN_SEL_CFG volatile.Register32 // 0x3A4
	FUNC158_IN_SEL_CFG volatile.Register32 // 0x3A8
	FUNC159_IN_SEL_CFG volatile.Register32 // 0x3AC
	FUNC160_IN_SEL_CFG volatile.Register32 // 0x3B0
	FUNC161_IN_SEL_CFG volatile.Register32 // 0x3B4
	FUNC162_IN_SEL_CFG volatile.Register32 // 0x3B8
	FUNC163_IN_SEL_CFG volatile.Register32 // 0x3BC
	FUNC164_IN_SEL_CFG volatile.Register32 // 0x3C0
	FUNC165_IN_SEL_CFG volatile.Register32 // 0x3C4
	FUNC166_IN_SEL_CFG volatile.Register32 // 0x3C8
	FUNC167_IN_SEL_CFG volatile.Register32 // 0x3CC
	FUNC168_IN_SEL_CFG volatile.Register32 // 0x3D0
	FUNC169_IN_SEL_CFG volatile.Register32 // 0x3D4
	FUNC170_IN_SEL_CFG volatile.Register32 // 0x3D8
	FUNC171_IN_SEL_CFG volatile.Register32 // 0x3DC
	FUNC172_IN_SEL_CFG volatile.Register32 // 0x3E0
	FUNC173_IN_SEL_CFG volatile.Register32 // 0x3E4
	FUNC174_IN_SEL_CFG volatile.Register32 // 0x3E8
	FUNC175_IN_SEL_CFG volatile.Register32 // 0x3EC
	FUNC176_IN_SEL_CFG volatile.Register32 // 0x3F0
	FUNC177_IN_SEL_CFG volatile.Register32 // 0x3F4
	FUNC178_IN_SEL_CFG volatile.Register32 // 0x3F8
	FUNC179_IN_SEL_CFG volatile.Register32 // 0x3FC
	FUNC180_IN_SEL_CFG volatile.Register32 // 0x400
	FUNC181_IN_SEL_CFG volatile.Register32 // 0x404
	FUNC182_IN_SEL_CFG volatile.Register32 // 0x408
	FUNC183_IN_SEL_CFG volatile.Register32 // 0x40C
	FUNC184_IN_SEL_CFG volatile.Register32 // 0x410
	FUNC185_IN_SEL_CFG volatile.Register32 // 0x414
	FUNC186_IN_SEL_CFG volatile.Register32 // 0x418
	FUNC187_IN_SEL_CFG volatile.Register32 // 0x41C
	FUNC188_IN_SEL_CFG volatile.Register32 // 0x420
	FUNC189_IN_SEL_CFG volatile.Register32 // 0x424
	FUNC190_IN_SEL_CFG volatile.Register32 // 0x428
	FUNC191_IN_SEL_CFG volatile.Register32 // 0x42C
	FUNC192_IN_SEL_CFG volatile.Register32 // 0x430
	FUNC193_IN_SEL_CFG volatile.Register32 // 0x434
	FUNC194_IN_SEL_CFG volatile.Register32 // 0x438
	FUNC195_IN_SEL_CFG volatile.Register32 // 0x43C
	FUNC196_IN_SEL_CFG volatile.Register32 // 0x440
	FUNC197_IN_SEL_CFG volatile.Register32 // 0x444
	FUNC198_IN_SEL_CFG volatile.Register32 // 0x448
	FUNC199_IN_SEL_CFG volatile.Register32 // 0x44C
	FUNC200_IN_SEL_CFG volatile.Register32 // 0x450
	FUNC201_IN_SEL_CFG volatile.Register32 // 0x454
	FUNC202_IN_SEL_CFG volatile.Register32 // 0x458
	FUNC203_IN_SEL_CFG volatile.Register32 // 0x45C
	FUNC204_IN_SEL_CFG volatile.Register32 // 0x460
	FUNC205_IN_SEL_CFG volatile.Register32 // 0x464
	FUNC206_IN_SEL_CFG volatile.Register32 // 0x468
	FUNC207_IN_SEL_CFG volatile.Register32 // 0x46C
	FUNC208_IN_SEL_CFG volatile.Register32 // 0x470
	FUNC209_IN_SEL_CFG volatile.Register32 // 0x474
	FUNC210_IN_SEL_CFG volatile.Register32 // 0x478
	FUNC211_IN_SEL_CFG volatile.Register32 // 0x47C
	FUNC212_IN_SEL_CFG volatile.Register32 // 0x480
	FUNC213_IN_SEL_CFG volatile.Register32 // 0x484
	FUNC214_IN_SEL_CFG volatile.Register32 // 0x488
	FUNC215_IN_SEL_CFG volatile.Register32 // 0x48C
	FUNC216_IN_SEL_CFG volatile.Register32 // 0x490
	FUNC217_IN_SEL_CFG volatile.Register32 // 0x494
	FUNC218_IN_SEL_CFG volatile.Register32 // 0x498
	FUNC219_IN_SEL_CFG volatile.Register32 // 0x49C
	FUNC220_IN_SEL_CFG volatile.Register32 // 0x4A0
	FUNC221_IN_SEL_CFG volatile.Register32 // 0x4A4
	FUNC222_IN_SEL_CFG volatile.Register32 // 0x4A8
	FUNC223_IN_SEL_CFG volatile.Register32 // 0x4AC
	FUNC224_IN_SEL_CFG volatile.Register32 // 0x4B0
	FUNC225_IN_SEL_CFG volatile.Register32 // 0x4B4
	FUNC226_IN_SEL_CFG volatile.Register32 // 0x4B8
	FUNC227_IN_SEL_CFG volatile.Register32 // 0x4BC
	FUNC228_IN_SEL_CFG volatile.Register32 // 0x4C0
	FUNC229_IN_SEL_CFG volatile.Register32 // 0x4C4
	FUNC230_IN_SEL_CFG volatile.Register32 // 0x4C8
	FUNC231_IN_SEL_CFG volatile.Register32 // 0x4CC
	FUNC232_IN_SEL_CFG volatile.Register32 // 0x4D0
	FUNC233_IN_SEL_CFG volatile.Register32 // 0x4D4
	FUNC234_IN_SEL_CFG volatile.Register32 // 0x4D8
	FUNC235_IN_SEL_CFG volatile.Register32 // 0x4DC
	FUNC236_IN_SEL_CFG volatile.Register32 // 0x4E0
	FUNC237_IN_SEL_CFG volatile.Register32 // 0x4E4
	FUNC238_IN_SEL_CFG volatile.Register32 // 0x4E8
	FUNC239_IN_SEL_CFG volatile.Register32 // 0x4EC
	FUNC240_IN_SEL_CFG volatile.Register32 // 0x4F0
	FUNC241_IN_SEL_CFG volatile.Register32 // 0x4F4
	FUNC242_IN_SEL_CFG volatile.Register32 // 0x4F8
	FUNC243_IN_SEL_CFG volatile.Register32 // 0x4FC
	FUNC244_IN_SEL_CFG volatile.Register32 // 0x500
	FUNC245_IN_SEL_CFG volatile.Register32 // 0x504
	FUNC246_IN_SEL_CFG volatile.Register32 // 0x508
	FUNC247_IN_SEL_CFG volatile.Register32 // 0x50C
	FUNC248_IN_SEL_CFG volatile.Register32 // 0x510
	FUNC249_IN_SEL_CFG volatile.Register32 // 0x514
	FUNC250_IN_SEL_CFG volatile.Register32 // 0x518
	FUNC251_IN_SEL_CFG volatile.Register32 // 0x51C
	FUNC252_IN_SEL_CFG volatile.Register32 // 0x520
	FUNC253_IN_SEL_CFG volatile.Register32 // 0x524
	FUNC254_IN_SEL_CFG volatile.Register32 // 0x528
	FUNC255_IN_SEL_CFG volatile.Register32 // 0x52C
	FUNC0_OUT_SEL_CFG  volatile.Register32 // 0x530
	FUNC1_OUT_SEL_CFG  volatile.Register32 // 0x534
	FUNC2_OUT_SEL_CFG  volatile.Register32 // 0x538
	FUNC3_OUT_SEL_CFG  volatile.Register32 // 0x53C
	FUNC4_OUT_SEL_CFG  volatile.Register32 // 0x540
	FUNC5_OUT_SEL_CFG  volatile.Register32 // 0x544
	FUNC6_OUT_SEL_CFG  volatile.Register32 // 0x548
	FUNC7_OUT_SEL_CFG  volatile.Register32 // 0x54C
	FUNC8_OUT_SEL_CFG  volatile.Register32 // 0x550
	FUNC9_OUT_SEL_CFG  volatile.Register32 // 0x554
	FUNC10_OUT_SEL_CFG volatile.Register32 // 0x558
	FUNC11_OUT_SEL_CFG volatile.Register32 // 0x55C
	FUNC12_OUT_SEL_CFG volatile.Register32 // 0x560
	FUNC13_OUT_SEL_CFG volatile.Register32 // 0x564
	FUNC14_OUT_SEL_CFG volatile.Register32 // 0x568
	FUNC15_OUT_SEL_CFG volatile.Register32 // 0x56C
	FUNC16_OUT_SEL_CFG volatile.Register32 // 0x570
	FUNC17_OUT_SEL_CFG volatile.Register32 // 0x574
	FUNC18_OUT_SEL_CFG volatile.Register32 // 0x578
	FUNC19_OUT_SEL_CFG volatile.Register32 // 0x57C
	FUNC20_OUT_SEL_CFG volatile.Register32 // 0x580
	FUNC21_OUT_SEL_CFG volatile.Register32 // 0x584
	FUNC22_OUT_SEL_CFG volatile.Register32 // 0x588
	FUNC23_OUT_SEL_CFG volatile.Register32 // 0x58C
	FUNC24_OUT_SEL_CFG volatile.Register32 // 0x590
	FUNC25_OUT_SEL_CFG volatile.Register32 // 0x594
	FUNC26_OUT_SEL_CFG volatile.Register32 // 0x598
	FUNC27_OUT_SEL_CFG volatile.Register32 // 0x59C
	FUNC28_OUT_SEL_CFG volatile.Register32 // 0x5A0
	FUNC29_OUT_SEL_CFG volatile.Register32 // 0x5A4
	FUNC30_OUT_SEL_CFG volatile.Register32 // 0x5A8
	FUNC31_OUT_SEL_CFG volatile.Register32 // 0x5AC
	FUNC32_OUT_SEL_CFG volatile.Register32 // 0x5B0
	FUNC33_OUT_SEL_CFG volatile.Register32 // 0x5B4
	FUNC34_OUT_SEL_CFG volatile.Register32 // 0x5B8
	FUNC35_OUT_SEL_CFG volatile.Register32 // 0x5BC
	FUNC36_OUT_SEL_CFG volatile.Register32 // 0x5C0
	FUNC37_OUT_SEL_CFG volatile.Register32 // 0x5C4
	FUNC38_OUT_SEL_CFG volatile.Register32 // 0x5C8
	FUNC39_OUT_SEL_CFG volatile.Register32 // 0x5CC
}

// GPIO.BT_SELECT
func (o *GPIO_Type) SetBT_SELECT(value uint32) {
	volatile.StoreUint32(&o.BT_SELECT.Reg, value)
}
func (o *GPIO_Type) GetBT_SELECT() uint32 {
	return volatile.LoadUint32(&o.BT_SELECT.Reg)
}

// GPIO.OUT
func (o *GPIO_Type) SetOUT(value uint32) {
	volatile.StoreUint32(&o.OUT.Reg, value)
}
func (o *GPIO_Type) GetOUT() uint32 {
	return volatile.LoadUint32(&o.OUT.Reg)
}

// GPIO.OUT_W1TS
func (o *GPIO_Type) SetOUT_W1TS(value uint32) {
	volatile.StoreUint32(&o.OUT_W1TS.Reg, value)
}
func (o *GPIO_Type) GetOUT_W1TS() uint32 {
	return volatile.LoadUint32(&o.OUT_W1TS.Reg)
}

// GPIO.OUT_W1TC
func (o *GPIO_Type) SetOUT_W1TC(value uint32) {
	volatile.StoreUint32(&o.OUT_W1TC.Reg, value)
}
func (o *GPIO_Type) GetOUT_W1TC() uint32 {
	return volatile.LoadUint32(&o.OUT_W1TC.Reg)
}

// GPIO.OUT1
func (o *GPIO_Type) SetOUT1_DATA(value uint32) {
	volatile.StoreUint32(&o.OUT1.Reg, volatile.LoadUint32(&o.OUT1.Reg)&^(0xff)|value)
}
func (o *GPIO_Type) GetOUT1_DATA() uint32 {
	return volatile.LoadUint32(&o.OUT1.Reg) & 0xff
}

// GPIO.OUT1_W1TS
func (o *GPIO_Type) SetOUT1_W1TS_OUT1_DATA_W1TS(value uint32) {
	volatile.StoreUint32(&o.OUT1_W1TS.Reg, volatile.LoadUint32(&o.OUT1_W1TS.Reg)&^(0xff)|value)
}
func (o *GPIO_Type) GetOUT1_W1TS_OUT1_DATA_W1TS() uint32 {
	return volatile.LoadUint32(&o.OUT1_W1TS.Reg) & 0xff
}

// GPIO.OUT1_W1TC
func (o *GPIO_Type) SetOUT1_W1TC_OUT1_DATA_W1TC(value uint32) {
	volatile.StoreUint32(&o.OUT1_W1TC.Reg, volatile.LoadUint32(&o.OUT1_W1TC.Reg)&^(0xff)|value)
}
func (o *GPIO_Type) GetOUT1_W1TC_OUT1_DATA_W1TC() uint32 {
	return volatile.LoadUint32(&o.OUT1_W1TC.Reg) & 0xff
}

// GPIO.SDIO_SELECT
func (o *GPIO_Type) SetSDIO_SELECT_SDIO_SEL(value uint32) {
	volatile.StoreUint32(&o.SDIO_SELECT.Reg, volatile.LoadUint32(&o.SDIO_SELECT.Reg)&^(0xff)|value)
}
func (o *GPIO_Type) GetSDIO_SELECT_SDIO_SEL() uint32 {
	return volatile.LoadUint32(&o.SDIO_SELECT.Reg) & 0xff
}

// GPIO.ENABLE
func (o *GPIO_Type) SetENABLE(value uint32) {
	volatile.StoreUint32(&o.ENABLE.Reg, value)
}
func (o *GPIO_Type) GetENABLE() uint32 {
	return volatile.LoadUint32(&o.ENABLE.Reg)
}

// GPIO.ENABLE_W1TS
func (o *GPIO_Type) SetENABLE_W1TS(value uint32) {
	volatile.StoreUint32(&o.ENABLE_W1TS.Reg, value)
}
func (o *GPIO_Type) GetENABLE_W1TS() uint32 {
	return volatile.LoadUint32(&o.ENABLE_W1TS.Reg)
}

// GPIO.ENABLE_W1TC
func (o *GPIO_Type) SetENABLE_W1TC(value uint32) {
	volatile.StoreUint32(&o.ENABLE_W1TC.Reg, value)
}
func (o *GPIO_Type) GetENABLE_W1TC() uint32 {
	return volatile.LoadUint32(&o.ENABLE_W1TC.Reg)
}

// GPIO.ENABLE1
func (o *GPIO_Type) SetENABLE1_DATA(value uint32) {
	volatile.StoreUint32(&o.ENABLE1.Reg, volatile.LoadUint32(&o.ENABLE1.Reg)&^(0xff)|value)
}
func (o *GPIO_Type) GetENABLE1_DATA() uint32 {
	return volatile.LoadUint32(&o.ENABLE1.Reg) & 0xff
}

// GPIO.ENABLE1_W1TS
func (o *GPIO_Type) SetENABLE1_W1TS_ENABLE1_DATA_W1TS(value uint32) {
	volatile.StoreUint32(&o.ENABLE1_W1TS.Reg, volatile.LoadUint32(&o.ENABLE1_W1TS.Reg)&^(0xff)|value)
}
func (o *GPIO_Type) GetENABLE1_W1TS_ENABLE1_DATA_W1TS() uint32 {
	return volatile.LoadUint32(&o.ENABLE1_W1TS.Reg) & 0xff
}

// GPIO.ENABLE1_W1TC
func (o *GPIO_Type) SetENABLE1_W1TC_ENABLE1_DATA_W1TC(value uint32) {
	volatile.StoreUint32(&o.ENABLE1_W1TC.Reg, volatile.LoadUint32(&o.ENABLE1_W1TC.Reg)&^(0xff)|value)
}
func (o *GPIO_Type) GetENABLE1_W1TC_ENABLE1_DATA_W1TC() uint32 {
	return volatile.LoadUint32(&o.ENABLE1_W1TC.Reg) & 0xff
}

// GPIO.STRAP
func (o *GPIO_Type) SetSTRAP_STRAPPING(value uint32) {
	volatile.StoreUint32(&o.STRAP.Reg, volatile.LoadUint32(&o.STRAP.Reg)&^(0xffff)|value)
}
func (o *GPIO_Type) GetSTRAP_STRAPPING() uint32 {
	return volatile.LoadUint32(&o.STRAP.Reg) & 0xffff
}

// GPIO.IN
func (o *GPIO_Type) SetIN(value uint32) {
	volatile.StoreUint32(&o.IN.Reg, value)
}
func (o *GPIO_Type) GetIN() uint32 {
	return volatile.LoadUint32(&o.IN.Reg)
}

// GPIO.IN1
func (o *GPIO_Type) SetIN1_DATA_NEXT(value uint32) {
	volatile.StoreUint32(&o.IN1.Reg, volatile.LoadUint32(&o.IN1.Reg)&^(0xff)|value)
}
func (o *GPIO_Type) GetIN1_DATA_NEXT() uint32 {
	return volatile.LoadUint32(&o.IN1.Reg) & 0xff
}

// GPIO.STATUS
func (o *GPIO_Type) SetSTATUS(value uint32) {
	volatile.StoreUint32(&o.STATUS.Reg, value)
}
func (o *GPIO_Type) GetSTATUS() uint32 {
	return volatile.LoadUint32(&o.STATUS.Reg)
}

// GPIO.STATUS_W1TS
func (o *GPIO_Type) SetSTATUS_W1TS(value uint32) {
	volatile.StoreUint32(&o.STATUS_W1TS.Reg, value)
}
func (o *GPIO_Type) GetSTATUS_W1TS() uint32 {
	return volatile.LoadUint32(&o.STATUS_W1TS.Reg)
}

// GPIO.STATUS_W1TC
func (o *GPIO_Type) SetSTATUS_W1TC(value uint32) {
	volatile.StoreUint32(&o.STATUS_W1TC.Reg, value)
}
func (o *GPIO_Type) GetSTATUS_W1TC() uint32 {
	return volatile.LoadUint32(&o.STATUS_W1TC.Reg)
}

// GPIO.STATUS1
func (o *GPIO_Type) SetSTATUS1_INT(value uint32) {
	volatile.StoreUint32(&o.STATUS1.Reg, volatile.LoadUint32(&o.STATUS1.Reg)&^(0xff)|value)
}
func (o *GPIO_Type) GetSTATUS1_INT() uint32 {
	return volatile.LoadUint32(&o.STATUS1.Reg) & 0xff
}

// GPIO.STATUS1_W1TS
func (o *GPIO_Type) SetSTATUS1_W1TS_STATUS1_INT_W1TS(value uint32) {
	volatile.StoreUint32(&o.STATUS1_W1TS.Reg, volatile.LoadUint32(&o.STATUS1_W1TS.Reg)&^(0xff)|value)
}
func (o *GPIO_Type) GetSTATUS1_W1TS_STATUS1_INT_W1TS() uint32 {
	return volatile.LoadUint32(&o.STATUS1_W1TS.Reg) & 0xff
}

// GPIO.STATUS1_W1TC
func (o *GPIO_Type) SetSTATUS1_W1TC_STATUS1_INT_W1TC(value uint32) {
	volatile.StoreUint32(&o.STATUS1_W1TC.Reg, volatile.LoadUint32(&o.STATUS1_W1TC.Reg)&^(0xff)|value)
}
func (o *GPIO_Type) GetSTATUS1_W1TC_STATUS1_INT_W1TC() uint32 {
	return volatile.LoadUint32(&o.STATUS1_W1TC.Reg) & 0xff
}

// GPIO.ACPU_INT
func (o *GPIO_Type) SetACPU_INT(value uint32) {
	volatile.StoreUint32(&o.ACPU_INT.Reg, value)
}
func (o *GPIO_Type) GetACPU_INT() uint32 {
	return volatile.LoadUint32(&o.ACPU_INT.Reg)
}

// GPIO.ACPU_NMI_INT
func (o *GPIO_Type) SetACPU_NMI_INT(value uint32) {
	volatile.StoreUint32(&o.ACPU_NMI_INT.Reg, value)
}
func (o *GPIO_Type) GetACPU_NMI_INT() uint32 {
	return volatile.LoadUint32(&o.ACPU_NMI_INT.Reg)
}

// GPIO.PCPU_INT
func (o *GPIO_Type) SetPCPU_INT(value uint32) {
	volatile.StoreUint32(&o.PCPU_INT.Reg, value)
}
func (o *GPIO_Type) GetPCPU_INT() uint32 {
	return volatile.LoadUint32(&o.PCPU_INT.Reg)
}

// GPIO.PCPU_NMI_INT
func (o *GPIO_Type) SetPCPU_NMI_INT(value uint32) {
	volatile.StoreUint32(&o.PCPU_NMI_INT.Reg, value)
}
func (o *GPIO_Type) GetPCPU_NMI_INT() uint32 {
	return volatile.LoadUint32(&o.PCPU_NMI_INT.Reg)
}

// GPIO.CPUSDIO_INT
func (o *GPIO_Type) SetCPUSDIO_INT(value uint32) {
	volatile.StoreUint32(&o.CPUSDIO_INT.Reg, value)
}
func (o *GPIO_Type) GetCPUSDIO_INT() uint32 {
	return volatile.LoadUint32(&o.CPUSDIO_INT.Reg)
}

// GPIO.ACPU_INT1
func (o *GPIO_Type) SetACPU_INT1_APPCPU_INT_H(value uint32) {
	volatile.StoreUint32(&o.ACPU_INT1.Reg, volatile.LoadUint32(&o.ACPU_INT1.Reg)&^(0xff)|value)
}
func (o *GPIO_Type) GetACPU_INT1_APPCPU_INT_H() uint32 {
	return volatile.LoadUint32(&o.ACPU_INT1.Reg) & 0xff
}

// GPIO.ACPU_NMI_INT1
func (o *GPIO_Type) SetACPU_NMI_INT1_APPCPU_NMI_INT_H(value uint32) {
	volatile.StoreUint32(&o.ACPU_NMI_INT1.Reg, volatile.LoadUint32(&o.ACPU_NMI_INT1.Reg)&^(0xff)|value)
}
func (o *GPIO_Type) GetACPU_NMI_INT1_APPCPU_NMI_INT_H() uint32 {
	return volatile.LoadUint32(&o.ACPU_NMI_INT1.Reg) & 0xff
}

// GPIO.PCPU_INT1
func (o *GPIO_Type) SetPCPU_INT1_PROCPU_INT_H(value uint32) {
	volatile.StoreUint32(&o.PCPU_INT1.Reg, volatile.LoadUint32(&o.PCPU_INT1.Reg)&^(0xff)|value)
}
func (o *GPIO_Type) GetPCPU_INT1_PROCPU_INT_H() uint32 {
	return volatile.LoadUint32(&o.PCPU_INT1.Reg) & 0xff
}

// GPIO.PCPU_NMI_INT1
func (o *GPIO_Type) SetPCPU_NMI_INT1_PROCPU_NMI_INT_H(value uint32) {
	volatile.StoreUint32(&o.PCPU_NMI_INT1.Reg, volatile.LoadUint32(&o.PCPU_NMI_INT1.Reg)&^(0xff)|value)
}
func (o *GPIO_Type) GetPCPU_NMI_INT1_PROCPU_NMI_INT_H() uint32 {
	return volatile.LoadUint32(&o.PCPU_NMI_INT1.Reg) & 0xff
}

// GPIO.CPUSDIO_INT1
func (o *GPIO_Type) SetCPUSDIO_INT1_SDIO_INT_H(value uint32) {
	volatile.StoreUint32(&o.CPUSDIO_INT1.Reg, volatile.LoadUint32(&o.CPUSDIO_INT1.Reg)&^(0xff)|value)
}
func (o *GPIO_Type) GetCPUSDIO_INT1_SDIO_INT_H() uint32 {
	return volatile.LoadUint32(&o.CPUSDIO_INT1.Reg) & 0xff
}
func (o *GPIO_Type) SetCPUSDIO_INT1_PIN_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.CPUSDIO_INT1.Reg, volatile.LoadUint32(&o.CPUSDIO_INT1.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetCPUSDIO_INT1_PIN_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.CPUSDIO_INT1.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetCPUSDIO_INT1_PIN_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.CPUSDIO_INT1.Reg, volatile.LoadUint32(&o.CPUSDIO_INT1.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetCPUSDIO_INT1_PIN_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.CPUSDIO_INT1.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetCPUSDIO_INT1_PIN_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.CPUSDIO_INT1.Reg, volatile.LoadUint32(&o.CPUSDIO_INT1.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetCPUSDIO_INT1_PIN_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.CPUSDIO_INT1.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetCPUSDIO_INT1_PIN_CONFIG(value uint32) {
	volatile.StoreUint32(&o.CPUSDIO_INT1.Reg, volatile.LoadUint32(&o.CPUSDIO_INT1.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetCPUSDIO_INT1_PIN_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.CPUSDIO_INT1.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetCPUSDIO_INT1_PIN_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.CPUSDIO_INT1.Reg, volatile.LoadUint32(&o.CPUSDIO_INT1.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetCPUSDIO_INT1_PIN_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.CPUSDIO_INT1.Reg) & 0x3e000) >> 13
}

// GPIO.PIN0
func (o *GPIO_Type) SetPIN0_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN0.Reg, volatile.LoadUint32(&o.PIN0.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN0_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN0.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN0_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN0.Reg, volatile.LoadUint32(&o.PIN0.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN0_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN0.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN0_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN0.Reg, volatile.LoadUint32(&o.PIN0.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN0_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN0.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN0_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN0.Reg, volatile.LoadUint32(&o.PIN0.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN0_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN0.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN0_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN0.Reg, volatile.LoadUint32(&o.PIN0.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN0_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN0.Reg) & 0x3e000) >> 13
}

// GPIO.PIN1
func (o *GPIO_Type) SetPIN1_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN1.Reg, volatile.LoadUint32(&o.PIN1.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN1_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN1.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN1_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN1.Reg, volatile.LoadUint32(&o.PIN1.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN1_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN1.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN1_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN1.Reg, volatile.LoadUint32(&o.PIN1.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN1_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN1.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN1_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN1.Reg, volatile.LoadUint32(&o.PIN1.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN1_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN1.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN1_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN1.Reg, volatile.LoadUint32(&o.PIN1.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN1_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN1.Reg) & 0x3e000) >> 13
}

// GPIO.PIN2
func (o *GPIO_Type) SetPIN2_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN2.Reg, volatile.LoadUint32(&o.PIN2.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN2_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN2.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN2_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN2.Reg, volatile.LoadUint32(&o.PIN2.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN2_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN2.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN2_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN2.Reg, volatile.LoadUint32(&o.PIN2.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN2_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN2.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN2_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN2.Reg, volatile.LoadUint32(&o.PIN2.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN2_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN2.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN2_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN2.Reg, volatile.LoadUint32(&o.PIN2.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN2_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN2.Reg) & 0x3e000) >> 13
}

// GPIO.PIN3
func (o *GPIO_Type) SetPIN3_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN3.Reg, volatile.LoadUint32(&o.PIN3.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN3_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN3.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN3_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN3.Reg, volatile.LoadUint32(&o.PIN3.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN3_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN3.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN3_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN3.Reg, volatile.LoadUint32(&o.PIN3.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN3_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN3.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN3_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN3.Reg, volatile.LoadUint32(&o.PIN3.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN3_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN3.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN3_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN3.Reg, volatile.LoadUint32(&o.PIN3.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN3_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN3.Reg) & 0x3e000) >> 13
}

// GPIO.PIN4
func (o *GPIO_Type) SetPIN4_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN4.Reg, volatile.LoadUint32(&o.PIN4.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN4_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN4.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN4_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN4.Reg, volatile.LoadUint32(&o.PIN4.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN4_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN4.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN4_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN4.Reg, volatile.LoadUint32(&o.PIN4.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN4_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN4.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN4_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN4.Reg, volatile.LoadUint32(&o.PIN4.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN4_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN4.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN4_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN4.Reg, volatile.LoadUint32(&o.PIN4.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN4_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN4.Reg) & 0x3e000) >> 13
}

// GPIO.PIN5
func (o *GPIO_Type) SetPIN5_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN5.Reg, volatile.LoadUint32(&o.PIN5.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN5_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN5.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN5_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN5.Reg, volatile.LoadUint32(&o.PIN5.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN5_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN5.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN5_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN5.Reg, volatile.LoadUint32(&o.PIN5.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN5_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN5.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN5_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN5.Reg, volatile.LoadUint32(&o.PIN5.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN5_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN5.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN5_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN5.Reg, volatile.LoadUint32(&o.PIN5.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN5_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN5.Reg) & 0x3e000) >> 13
}

// GPIO.PIN6
func (o *GPIO_Type) SetPIN6_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN6.Reg, volatile.LoadUint32(&o.PIN6.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN6_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN6.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN6_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN6.Reg, volatile.LoadUint32(&o.PIN6.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN6_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN6.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN6_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN6.Reg, volatile.LoadUint32(&o.PIN6.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN6_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN6.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN6_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN6.Reg, volatile.LoadUint32(&o.PIN6.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN6_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN6.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN6_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN6.Reg, volatile.LoadUint32(&o.PIN6.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN6_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN6.Reg) & 0x3e000) >> 13
}

// GPIO.PIN7
func (o *GPIO_Type) SetPIN7_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN7.Reg, volatile.LoadUint32(&o.PIN7.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN7_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN7.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN7_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN7.Reg, volatile.LoadUint32(&o.PIN7.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN7_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN7.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN7_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN7.Reg, volatile.LoadUint32(&o.PIN7.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN7_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN7.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN7_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN7.Reg, volatile.LoadUint32(&o.PIN7.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN7_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN7.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN7_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN7.Reg, volatile.LoadUint32(&o.PIN7.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN7_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN7.Reg) & 0x3e000) >> 13
}

// GPIO.PIN8
func (o *GPIO_Type) SetPIN8_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN8.Reg, volatile.LoadUint32(&o.PIN8.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN8_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN8.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN8_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN8.Reg, volatile.LoadUint32(&o.PIN8.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN8_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN8.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN8_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN8.Reg, volatile.LoadUint32(&o.PIN8.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN8_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN8.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN8_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN8.Reg, volatile.LoadUint32(&o.PIN8.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN8_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN8.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN8_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN8.Reg, volatile.LoadUint32(&o.PIN8.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN8_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN8.Reg) & 0x3e000) >> 13
}

// GPIO.PIN9
func (o *GPIO_Type) SetPIN9_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN9.Reg, volatile.LoadUint32(&o.PIN9.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN9_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN9.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN9_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN9.Reg, volatile.LoadUint32(&o.PIN9.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN9_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN9.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN9_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN9.Reg, volatile.LoadUint32(&o.PIN9.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN9_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN9.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN9_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN9.Reg, volatile.LoadUint32(&o.PIN9.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN9_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN9.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN9_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN9.Reg, volatile.LoadUint32(&o.PIN9.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN9_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN9.Reg) & 0x3e000) >> 13
}

// GPIO.PIN10
func (o *GPIO_Type) SetPIN10_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN10.Reg, volatile.LoadUint32(&o.PIN10.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN10_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN10.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN10_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN10.Reg, volatile.LoadUint32(&o.PIN10.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN10_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN10.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN10_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN10.Reg, volatile.LoadUint32(&o.PIN10.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN10_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN10.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN10_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN10.Reg, volatile.LoadUint32(&o.PIN10.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN10_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN10.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN10_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN10.Reg, volatile.LoadUint32(&o.PIN10.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN10_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN10.Reg) & 0x3e000) >> 13
}

// GPIO.PIN11
func (o *GPIO_Type) SetPIN11_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN11.Reg, volatile.LoadUint32(&o.PIN11.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN11_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN11.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN11_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN11.Reg, volatile.LoadUint32(&o.PIN11.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN11_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN11.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN11_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN11.Reg, volatile.LoadUint32(&o.PIN11.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN11_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN11.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN11_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN11.Reg, volatile.LoadUint32(&o.PIN11.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN11_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN11.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN11_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN11.Reg, volatile.LoadUint32(&o.PIN11.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN11_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN11.Reg) & 0x3e000) >> 13
}

// GPIO.PIN12
func (o *GPIO_Type) SetPIN12_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN12.Reg, volatile.LoadUint32(&o.PIN12.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN12_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN12.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN12_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN12.Reg, volatile.LoadUint32(&o.PIN12.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN12_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN12.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN12_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN12.Reg, volatile.LoadUint32(&o.PIN12.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN12_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN12.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN12_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN12.Reg, volatile.LoadUint32(&o.PIN12.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN12_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN12.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN12_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN12.Reg, volatile.LoadUint32(&o.PIN12.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN12_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN12.Reg) & 0x3e000) >> 13
}

// GPIO.PIN13
func (o *GPIO_Type) SetPIN13_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN13.Reg, volatile.LoadUint32(&o.PIN13.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN13_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN13.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN13_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN13.Reg, volatile.LoadUint32(&o.PIN13.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN13_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN13.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN13_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN13.Reg, volatile.LoadUint32(&o.PIN13.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN13_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN13.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN13_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN13.Reg, volatile.LoadUint32(&o.PIN13.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN13_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN13.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN13_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN13.Reg, volatile.LoadUint32(&o.PIN13.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN13_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN13.Reg) & 0x3e000) >> 13
}

// GPIO.PIN14
func (o *GPIO_Type) SetPIN14_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN14.Reg, volatile.LoadUint32(&o.PIN14.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN14_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN14.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN14_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN14.Reg, volatile.LoadUint32(&o.PIN14.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN14_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN14.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN14_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN14.Reg, volatile.LoadUint32(&o.PIN14.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN14_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN14.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN14_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN14.Reg, volatile.LoadUint32(&o.PIN14.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN14_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN14.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN14_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN14.Reg, volatile.LoadUint32(&o.PIN14.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN14_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN14.Reg) & 0x3e000) >> 13
}

// GPIO.PIN15
func (o *GPIO_Type) SetPIN15_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN15.Reg, volatile.LoadUint32(&o.PIN15.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN15_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN15.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN15_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN15.Reg, volatile.LoadUint32(&o.PIN15.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN15_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN15.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN15_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN15.Reg, volatile.LoadUint32(&o.PIN15.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN15_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN15.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN15_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN15.Reg, volatile.LoadUint32(&o.PIN15.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN15_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN15.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN15_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN15.Reg, volatile.LoadUint32(&o.PIN15.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN15_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN15.Reg) & 0x3e000) >> 13
}

// GPIO.PIN16
func (o *GPIO_Type) SetPIN16_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN16.Reg, volatile.LoadUint32(&o.PIN16.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN16_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN16.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN16_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN16.Reg, volatile.LoadUint32(&o.PIN16.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN16_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN16.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN16_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN16.Reg, volatile.LoadUint32(&o.PIN16.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN16_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN16.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN16_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN16.Reg, volatile.LoadUint32(&o.PIN16.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN16_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN16.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN16_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN16.Reg, volatile.LoadUint32(&o.PIN16.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN16_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN16.Reg) & 0x3e000) >> 13
}

// GPIO.PIN17
func (o *GPIO_Type) SetPIN17_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN17.Reg, volatile.LoadUint32(&o.PIN17.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN17_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN17.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN17_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN17.Reg, volatile.LoadUint32(&o.PIN17.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN17_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN17.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN17_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN17.Reg, volatile.LoadUint32(&o.PIN17.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN17_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN17.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN17_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN17.Reg, volatile.LoadUint32(&o.PIN17.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN17_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN17.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN17_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN17.Reg, volatile.LoadUint32(&o.PIN17.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN17_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN17.Reg) & 0x3e000) >> 13
}

// GPIO.PIN18
func (o *GPIO_Type) SetPIN18_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN18.Reg, volatile.LoadUint32(&o.PIN18.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN18_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN18.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN18_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN18.Reg, volatile.LoadUint32(&o.PIN18.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN18_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN18.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN18_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN18.Reg, volatile.LoadUint32(&o.PIN18.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN18_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN18.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN18_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN18.Reg, volatile.LoadUint32(&o.PIN18.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN18_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN18.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN18_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN18.Reg, volatile.LoadUint32(&o.PIN18.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN18_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN18.Reg) & 0x3e000) >> 13
}

// GPIO.PIN19
func (o *GPIO_Type) SetPIN19_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN19.Reg, volatile.LoadUint32(&o.PIN19.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN19_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN19.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN19_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN19.Reg, volatile.LoadUint32(&o.PIN19.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN19_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN19.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN19_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN19.Reg, volatile.LoadUint32(&o.PIN19.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN19_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN19.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN19_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN19.Reg, volatile.LoadUint32(&o.PIN19.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN19_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN19.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN19_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN19.Reg, volatile.LoadUint32(&o.PIN19.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN19_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN19.Reg) & 0x3e000) >> 13
}

// GPIO.PIN20
func (o *GPIO_Type) SetPIN20_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN20.Reg, volatile.LoadUint32(&o.PIN20.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN20_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN20.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN20_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN20.Reg, volatile.LoadUint32(&o.PIN20.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN20_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN20.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN20_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN20.Reg, volatile.LoadUint32(&o.PIN20.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN20_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN20.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN20_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN20.Reg, volatile.LoadUint32(&o.PIN20.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN20_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN20.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN20_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN20.Reg, volatile.LoadUint32(&o.PIN20.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN20_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN20.Reg) & 0x3e000) >> 13
}

// GPIO.PIN21
func (o *GPIO_Type) SetPIN21_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN21.Reg, volatile.LoadUint32(&o.PIN21.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN21_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN21.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN21_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN21.Reg, volatile.LoadUint32(&o.PIN21.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN21_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN21.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN21_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN21.Reg, volatile.LoadUint32(&o.PIN21.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN21_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN21.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN21_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN21.Reg, volatile.LoadUint32(&o.PIN21.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN21_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN21.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN21_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN21.Reg, volatile.LoadUint32(&o.PIN21.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN21_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN21.Reg) & 0x3e000) >> 13
}

// GPIO.PIN22
func (o *GPIO_Type) SetPIN22_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN22.Reg, volatile.LoadUint32(&o.PIN22.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN22_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN22.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN22_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN22.Reg, volatile.LoadUint32(&o.PIN22.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN22_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN22.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN22_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN22.Reg, volatile.LoadUint32(&o.PIN22.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN22_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN22.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN22_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN22.Reg, volatile.LoadUint32(&o.PIN22.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN22_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN22.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN22_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN22.Reg, volatile.LoadUint32(&o.PIN22.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN22_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN22.Reg) & 0x3e000) >> 13
}

// GPIO.PIN23
func (o *GPIO_Type) SetPIN23_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN23.Reg, volatile.LoadUint32(&o.PIN23.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN23_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN23.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN23_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN23.Reg, volatile.LoadUint32(&o.PIN23.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN23_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN23.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN23_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN23.Reg, volatile.LoadUint32(&o.PIN23.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN23_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN23.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN23_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN23.Reg, volatile.LoadUint32(&o.PIN23.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN23_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN23.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN23_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN23.Reg, volatile.LoadUint32(&o.PIN23.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN23_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN23.Reg) & 0x3e000) >> 13
}

// GPIO.PIN24
func (o *GPIO_Type) SetPIN24_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN24.Reg, volatile.LoadUint32(&o.PIN24.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN24_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN24.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN24_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN24.Reg, volatile.LoadUint32(&o.PIN24.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN24_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN24.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN24_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN24.Reg, volatile.LoadUint32(&o.PIN24.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN24_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN24.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN24_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN24.Reg, volatile.LoadUint32(&o.PIN24.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN24_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN24.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN24_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN24.Reg, volatile.LoadUint32(&o.PIN24.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN24_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN24.Reg) & 0x3e000) >> 13
}

// GPIO.PIN25
func (o *GPIO_Type) SetPIN25_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN25.Reg, volatile.LoadUint32(&o.PIN25.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN25_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN25.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN25_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN25.Reg, volatile.LoadUint32(&o.PIN25.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN25_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN25.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN25_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN25.Reg, volatile.LoadUint32(&o.PIN25.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN25_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN25.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN25_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN25.Reg, volatile.LoadUint32(&o.PIN25.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN25_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN25.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN25_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN25.Reg, volatile.LoadUint32(&o.PIN25.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN25_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN25.Reg) & 0x3e000) >> 13
}

// GPIO.PIN26
func (o *GPIO_Type) SetPIN26_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN26.Reg, volatile.LoadUint32(&o.PIN26.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN26_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN26.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN26_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN26.Reg, volatile.LoadUint32(&o.PIN26.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN26_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN26.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN26_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN26.Reg, volatile.LoadUint32(&o.PIN26.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN26_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN26.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN26_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN26.Reg, volatile.LoadUint32(&o.PIN26.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN26_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN26.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN26_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN26.Reg, volatile.LoadUint32(&o.PIN26.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN26_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN26.Reg) & 0x3e000) >> 13
}

// GPIO.PIN27
func (o *GPIO_Type) SetPIN27_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN27.Reg, volatile.LoadUint32(&o.PIN27.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN27_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN27.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN27_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN27.Reg, volatile.LoadUint32(&o.PIN27.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN27_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN27.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN27_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN27.Reg, volatile.LoadUint32(&o.PIN27.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN27_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN27.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN27_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN27.Reg, volatile.LoadUint32(&o.PIN27.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN27_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN27.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN27_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN27.Reg, volatile.LoadUint32(&o.PIN27.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN27_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN27.Reg) & 0x3e000) >> 13
}

// GPIO.PIN28
func (o *GPIO_Type) SetPIN28_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN28.Reg, volatile.LoadUint32(&o.PIN28.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN28_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN28.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN28_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN28.Reg, volatile.LoadUint32(&o.PIN28.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN28_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN28.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN28_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN28.Reg, volatile.LoadUint32(&o.PIN28.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN28_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN28.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN28_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN28.Reg, volatile.LoadUint32(&o.PIN28.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN28_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN28.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN28_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN28.Reg, volatile.LoadUint32(&o.PIN28.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN28_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN28.Reg) & 0x3e000) >> 13
}

// GPIO.PIN29
func (o *GPIO_Type) SetPIN29_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN29.Reg, volatile.LoadUint32(&o.PIN29.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN29_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN29.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN29_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN29.Reg, volatile.LoadUint32(&o.PIN29.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN29_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN29.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN29_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN29.Reg, volatile.LoadUint32(&o.PIN29.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN29_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN29.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN29_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN29.Reg, volatile.LoadUint32(&o.PIN29.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN29_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN29.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN29_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN29.Reg, volatile.LoadUint32(&o.PIN29.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN29_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN29.Reg) & 0x3e000) >> 13
}

// GPIO.PIN30
func (o *GPIO_Type) SetPIN30_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN30.Reg, volatile.LoadUint32(&o.PIN30.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN30_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN30.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN30_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN30.Reg, volatile.LoadUint32(&o.PIN30.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN30_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN30.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN30_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN30.Reg, volatile.LoadUint32(&o.PIN30.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN30_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN30.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN30_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN30.Reg, volatile.LoadUint32(&o.PIN30.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN30_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN30.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN30_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN30.Reg, volatile.LoadUint32(&o.PIN30.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN30_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN30.Reg) & 0x3e000) >> 13
}

// GPIO.PIN31
func (o *GPIO_Type) SetPIN31_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN31.Reg, volatile.LoadUint32(&o.PIN31.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN31_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN31.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN31_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN31.Reg, volatile.LoadUint32(&o.PIN31.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN31_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN31.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN31_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN31.Reg, volatile.LoadUint32(&o.PIN31.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN31_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN31.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN31_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN31.Reg, volatile.LoadUint32(&o.PIN31.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN31_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN31.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN31_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN31.Reg, volatile.LoadUint32(&o.PIN31.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN31_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN31.Reg) & 0x3e000) >> 13
}

// GPIO.PIN32
func (o *GPIO_Type) SetPIN32_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN32.Reg, volatile.LoadUint32(&o.PIN32.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN32_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN32.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN32_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN32.Reg, volatile.LoadUint32(&o.PIN32.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN32_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN32.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN32_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN32.Reg, volatile.LoadUint32(&o.PIN32.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN32_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN32.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN32_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN32.Reg, volatile.LoadUint32(&o.PIN32.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN32_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN32.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN32_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN32.Reg, volatile.LoadUint32(&o.PIN32.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN32_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN32.Reg) & 0x3e000) >> 13
}

// GPIO.PIN33
func (o *GPIO_Type) SetPIN33_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN33.Reg, volatile.LoadUint32(&o.PIN33.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN33_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN33.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN33_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN33.Reg, volatile.LoadUint32(&o.PIN33.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN33_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN33.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN33_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN33.Reg, volatile.LoadUint32(&o.PIN33.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN33_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN33.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN33_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN33.Reg, volatile.LoadUint32(&o.PIN33.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN33_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN33.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN33_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN33.Reg, volatile.LoadUint32(&o.PIN33.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN33_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN33.Reg) & 0x3e000) >> 13
}

// GPIO.PIN34
func (o *GPIO_Type) SetPIN34_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN34.Reg, volatile.LoadUint32(&o.PIN34.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN34_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN34.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN34_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN34.Reg, volatile.LoadUint32(&o.PIN34.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN34_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN34.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN34_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN34.Reg, volatile.LoadUint32(&o.PIN34.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN34_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN34.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN34_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN34.Reg, volatile.LoadUint32(&o.PIN34.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN34_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN34.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN34_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN34.Reg, volatile.LoadUint32(&o.PIN34.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN34_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN34.Reg) & 0x3e000) >> 13
}

// GPIO.PIN35
func (o *GPIO_Type) SetPIN35_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN35.Reg, volatile.LoadUint32(&o.PIN35.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN35_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN35.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN35_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN35.Reg, volatile.LoadUint32(&o.PIN35.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN35_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN35.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN35_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN35.Reg, volatile.LoadUint32(&o.PIN35.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN35_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN35.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN35_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN35.Reg, volatile.LoadUint32(&o.PIN35.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN35_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN35.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN35_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN35.Reg, volatile.LoadUint32(&o.PIN35.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN35_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN35.Reg) & 0x3e000) >> 13
}

// GPIO.PIN36
func (o *GPIO_Type) SetPIN36_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN36.Reg, volatile.LoadUint32(&o.PIN36.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN36_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN36.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN36_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN36.Reg, volatile.LoadUint32(&o.PIN36.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN36_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN36.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN36_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN36.Reg, volatile.LoadUint32(&o.PIN36.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN36_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN36.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN36_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN36.Reg, volatile.LoadUint32(&o.PIN36.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN36_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN36.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN36_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN36.Reg, volatile.LoadUint32(&o.PIN36.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN36_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN36.Reg) & 0x3e000) >> 13
}

// GPIO.PIN37
func (o *GPIO_Type) SetPIN37_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN37.Reg, volatile.LoadUint32(&o.PIN37.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN37_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN37.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN37_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN37.Reg, volatile.LoadUint32(&o.PIN37.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN37_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN37.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN37_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN37.Reg, volatile.LoadUint32(&o.PIN37.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN37_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN37.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN37_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN37.Reg, volatile.LoadUint32(&o.PIN37.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN37_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN37.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN37_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN37.Reg, volatile.LoadUint32(&o.PIN37.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN37_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN37.Reg) & 0x3e000) >> 13
}

// GPIO.PIN38
func (o *GPIO_Type) SetPIN38_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN38.Reg, volatile.LoadUint32(&o.PIN38.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN38_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN38.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN38_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN38.Reg, volatile.LoadUint32(&o.PIN38.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN38_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN38.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN38_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN38.Reg, volatile.LoadUint32(&o.PIN38.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN38_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN38.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN38_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN38.Reg, volatile.LoadUint32(&o.PIN38.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN38_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN38.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN38_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN38.Reg, volatile.LoadUint32(&o.PIN38.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN38_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN38.Reg) & 0x3e000) >> 13
}

// GPIO.PIN39
func (o *GPIO_Type) SetPIN39_PAD_DRIVER(value uint32) {
	volatile.StoreUint32(&o.PIN39.Reg, volatile.LoadUint32(&o.PIN39.Reg)&^(0x4)|value<<2)
}
func (o *GPIO_Type) GetPIN39_PAD_DRIVER() uint32 {
	return (volatile.LoadUint32(&o.PIN39.Reg) & 0x4) >> 2
}
func (o *GPIO_Type) SetPIN39_INT_TYPE(value uint32) {
	volatile.StoreUint32(&o.PIN39.Reg, volatile.LoadUint32(&o.PIN39.Reg)&^(0x380)|value<<7)
}
func (o *GPIO_Type) GetPIN39_INT_TYPE() uint32 {
	return (volatile.LoadUint32(&o.PIN39.Reg) & 0x380) >> 7
}
func (o *GPIO_Type) SetPIN39_WAKEUP_ENABLE(value uint32) {
	volatile.StoreUint32(&o.PIN39.Reg, volatile.LoadUint32(&o.PIN39.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetPIN39_WAKEUP_ENABLE() uint32 {
	return (volatile.LoadUint32(&o.PIN39.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetPIN39_CONFIG(value uint32) {
	volatile.StoreUint32(&o.PIN39.Reg, volatile.LoadUint32(&o.PIN39.Reg)&^(0x1800)|value<<11)
}
func (o *GPIO_Type) GetPIN39_CONFIG() uint32 {
	return (volatile.LoadUint32(&o.PIN39.Reg) & 0x1800) >> 11
}
func (o *GPIO_Type) SetPIN39_INT_ENA(value uint32) {
	volatile.StoreUint32(&o.PIN39.Reg, volatile.LoadUint32(&o.PIN39.Reg)&^(0x3e000)|value<<13)
}
func (o *GPIO_Type) GetPIN39_INT_ENA() uint32 {
	return (volatile.LoadUint32(&o.PIN39.Reg) & 0x3e000) >> 13
}

// GPIO.CALI_CONF
func (o *GPIO_Type) SetCALI_CONF_CALI_RTC_MAX(value uint32) {
	volatile.StoreUint32(&o.CALI_CONF.Reg, volatile.LoadUint32(&o.CALI_CONF.Reg)&^(0x3ff)|value)
}
func (o *GPIO_Type) GetCALI_CONF_CALI_RTC_MAX() uint32 {
	return volatile.LoadUint32(&o.CALI_CONF.Reg) & 0x3ff
}
func (o *GPIO_Type) SetCALI_CONF_CALI_START(value uint32) {
	volatile.StoreUint32(&o.CALI_CONF.Reg, volatile.LoadUint32(&o.CALI_CONF.Reg)&^(0x80000000)|value<<31)
}
func (o *GPIO_Type) GetCALI_CONF_CALI_START() uint32 {
	return (volatile.LoadUint32(&o.CALI_CONF.Reg) & 0x80000000) >> 31
}

// GPIO.CALI_DATA
func (o *GPIO_Type) SetCALI_DATA_CALI_VALUE_SYNC2(value uint32) {
	volatile.StoreUint32(&o.CALI_DATA.Reg, volatile.LoadUint32(&o.CALI_DATA.Reg)&^(0xfffff)|value)
}
func (o *GPIO_Type) GetCALI_DATA_CALI_VALUE_SYNC2() uint32 {
	return volatile.LoadUint32(&o.CALI_DATA.Reg) & 0xfffff
}
func (o *GPIO_Type) SetCALI_DATA_CALI_RDY_REAL(value uint32) {
	volatile.StoreUint32(&o.CALI_DATA.Reg, volatile.LoadUint32(&o.CALI_DATA.Reg)&^(0x40000000)|value<<30)
}
func (o *GPIO_Type) GetCALI_DATA_CALI_RDY_REAL() uint32 {
	return (volatile.LoadUint32(&o.CALI_DATA.Reg) & 0x40000000) >> 30
}
func (o *GPIO_Type) SetCALI_DATA_CALI_RDY_SYNC2(value uint32) {
	volatile.StoreUint32(&o.CALI_DATA.Reg, volatile.LoadUint32(&o.CALI_DATA.Reg)&^(0x80000000)|value<<31)
}
func (o *GPIO_Type) GetCALI_DATA_CALI_RDY_SYNC2() uint32 {
	return (volatile.LoadUint32(&o.CALI_DATA.Reg) & 0x80000000) >> 31
}

// GPIO.FUNC0_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC0_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC0_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC0_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC0_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC0_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC0_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC0_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC0_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC0_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC0_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC0_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC0_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC0_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC0_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC0_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC1_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC1_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC1_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC1_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC1_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC1_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC1_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC1_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC1_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC1_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC1_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC1_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC1_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC1_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC1_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC1_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC2_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC2_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC2_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC2_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC2_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC2_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC2_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC2_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC2_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC2_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC2_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC2_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC2_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC2_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC2_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC2_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC3_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC3_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC3_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC3_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC3_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC3_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC3_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC3_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC3_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC3_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC3_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC3_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC3_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC3_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC3_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC3_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC4_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC4_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC4_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC4_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC4_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC4_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC4_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC4_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC4_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC4_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC4_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC4_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC4_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC4_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC4_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC4_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC5_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC5_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC5_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC5_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC5_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC5_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC5_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC5_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC5_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC5_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC5_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC5_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC5_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC5_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC5_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC5_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC6_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC6_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC6_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC6_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC6_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC6_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC6_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC6_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC6_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC6_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC6_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC6_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC6_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC6_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC6_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC6_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC7_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC7_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC7_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC7_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC7_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC7_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC7_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC7_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC7_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC7_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC7_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC7_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC7_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC7_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC7_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC7_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC8_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC8_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC8_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC8_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC8_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC8_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC8_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC8_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC8_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC8_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC8_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC8_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC8_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC8_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC8_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC8_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC9_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC9_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC9_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC9_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC9_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC9_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC9_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC9_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC9_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC9_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC9_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC9_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC9_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC9_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC9_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC9_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC10_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC10_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC10_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC10_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC10_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC10_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC10_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC10_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC10_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC10_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC10_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC10_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC10_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC10_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC10_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC10_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC11_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC11_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC11_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC11_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC11_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC11_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC11_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC11_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC11_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC11_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC11_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC11_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC11_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC11_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC11_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC11_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC12_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC12_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC12_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC12_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC12_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC12_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC12_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC12_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC12_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC12_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC12_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC12_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC12_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC12_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC12_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC12_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC13_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC13_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC13_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC13_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC13_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC13_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC13_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC13_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC13_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC13_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC13_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC13_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC13_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC13_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC13_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC13_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC14_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC14_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC14_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC14_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC14_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC14_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC14_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC14_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC14_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC14_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC14_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC14_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC14_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC14_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC14_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC14_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC15_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC15_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC15_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC15_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC15_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC15_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC15_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC15_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC15_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC15_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC15_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC15_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC15_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC15_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC15_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC15_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC16_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC16_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC16_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC16_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC16_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC16_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC16_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC16_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC16_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC16_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC16_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC16_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC16_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC16_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC16_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC16_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC17_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC17_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC17_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC17_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC17_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC17_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC17_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC17_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC17_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC17_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC17_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC17_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC17_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC17_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC17_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC17_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC18_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC18_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC18_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC18_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC18_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC18_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC18_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC18_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC18_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC18_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC18_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC18_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC18_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC18_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC18_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC18_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC19_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC19_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC19_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC19_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC19_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC19_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC19_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC19_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC19_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC19_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC19_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC19_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC19_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC19_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC19_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC19_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC20_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC20_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC20_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC20_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC20_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC20_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC20_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC20_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC20_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC20_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC20_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC20_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC20_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC20_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC20_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC20_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC21_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC21_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC21_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC21_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC21_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC21_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC21_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC21_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC21_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC21_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC21_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC21_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC21_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC21_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC21_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC21_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC22_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC22_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC22_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC22_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC22_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC22_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC22_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC22_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC22_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC22_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC22_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC22_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC22_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC22_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC22_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC22_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC23_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC23_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC23_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC23_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC23_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC23_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC23_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC23_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC23_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC23_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC23_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC23_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC23_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC23_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC23_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC23_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC24_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC24_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC24_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC24_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC24_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC24_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC24_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC24_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC24_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC24_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC24_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC24_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC24_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC24_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC24_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC24_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC25_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC25_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC25_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC25_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC25_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC25_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC25_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC25_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC25_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC25_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC25_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC25_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC25_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC25_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC25_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC25_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC26_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC26_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC26_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC26_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC26_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC26_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC26_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC26_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC26_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC26_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC26_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC26_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC26_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC26_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC26_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC26_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC27_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC27_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC27_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC27_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC27_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC27_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC27_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC27_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC27_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC27_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC27_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC27_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC27_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC27_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC27_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC27_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC28_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC28_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC28_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC28_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC28_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC28_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC28_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC28_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC28_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC28_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC28_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC28_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC28_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC28_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC28_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC28_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC29_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC29_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC29_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC29_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC29_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC29_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC29_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC29_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC29_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC29_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC29_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC29_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC29_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC29_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC29_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC29_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC30_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC30_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC30_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC30_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC30_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC30_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC30_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC30_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC30_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC30_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC30_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC30_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC30_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC30_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC30_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC30_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC31_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC31_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC31_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC31_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC31_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC31_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC31_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC31_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC31_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC31_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC31_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC31_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC31_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC31_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC31_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC31_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC32_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC32_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC32_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC32_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC32_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC32_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC32_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC32_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC32_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC32_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC32_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC32_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC32_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC32_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC32_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC32_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC33_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC33_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC33_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC33_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC33_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC33_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC33_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC33_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC33_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC33_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC33_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC33_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC33_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC33_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC33_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC33_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC34_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC34_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC34_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC34_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC34_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC34_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC34_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC34_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC34_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC34_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC34_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC34_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC34_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC34_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC34_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC34_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC35_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC35_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC35_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC35_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC35_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC35_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC35_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC35_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC35_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC35_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC35_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC35_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC35_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC35_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC35_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC35_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC36_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC36_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC36_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC36_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC36_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC36_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC36_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC36_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC36_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC36_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC36_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC36_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC36_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC36_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC36_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC36_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC37_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC37_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC37_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC37_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC37_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC37_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC37_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC37_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC37_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC37_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC37_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC37_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC37_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC37_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC37_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC37_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC38_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC38_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC38_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC38_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC38_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC38_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC38_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC38_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC38_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC38_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC38_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC38_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC38_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC38_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC38_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC38_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC39_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC39_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC39_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC39_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC39_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC39_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC39_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC39_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC39_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC39_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC39_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC39_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC39_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC39_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC39_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC39_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC40_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC40_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC40_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC40_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC40_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC40_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC40_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC40_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC40_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC40_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC40_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC40_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC40_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC40_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC40_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC40_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC41_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC41_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC41_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC41_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC41_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC41_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC41_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC41_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC41_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC41_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC41_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC41_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC41_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC41_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC41_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC41_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC42_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC42_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC42_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC42_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC42_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC42_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC42_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC42_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC42_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC42_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC42_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC42_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC42_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC42_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC42_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC42_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC43_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC43_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC43_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC43_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC43_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC43_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC43_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC43_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC43_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC43_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC43_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC43_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC43_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC43_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC43_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC43_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC44_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC44_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC44_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC44_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC44_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC44_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC44_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC44_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC44_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC44_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC44_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC44_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC44_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC44_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC44_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC44_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC45_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC45_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC45_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC45_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC45_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC45_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC45_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC45_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC45_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC45_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC45_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC45_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC45_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC45_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC45_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC45_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC46_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC46_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC46_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC46_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC46_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC46_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC46_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC46_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC46_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC46_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC46_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC46_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC46_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC46_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC46_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC46_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC47_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC47_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC47_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC47_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC47_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC47_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC47_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC47_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC47_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC47_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC47_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC47_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC47_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC47_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC47_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC47_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC48_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC48_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC48_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC48_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC48_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC48_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC48_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC48_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC48_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC48_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC48_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC48_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC48_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC48_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC48_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC48_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC49_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC49_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC49_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC49_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC49_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC49_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC49_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC49_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC49_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC49_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC49_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC49_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC49_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC49_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC49_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC49_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC50_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC50_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC50_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC50_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC50_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC50_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC50_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC50_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC50_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC50_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC50_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC50_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC50_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC50_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC50_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC50_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC51_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC51_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC51_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC51_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC51_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC51_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC51_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC51_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC51_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC51_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC51_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC51_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC51_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC51_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC51_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC51_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC52_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC52_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC52_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC52_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC52_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC52_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC52_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC52_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC52_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC52_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC52_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC52_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC52_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC52_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC52_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC52_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC53_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC53_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC53_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC53_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC53_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC53_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC53_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC53_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC53_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC53_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC53_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC53_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC53_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC53_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC53_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC53_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC54_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC54_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC54_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC54_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC54_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC54_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC54_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC54_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC54_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC54_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC54_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC54_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC54_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC54_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC54_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC54_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC55_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC55_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC55_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC55_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC55_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC55_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC55_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC55_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC55_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC55_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC55_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC55_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC55_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC55_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC55_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC55_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC56_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC56_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC56_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC56_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC56_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC56_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC56_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC56_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC56_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC56_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC56_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC56_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC56_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC56_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC56_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC56_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC57_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC57_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC57_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC57_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC57_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC57_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC57_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC57_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC57_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC57_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC57_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC57_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC57_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC57_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC57_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC57_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC58_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC58_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC58_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC58_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC58_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC58_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC58_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC58_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC58_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC58_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC58_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC58_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC58_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC58_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC58_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC58_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC59_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC59_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC59_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC59_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC59_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC59_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC59_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC59_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC59_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC59_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC59_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC59_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC59_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC59_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC59_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC59_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC60_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC60_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC60_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC60_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC60_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC60_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC60_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC60_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC60_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC60_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC60_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC60_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC60_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC60_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC60_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC60_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC61_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC61_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC61_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC61_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC61_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC61_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC61_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC61_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC61_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC61_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC61_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC61_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC61_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC61_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC61_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC61_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC62_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC62_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC62_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC62_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC62_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC62_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC62_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC62_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC62_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC62_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC62_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC62_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC62_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC62_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC62_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC62_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC63_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC63_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC63_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC63_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC63_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC63_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC63_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC63_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC63_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC63_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC63_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC63_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC63_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC63_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC63_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC63_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC64_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC64_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC64_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC64_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC64_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC64_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC64_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC64_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC64_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC64_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC64_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC64_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC64_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC64_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC64_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC64_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC65_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC65_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC65_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC65_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC65_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC65_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC65_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC65_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC65_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC65_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC65_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC65_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC65_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC65_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC65_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC65_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC66_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC66_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC66_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC66_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC66_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC66_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC66_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC66_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC66_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC66_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC66_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC66_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC66_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC66_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC66_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC66_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC67_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC67_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC67_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC67_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC67_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC67_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC67_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC67_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC67_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC67_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC67_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC67_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC67_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC67_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC67_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC67_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC68_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC68_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC68_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC68_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC68_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC68_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC68_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC68_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC68_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC68_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC68_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC68_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC68_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC68_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC68_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC68_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC69_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC69_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC69_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC69_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC69_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC69_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC69_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC69_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC69_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC69_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC69_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC69_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC69_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC69_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC69_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC69_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC70_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC70_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC70_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC70_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC70_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC70_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC70_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC70_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC70_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC70_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC70_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC70_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC70_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC70_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC70_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC70_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC71_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC71_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC71_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC71_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC71_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC71_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC71_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC71_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC71_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC71_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC71_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC71_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC71_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC71_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC71_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC71_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC72_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC72_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC72_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC72_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC72_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC72_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC72_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC72_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC72_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC72_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC72_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC72_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC72_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC72_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC72_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC72_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC73_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC73_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC73_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC73_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC73_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC73_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC73_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC73_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC73_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC73_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC73_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC73_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC73_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC73_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC73_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC73_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC74_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC74_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC74_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC74_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC74_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC74_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC74_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC74_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC74_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC74_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC74_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC74_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC74_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC74_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC74_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC74_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC75_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC75_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC75_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC75_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC75_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC75_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC75_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC75_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC75_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC75_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC75_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC75_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC75_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC75_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC75_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC75_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC76_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC76_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC76_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC76_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC76_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC76_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC76_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC76_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC76_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC76_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC76_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC76_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC76_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC76_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC76_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC76_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC77_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC77_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC77_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC77_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC77_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC77_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC77_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC77_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC77_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC77_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC77_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC77_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC77_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC77_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC77_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC77_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC78_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC78_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC78_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC78_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC78_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC78_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC78_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC78_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC78_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC78_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC78_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC78_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC78_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC78_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC78_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC78_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC79_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC79_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC79_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC79_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC79_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC79_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC79_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC79_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC79_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC79_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC79_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC79_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC79_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC79_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC79_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC79_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC80_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC80_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC80_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC80_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC80_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC80_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC80_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC80_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC80_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC80_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC80_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC80_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC80_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC80_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC80_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC80_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC81_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC81_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC81_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC81_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC81_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC81_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC81_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC81_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC81_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC81_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC81_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC81_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC81_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC81_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC81_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC81_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC82_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC82_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC82_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC82_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC82_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC82_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC82_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC82_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC82_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC82_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC82_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC82_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC82_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC82_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC82_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC82_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC83_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC83_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC83_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC83_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC83_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC83_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC83_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC83_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC83_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC83_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC83_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC83_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC83_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC83_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC83_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC83_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC84_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC84_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC84_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC84_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC84_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC84_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC84_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC84_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC84_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC84_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC84_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC84_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC84_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC84_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC84_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC84_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC85_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC85_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC85_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC85_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC85_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC85_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC85_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC85_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC85_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC85_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC85_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC85_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC85_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC85_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC85_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC85_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC86_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC86_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC86_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC86_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC86_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC86_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC86_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC86_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC86_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC86_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC86_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC86_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC86_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC86_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC86_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC86_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC87_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC87_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC87_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC87_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC87_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC87_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC87_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC87_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC87_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC87_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC87_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC87_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC87_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC87_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC87_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC87_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC88_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC88_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC88_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC88_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC88_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC88_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC88_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC88_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC88_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC88_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC88_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC88_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC88_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC88_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC88_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC88_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC89_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC89_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC89_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC89_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC89_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC89_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC89_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC89_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC89_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC89_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC89_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC89_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC89_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC89_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC89_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC89_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC90_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC90_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC90_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC90_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC90_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC90_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC90_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC90_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC90_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC90_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC90_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC90_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC90_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC90_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC90_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC90_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC91_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC91_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC91_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC91_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC91_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC91_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC91_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC91_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC91_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC91_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC91_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC91_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC91_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC91_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC91_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC91_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC92_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC92_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC92_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC92_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC92_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC92_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC92_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC92_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC92_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC92_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC92_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC92_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC92_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC92_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC92_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC92_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC93_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC93_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC93_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC93_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC93_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC93_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC93_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC93_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC93_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC93_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC93_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC93_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC93_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC93_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC93_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC93_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC94_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC94_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC94_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC94_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC94_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC94_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC94_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC94_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC94_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC94_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC94_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC94_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC94_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC94_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC94_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC94_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC95_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC95_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC95_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC95_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC95_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC95_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC95_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC95_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC95_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC95_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC95_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC95_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC95_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC95_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC95_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC95_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC96_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC96_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC96_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC96_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC96_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC96_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC96_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC96_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC96_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC96_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC96_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC96_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC96_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC96_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC96_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC96_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC97_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC97_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC97_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC97_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC97_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC97_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC97_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC97_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC97_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC97_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC97_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC97_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC97_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC97_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC97_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC97_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC98_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC98_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC98_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC98_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC98_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC98_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC98_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC98_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC98_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC98_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC98_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC98_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC98_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC98_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC98_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC98_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC99_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC99_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC99_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC99_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC99_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC99_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC99_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC99_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC99_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC99_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC99_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC99_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC99_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC99_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC99_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC99_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC100_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC100_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC100_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC100_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC100_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC100_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC100_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC100_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC100_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC100_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC100_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC100_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC100_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC100_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC100_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC100_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC101_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC101_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC101_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC101_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC101_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC101_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC101_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC101_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC101_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC101_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC101_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC101_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC101_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC101_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC101_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC101_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC102_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC102_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC102_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC102_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC102_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC102_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC102_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC102_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC102_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC102_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC102_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC102_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC102_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC102_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC102_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC102_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC103_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC103_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC103_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC103_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC103_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC103_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC103_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC103_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC103_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC103_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC103_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC103_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC103_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC103_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC103_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC103_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC104_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC104_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC104_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC104_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC104_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC104_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC104_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC104_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC104_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC104_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC104_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC104_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC104_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC104_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC104_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC104_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC105_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC105_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC105_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC105_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC105_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC105_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC105_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC105_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC105_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC105_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC105_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC105_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC105_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC105_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC105_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC105_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC106_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC106_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC106_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC106_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC106_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC106_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC106_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC106_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC106_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC106_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC106_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC106_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC106_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC106_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC106_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC106_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC107_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC107_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC107_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC107_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC107_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC107_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC107_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC107_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC107_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC107_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC107_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC107_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC107_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC107_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC107_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC107_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC108_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC108_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC108_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC108_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC108_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC108_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC108_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC108_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC108_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC108_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC108_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC108_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC108_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC108_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC108_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC108_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC109_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC109_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC109_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC109_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC109_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC109_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC109_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC109_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC109_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC109_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC109_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC109_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC109_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC109_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC109_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC109_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC110_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC110_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC110_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC110_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC110_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC110_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC110_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC110_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC110_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC110_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC110_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC110_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC110_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC110_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC110_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC110_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC111_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC111_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC111_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC111_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC111_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC111_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC111_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC111_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC111_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC111_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC111_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC111_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC111_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC111_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC111_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC111_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC112_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC112_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC112_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC112_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC112_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC112_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC112_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC112_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC112_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC112_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC112_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC112_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC112_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC112_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC112_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC112_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC113_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC113_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC113_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC113_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC113_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC113_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC113_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC113_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC113_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC113_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC113_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC113_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC113_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC113_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC113_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC113_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC114_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC114_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC114_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC114_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC114_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC114_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC114_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC114_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC114_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC114_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC114_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC114_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC114_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC114_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC114_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC114_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC115_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC115_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC115_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC115_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC115_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC115_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC115_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC115_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC115_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC115_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC115_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC115_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC115_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC115_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC115_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC115_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC116_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC116_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC116_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC116_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC116_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC116_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC116_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC116_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC116_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC116_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC116_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC116_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC116_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC116_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC116_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC116_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC117_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC117_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC117_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC117_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC117_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC117_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC117_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC117_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC117_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC117_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC117_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC117_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC117_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC117_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC117_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC117_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC118_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC118_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC118_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC118_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC118_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC118_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC118_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC118_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC118_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC118_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC118_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC118_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC118_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC118_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC118_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC118_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC119_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC119_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC119_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC119_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC119_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC119_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC119_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC119_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC119_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC119_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC119_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC119_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC119_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC119_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC119_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC119_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC120_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC120_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC120_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC120_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC120_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC120_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC120_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC120_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC120_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC120_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC120_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC120_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC120_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC120_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC120_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC120_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC121_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC121_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC121_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC121_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC121_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC121_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC121_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC121_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC121_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC121_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC121_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC121_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC121_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC121_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC121_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC121_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC122_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC122_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC122_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC122_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC122_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC122_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC122_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC122_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC122_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC122_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC122_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC122_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC122_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC122_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC122_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC122_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC123_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC123_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC123_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC123_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC123_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC123_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC123_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC123_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC123_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC123_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC123_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC123_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC123_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC123_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC123_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC123_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC124_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC124_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC124_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC124_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC124_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC124_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC124_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC124_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC124_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC124_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC124_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC124_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC124_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC124_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC124_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC124_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC125_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC125_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC125_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC125_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC125_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC125_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC125_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC125_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC125_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC125_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC125_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC125_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC125_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC125_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC125_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC125_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC126_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC126_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC126_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC126_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC126_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC126_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC126_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC126_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC126_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC126_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC126_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC126_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC126_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC126_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC126_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC126_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC127_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC127_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC127_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC127_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC127_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC127_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC127_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC127_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC127_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC127_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC127_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC127_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC127_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC127_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC127_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC127_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC128_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC128_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC128_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC128_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC128_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC128_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC128_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC128_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC128_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC128_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC128_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC128_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC128_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC128_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC128_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC128_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC129_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC129_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC129_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC129_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC129_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC129_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC129_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC129_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC129_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC129_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC129_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC129_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC129_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC129_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC129_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC129_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC130_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC130_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC130_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC130_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC130_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC130_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC130_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC130_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC130_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC130_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC130_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC130_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC130_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC130_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC130_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC130_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC131_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC131_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC131_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC131_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC131_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC131_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC131_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC131_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC131_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC131_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC131_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC131_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC131_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC131_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC131_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC131_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC132_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC132_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC132_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC132_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC132_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC132_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC132_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC132_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC132_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC132_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC132_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC132_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC132_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC132_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC132_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC132_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC133_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC133_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC133_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC133_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC133_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC133_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC133_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC133_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC133_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC133_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC133_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC133_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC133_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC133_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC133_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC133_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC134_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC134_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC134_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC134_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC134_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC134_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC134_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC134_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC134_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC134_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC134_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC134_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC134_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC134_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC134_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC134_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC135_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC135_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC135_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC135_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC135_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC135_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC135_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC135_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC135_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC135_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC135_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC135_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC135_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC135_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC135_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC135_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC136_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC136_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC136_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC136_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC136_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC136_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC136_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC136_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC136_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC136_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC136_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC136_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC136_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC136_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC136_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC136_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC137_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC137_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC137_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC137_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC137_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC137_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC137_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC137_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC137_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC137_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC137_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC137_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC137_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC137_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC137_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC137_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC138_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC138_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC138_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC138_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC138_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC138_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC138_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC138_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC138_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC138_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC138_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC138_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC138_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC138_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC138_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC138_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC139_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC139_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC139_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC139_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC139_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC139_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC139_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC139_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC139_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC139_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC139_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC139_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC139_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC139_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC139_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC139_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC140_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC140_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC140_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC140_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC140_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC140_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC140_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC140_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC140_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC140_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC140_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC140_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC140_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC140_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC140_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC140_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC141_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC141_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC141_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC141_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC141_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC141_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC141_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC141_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC141_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC141_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC141_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC141_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC141_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC141_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC141_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC141_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC142_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC142_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC142_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC142_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC142_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC142_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC142_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC142_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC142_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC142_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC142_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC142_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC142_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC142_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC142_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC142_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC143_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC143_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC143_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC143_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC143_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC143_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC143_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC143_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC143_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC143_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC143_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC143_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC143_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC143_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC143_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC143_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC144_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC144_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC144_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC144_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC144_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC144_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC144_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC144_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC144_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC144_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC144_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC144_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC144_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC144_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC144_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC144_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC145_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC145_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC145_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC145_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC145_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC145_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC145_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC145_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC145_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC145_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC145_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC145_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC145_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC145_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC145_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC145_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC146_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC146_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC146_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC146_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC146_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC146_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC146_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC146_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC146_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC146_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC146_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC146_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC146_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC146_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC146_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC146_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC147_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC147_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC147_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC147_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC147_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC147_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC147_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC147_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC147_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC147_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC147_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC147_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC147_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC147_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC147_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC147_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC148_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC148_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC148_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC148_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC148_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC148_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC148_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC148_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC148_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC148_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC148_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC148_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC148_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC148_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC148_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC148_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC149_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC149_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC149_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC149_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC149_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC149_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC149_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC149_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC149_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC149_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC149_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC149_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC149_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC149_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC149_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC149_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC150_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC150_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC150_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC150_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC150_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC150_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC150_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC150_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC150_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC150_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC150_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC150_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC150_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC150_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC150_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC150_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC151_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC151_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC151_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC151_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC151_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC151_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC151_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC151_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC151_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC151_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC151_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC151_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC151_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC151_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC151_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC151_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC152_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC152_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC152_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC152_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC152_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC152_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC152_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC152_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC152_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC152_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC152_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC152_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC152_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC152_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC152_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC152_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC153_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC153_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC153_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC153_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC153_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC153_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC153_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC153_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC153_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC153_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC153_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC153_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC153_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC153_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC153_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC153_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC154_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC154_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC154_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC154_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC154_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC154_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC154_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC154_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC154_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC154_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC154_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC154_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC154_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC154_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC154_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC154_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC155_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC155_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC155_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC155_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC155_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC155_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC155_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC155_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC155_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC155_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC155_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC155_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC155_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC155_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC155_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC155_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC156_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC156_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC156_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC156_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC156_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC156_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC156_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC156_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC156_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC156_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC156_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC156_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC156_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC156_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC156_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC156_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC157_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC157_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC157_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC157_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC157_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC157_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC157_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC157_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC157_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC157_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC157_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC157_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC157_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC157_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC157_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC157_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC158_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC158_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC158_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC158_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC158_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC158_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC158_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC158_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC158_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC158_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC158_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC158_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC158_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC158_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC158_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC158_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC159_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC159_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC159_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC159_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC159_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC159_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC159_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC159_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC159_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC159_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC159_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC159_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC159_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC159_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC159_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC159_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC160_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC160_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC160_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC160_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC160_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC160_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC160_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC160_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC160_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC160_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC160_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC160_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC160_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC160_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC160_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC160_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC161_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC161_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC161_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC161_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC161_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC161_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC161_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC161_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC161_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC161_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC161_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC161_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC161_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC161_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC161_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC161_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC162_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC162_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC162_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC162_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC162_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC162_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC162_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC162_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC162_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC162_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC162_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC162_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC162_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC162_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC162_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC162_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC163_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC163_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC163_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC163_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC163_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC163_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC163_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC163_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC163_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC163_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC163_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC163_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC163_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC163_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC163_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC163_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC164_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC164_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC164_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC164_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC164_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC164_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC164_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC164_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC164_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC164_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC164_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC164_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC164_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC164_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC164_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC164_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC165_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC165_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC165_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC165_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC165_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC165_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC165_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC165_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC165_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC165_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC165_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC165_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC165_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC165_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC165_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC165_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC166_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC166_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC166_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC166_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC166_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC166_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC166_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC166_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC166_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC166_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC166_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC166_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC166_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC166_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC166_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC166_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC167_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC167_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC167_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC167_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC167_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC167_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC167_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC167_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC167_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC167_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC167_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC167_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC167_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC167_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC167_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC167_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC168_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC168_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC168_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC168_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC168_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC168_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC168_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC168_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC168_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC168_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC168_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC168_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC168_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC168_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC168_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC168_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC169_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC169_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC169_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC169_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC169_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC169_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC169_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC169_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC169_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC169_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC169_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC169_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC169_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC169_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC169_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC169_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC170_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC170_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC170_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC170_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC170_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC170_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC170_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC170_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC170_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC170_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC170_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC170_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC170_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC170_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC170_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC170_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC171_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC171_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC171_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC171_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC171_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC171_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC171_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC171_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC171_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC171_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC171_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC171_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC171_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC171_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC171_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC171_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC172_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC172_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC172_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC172_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC172_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC172_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC172_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC172_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC172_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC172_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC172_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC172_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC172_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC172_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC172_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC172_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC173_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC173_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC173_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC173_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC173_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC173_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC173_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC173_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC173_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC173_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC173_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC173_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC173_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC173_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC173_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC173_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC174_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC174_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC174_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC174_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC174_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC174_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC174_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC174_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC174_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC174_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC174_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC174_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC174_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC174_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC174_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC174_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC175_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC175_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC175_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC175_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC175_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC175_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC175_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC175_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC175_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC175_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC175_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC175_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC175_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC175_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC175_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC175_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC176_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC176_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC176_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC176_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC176_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC176_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC176_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC176_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC176_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC176_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC176_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC176_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC176_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC176_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC176_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC176_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC177_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC177_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC177_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC177_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC177_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC177_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC177_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC177_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC177_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC177_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC177_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC177_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC177_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC177_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC177_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC177_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC178_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC178_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC178_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC178_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC178_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC178_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC178_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC178_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC178_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC178_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC178_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC178_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC178_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC178_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC178_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC178_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC179_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC179_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC179_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC179_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC179_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC179_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC179_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC179_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC179_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC179_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC179_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC179_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC179_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC179_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC179_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC179_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC180_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC180_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC180_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC180_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC180_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC180_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC180_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC180_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC180_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC180_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC180_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC180_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC180_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC180_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC180_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC180_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC181_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC181_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC181_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC181_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC181_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC181_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC181_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC181_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC181_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC181_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC181_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC181_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC181_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC181_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC181_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC181_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC182_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC182_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC182_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC182_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC182_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC182_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC182_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC182_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC182_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC182_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC182_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC182_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC182_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC182_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC182_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC182_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC183_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC183_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC183_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC183_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC183_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC183_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC183_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC183_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC183_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC183_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC183_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC183_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC183_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC183_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC183_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC183_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC184_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC184_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC184_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC184_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC184_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC184_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC184_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC184_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC184_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC184_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC184_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC184_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC184_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC184_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC184_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC184_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC185_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC185_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC185_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC185_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC185_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC185_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC185_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC185_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC185_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC185_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC185_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC185_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC185_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC185_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC185_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC185_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC186_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC186_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC186_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC186_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC186_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC186_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC186_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC186_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC186_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC186_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC186_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC186_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC186_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC186_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC186_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC186_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC187_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC187_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC187_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC187_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC187_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC187_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC187_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC187_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC187_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC187_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC187_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC187_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC187_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC187_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC187_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC187_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC188_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC188_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC188_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC188_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC188_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC188_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC188_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC188_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC188_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC188_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC188_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC188_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC188_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC188_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC188_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC188_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC189_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC189_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC189_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC189_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC189_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC189_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC189_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC189_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC189_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC189_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC189_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC189_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC189_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC189_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC189_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC189_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC190_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC190_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC190_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC190_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC190_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC190_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC190_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC190_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC190_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC190_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC190_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC190_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC190_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC190_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC190_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC190_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC191_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC191_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC191_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC191_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC191_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC191_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC191_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC191_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC191_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC191_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC191_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC191_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC191_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC191_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC191_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC191_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC192_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC192_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC192_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC192_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC192_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC192_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC192_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC192_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC192_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC192_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC192_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC192_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC192_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC192_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC192_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC192_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC193_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC193_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC193_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC193_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC193_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC193_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC193_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC193_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC193_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC193_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC193_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC193_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC193_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC193_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC193_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC193_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC194_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC194_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC194_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC194_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC194_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC194_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC194_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC194_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC194_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC194_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC194_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC194_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC194_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC194_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC194_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC194_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC195_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC195_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC195_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC195_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC195_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC195_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC195_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC195_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC195_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC195_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC195_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC195_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC195_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC195_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC195_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC195_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC196_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC196_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC196_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC196_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC196_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC196_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC196_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC196_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC196_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC196_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC196_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC196_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC196_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC196_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC196_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC196_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC197_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC197_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC197_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC197_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC197_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC197_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC197_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC197_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC197_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC197_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC197_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC197_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC197_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC197_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC197_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC197_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC198_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC198_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC198_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC198_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC198_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC198_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC198_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC198_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC198_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC198_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC198_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC198_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC198_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC198_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC198_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC198_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC199_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC199_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC199_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC199_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC199_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC199_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC199_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC199_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC199_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC199_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC199_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC199_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC199_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC199_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC199_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC199_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC200_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC200_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC200_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC200_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC200_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC200_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC200_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC200_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC200_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC200_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC200_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC200_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC200_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC200_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC200_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC200_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC201_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC201_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC201_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC201_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC201_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC201_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC201_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC201_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC201_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC201_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC201_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC201_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC201_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC201_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC201_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC201_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC202_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC202_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC202_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC202_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC202_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC202_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC202_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC202_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC202_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC202_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC202_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC202_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC202_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC202_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC202_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC202_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC203_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC203_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC203_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC203_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC203_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC203_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC203_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC203_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC203_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC203_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC203_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC203_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC203_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC203_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC203_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC203_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC204_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC204_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC204_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC204_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC204_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC204_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC204_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC204_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC204_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC204_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC204_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC204_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC204_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC204_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC204_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC204_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC205_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC205_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC205_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC205_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC205_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC205_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC205_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC205_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC205_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC205_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC205_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC205_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC205_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC205_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC205_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC205_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC206_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC206_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC206_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC206_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC206_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC206_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC206_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC206_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC206_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC206_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC206_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC206_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC206_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC206_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC206_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC206_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC207_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC207_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC207_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC207_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC207_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC207_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC207_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC207_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC207_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC207_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC207_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC207_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC207_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC207_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC207_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC207_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC208_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC208_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC208_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC208_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC208_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC208_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC208_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC208_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC208_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC208_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC208_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC208_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC208_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC208_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC208_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC208_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC209_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC209_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC209_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC209_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC209_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC209_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC209_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC209_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC209_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC209_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC209_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC209_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC209_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC209_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC209_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC209_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC210_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC210_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC210_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC210_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC210_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC210_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC210_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC210_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC210_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC210_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC210_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC210_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC210_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC210_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC210_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC210_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC211_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC211_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC211_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC211_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC211_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC211_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC211_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC211_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC211_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC211_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC211_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC211_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC211_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC211_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC211_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC211_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC212_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC212_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC212_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC212_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC212_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC212_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC212_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC212_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC212_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC212_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC212_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC212_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC212_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC212_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC212_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC212_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC213_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC213_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC213_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC213_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC213_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC213_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC213_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC213_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC213_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC213_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC213_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC213_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC213_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC213_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC213_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC213_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC214_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC214_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC214_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC214_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC214_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC214_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC214_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC214_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC214_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC214_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC214_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC214_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC214_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC214_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC214_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC214_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC215_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC215_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC215_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC215_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC215_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC215_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC215_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC215_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC215_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC215_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC215_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC215_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC215_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC215_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC215_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC215_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC216_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC216_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC216_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC216_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC216_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC216_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC216_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC216_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC216_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC216_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC216_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC216_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC216_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC216_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC216_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC216_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC217_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC217_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC217_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC217_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC217_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC217_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC217_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC217_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC217_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC217_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC217_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC217_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC217_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC217_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC217_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC217_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC218_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC218_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC218_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC218_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC218_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC218_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC218_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC218_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC218_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC218_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC218_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC218_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC218_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC218_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC218_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC218_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC219_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC219_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC219_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC219_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC219_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC219_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC219_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC219_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC219_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC219_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC219_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC219_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC219_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC219_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC219_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC219_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC220_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC220_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC220_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC220_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC220_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC220_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC220_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC220_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC220_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC220_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC220_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC220_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC220_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC220_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC220_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC220_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC221_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC221_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC221_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC221_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC221_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC221_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC221_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC221_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC221_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC221_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC221_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC221_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC221_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC221_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC221_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC221_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC222_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC222_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC222_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC222_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC222_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC222_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC222_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC222_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC222_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC222_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC222_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC222_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC222_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC222_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC222_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC222_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC223_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC223_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC223_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC223_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC223_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC223_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC223_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC223_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC223_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC223_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC223_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC223_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC223_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC223_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC223_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC223_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC224_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC224_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC224_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC224_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC224_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC224_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC224_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC224_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC224_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC224_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC224_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC224_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC224_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC224_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC224_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC224_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC225_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC225_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC225_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC225_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC225_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC225_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC225_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC225_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC225_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC225_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC225_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC225_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC225_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC225_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC225_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC225_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC226_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC226_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC226_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC226_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC226_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC226_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC226_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC226_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC226_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC226_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC226_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC226_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC226_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC226_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC226_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC226_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC227_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC227_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC227_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC227_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC227_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC227_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC227_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC227_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC227_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC227_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC227_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC227_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC227_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC227_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC227_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC227_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC228_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC228_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC228_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC228_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC228_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC228_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC228_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC228_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC228_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC228_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC228_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC228_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC228_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC228_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC228_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC228_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC229_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC229_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC229_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC229_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC229_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC229_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC229_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC229_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC229_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC229_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC229_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC229_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC229_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC229_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC229_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC229_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC230_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC230_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC230_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC230_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC230_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC230_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC230_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC230_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC230_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC230_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC230_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC230_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC230_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC230_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC230_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC230_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC231_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC231_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC231_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC231_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC231_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC231_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC231_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC231_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC231_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC231_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC231_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC231_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC231_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC231_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC231_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC231_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC232_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC232_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC232_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC232_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC232_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC232_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC232_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC232_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC232_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC232_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC232_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC232_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC232_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC232_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC232_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC232_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC233_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC233_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC233_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC233_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC233_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC233_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC233_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC233_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC233_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC233_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC233_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC233_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC233_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC233_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC233_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC233_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC234_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC234_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC234_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC234_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC234_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC234_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC234_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC234_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC234_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC234_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC234_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC234_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC234_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC234_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC234_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC234_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC235_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC235_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC235_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC235_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC235_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC235_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC235_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC235_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC235_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC235_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC235_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC235_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC235_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC235_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC235_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC235_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC236_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC236_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC236_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC236_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC236_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC236_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC236_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC236_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC236_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC236_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC236_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC236_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC236_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC236_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC236_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC236_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC237_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC237_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC237_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC237_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC237_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC237_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC237_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC237_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC237_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC237_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC237_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC237_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC237_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC237_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC237_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC237_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC238_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC238_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC238_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC238_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC238_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC238_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC238_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC238_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC238_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC238_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC238_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC238_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC238_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC238_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC238_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC238_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC239_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC239_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC239_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC239_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC239_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC239_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC239_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC239_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC239_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC239_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC239_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC239_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC239_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC239_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC239_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC239_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC240_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC240_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC240_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC240_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC240_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC240_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC240_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC240_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC240_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC240_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC240_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC240_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC240_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC240_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC240_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC240_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC241_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC241_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC241_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC241_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC241_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC241_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC241_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC241_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC241_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC241_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC241_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC241_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC241_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC241_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC241_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC241_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC242_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC242_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC242_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC242_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC242_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC242_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC242_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC242_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC242_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC242_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC242_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC242_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC242_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC242_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC242_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC242_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC243_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC243_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC243_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC243_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC243_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC243_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC243_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC243_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC243_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC243_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC243_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC243_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC243_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC243_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC243_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC243_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC244_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC244_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC244_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC244_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC244_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC244_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC244_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC244_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC244_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC244_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC244_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC244_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC244_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC244_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC244_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC244_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC245_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC245_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC245_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC245_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC245_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC245_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC245_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC245_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC245_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC245_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC245_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC245_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC245_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC245_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC245_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC245_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC246_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC246_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC246_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC246_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC246_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC246_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC246_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC246_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC246_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC246_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC246_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC246_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC246_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC246_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC246_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC246_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC247_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC247_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC247_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC247_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC247_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC247_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC247_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC247_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC247_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC247_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC247_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC247_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC247_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC247_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC247_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC247_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC248_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC248_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC248_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC248_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC248_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC248_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC248_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC248_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC248_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC248_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC248_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC248_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC248_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC248_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC248_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC248_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC249_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC249_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC249_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC249_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC249_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC249_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC249_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC249_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC249_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC249_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC249_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC249_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC249_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC249_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC249_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC249_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC250_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC250_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC250_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC250_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC250_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC250_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC250_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC250_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC250_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC250_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC250_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC250_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC250_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC250_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC250_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC250_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC251_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC251_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC251_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC251_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC251_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC251_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC251_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC251_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC251_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC251_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC251_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC251_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC251_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC251_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC251_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC251_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC252_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC252_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC252_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC252_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC252_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC252_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC252_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC252_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC252_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC252_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC252_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC252_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC252_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC252_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC252_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC252_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC253_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC253_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC253_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC253_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC253_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC253_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC253_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC253_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC253_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC253_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC253_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC253_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC253_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC253_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC253_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC253_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC254_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC254_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC254_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC254_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC254_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC254_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC254_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC254_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC254_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC254_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC254_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC254_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC254_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC254_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC254_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC254_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC255_IN_SEL_CFG
func (o *GPIO_Type) SetFUNC255_IN_SEL_CFG_IN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC255_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC255_IN_SEL_CFG.Reg)&^(0x3f)|value)
}
func (o *GPIO_Type) GetFUNC255_IN_SEL_CFG_IN_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC255_IN_SEL_CFG.Reg) & 0x3f
}
func (o *GPIO_Type) SetFUNC255_IN_SEL_CFG_IN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC255_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC255_IN_SEL_CFG.Reg)&^(0x40)|value<<6)
}
func (o *GPIO_Type) GetFUNC255_IN_SEL_CFG_IN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC255_IN_SEL_CFG.Reg) & 0x40) >> 6
}
func (o *GPIO_Type) SetFUNC255_IN_SEL_CFG_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC255_IN_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC255_IN_SEL_CFG.Reg)&^(0x80)|value<<7)
}
func (o *GPIO_Type) GetFUNC255_IN_SEL_CFG_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC255_IN_SEL_CFG.Reg) & 0x80) >> 7
}

// GPIO.FUNC0_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC0_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC0_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC0_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC0_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC0_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC0_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC0_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC0_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC0_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC0_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC0_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC0_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC0_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC0_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC0_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC0_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC0_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC0_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC0_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC0_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC1_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC1_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC1_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC1_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC1_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC1_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC1_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC1_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC1_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC1_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC1_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC1_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC1_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC1_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC1_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC1_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC1_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC1_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC1_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC1_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC1_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC2_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC2_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC2_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC2_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC2_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC2_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC2_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC2_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC2_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC2_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC2_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC2_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC2_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC2_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC2_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC2_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC2_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC2_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC2_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC2_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC2_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC3_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC3_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC3_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC3_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC3_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC3_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC3_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC3_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC3_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC3_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC3_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC3_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC3_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC3_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC3_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC3_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC3_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC3_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC3_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC3_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC3_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC4_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC4_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC4_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC4_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC4_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC4_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC4_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC4_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC4_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC4_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC4_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC4_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC4_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC4_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC4_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC4_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC4_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC4_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC4_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC4_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC4_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC5_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC5_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC5_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC5_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC5_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC5_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC5_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC5_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC5_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC5_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC5_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC5_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC5_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC5_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC5_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC5_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC5_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC5_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC5_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC5_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC5_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC6_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC6_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC6_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC6_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC6_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC6_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC6_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC6_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC6_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC6_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC6_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC6_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC6_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC6_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC6_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC6_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC6_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC6_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC6_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC6_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC6_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC7_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC7_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC7_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC7_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC7_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC7_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC7_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC7_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC7_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC7_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC7_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC7_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC7_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC7_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC7_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC7_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC7_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC7_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC7_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC7_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC7_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC8_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC8_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC8_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC8_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC8_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC8_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC8_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC8_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC8_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC8_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC8_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC8_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC8_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC8_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC8_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC8_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC8_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC8_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC8_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC8_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC8_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC9_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC9_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC9_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC9_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC9_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC9_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC9_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC9_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC9_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC9_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC9_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC9_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC9_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC9_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC9_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC9_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC9_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC9_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC9_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC9_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC9_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC10_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC10_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC10_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC10_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC10_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC10_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC10_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC10_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC10_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC10_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC10_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC10_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC10_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC10_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC10_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC10_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC10_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC10_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC10_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC10_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC10_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC11_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC11_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC11_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC11_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC11_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC11_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC11_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC11_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC11_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC11_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC11_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC11_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC11_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC11_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC11_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC11_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC11_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC11_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC11_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC11_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC11_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC12_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC12_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC12_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC12_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC12_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC12_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC12_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC12_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC12_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC12_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC12_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC12_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC12_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC12_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC12_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC12_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC12_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC12_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC12_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC12_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC12_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC13_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC13_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC13_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC13_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC13_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC13_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC13_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC13_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC13_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC13_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC13_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC13_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC13_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC13_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC13_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC13_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC13_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC13_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC13_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC13_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC13_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC14_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC14_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC14_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC14_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC14_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC14_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC14_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC14_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC14_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC14_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC14_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC14_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC14_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC14_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC14_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC14_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC14_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC14_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC14_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC14_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC14_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC15_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC15_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC15_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC15_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC15_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC15_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC15_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC15_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC15_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC15_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC15_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC15_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC15_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC15_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC15_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC15_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC15_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC15_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC15_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC15_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC15_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC16_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC16_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC16_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC16_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC16_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC16_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC16_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC16_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC16_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC16_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC16_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC16_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC16_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC16_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC16_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC16_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC16_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC16_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC16_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC16_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC16_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC17_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC17_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC17_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC17_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC17_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC17_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC17_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC17_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC17_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC17_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC17_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC17_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC17_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC17_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC17_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC17_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC17_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC17_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC17_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC17_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC17_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC18_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC18_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC18_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC18_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC18_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC18_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC18_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC18_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC18_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC18_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC18_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC18_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC18_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC18_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC18_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC18_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC18_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC18_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC18_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC18_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC18_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC19_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC19_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC19_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC19_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC19_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC19_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC19_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC19_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC19_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC19_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC19_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC19_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC19_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC19_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC19_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC19_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC19_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC19_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC19_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC19_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC19_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC20_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC20_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC20_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC20_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC20_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC20_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC20_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC20_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC20_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC20_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC20_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC20_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC20_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC20_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC20_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC20_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC20_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC20_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC20_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC20_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC20_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC21_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC21_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC21_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC21_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC21_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC21_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC21_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC21_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC21_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC21_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC21_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC21_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC21_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC21_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC21_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC21_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC21_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC21_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC21_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC21_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC21_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC22_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC22_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC22_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC22_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC22_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC22_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC22_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC22_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC22_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC22_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC22_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC22_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC22_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC22_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC22_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC22_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC22_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC22_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC22_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC22_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC22_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC23_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC23_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC23_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC23_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC23_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC23_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC23_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC23_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC23_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC23_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC23_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC23_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC23_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC23_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC23_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC23_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC23_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC23_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC23_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC23_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC23_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC24_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC24_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC24_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC24_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC24_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC24_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC24_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC24_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC24_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC24_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC24_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC24_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC24_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC24_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC24_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC24_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC24_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC24_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC24_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC24_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC24_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC25_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC25_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC25_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC25_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC25_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC25_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC25_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC25_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC25_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC25_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC25_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC25_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC25_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC25_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC25_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC25_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC25_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC25_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC25_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC25_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC25_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC26_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC26_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC26_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC26_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC26_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC26_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC26_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC26_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC26_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC26_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC26_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC26_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC26_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC26_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC26_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC26_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC26_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC26_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC26_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC26_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC26_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC27_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC27_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC27_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC27_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC27_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC27_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC27_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC27_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC27_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC27_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC27_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC27_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC27_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC27_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC27_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC27_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC27_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC27_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC27_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC27_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC27_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC28_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC28_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC28_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC28_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC28_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC28_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC28_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC28_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC28_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC28_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC28_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC28_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC28_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC28_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC28_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC28_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC28_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC28_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC28_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC28_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC28_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC29_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC29_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC29_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC29_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC29_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC29_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC29_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC29_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC29_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC29_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC29_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC29_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC29_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC29_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC29_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC29_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC29_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC29_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC29_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC29_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC29_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC30_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC30_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC30_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC30_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC30_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC30_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC30_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC30_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC30_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC30_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC30_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC30_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC30_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC30_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC30_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC30_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC30_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC30_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC30_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC30_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC30_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC31_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC31_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC31_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC31_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC31_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC31_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC31_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC31_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC31_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC31_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC31_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC31_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC31_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC31_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC31_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC31_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC31_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC31_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC31_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC31_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC31_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC32_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC32_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC32_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC32_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC32_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC32_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC32_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC32_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC32_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC32_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC32_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC32_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC32_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC32_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC32_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC32_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC32_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC32_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC32_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC32_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC32_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC33_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC33_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC33_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC33_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC33_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC33_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC33_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC33_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC33_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC33_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC33_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC33_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC33_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC33_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC33_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC33_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC33_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC33_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC33_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC33_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC33_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC34_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC34_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC34_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC34_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC34_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC34_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC34_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC34_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC34_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC34_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC34_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC34_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC34_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC34_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC34_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC34_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC34_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC34_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC34_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC34_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC34_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC35_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC35_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC35_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC35_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC35_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC35_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC35_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC35_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC35_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC35_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC35_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC35_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC35_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC35_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC35_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC35_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC35_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC35_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC35_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC35_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC35_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC36_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC36_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC36_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC36_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC36_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC36_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC36_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC36_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC36_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC36_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC36_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC36_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC36_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC36_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC36_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC36_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC36_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC36_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC36_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC36_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC36_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC37_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC37_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC37_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC37_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC37_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC37_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC37_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC37_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC37_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC37_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC37_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC37_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC37_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC37_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC37_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC37_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC37_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC37_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC37_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC37_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC37_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC38_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC38_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC38_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC38_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC38_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC38_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC38_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC38_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC38_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC38_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC38_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC38_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC38_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC38_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC38_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC38_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC38_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC38_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC38_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC38_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC38_OUT_SEL_CFG.Reg) & 0x800) >> 11
}

// GPIO.FUNC39_OUT_SEL_CFG
func (o *GPIO_Type) SetFUNC39_OUT_SEL_CFG_OUT_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC39_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC39_OUT_SEL_CFG.Reg)&^(0x1ff)|value)
}
func (o *GPIO_Type) GetFUNC39_OUT_SEL_CFG_OUT_SEL() uint32 {
	return volatile.LoadUint32(&o.FUNC39_OUT_SEL_CFG.Reg) & 0x1ff
}
func (o *GPIO_Type) SetFUNC39_OUT_SEL_CFG_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC39_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC39_OUT_SEL_CFG.Reg)&^(0x200)|value<<9)
}
func (o *GPIO_Type) GetFUNC39_OUT_SEL_CFG_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC39_OUT_SEL_CFG.Reg) & 0x200) >> 9
}
func (o *GPIO_Type) SetFUNC39_OUT_SEL_CFG_OEN_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC39_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC39_OUT_SEL_CFG.Reg)&^(0x400)|value<<10)
}
func (o *GPIO_Type) GetFUNC39_OUT_SEL_CFG_OEN_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC39_OUT_SEL_CFG.Reg) & 0x400) >> 10
}
func (o *GPIO_Type) SetFUNC39_OUT_SEL_CFG_OEN_INV_SEL(value uint32) {
	volatile.StoreUint32(&o.FUNC39_OUT_SEL_CFG.Reg, volatile.LoadUint32(&o.FUNC39_OUT_SEL_CFG.Reg)&^(0x800)|value<<11)
}
func (o *GPIO_Type) GetFUNC39_OUT_SEL_CFG_OEN_INV_SEL() uint32 {
	return (volatile.LoadUint32(&o.FUNC39_OUT_SEL_CFG.Reg) & 0x800) >> 11
}
