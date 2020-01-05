package aes2

import (
	"crypto/aes"
	"crypto/cipher"
)

var (
	AesKey0 = []byte{
		0xD6, 0x2A, 0xB2, 0xC1, 0x0C, 0xC0, 0x1B, 0xC5, 0x35, 0xDB, 0x7B, 0x86, 0x55, 0xC7, 0xDC, 0x3B,
	}

	AesKey1 = []byte{
		0x3A, 0x4A, 0x5D, 0x36, 0x73, 0xA6, 0x60, 0x58, 0x7E, 0x63, 0xE6, 0x76, 0xE4, 0x08, 0x92, 0xB5,
	}
)

func Decrypt(src, nonce, gcmtag, key []byte) []byte {
	auth, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	cng, err := cipher.NewGCM(auth)
	if err != nil {
		panic(err)
	}

	next := make([]byte, len(src)+len(gcmtag))
	copy(next, src)
	copy(next[len(src):], gcmtag)

	data, err := cng.Open(nil, nonce, next, nil)
	if err != nil {
		panic(err)
	}

	return data
}

const blockCount = 0x40000

var nonce = []byte{
	0x84, 0xDF, 0x11, 0xC0,
	0xAC, 0xAB, 0xFA, 0x20,
	0x33, 0x11, 0x26, 0x99,
}

func Nonce(pkgID uint) []byte {
	pkgnonce := make([]byte, len(nonce))
	copy(pkgnonce, nonce)

	pkgnonce[11] ^= uint8(pkgID & 0xFF)
	pkgnonce[1] ^= 0x26
	pkgnonce[0] ^= uint8(pkgID >> 8 & 0xFF)

	return pkgnonce
}
