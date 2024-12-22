package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"text/tabwriter"

	"golang.org/x/arch/x86/x86asm"
)

var (
	bits = flag.Int("bits", 32, "16, 32 or 64")
	org  = flag.Uint64("org", 0x0, "set origin")
)

func main() {
	flag.Parse()

	if *bits != 16 && *bits != 32 && *bits != 64 {
		flag.Usage()
		log.Fatalf("invalid bits: %d", *bits)
	}

	in := os.Stdin
	if len(flag.Args()) > 0 {
		f, err := os.Open(flag.Args()[0])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		in = f
	}

	CODE, err := io.ReadAll(in)
	if err != nil {
		log.Fatal(err)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer tw.Flush()
	var pc uint64 = *org
	for len(CODE) > 0 {
		instr, err := x86asm.Decode(CODE, *bits)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(tw, "%04X:%04X\t% x\t%s\n", pc>>16, pc&0x0000FFFF, CODE[:instr.Len], x86asm.IntelSyntax(instr, pc, nil))
		CODE = CODE[instr.Len:]
		pc += uint64(instr.Len)
	}
}
