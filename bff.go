package bff

import (
	"io"
	"strings"
)

// Magic is the magic number of BFF files.
const Magic uint32 = 0x09006BEA

// FileHeader is at the start of a BFF file.
//
// Size: 0x48
type FileHeader struct {
	Magic  uint32  // [0x00] =0x09006BEA
	Unk04  uint32  // [0x04]
	Time08 uint32  // [0x08]
	Time0C uint32  // [0x0C]
	Unk10  uint32  // [0x10]
	Unk14  [8]byte // [0x14] "by name\0"
	Unk1C  uint32  // [0x1C]
	Unk20  uint32  // [0x20]
	Unk24  [8]byte // [0x24] "by name\0"
	Unk2C  uint32  // [0x2C]
	Unk30  uint32  // [0x30]
	Unk34  [8]byte // [0x24] "BUILD\0\0\0"
	Unk3C  uint32  // [0x3C]
	Unk40  uint32  // [0x40]
	Unk44  uint32  // [0x44]
}

// RecordHeader is the first structure of a record.
// It precedes the record name (variable-length string).
//
// Size: 0x40
type RecordHeader struct {
	Unk00  uint32 // [0x00]
	Unk04  uint32 // [0x04]
	Unk08  uint32 // [0x08]
	Unk0C  uint32 // [0x0C]
	Unk10  uint32 // [0x10]
	Unk14  uint32 // [0x14]
	Unk18  uint32 // [0x18] Mask of file size?
	Time1C uint32 // [0x1C]
	Time20 uint32 // [0x20]
	Time24 uint32 // [0x24]
	Unk28  uint32 // [0x28]
	Unk2C  uint32 // [0x2C]
	Unk30  uint32 // [0x30]
	Unk34  uint32 // [0x34]
	Size   uint32 // [0x38] File size
	Unk3C  uint32 // [0x3C]
}

func ReadAlignedString(rd io.Reader) (string, error) {
	var str strings.Builder
	for {
		var line [8]byte
		if _, err := io.ReadFull(rd, line[:]); err != nil {
			return "", err
		}
		for _, c := range line {
			if c == 0 {
				return str.String(), nil
			}
			str.WriteByte(c)
		}
	}
}

type RecordTrailer struct {
	Unk00 uint32 // [0x00]
	Unk04 uint32 // [0x04]
	Unk08 uint32 // [0x08]
	Unk0C uint32 // [0x0C]
	Unk10 uint32 // [0x10]
	Unk14 uint32 // [0x14]
	Unk18 uint32 // [0x18]
	Unk1C uint32 // [0x1C]
	Unk20 uint32 // [0x20]
	Unk24 uint32 // [0x24]
}
