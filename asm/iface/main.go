package main

//IFACE

/*
GOOS=linux GOARCH=amd64 go tool compile -S asm/iface/main.go
"".Print STEXT size=148 args=0x10 locals=0x30 funcid=0x0 align=0x0
        0x0000 00000 (asm/iface/main.go:153)    TEXT    "".Print(SB), ABIInternal, $48-16
        0x0000 00000 (asm/iface/main.go:153)    CMPQ    SP, 16(R14)
        0x0004 00004 (asm/iface/main.go:153)    PCDATA  $0, $-2
        0x0004 00004 (asm/iface/main.go:153)    JLS     115
        0x0006 00006 (asm/iface/main.go:153)    PCDATA  $0, $-1
        0x0006 00006 (asm/iface/main.go:153)    SUBQ    $48, SP
        0x000a 00010 (asm/iface/main.go:153)    MOVQ    BP, 40(SP)
        0x000f 00015 (asm/iface/main.go:153)    LEAQ    40(SP), BP
        0x0014 00020 (asm/iface/main.go:153)    MOVQ    AX, "".t+56(FP)
        0x0019 00025 (asm/iface/main.go:153)    MOVQ    BX, "".t+64(FP)
        0x001e 00030 (asm/iface/main.go:153)    FUNCDATA        $0, gclocals·ba30782f8935b28ed1adaec603e72627(SB)
        0x001e 00030 (asm/iface/main.go:153)    FUNCDATA        $1, gclocals·663f8c6bfa83aa777198789ce63d9ab4(SB)
        0x001e 00030 (asm/iface/main.go:153)    FUNCDATA        $5, "".Print.arginfo1(SB)
        0x001e 00030 (asm/iface/main.go:153)    FUNCDATA        $6, "".Print.argliveinfo(SB)
        0x001e 00030 (asm/iface/main.go:153)    PCDATA  $3, $1
        0x001e 00030 (asm/iface/main.go:154)    LEAQ    type.string(SB), DX
        0x0025 00037 (asm/iface/main.go:154)    CMPQ    AX, DX
        0x0028 00040 (asm/iface/main.go:154)    JNE     99
        0x002a 00042 (asm/iface/main.go:154)    MOVQ    (BX), AX
        0x002d 00045 (asm/iface/main.go:154)    MOVQ    AX, ""..autotmp_2+32(SP)
        0x0032 00050 (asm/iface/main.go:154)    MOVQ    8(BX), CX
        0x0036 00054 (asm/iface/main.go:154)    MOVQ    CX, ""..autotmp_3+24(SP)
        0x003b 00059 (asm/iface/main.go:154)    PCDATA  $1, $1
        0x003b 00059 (asm/iface/main.go:154)    NOP
        0x0040 00064 (asm/iface/main.go:154)    CALL    runtime.printlock(SB)
        0x0045 00069 (asm/iface/main.go:154)    MOVQ    ""..autotmp_2+32(SP), AX
        0x004a 00074 (asm/iface/main.go:154)    MOVQ    ""..autotmp_3+24(SP), BX
        0x004f 00079 (asm/iface/main.go:154)    PCDATA  $1, $2
        0x004f 00079 (asm/iface/main.go:154)    CALL    runtime.printstring(SB)
        0x0054 00084 (asm/iface/main.go:154)    CALL    runtime.printunlock(SB)
        0x0059 00089 (asm/iface/main.go:155)    MOVQ    40(SP), BP
        0x005e 00094 (asm/iface/main.go:155)    ADDQ    $48, SP
        0x0062 00098 (asm/iface/main.go:155)    RET
        0x0063 00099 (asm/iface/main.go:154)    MOVQ    DX, BX
        0x0066 00102 (asm/iface/main.go:154)    LEAQ    type.interface {}(SB), CX
        0x006d 00109 (asm/iface/main.go:154)    CALL    runtime.panicdottypeE(SB)
        0x0072 00114 (asm/iface/main.go:154)    XCHGL   AX, AX
        0x0073 00115 (asm/iface/main.go:154)    NOP
        0x0073 00115 (asm/iface/main.go:153)    PCDATA  $1, $-1
        0x0073 00115 (asm/iface/main.go:153)    PCDATA  $0, $-2
        0x0073 00115 (asm/iface/main.go:153)    MOVQ    AX, 8(SP)
        0x0078 00120 (asm/iface/main.go:153)    MOVQ    BX, 16(SP)
        0x007d 00125 (asm/iface/main.go:153)    NOP
        0x0080 00128 (asm/iface/main.go:153)    CALL    runtime.morestack_noctxt(SB)
        0x0085 00133 (asm/iface/main.go:153)    MOVQ    8(SP), AX
        0x008a 00138 (asm/iface/main.go:153)    MOVQ    16(SP), BX
        0x008f 00143 (asm/iface/main.go:153)    PCDATA  $0, $-1
        0x008f 00143 (asm/iface/main.go:153)    JMP     0
        0x0000 49 3b 66 10 76 6d 48 83 ec 30 48 89 6c 24 28 48  I;f.vmH..0H.l$(H
        0x0010 8d 6c 24 28 48 89 44 24 38 48 89 5c 24 40 48 8d  .l$(H.D$8H.\$@H.
        0x0020 15 00 00 00 00 48 39 d0 75 39 48 8b 03 48 89 44  .....H9.u9H..H.D
        0x0030 24 20 48 8b 4b 08 48 89 4c 24 18 0f 1f 44 00 00  $ H.K.H.L$...D..
        0x0040 e8 00 00 00 00 48 8b 44 24 20 48 8b 5c 24 18 e8  .....H.D$ H.\$..
        0x0050 00 00 00 00 e8 00 00 00 00 48 8b 6c 24 28 48 83  .........H.l$(H.
        0x0060 c4 30 c3 48 89 d3 48 8d 0d 00 00 00 00 e8 00 00  .0.H..H.........
        0x0070 00 00 90 48 89 44 24 08 48 89 5c 24 10 0f 1f 00  ...H.D$.H.\$....
        0x0080 e8 00 00 00 00 48 8b 44 24 08 48 8b 5c 24 10 e9  .....H.D$.H.\$..
        0x0090 6c ff ff ff                                      l...
        rel 33+4 t=14 type.string+0
        rel 65+4 t=7 runtime.printlock+0
        rel 80+4 t=7 runtime.printstring+0
        rel 85+4 t=7 runtime.printunlock+0
        rel 105+4 t=14 type.interface {}+0
        rel 110+4 t=7 runtime.panicdottypeE+0
        rel 129+4 t=7 runtime.morestack_noctxt+0
"".main STEXT size=91 args=0x0 locals=0x28 funcid=0x0 align=0x0
        0x0000 00000 (asm/iface/main.go:157)    TEXT    "".main(SB), ABIInternal, $40-0
        0x0000 00000 (asm/iface/main.go:157)    CMPQ    SP, 16(R14)
        0x0004 00004 (asm/iface/main.go:157)    PCDATA  $0, $-2
        0x0004 00004 (asm/iface/main.go:157)    JLS     84
        0x0006 00006 (asm/iface/main.go:157)    PCDATA  $0, $-1
        0x0006 00006 (asm/iface/main.go:157)    SUBQ    $40, SP
        0x000a 00010 (asm/iface/main.go:157)    MOVQ    BP, 32(SP)
        0x000f 00015 (asm/iface/main.go:157)    LEAQ    32(SP), BP
        0x0014 00020 (asm/iface/main.go:157)    FUNCDATA        $0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
        0x0014 00020 (asm/iface/main.go:157)    FUNCDATA        $1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
        0x0014 00020 (<unknown line number>)    NOP
        0x0014 00020 (asm/iface/main.go:158)    MOVQ    ""..stmp_0(SB), AX
        0x001b 00027 (asm/iface/main.go:154)    MOVQ    AX, ""..autotmp_4+24(SP)
        0x0020 00032 (asm/iface/main.go:154)    MOVQ    ""..stmp_0+8(SB), CX
        0x0027 00039 (asm/iface/main.go:154)    MOVQ    CX, ""..autotmp_5+16(SP)
        0x002c 00044 (asm/iface/main.go:154)    PCDATA  $1, $1
        0x002c 00044 (asm/iface/main.go:154)    CALL    runtime.printlock(SB)
        0x0031 00049 (asm/iface/main.go:154)    MOVQ    ""..autotmp_4+24(SP), AX
        0x0036 00054 (asm/iface/main.go:154)    MOVQ    ""..autotmp_5+16(SP), BX
        0x003b 00059 (asm/iface/main.go:154)    PCDATA  $1, $0
        0x003b 00059 (asm/iface/main.go:154)    NOP
        0x0040 00064 (asm/iface/main.go:154)    CALL    runtime.printstring(SB)
        0x0045 00069 (asm/iface/main.go:154)    CALL    runtime.printunlock(SB)
        0x004a 00074 (asm/iface/main.go:159)    MOVQ    32(SP), BP
        0x004f 00079 (asm/iface/main.go:159)    ADDQ    $40, SP
        0x0053 00083 (asm/iface/main.go:159)    RET
        0x0054 00084 (asm/iface/main.go:159)    NOP
        0x0054 00084 (asm/iface/main.go:157)    PCDATA  $1, $-1
        0x0054 00084 (asm/iface/main.go:157)    PCDATA  $0, $-2
        0x0054 00084 (asm/iface/main.go:157)    CALL    runtime.morestack_noctxt(SB)
        0x0059 00089 (asm/iface/main.go:157)    PCDATA  $0, $-1
        0x0059 00089 (asm/iface/main.go:157)    JMP     0
        0x0000 49 3b 66 10 76 4e 48 83 ec 28 48 89 6c 24 20 48  I;f.vNH..(H.l$ H
        0x0010 8d 6c 24 20 48 8b 05 00 00 00 00 48 89 44 24 18  .l$ H......H.D$.
        0x0020 48 8b 0d 00 00 00 00 48 89 4c 24 10 e8 00 00 00  H......H.L$.....
        0x0030 00 48 8b 44 24 18 48 8b 5c 24 10 0f 1f 44 00 00  .H.D$.H.\$...D..
        0x0040 e8 00 00 00 00 e8 00 00 00 00 48 8b 6c 24 20 48  ..........H.l$ H
        0x0050 83 c4 28 c3 e8 00 00 00 00 eb a5                 ..(........
        rel 2+0 t=23 type.string+0
        rel 23+4 t=14 ""..stmp_0+0
        rel 35+4 t=14 ""..stmp_0+8
        rel 45+4 t=7 runtime.printlock+0
        rel 65+4 t=7 runtime.printstring+0
        rel 70+4 t=7 runtime.printunlock+0
        rel 85+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
        0x0000 6d 61 69 6e                                      main
go.info."".Print$abstract SDWARFABSFCN dupok size=19
        0x0000 05 2e 50 72 69 6e 74 00 01 01 13 74 00 00 00 00  ..Print....t....
        0x0010 00 00 00                                         ...
        rel 0+0 t=22 type.interface {}+0
        rel 14+4 t=31 go.info.interface {}+0
""..inittask SNOPTRDATA size=24
        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0010 00 00 00 00 00 00 00 00                          ........
go.string."TEST" SRODATA dupok size=4
        0x0000 54 45 53 54                                      TEST
""..stmp_0 SRODATA static size=16
        0x0000 00 00 00 00 00 00 00 00 04 00 00 00 00 00 00 00  ................
        rel 0+8 t=1 go.string."TEST"+0
runtime.nilinterequal·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.nilinterequal+0
runtime.memequal64·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
        0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=15
        0x0000 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d     ..*interface {}
type.*interface {} SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 4f 0f 96 9d 08 08 08 36 00 00 00 00 00 00 00 00  O......6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*interface {}-+0
        rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
        0x0000 02                                               .
type.interface {} SRODATA dupok size=80
        0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
        0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        rel 24+8 t=1 runtime.nilinterequal·f+0
        rel 32+8 t=1 runtime.gcbits.02+0
        rel 40+4 t=5 type..namedata.*interface {}-+0
        rel 44+4 t=-32763 type.*interface {}+0
        rel 56+8 t=1 type.interface {}+80
gclocals·ba30782f8935b28ed1adaec603e72627 SRODATA dupok size=11
        0x0000 03 00 00 00 02 00 00 00 02 00 00                 ...........
gclocals·663f8c6bfa83aa777198789ce63d9ab4 SRODATA dupok size=11
        0x0000 03 00 00 00 01 00 00 00 00 01 00                 ...........
"".Print.arginfo1 SRODATA static dupok size=7
        0x0000 fe 00 08 08 08 fd ff                             .......
"".Print.argliveinfo SRODATA static dupok size=2
        0x0000 00 00                                            ..
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
        0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
        0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
*/
func Print(t interface{}) {
	print(t.(string))
}

func main() {
	Print("TEST")
}
