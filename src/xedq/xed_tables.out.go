// Code generated by xed_tables.go. DO NOT EDIT.

package xedq

var registerByName = map[string]xedRegister{
	"INVALID":    0,
	"BNDCFGU":    1,
	"BNDSTATUS":  2,
	"BND0":       3,
	"BND1":       4,
	"BND2":       5,
	"BND3":       6,
	"CR0":        7,
	"CR1":        8,
	"CR2":        9,
	"CR3":        10,
	"CR4":        11,
	"CR5":        12,
	"CR6":        13,
	"CR7":        14,
	"CR8":        15,
	"CR9":        16,
	"CR10":       17,
	"CR11":       18,
	"CR12":       19,
	"CR13":       20,
	"CR14":       21,
	"CR15":       22,
	"DR0":        23,
	"DR1":        24,
	"DR2":        25,
	"DR3":        26,
	"DR4":        27,
	"DR5":        28,
	"DR6":        29,
	"DR7":        30,
	"DR8":        31,
	"DR9":        32,
	"DR10":       33,
	"DR11":       34,
	"DR12":       35,
	"DR13":       36,
	"DR14":       37,
	"DR15":       38,
	"FLAGS":      39,
	"EFLAGS":     40,
	"RFLAGS":     41,
	"AX":         42,
	"CX":         43,
	"DX":         44,
	"BX":         45,
	"SP":         46,
	"BP":         47,
	"SI":         48,
	"DI":         49,
	"R8W":        50,
	"R9W":        51,
	"R10W":       52,
	"R11W":       53,
	"R12W":       54,
	"R13W":       55,
	"R14W":       56,
	"R15W":       57,
	"EAX":        58,
	"ECX":        59,
	"EDX":        60,
	"EBX":        61,
	"ESP":        62,
	"EBP":        63,
	"ESI":        64,
	"EDI":        65,
	"R8D":        66,
	"R9D":        67,
	"R10D":       68,
	"R11D":       69,
	"R12D":       70,
	"R13D":       71,
	"R14D":       72,
	"R15D":       73,
	"RAX":        74,
	"RCX":        75,
	"RDX":        76,
	"RBX":        77,
	"RSP":        78,
	"RBP":        79,
	"RSI":        80,
	"RDI":        81,
	"R8":         82,
	"R9":         83,
	"R10":        84,
	"R11":        85,
	"R12":        86,
	"R13":        87,
	"R14":        88,
	"R15":        89,
	"AL":         90,
	"CL":         91,
	"DL":         92,
	"BL":         93,
	"SPL":        94,
	"BPL":        95,
	"SIL":        96,
	"DIL":        97,
	"R8B":        98,
	"R9B":        99,
	"R10B":       100,
	"R11B":       101,
	"R12B":       102,
	"R13B":       103,
	"R14B":       104,
	"R15B":       105,
	"AH":         106,
	"CH":         107,
	"DH":         108,
	"BH":         109,
	"ERROR":      110,
	"RIP":        111,
	"EIP":        112,
	"IP":         113,
	"K0":         114,
	"K1":         115,
	"K2":         116,
	"K3":         117,
	"K4":         118,
	"K5":         119,
	"K6":         120,
	"K7":         121,
	"MMX0":       122,
	"MMX1":       123,
	"MMX2":       124,
	"MMX3":       125,
	"MMX4":       126,
	"MMX5":       127,
	"MMX6":       128,
	"MMX7":       129,
	"SSP":        130,
	"IA32_U_CET": 131,
	"MXCSR":      132,
	"STACKPUSH":  133,
	"STACKPOP":   134,
	"GDTR":       135,
	"LDTR":       136,
	"IDTR":       137,
	"TR":         138,
	"TSC":        139,
	"TSCAUX":     140,
	"MSRS":       141,
	"FSBASE":     142,
	"GSBASE":     143,
	"X87CONTROL": 144,
	"X87STATUS":  145,
	"X87TAG":     146,
	"X87PUSH":    147,
	"X87POP":     148,
	"X87POP2":    149,
	"X87OPCODE":  150,
	"X87LASTCS":  151,
	"X87LASTIP":  152,
	"X87LASTDS":  153,
	"X87LASTDP":  154,
	"CS":         155,
	"DS":         156,
	"ES":         157,
	"SS":         158,
	"FS":         159,
	"GS":         160,
	"TMP0":       161,
	"TMP1":       162,
	"TMP2":       163,
	"TMP3":       164,
	"TMP4":       165,
	"TMP5":       166,
	"TMP6":       167,
	"TMP7":       168,
	"TMP8":       169,
	"TMP9":       170,
	"TMP10":      171,
	"TMP11":      172,
	"TMP12":      173,
	"TMP13":      174,
	"TMP14":      175,
	"TMP15":      176,
	"st(0)":      177,
	"st(1)":      178,
	"st(2)":      179,
	"st(3)":      180,
	"st(4)":      181,
	"st(5)":      182,
	"st(6)":      183,
	"st(7)":      184,
	"XCR0":       185,
	"XMM0":       186,
	"XMM1":       187,
	"XMM2":       188,
	"XMM3":       189,
	"XMM4":       190,
	"XMM5":       191,
	"XMM6":       192,
	"XMM7":       193,
	"XMM8":       194,
	"XMM9":       195,
	"XMM10":      196,
	"XMM11":      197,
	"XMM12":      198,
	"XMM13":      199,
	"XMM14":      200,
	"XMM15":      201,
	"XMM16":      202,
	"XMM17":      203,
	"XMM18":      204,
	"XMM19":      205,
	"XMM20":      206,
	"XMM21":      207,
	"XMM22":      208,
	"XMM23":      209,
	"XMM24":      210,
	"XMM25":      211,
	"XMM26":      212,
	"XMM27":      213,
	"XMM28":      214,
	"XMM29":      215,
	"XMM30":      216,
	"XMM31":      217,
	"YMM0":       218,
	"YMM1":       219,
	"YMM2":       220,
	"YMM3":       221,
	"YMM4":       222,
	"YMM5":       223,
	"YMM6":       224,
	"YMM7":       225,
	"YMM8":       226,
	"YMM9":       227,
	"YMM10":      228,
	"YMM11":      229,
	"YMM12":      230,
	"YMM13":      231,
	"YMM14":      232,
	"YMM15":      233,
	"YMM16":      234,
	"YMM17":      235,
	"YMM18":      236,
	"YMM19":      237,
	"YMM20":      238,
	"YMM21":      239,
	"YMM22":      240,
	"YMM23":      241,
	"YMM24":      242,
	"YMM25":      243,
	"YMM26":      244,
	"YMM27":      245,
	"YMM28":      246,
	"YMM29":      247,
	"YMM30":      248,
	"YMM31":      249,
	"ZMM0":       250,
	"ZMM1":       251,
	"ZMM2":       252,
	"ZMM3":       253,
	"ZMM4":       254,
	"ZMM5":       255,
	"ZMM6":       256,
	"ZMM7":       257,
	"ZMM8":       258,
	"ZMM9":       259,
	"ZMM10":      260,
	"ZMM11":      261,
	"ZMM12":      262,
	"ZMM13":      263,
	"ZMM14":      264,
	"ZMM15":      265,
	"ZMM16":      266,
	"ZMM17":      267,
	"ZMM18":      268,
	"ZMM19":      269,
	"ZMM20":      270,
	"ZMM21":      271,
	"ZMM22":      272,
	"ZMM23":      273,
	"ZMM24":      274,
	"ZMM25":      275,
	"ZMM26":      276,
	"ZMM27":      277,
	"ZMM28":      278,
	"ZMM29":      279,
	"ZMM30":      280,
	"ZMM31":      281,
}
