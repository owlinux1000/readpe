package header

import (
    "fmt"
    "strings"
    "debug/pe"
)

type ImageOptionalHeader32 pe.OptionalHeader32 

func (h ImageOptionalHeader32) String() (result string) {
    result = fmt.Sprintf(`Optional Header:
  Magic:                           %s
  Major linker version:            %d
  Minor linker version:            %d
  Size of code:                    %d (bytes)
  Size of initialized data:        %d (bytes)
  Size of uninitialized data:      %d (bytes)
  Address of entry point:          0x%x
  Base of code:                    0x%x
  Base of data:                    0x%x
  Image base:                      0x%x
  Section alignment:               0x%x
  File alignment:                  0x%x
  Major operating system version:  %d
  Minor operating system version:  %d
  Major image version:             %d
  Minor image version:             %d
  Major subsystem version:         %d
  Minor subsystem version:         %d
  Win32 version value:             %d
  Size of image:                   %d (bytes)
  Size of headers:                 %d (bytes)
  CheckSum:                        %d
  Subsystem:                       %s
  Dll characteristics:             %s
  Size of stack reserve:           %d (bytes)
  Size of stack commit:            %d (bytes)
  Size of heap reserve:            %d (bytes)
  Size of heap commit:             %d (bytes)
  Loader flags:                    %d
  Number of RVA and sizes:         %d
`,
        descByMagic(h.Magic),
        h.MajorLinkerVersion,
        h.MinorLinkerVersion,
        h.SizeOfCode,
        h.SizeOfInitializedData,
        h.SizeOfUninitializedData,
        h.AddressOfEntryPoint,
        h.BaseOfCode,
        h.BaseOfData,
        h.ImageBase,
        h.SectionAlignment,
        h.FileAlignment,
        h.MajorOperatingSystemVersion,
        h.MinorOperatingSystemVersion,
        h.MajorImageVersion,
        h.MinorImageVersion,
        h.MajorSubsystemVersion,
        h.MinorSubsystemVersion,
        h.Win32VersionValue,
        h.SizeOfImage,
        h.SizeOfHeaders,
        h.CheckSum,
        descBySubsystem(h.Subsystem),
        descByDllCharacteristics(h.DllCharacteristics),
        h.SizeOfStackReserve,
        h.SizeOfStackCommit,
        h.SizeOfHeapReserve,
        h.SizeOfHeapCommit,
        h.LoaderFlags,
        h.NumberOfRvaAndSizes,
    )
    return result
}

func descBySubsystem(subsystem uint16) (desc string) {
    switch subsystem {
    case 0:
    default:
        desc = "Unknown"
    case 1:
        desc = "Native"
    case 2:
        desc = "Windows GUI"
    case 3:
        desc = "Windows CUI"
    case 5:
        desc = "OS/2 CUI"
    case 7:
        desc = "POSIX CUI"
    case 9:
        desc = "Windows CE GUI"
    case 10:
        desc = "EFI application"
    case 11:
        desc = "EFI boot service driver"
    case 12:
        desc = "EFI runtime driver"
    case 13:
        desc = "EFI ROM image"
    case 14:
        desc = "Xbox"
    case 16:
        desc = "Windows boot applicatoin"
    }
    return desc
}

func descByDllCharacteristics(characteristics uint16) (desc string) {
    d := []string{}
    if (characteristics & 0x40) == 0x40 {
        d = append(d, "DYNAMIC_BASE")
    }
    if (characteristics & 0x80) == 0x80 {
        d = append(d, "FORCE_INTEGRITY")
    }
    if (characteristics & 0x100) == 0x100 {
        d = append(d, "NX_COMPAT")
    }
    if (characteristics & 0x200) == 0x200 {
        d = append(d, "NO_ISOLATION")
    }
    if (characteristics & 0x400) == 0x400 {
        d = append(d, "NO_SEH")
    }
    if (characteristics & 0x800) == 0x800 {
        d = append(d, "NO_BIND")
    }
    if (characteristics & 0x2000) == 0x2000 {
        d = append(d, "WDM_DRIVER")
    }
    if (characteristics & 0x8000) == 0x8000 {
        d = append(d, "TERMINAL_SERVER_AWARE")
    }

    if len(d) == 0 {
        return fmt.Sprintf("%d", characteristics)
    }
    return strings.Join(d, " | ")
}

func descByMagic(magic uint16) (desc string) {
    switch magic {
    case 0x10b:
        desc = "PE32"
    case 0x20b:
        desc = "PE+"
    case 0x107:
        desc = "ROM"
    }
    return desc
}
