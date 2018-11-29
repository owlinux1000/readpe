package header

import (
    "fmt"
    "debug/pe"
)

type ImageOptionalHeader64 pe.OptionalHeader64

func (h ImageOptionalHeader64) String() (result string) {
    result = fmt.Sprintf(`Optional Header:
  Magic:                           %s
  Major linker version:            %d
  Minor linker version:            %d
  Size of code:                    %d (bytes)
  Size of initialized data:        %d (bytes)
  Size of uninitialized data:      %d (bytes)
  Address of entry point:          0x%x
  Base of code:                    0x%x
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
