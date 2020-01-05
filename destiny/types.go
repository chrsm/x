package destiny

const (
	typeStd = 0x00
	typeNew = 0x01
)

type Header struct {
	// only shit we care about, see SPkgHeader.txt for full list
	Version   uint16 // 0x00
	Platform  uint16 // 0x02
	PackageID uint16 // 0x04
	BuildDate uint32 // 0x10
	Type      uint16 // 0x1A
	PatchID   uint16 // 0x20
	LangID    uint16 // 0x22

	EntryTableOffset uint32 // type != 1: 0xB8; type == 1: read 4 bytes after size, add 0x28 and result from size
	EntryTableSize   uint32 // type != 1: 0xB4; type == 1: read 0x110, add 0x10

	BlockTableOffset uint32 // type != 1: 0xD0; type == 1: read 4 bytes after size, add 0x38 and result from size
	BlockTableSize   uint32 // type != 1: 0xD4; type == 1: (4byte 0x110) + 0x20
}

const (
	blockEncrypted       = 0x01
	blockCompressed      = 0x02
	blockEncryptedAltKey = 0x04
)

type Block struct {
	Offset  uint32
	Size    uint32
	PatchID uint16
	Flags   uint16
	Hash    [0x14]byte
	GCMTag  [0x10]byte
}

func (b *Block) IsCompressed() bool {
	return b.Flags&blockCompressed != 0
}

func (b *Block) IsEncrypted() bool {
	return b.Flags&blockEncrypted != 0
}

func (b *Block) IsEncryptedAlt() bool {
	return b.Flags&blockEncryptedAltKey != 0
}

type Entry struct {
	Unknown32 uint32
	Type      uint32

	Unknown_B  uint64 // bit array
	StartBlock uint64 // bits:14
	Offset     uint64 // bits:14
	Size       uint64 // bits:30
	Unknown    uint64 // bits:6

	_BlockCount uint64
}

const blockSize = 0x40000

func (e *Entry) BlockCount() uint64 {
	return e._BlockCount // precomputed
	//((e.Offset * 0x10) + e.Size + blockSize - 1) / blockSize
}

func (e *Entry) Contains() string {
	if e.Unknown32 == EntryTypeText || e.Type == EntryTypeText {
		return "Text"
	}

	if e.Unknown32 == EntryTypeTexts || e.Type == EntryTypeTexts {
		return "Texts"
	}

	return "unk"
}

const (
	// need to figure out what the rest of the types are..
	EntryTypeText  = 0x80809A8A
	EntryTypeTexts = 0x80809A88
)
