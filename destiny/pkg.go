package destiny

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/davecgh/go-spew/spew"

	aes "bits.chrsm.org/x/destiny/aes2"
)

type Package struct {
	r io.ReadSeeker

	Header  *Header
	Blocks  []*Block
	Entries []*Entry
	Nonce   []byte
}

func New(r io.ReadSeeker) *Package {
	pkg := &Package{
		r: r,

		Header: &Header{},
	}

	pkg.header()
	pkg.entries()
	pkg.blocks()
	pkg.Nonce = aes.Nonce(uint(pkg.Header.PackageID))

	return pkg
}

func (pkg *Package) Dump() {
	spew.Dump(pkg)
}

func (pkg *Package) header() error {
	// all manual from now, possibly impl struct tag to help out here
	brd := func(v interface{}) {
		if err := binary.Read(pkg.r, binary.LittleEndian, v); err != nil {
			panic(fmt.Sprintf("header(): could not read: %s", err))
		}
	}

	h := &Header{}

	pkg.r.Seek(0, io.SeekStart)
	brd(&h.Version)
	brd(&h.Platform)
	brd(&h.PackageID)

	pkg.r.Seek(0x10, io.SeekStart)
	brd(&h.BuildDate)

	pkg.r.Seek(0x1A, io.SeekStart)
	brd(&h.Type)

	pkg.r.Seek(0x20, io.SeekStart)
	brd(&h.PatchID)
	brd(&h.LangID)

	if h.Type != typeNew {
		pkg.r.Seek(0xB4, io.SeekStart)
		brd(&h.EntryTableSize)
		brd(&h.EntryTableOffset)

		pkg.r.Seek(0xD0, io.SeekStart)
		brd(&h.BlockTableSize)
		brd(&h.BlockTableOffset)
	} else {
		var ofs, oofs uint32

		pkg.r.Seek(0x110, io.SeekStart)
		brd(&ofs)

		pkg.r.Seek(int64(ofs+0x10), io.SeekStart)
		brd(&h.EntryTableSize)

		pkg.r.Seek(0x04, io.SeekCurrent)
		brd(&oofs)
		h.EntryTableOffset = ofs + oofs + 0x28

		pkg.r.Seek(0x110, io.SeekStart)
		brd(&ofs) // num

		pkg.r.Seek(int64(ofs+0x20), io.SeekStart)
		brd(&h.BlockTableSize)

		pkg.r.Seek(0x04, io.SeekCurrent)
		brd(&oofs)

		h.BlockTableOffset = ofs + oofs + 0x38
	}

	pkg.Header = h

	return nil
}

func (pkg *Package) blocks() {
	brd := func(v interface{}) {
		if err := binary.Read(pkg.r, binary.LittleEndian, v); err != nil {
			panic(fmt.Sprintf("blocks(): could not read: %s", err))
		}
	}

	pkg.Blocks = make([]*Block, pkg.Header.BlockTableSize)
	pkg.r.Seek(int64(pkg.Header.BlockTableOffset), io.SeekStart)
	for i := 0; i < int(pkg.Header.BlockTableSize); i++ {
		b := &Block{}

		brd(&b.Offset)
		brd(&b.Size)
		brd(&b.PatchID)
		brd(&b.Flags)

		pkg.r.Read(b.Hash[:])
		pkg.r.Read(b.GCMTag[:])

		pkg.Blocks[i] = b
	}
}

func (pkg *Package) entries() {
	brd := func(v interface{}) {
		if err := binary.Read(pkg.r, binary.LittleEndian, v); err != nil {
			panic(fmt.Sprintf("entries(): could not read: %s", err))
		}
	}

	pkg.Entries = make([]*Entry, pkg.Header.EntryTableSize)
	pkg.r.Seek(int64(pkg.Header.EntryTableOffset), io.SeekStart)
	for i := 0; i < int(pkg.Header.EntryTableSize); i++ {
		e := &Entry{}

		brd(&e.Unknown32)
		brd(&e.Type)
		brd(&e.Unknown_B)

		e.StartBlock = uint64(uint32(e.Unknown_B & 0x3FFF))
		e.Offset = uint64(pkg.Header.BlockTableOffset) + e.StartBlock*0x30
		e.Size = (e.Unknown_B >> 0x1C & 0x3FFFFFFF)

		ibo := (e.Unknown_B >> 0x0E & 0x3FFF) * 0x10

		e._BlockCount = (ibo + e.Size + blockSize - 1) / blockSize

		pkg.Entries[i] = e
	}
}
