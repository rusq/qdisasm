# QDISASM

Quick disassembler for x86 architecture using [x86asm][1] package.

Why - because sometimes you get the urge to disassemble things.

## Usage
```bash
qdisasm [-bits 16|32|64] [binary]
```
if no binary is given, it reads from stdin.

[x86asm][1] is MIT-licensed: Copyright 2015 The Go Authors.

[1]: golang.org/x/arch/x86/x86asm 
