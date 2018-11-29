package header

import (
    "fmt"
    "strings"
    "debug/pe"
)

type ImageSectionHeader pe.SectionHeader

func (h ImageSectionHeader) String() (result string) {
    result = fmt.Sprintf(`Section Header:
  Name:                     %s
  Virtual size:             %d (bytes)
  Virtual address:          0x%x
  Size of raw data:         %d (bytes)
  Pointer to raw data:      %d (bytes)
  Pointer to relocations:   0x%x
  Pointer to line numbers:  0x%x
  Number of relocations:    %d
  Number of line numbers:   %d
  Characteristics:          %s
`,
        h.Name,
        h.VirtualSize,
        h.VirtualAddress,
        h.Size,
        h.Offset,
        h.PointerToRelocations,
        h.PointerToLineNumbers,
        h.NumberOfRelocations,
        h.NumberOfLineNumbers,
        descBySectionCharacteristics(h.Characteristics),
    )
    return result
}

func descBySectionCharacteristics(characteristics uint32) (desc string) {
    d := []string{}
    if (characteristics & 0x8) == 0x8 {
        d = append(d, "TYPE_NO_PAD")
    }
    if (characteristics & 0x20) == 0x20 {
        d = append(d, "CNT_CODE")
    }
    if (characteristics & 0x40) == 0x40 {
        d = append(d, "CNT_INITIALIZED_DATA")
    }
    if (characteristics & 0x80) == 0x80 {
        d = append(d, "CNT_UNINITIALIZED_DATA")
    }
    if (characteristics & 0x200) == 0x200 {
        d = append(d, "LNK_INFO")
    }
    if (characteristics & 0x800) == 0x800 {
        d = append(d, "LNK_REMOVE")
    }
    if (characteristics & 0x1000) == 0x1000 {
        d = append(d, "LNK_COMDAT")
    }
    if (characteristics & 0x4000) == 0x4000 {
        d = append(d, "NO_DEFER_SPEC_EXC")
    }
    if (characteristics & 0x8000) == 0x8000 {
        d = append(d, "GPREL")
    }
    if (characteristics & 0x100000) == 0x100000 {
        d = append(d, "ALIGN_1BYTES")
    }
    if (characteristics & 0x200000) == 0x200000 {
        d = append(d, "ALIGN_2BYTES")
    }
    if (characteristics & 0x300000) == 0x300000 {
        d = append(d, "ALIGN_4BYTES")
    }
    if (characteristics & 0x400000) == 0x400000 {
        d = append(d, "ALIGN_8BYTES")
    }
    if (characteristics & 0x500000) == 0x500000 {
        d = append(d, "ALIGN_16BYTES")
    }
    if (characteristics & 0x600000) == 0x600000 {
        d = append(d, "ALIGN_32BYTES")
    }
    if (characteristics & 0x700000) == 0x700000 {
        d = append(d, "ALIGN_64BYTES")
    }
    if (characteristics & 0x800000) == 0x800000 {
        d = append(d, "ALIGN_128BYTES")
    }
    if (characteristics & 0x900000) == 0x900000 {
        d = append(d, "ALIGN_256BYTES")
    }
    if (characteristics & 0xa00000) == 0xa00000 {
        d = append(d, "ALIGN_512BYTES")
    }
    if (characteristics & 0xb00000) == 0xb00000 {
        d = append(d, "ALIGN_1024BYTES")
    }
    if (characteristics & 0xc00000) == 0xc00000 {
        d = append(d, "ALIGN_2048BYTES")
    }
    if (characteristics & 0xd00000) == 0xd00000 {
        d = append(d, "ALIGN_4096BYTES")
    }
    if (characteristics & 0xe00000) == 0xe00000 {
        d = append(d, "ALIGN_8192BYTES")
    }
    if (characteristics & 0x1000000) == 0x1000000 {
        d = append(d, "LNK_NRELOC_OVFL")
    }
    if (characteristics & 0x2000000) == 0x2000000 {
        d = append(d, "MEM_DISCARDABLE")
    }
    if (characteristics & 0x4000000) == 0x4000000 {
        d = append(d, "MEM_NOT_CACHED")
    }
    if (characteristics & 0x8000000) == 0x8000000 {
        d = append(d, "MEM_NOT_PAGED")
    }
    if (characteristics & 0x10000000) == 0x10000000 {
        d = append(d, "MEM_SHARED")
    }
    if (characteristics & 0x20000000) == 0x20000000 {
        d = append(d, "MEM_EXECUTE")
    }
    if (characteristics & 0x40000000) == 0x40000000 {
        d = append(d, "MEM_READ")
    }
    if (characteristics & 0x80000000) == 0x80000000 {
        d = append(d, "MEM_WRITE")
    }
    return strings.Join(d, " | ")
}
