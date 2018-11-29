package header

import (
    "fmt"
    "encoding/binary"
)

type ImageDosHeader struct {
    E_magic uint16
    E_cblp uint16
    E_cp uint16;
    E_crlc uint16
    E_cparhdr uint16
    Eminalloc uint16
    E_maxalloc uint16
    E_ss uint16
    E_sp uint16
    E_csum uint16
    Eip uint16
    E_cs uint16
    E_lfarlc uint16
    E_ovno uint16
    E_res []uint16
    E_oemid uint16
    E_oeminfo uint16
    E_res2 []uint16
    E_lfanew uint32
}
type PImageDosHeader *ImageDosHeader

func NewImageDosHeader(data []byte) *ImageDosHeader {
    image_dos_header := new(ImageDosHeader)
    image_dos_header.Parse(data)
    return image_dos_header
}

func (h *ImageDosHeader) Parse(data []byte) {
    h.E_magic = binary.LittleEndian.Uint16(data[0:2])
    h.E_cblp = binary.LittleEndian.Uint16(data[2:4])
    h.E_cp = binary.LittleEndian.Uint16(data[4:6])
    h.E_crlc = binary.LittleEndian.Uint16(data[6:8])
    h.E_cparhdr = binary.LittleEndian.Uint16(data[8:10])
    h.Eminalloc = binary.LittleEndian.Uint16(data[10:12])
    h.E_maxalloc = binary.LittleEndian.Uint16(data[12:14])
    h.E_ss = binary.LittleEndian.Uint16(data[14:16])
    h.E_sp = binary.LittleEndian.Uint16(data[16:18])
    h.E_csum = binary.LittleEndian.Uint16(data[18:20])
    h.Eip = binary.LittleEndian.Uint16(data[20:22])
    h.E_cs = binary.LittleEndian.Uint16(data[22:24])
    h.E_lfarlc = binary.LittleEndian.Uint16(data[24:26])
    h.E_ovno = binary.LittleEndian.Uint16(data[26:28])
    for i := 0; i < 8; i+=2 {
        h.E_res = append(
            h.E_res,
            binary.LittleEndian.Uint16(data[28+i:30+i]),
        )
    }
    
    h.E_oemid = binary.LittleEndian.Uint16(data[36:38])
    h.E_oeminfo = binary.LittleEndian.Uint16(data[38:40])
    for i := 0; i < 20; i+=2 {
        h.E_res2 = append(
            h.E_res2,
            binary.LittleEndian.Uint16(data[40+i:42+i]),
        )
    }
    h.E_lfanew = binary.LittleEndian.Uint32(data[60:64])
}

func (h ImageDosHeader) String() (result string){
    result = fmt.Sprintf(`DOS Header:
  Magic:                                 %s
  Bytes on last page of file:            %d (bytes)
  Pages in file:                         %d
  Relocations:                           %d
  Minimum extra paragraph needed:        %d
  Maximum extra paragraph needed:        %d
  Initial (relative) SS value:           %d
  Initial SP value:                      %d
  Checksum:                              %d
  Initial IP Value:                      0x%x
  Initial (relative) CS value:           %d
  File address of relocation table:      0x%x
  Overlay number:                        %d
  Reserved words:                        %s
  OEM identifier:                        %d
  OEM information:                       %d
  Reserved words:                        %s
  File address of new exe header:        0x%x
`,
        reverse(decodeStringFromHex(h.E_magic)),
        h.E_cblp,
        h.E_cp,
        h.E_crlc,
        h.Eminalloc,
        h.E_maxalloc,
        h.E_ss,
        h.E_sp,
        h.E_csum,
        h.Eip,
        h.E_cs,
        h.E_lfarlc,
        h.E_ovno,
        hexString(h.E_res),
        h.E_oemid,
        h.E_oeminfo,
        hexString(h.E_res2),
        h.E_lfanew,
    )
    return result
}
