# QDISASM

Quick disassembler for x86 architecture using [x86asm][1] package.

Why - because sometimes you get the urge to disassemble things.

## Usage
```bash
qdisasm [-bits 16|32|64] [binary]
```
if no binary is given, it reads from stdin.

Sample output:
```asm
[7:42:55] 0:~/Downloads/grp_plugin> 7z x grp.wcx CODE -so | qdisasm -org 0x100 |head
0000:0100 c8 00 00 00                   enter 0x0, 0x0
0000:0104 33 c0                         xor eax, eax
0000:0106 40                            inc eax
0000:0107 c9                            leave
0000:0108 c2 0c 00                      ret 0xc
0000:010B 00 c8                         add al, cl
0000:010D 00 00                         add byte ptr [eax], al
0000:010F 00 51 55                      add byte ptr [ecx+0x55], dl
0000:0112 56                            push esi
0000:0113 57                            push edi
```

[x86asm][1] is MIT-licensed: Copyright 2015 The Go Authors.

[1]: golang.org/x/arch/x86/x86asm 
