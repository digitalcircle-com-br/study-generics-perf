package main

//GENERICS

/*
GOOS=linux GOARCH=amd64 go tool compile -S asm/generics/main.go
"".main STEXT size=66 args=0x0 locals=0x18 funcid=0x0 align=0x0
        0x0000 00000 (asm/generics/main.go:7)   TEXT    "".main(SB), ABIInternal, $24-0
        0x0000 00000 (asm/generics/main.go:7)   CMPQ    SP, 16(R14)
        0x0004 00004 (asm/generics/main.go:7)   PCDATA  $0, $-2
        0x0004 00004 (asm/generics/main.go:7)   JLS     58
        0x0006 00006 (asm/generics/main.go:7)   PCDATA  $0, $-1
        0x0006 00006 (asm/generics/main.go:7)   SUBQ    $24, SP
        0x000a 00010 (asm/generics/main.go:7)   MOVQ    BP, 16(SP)
        0x000f 00015 (asm/generics/main.go:7)   LEAQ    16(SP), BP
        0x0014 00020 (asm/generics/main.go:7)   FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0014 00020 (asm/generics/main.go:7)   FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0014 00020 (asm/generics/main.go:8)   XCHGL   AX, AX
        0x0015 00021 (asm/generics/main.go:8)   PCDATA  $1, $0
        0x0015 00021 (asm/generics/main.go:4)   CALL    runtime.printlock(SB)
        0x001a 00026 (asm/generics/main.go:4)   LEAQ    go.string."TEST"(SB), AX
        0x0021 00033 (asm/generics/main.go:4)   MOVL    $4, BX
        0x0026 00038 (asm/generics/main.go:4)   CALL    runtime.printstring(SB)
        0x002b 00043 (asm/generics/main.go:4)   CALL    runtime.printunlock(SB)
        0x0030 00048 (asm/generics/main.go:9)   MOVQ    16(SP), BP
        0x0035 00053 (asm/generics/main.go:9)   ADDQ    $24, SP
        0x0039 00057 (asm/generics/main.go:9)   RET
        0x003a 00058 (asm/generics/main.go:9)   NOP
        0x003a 00058 (asm/generics/main.go:7)   PCDATA  $1, $-1
        0x003a 00058 (asm/generics/main.go:7)   PCDATA  $0, $-2
        0x003a 00058 (asm/generics/main.go:7)   CALL    runtime.morestack_noctxt(SB)
        0x003f 00063 (asm/generics/main.go:7)   PCDATA  $0, $-1
        0x003f 00063 (asm/generics/main.go:7)   NOP
        0x0040 00064 (asm/generics/main.go:7)   JMP     0
        0x0000 49 3b 66 10 76 34 48 83 ec 18 48 89 6c 24 10 48  I;f.v4H...H.l$.H
        0x0010 8d 6c 24 10 90 e8 00 00 00 00 48 8d 05 00 00 00  .l$.......H.....
        0x0020 00 bb 04 00 00 00 e8 00 00 00 00 e8 00 00 00 00  ................
        0x0030 48 8b 6c 24 10 48 83 c4 18 c3 e8 00 00 00 00 90  H.l$.H..........
        0x0040 eb be                                            ..
        rel 22+4 t=7 runtime.printlock+0
        rel 29+4 t=14 go.string."TEST"+0
        rel 39+4 t=7 runtime.printstring+0
        rel 44+4 t=7 runtime.printunlock+0
        rel 59+4 t=7 runtime.morestack_noctxt+0
"".Print[go.shape.string_0] STEXT dupok size=104 args=0x18 locals=0x18 funcid=0x0 align=0x0
        0x0000 00000 (asm/generics/main.go:3)   TEXT    "".Print[go.shape.string_0](SB), DUPOK|ABIInternal, $24-24
        0x0000 00000 (asm/generics/main.go:3)   CMPQ    SP, 16(R14)
        0x0004 00004 (asm/generics/main.go:3)   PCDATA  $0, $-2
        0x0004 00004 (asm/generics/main.go:3)   JLS     67
        0x0006 00006 (asm/generics/main.go:3)   PCDATA  $0, $-1
        0x0006 00006 (asm/generics/main.go:3)   SUBQ    $24, SP
        0x000a 00010 (asm/generics/main.go:3)   MOVQ    BP, 16(SP)
        0x000f 00015 (asm/generics/main.go:3)   LEAQ    16(SP), BP
        0x0014 00020 (asm/generics/main.go:3)   FUNCDATA        $0, gclocals·09cf9819fc716118c209c2d2155a3632(SB)
        0x0014 00020 (asm/generics/main.go:3)   FUNCDATA        $1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
        0x0014 00020 (asm/generics/main.go:3)   FUNCDATA        $5, "".Print[go.shape.string_0].arginfo1(SB)
        0x0014 00020 (asm/generics/main.go:3)   FUNCDATA        $6, "".Print[go.shape.string_0].argliveinfo(SB)
        0x0014 00020 (asm/generics/main.go:3)   PCDATA  $3, $1
        0x0014 00020 (asm/generics/main.go:3)   MOVQ    CX, "".t+48(SP)
        0x0019 00025 (asm/generics/main.go:3)   MOVQ    BX, "".t+40(SP)
        0x001e 00030 (asm/generics/main.go:3)   PCDATA  $3, $2
        0x001e 00030 (asm/generics/main.go:4)   PCDATA  $1, $0
        0x001e 00030 (asm/generics/main.go:4)   NOP
        0x0020 00032 (asm/generics/main.go:4)   CALL    runtime.printlock(SB)
        0x0025 00037 (asm/generics/main.go:4)   MOVQ    "".t+40(SP), AX
        0x002a 00042 (asm/generics/main.go:4)   MOVQ    "".t+48(SP), BX
        0x002f 00047 (asm/generics/main.go:4)   PCDATA  $1, $1
        0x002f 00047 (asm/generics/main.go:4)   CALL    runtime.printstring(SB)
        0x0034 00052 (asm/generics/main.go:4)   CALL    runtime.printunlock(SB)
        0x0039 00057 (asm/generics/main.go:4)   MOVQ    16(SP), BP
        0x003e 00062 (asm/generics/main.go:4)   ADDQ    $24, SP
        0x0042 00066 (asm/generics/main.go:4)   RET
        0x0043 00067 (asm/generics/main.go:4)   NOP
        0x0043 00067 (asm/generics/main.go:3)   PCDATA  $1, $-1
        0x0043 00067 (asm/generics/main.go:3)   PCDATA  $0, $-2
        0x0043 00067 (asm/generics/main.go:3)   MOVQ    AX, 8(SP)
        0x0048 00072 (asm/generics/main.go:3)   MOVQ    BX, 16(SP)
        0x004d 00077 (asm/generics/main.go:3)   MOVQ    CX, 24(SP)
        0x0052 00082 (asm/generics/main.go:3)   CALL    runtime.morestack_noctxt(SB)
        0x0057 00087 (asm/generics/main.go:3)   MOVQ    8(SP), AX
        0x005c 00092 (asm/generics/main.go:3)   MOVQ    16(SP), BX
        0x0061 00097 (asm/generics/main.go:3)   MOVQ    24(SP), CX
        0x0066 00102 (asm/generics/main.go:3)   PCDATA  $0, $-1
        0x0066 00102 (asm/generics/main.go:3)   JMP     0
        0x0000 49 3b 66 10 76 3d 48 83 ec 18 48 89 6c 24 10 48  I;f.v=H...H.l$.H
        0x0010 8d 6c 24 10 48 89 4c 24 30 48 89 5c 24 28 66 90  .l$.H.L$0H.\$(f.
        0x0020 e8 00 00 00 00 48 8b 44 24 28 48 8b 5c 24 30 e8  .....H.D$(H.\$0.
        0x0030 00 00 00 00 e8 00 00 00 00 48 8b 6c 24 10 48 83  .........H.l$.H.
        0x0040 c4 18 c3 48 89 44 24 08 48 89 5c 24 10 48 89 4c  ...H.D$.H.\$.H.L
        0x0050 24 18 e8 00 00 00 00 48 8b 44 24 08 48 8b 5c 24  $......H.D$.H.\$
        0x0060 10 48 8b 4c 24 18 eb 98                          .H.L$...
        rel 33+4 t=7 runtime.printlock+0
        rel 48+4 t=7 runtime.printstring+0
        rel 53+4 t=7 runtime.printunlock+0
        rel 83+4 t=7 runtime.morestack_noctxt+0
""..dict.Print[string] SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 type.string+0
        rel 0+0 t=23 type.string+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
        0x0000 6d 61 69 6e                                      main
go.info."".Print[go.shape.string_0]$abstract SDWARFABSFCN dupok size=38
        0x0000 05 2e 50 72 69 6e 74 5b 67 6f 2e 73 68 61 70 65  ..Print[go.shape
        0x0010 2e 73 74 72 69 6e 67 5f 30 5d 00 01 01 13 74 00  .string_0]....t.
        0x0020 00 00 00 00 00 00                                ......
        rel 0+0 t=22 type.go.shape.string_0+0
        rel 33+4 t=31 go.info.go.shape.string_0+0
""..inittask SNOPTRDATA size=24
        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0010 00 00 00 00 00 00 00 00                          ........
go.string."TEST" SRODATA dupok size=4
        0x0000 54 45 53 54                                      TEST
runtime.strequal·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.strequal+0
runtime.memequal64·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
        0x0000 01                                               .
type..namedata.*go.shape.string_0- SRODATA dupok size=20
        0x0000 00 12 2a 67 6f 2e 73 68 61 70 65 2e 73 74 72 69  ..*go.shape.stri
        0x0010 6e 67 5f 30                                      ng_0
type.*go.shape.string_0 SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 54 6f 54 9a 08 08 08 36 00 00 00 00 00 00 00 00  ToT....6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*go.shape.string_0-+0
        rel 48+8 t=1 type.go.shape.string_0+0
type..importpath.go.shape. SRODATA dupok size=10
        0x0000 00 08 67 6f 2e 73 68 61 70 65                    ..go.shape
type.go.shape.string_0 SRODATA dupok size=64
        0x0000 10 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 b4 93 c9 2b 07 08 08 18 00 00 00 00 00 00 00 00  ...+............
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
        rel 24+8 t=1 runtime.strequal·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*go.shape.string_0-+0
        rel 44+4 t=5 type.*go.shape.string_0+0
        rel 48+4 t=5 type..importpath.go.shape.+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·09cf9819fc716118c209c2d2155a3632 SRODATA dupok size=10
        0x0000 02 00 00 00 02 00 00 00 02 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
        0x0000 02 00 00 00 00 00 00 00                          ........
"".Print[go.shape.string_0].arginfo1 SRODATA static dupok size=7
        0x0000 fe 08 08 10 08 fd ff                             .......
"".Print[go.shape.string_0].argliveinfo SRODATA static dupok size=3
        0x0000 00 00 06
*/
func Print[T any](t T) {
	print(t)
}

func main() {
	Print("TEST")
}
