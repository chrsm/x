// +build windows

package oodle

import (
	"log"
	"syscall"
	"unsafe"
)

const CanDecompress = true

var (
	oodle = syscall.NewLazyDLL("oo2core_3_win64.dll")

	/*pOodleLZDecoderCreate           = oodle.NewProc("OodleLZDecoder_Create")
	pOodleLZDecoderDestroy          = oodle.NewProc("OodleLZDecoder_Destroy")
	pOodleLZDecoderDecodeSome       = oodle.NewProc("OodleLZDecoder_DecodeSome")
	pOodleLZDecoderMemorySizeNeeded = oodle.NewProc("OodleLZDecoder_MemorySizeNeeded")*/

	pOodleLZCompress   = oodle.NewProc("OodleLZ_Compress")
	pOodleLZDecompress = oodle.NewProc("OodleLZ_Decompress")
)

func lzcompress(src []byte) []byte {
	// dst must be big enough
	dst := make([]byte, len(src)+274*((len(src)+0x3FFFF)/0x40000))

	ret, x1, x2 := pOodleLZCompress.Call(
		0,                                // FORMAT(LZH)
		uintptr(unsafe.Pointer(&src[0])), // to be compressed
		uintptr(len(src)),
		uintptr(unsafe.Pointer(&dst[0])), // destination of compressed data
		1,                                // LEVEL (VeryFast),
		0,
		0,
		0,
	)

	log.Printf("lzcompress(%p, %d, %p, %d): %v %v %v", &src[0], len(src), &dst[0], len(dst), ret, x1, x2)
	// log.Printf("dst[0:10] = %X", dst[0:10])
	_, _ = x1, x2

	return dst[:int(ret)]
}

// lzdecompress wraps OodleLZ_Decompress, returning the number of bytes written into dst
// OodleLZ_Decompress(src []byte, srcLen int, dst []byte, dstLen int, 9 uints (0), uint(3))
func lzdecompress(src []byte, dst []byte) int {
	ret, x1, x2 := pOodleLZDecompress.Call(
		uintptr(unsafe.Pointer(&src[0])),
		uintptr(len(src)),
		uintptr(unsafe.Pointer(&dst[0])),
		uintptr(len(dst)),
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		3,
	)

	log.Printf("lzdecompress(%p, %d, %p, %d): %v %v %v", &src[0], len(src), &dst[0], len(dst), ret, x1, x2)
	// log.Printf("dst[0:10] = %X", dst[0:10])
	_, _ = x1, x2

	return int(ret)
}
