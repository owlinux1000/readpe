package header

import (
    "fmt"
    "time"
    "strings"
    "debug/pe"
)
type ImageFileHeader pe.FileHeader

func (h ImageFileHeader) String() (result string) {
    result = fmt.Sprintf(`FILE Header:
  Machine:                   %s
  Number of Sections:        %d
  Time Date Stamp:           %s
  Pointer of Symbol Table:   0x%x
  Number of Symbols:         %d
  Size of Option Header:     %d (bytes)
  Characteristics:           %s`,
        archByMachine(h.Machine),
        h.NumberOfSections,
        timeDateFormat(h.TimeDateStamp),
        h.PointerToSymbolTable,
        h.NumberOfSymbols,
        h.SizeOfOptionalHeader,
        descByCharacteristics(h.Characteristics),
        
    )
    return result
}

func descByCharacteristics(characteristics uint16) (desc string) {
    d := []string{}
    if (characteristics & 0x1) == 0x1 {
        d = append(d, "RELOCS_STRIPPED")
    }
    if (characteristics & 0x2) == 0x2 {
        d = append(d, "EXECUTABLE_IMAGE")
    }
    if (characteristics & 0x4) == 0x4 {
        d = append(d, "LINE_NUMS_STRIPPED")
    }
    if (characteristics & 0x8) == 0x8 {
        d = append(d, "LOCAL_SYMS_STRIPPED")
    }
    if (characteristics & 0x10) == 0x10 {
        d = append(d, "AGGRESIVE_WS_TRIM")
    }
    if (characteristics & 0x20) == 0x20 {
        d = append(d, "LARGE_ADDRESS_AWARE")
    }    
    if (characteristics & 0x80) == 0x80 {
        d = append(d, "BYTES_REVERSED_LO")
    }
    if (characteristics & 0x100) == 0x100 {
        d = append(d, "32BIT_MACHINE")
    }
    if (characteristics & 0x200) == 0x200 {
        d = append(d, "DEBUG_STRIPPED")
    }
    if (characteristics & 0x400) == 0x400 {
        d = append(d, "REMOVABLE_RUN_FROM_SWAP")
    }
    if len(d) == 0 {
        return fmt.Sprintf("%d", characteristics)
    }
    return strings.Join(d, " | ")
}

func archByMachine(machine uint16) (arch string) {
    switch machine {
    case 0x14c:
        arch = "Intel 386"
    case 0x8664:
        arch = "x64"
    case 0x162:
        arch = "MIPS R3000"
    case 0x168:
        arch = "MIPS R1000"
    case 0x169:
        arch = "MIPS little endian WCI v2"
    case 0x183:
        arch = "old Alpha AXP"
    case 0x184:
        arch = "Alpha AXP"
    case 0x1a2:
        arch = "Hitachi SH3"
    case 0x1a3:
        arch = "Hitachi SH3 DSP"
    case 0x1a6:
        arch = "Hitachi SH4"
    case 0x1a8:
        arch = "Hitachi SH5"
    case 0x1c0:
        arch = "ARM little endian"
    case 0x1c2:
        arch = "Thumb"
    case 0x1c4:
        arch = "ARMv7"
    case 0x1d3:
        arch = "Matsushita AM33"
    case 0x1f0:
        arch = "PowerPC little endian"
    case 0x1f1:
        arch = "PowerPC with floating point support"
    case 0x200:
        arch = "Intel IA64"
    case 0x266:
        arch = "MIPS16"
    case 0x268:
        arch = "Motorola 68000 series"
    case 0x284:
        arch = "Alpha AXP 64-bit"
    case 0x366:
        arch = "MIPS with FPU"
    case 0x466:
        arch = "MIPS16 with FPU"
    case 0xebc:
        arch = "EFI Byte Code"
    case 0x9041:
        arch = "Mitsubishi M32R little endian"
    case 0xaa64:
        arch = "ARM64 little endian"
    case 0xc0ee:
        arch = "clr pure MSIL"
    default:
        arch = "unknown"
    }
    return arch
}

func timeDateFormat(timeDateStamp uint32) string {
    return fmt.Sprintf("%s", time.Unix(int64(timeDateStamp), 0))
}
