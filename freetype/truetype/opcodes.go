// Copyright 2012 The Freetype-Go Authors. All rights reserved.
// Use of this source code is governed by your choice of either the
// FreeType License or the GNU General Public License version 2 (or
// any later version), both of which can be found in the LICENSE file.

package truetype

// The Truetype opcodes are summarized at https://developer.apple.com/fonts/TTRefMan/RM07/appendixA.html

// TODO: the opXXX constants without end-of-line comments are not yet implemented.

const (
	opSVTCA0    = 0x00
	opSVTCA1    = 0x01
	opSPVTCA0   = 0x02
	opSPVTCA1   = 0x03
	opSFVTCA0   = 0x04
	opSFVTCA1   = 0x05
	opSPVTL0    = 0x06
	opSPVTL1    = 0x07
	opSFVTL0    = 0x08
	opSFVTL1    = 0x09
	opSFVFS     = 0x0b
	opSPVFS     = 0x0a
	opGPV       = 0x0c
	opGFV       = 0x0d
	opSFVTPV    = 0x0e
	opISECT     = 0x0f
	opSRP0      = 0x10
	opSRP1      = 0x11
	opSRP2      = 0x12
	opSZP0      = 0x13
	opSZP1      = 0x14
	opSZP2      = 0x15
	opSZPS      = 0x16
	opSLOOP     = 0x17
	opRTG       = 0x18
	opRTHG      = 0x19
	opSMD       = 0x1a
	opELSE      = 0x1b // ELSE clause
	opJMPR      = 0x1c
	opSCVTCI    = 0x1d
	opSSWCI     = 0x1e
	opSSW       = 0x1f
	opDUP       = 0x20 // DUPlicate top stack element
	opPOP       = 0x21 // POP top stack element
	opCLEAR     = 0x22 // CLEAR the stack
	opSWAP      = 0x23 // SWAP the top two elements on the stack
	opDEPTH     = 0x24 // DEPTH of the stack
	opCINDEX    = 0x25 // Copy the INDEXed element to the top of the stack
	opMINDEX    = 0x26 // Move the INDEXed element to the top of the stack
	opALIGNPTS  = 0x27
	op_0x28     = 0x28
	opUTP       = 0x29
	opLOOPCALL  = 0x2a
	opCALL      = 0x2b
	opFDEF      = 0x2c
	opENDF      = 0x2d
	opMDAP0     = 0x2e
	opMDAP1     = 0x2f
	opIUP0      = 0x30
	opIUP1      = 0x31
	opSHP0      = 0x32
	opSHP1      = 0x33
	opSHC0      = 0x34
	opSHC1      = 0x35
	opSHZ0      = 0x36
	opSHZ1      = 0x37
	opSHPIX     = 0x38
	opIP        = 0x39
	opMSIRP0    = 0x3a
	opMSIRP1    = 0x3b
	opALIGNRP   = 0x3c
	opRTGD      = 0x3d
	opMIAP0     = 0x3e
	opMIAP1     = 0x3f
	opNPUSHB    = 0x40 // PUSH N Bytes
	opNPUSHW    = 0x41 // PUSH N Words
	opWS        = 0x42
	opRS        = 0x43
	opWCVTP     = 0x44
	opRCVT      = 0x45
	opGC0       = 0x46
	opGC1       = 0x47
	opSCFS      = 0x48
	opMD0       = 0x49
	opMD1       = 0x4a
	opMPPEM     = 0x4b
	opMPS       = 0x4c
	opFLIPON    = 0x4d
	opFLIPOFF   = 0x4e
	opDEBUG     = 0x4f // DEBUG call
	opLT        = 0x50 // Less Than
	opLTEQ      = 0x51 // Less Than or EQual
	opGT        = 0x52 // Greater Than
	opGTEQ      = 0x53 // Greater Than or EQual
	opEQ        = 0x54 // EQual
	opNEQ       = 0x55 // Not EQual
	opODD       = 0x56
	opEVEN      = 0x57
	opIF        = 0x58 // IF test
	opEIF       = 0x59 // End IF
	opAND       = 0x5a // logical AND
	opOR        = 0x5b // logical OR
	opNOT       = 0x5c // logical NOT
	opDELTAP1   = 0x5d
	opSDB       = 0x5e
	opSDS       = 0x5f
	opADD       = 0x60 // ADD
	opSUB       = 0x61 // SUBtract
	opDIV       = 0x62 // DIVide
	opMUL       = 0x63 // MULtiply
	opABS       = 0x64 // ABSolute value
	opNEG       = 0x65 // NEGate
	opFLOOR     = 0x66 // FLOOR
	opCEILING   = 0x67 // CEILING
	opROUND00   = 0x68
	opROUND01   = 0x69
	opROUND10   = 0x6a
	opROUND11   = 0x6b
	opNROUND00  = 0x6c
	opNROUND01  = 0x6d
	opNROUND10  = 0x6e
	opNROUND11  = 0x6f
	opWCVTF     = 0x70
	opDELTAP2   = 0x71
	opDELTAP3   = 0x72
	opDELTAC1   = 0x73
	opDELTAC2   = 0x74
	opDELTAC3   = 0x75
	opSROUND    = 0x76
	opS45ROUND  = 0x77
	opJROT      = 0x78
	opJROF      = 0x79
	opROFF      = 0x7a
	op_0x7b     = 0x7b
	opRUTG      = 0x7c
	opRDTG      = 0x7d
	opSANGW     = 0x7e
	opAA        = 0x7f
	opFLIPPT    = 0x80
	opFLIPRGON  = 0x81
	opFLIPRGOFF = 0x82
	op_0x83     = 0x83
	op_0x84     = 0x84
	opSCANCTRL  = 0x85
	opSDPVTL0   = 0x86
	opSDPVTL1   = 0x87
	opGETINFO   = 0x88
	opIFDEF     = 0x89
	opROLL      = 0x8a
	opMAX       = 0x8b
	opMIN       = 0x8c
	opSCANTYPE  = 0x8d
	opINSTCTRL  = 0x8e
	op_0x8f     = 0x8f
	op_0x90     = 0x90
	op_0x91     = 0x91
	op_0x92     = 0x92
	op_0x93     = 0x93
	op_0x94     = 0x94
	op_0x95     = 0x95
	op_0x96     = 0x96
	op_0x97     = 0x97
	op_0x98     = 0x98
	op_0x99     = 0x99
	op_0x9a     = 0x9a
	op_0x9b     = 0x9b
	op_0x9c     = 0x9c
	op_0x9d     = 0x9d
	op_0x9e     = 0x9e
	op_0x9f     = 0x9f
	op_0xa0     = 0xa0
	op_0xa1     = 0xa1
	op_0xa2     = 0xa2
	op_0xa3     = 0xa3
	op_0xa4     = 0xa4
	op_0xa5     = 0xa5
	op_0xa6     = 0xa6
	op_0xa7     = 0xa7
	op_0xa8     = 0xa8
	op_0xa9     = 0xa9
	op_0xaa     = 0xaa
	op_0xab     = 0xab
	op_0xac     = 0xac
	op_0xad     = 0xad
	op_0xae     = 0xae
	op_0xaf     = 0xaf
	opPUSHB000  = 0xb0 // PUSH Bytes
	opPUSHB001  = 0xb1 // .
	opPUSHB010  = 0xb2 // .
	opPUSHB011  = 0xb3 // .
	opPUSHB100  = 0xb4 // .
	opPUSHB101  = 0xb5 // .
	opPUSHB110  = 0xb6 // .
	opPUSHB111  = 0xb7 // .
	opPUSHW000  = 0xb8 // PUSH Words
	opPUSHW001  = 0xb9 // .
	opPUSHW010  = 0xba // .
	opPUSHW011  = 0xbb // .
	opPUSHW100  = 0xbc // .
	opPUSHW101  = 0xbd // .
	opPUSHW110  = 0xbe // .
	opPUSHW111  = 0xbf // .
	opMDRP00000 = 0xc0
	opMDRP00001 = 0xc1
	opMDRP00010 = 0xc2
	opMDRP00011 = 0xc3
	opMDRP00100 = 0xc4
	opMDRP00101 = 0xc5
	opMDRP00110 = 0xc6
	opMDRP00111 = 0xc7
	opMDRP01000 = 0xc8
	opMDRP01001 = 0xc9
	opMDRP01010 = 0xca
	opMDRP01011 = 0xcb
	opMDRP01100 = 0xcc
	opMDRP01101 = 0xcd
	opMDRP01110 = 0xce
	opMDRP01111 = 0xcf
	opMDRP10000 = 0xd0
	opMDRP10001 = 0xd1
	opMDRP10010 = 0xd2
	opMDRP10011 = 0xd3
	opMDRP10100 = 0xd4
	opMDRP10101 = 0xd5
	opMDRP10110 = 0xd6
	opMDRP10111 = 0xd7
	opMDRP11000 = 0xd8
	opMDRP11001 = 0xd9
	opMDRP11010 = 0xda
	opMDRP11011 = 0xdb
	opMDRP11100 = 0xdd
	opMDRP11101 = 0xdc
	opMDRP11110 = 0xde
	opMDRP11111 = 0xdf
	opMIRP00000 = 0xe0
	opMIRP00001 = 0xe1
	opMIRP00010 = 0xe2
	opMIRP00011 = 0xe3
	opMIRP00100 = 0xe4
	opMIRP00101 = 0xe5
	opMIRP00110 = 0xe6
	opMIRP00111 = 0xe7
	opMIRP01000 = 0xe8
	opMIRP01001 = 0xe9
	opMIRP01010 = 0xea
	opMIRP01011 = 0xeb
	opMIRP01100 = 0xec
	opMIRP01101 = 0xed
	opMIRP01110 = 0xee
	opMIRP01111 = 0xef
	opMIRP10000 = 0xf0
	opMIRP10001 = 0xf1
	opMIRP10010 = 0xf2
	opMIRP10011 = 0xf3
	opMIRP10100 = 0xf4
	opMIRP10101 = 0xf5
	opMIRP10110 = 0xf6
	opMIRP10111 = 0xf7
	opMIRP11000 = 0xf8
	opMIRP11001 = 0xf9
	opMIRP11010 = 0xfa
	opMIRP11011 = 0xfb
	opMIRP11100 = 0xfd
	opMIRP11101 = 0xfc
	opMIRP11110 = 0xfe
	opMIRP11111 = 0xff
)

// popCount is the number of stack elements that each opcode pops.
var popCount = [256]uint8{
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, a, b, c, d, e, f
	q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, // 0x00 - 0x0f
	q, q, q, q, q, q, q, q, q, q, q, 0, q, q, q, q, // 0x10 - 0x1f
	1, 1, 0, 2, 0, 1, 1, q, q, q, q, q, q, q, q, q, // 0x20 - 0x2f
	q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, // 0x30 - 0x3f
	0, 0, q, q, q, q, q, q, q, q, q, q, q, q, q, 0, // 0x40 - 0x4f
	2, 2, 2, 2, 2, 2, q, q, 1, 0, 2, 2, 1, q, q, q, // 0x50 - 0x5f
	2, 2, 2, 2, 1, 1, 1, 1, q, q, q, q, q, q, q, q, // 0x60 - 0x6f
	q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, // 0x70 - 0x7f
	q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, // 0x80 - 0x8f
	q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, // 0x90 - 0x9f
	q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, // 0xa0 - 0xaf
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, // 0xb0 - 0xbf
	q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, // 0xc0 - 0xcf
	q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, // 0xd0 - 0xdf
	q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, // 0xe0 - 0xef
	q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, q, // 0xf0 - 0xff
}

// popCount[opcode] == q means that that opcode is not yet implemented.
const q = 255
