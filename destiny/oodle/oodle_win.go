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

	/*
		Oodle_CheckVersion	0x0000000180067ec0	0x00067ec0	58 (0x3a)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		Oodle_GetConfigValues	0x000000018005d030	0x0005d030	59 (0x3b)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		Oodle_SetConfigValues	0x000000018005d070	0x0005d070	60 (0x3c)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_CheckSeekTableCRCs	0x000000018005f280	0x0005f280	7 (0x7)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_Compress	0x000000018005f350	0x0005f350	8 (0x8)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_CompressContext_Alloc	0x0000000180060770	0x00060770	9 (0x9)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_CompressContext_Free	0x0000000180060980	0x00060980	10 (0xa)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_CompressContext_Reset	0x0000000180060a40	0x00060a40	11 (0xb)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_CompressionLevel_GetName	0x000000018005f710	0x0005f710	15 (0xf)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_CompressOptions_GetDefault	0x000000018005f5d0	0x0005f5d0	12 (0xc)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_CompressOptions_Validate	0x000000018005f610	0x0005f610	13 (0xd)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_Compressor_GetName	0x000000018005f730	0x0005f730	16 (0x10)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_CompressWithContext	0x0000000180060bb0	0x00060bb0	14 (0xe)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_CreateSeekTable	0x000000018005f750	0x0005f750	17 (0x11)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_Decompress	0x000000018005f8b0	0x0005f8b0	18 (0x12)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_FillSeekTable	0x000000018005fd00	0x0005fd00	19 (0x13)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_FindSeekEntry	0x0000000180060070	0x00060070	20 (0x14)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_FreeSeekTable	0x0000000180060080	0x00060080	21 (0x15)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_GetChunkCompressor	0x0000000180060090	0x00060090	22 (0x16)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_GetCompressedBufferSizeNeeded	0x00000001800600b0	0x000600b0	23 (0x17)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_GetCompressedStepForRawStep	0x00000001800600e0	0x000600e0	24 (0x18)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_GetDecodeBufferSize	0x00000001800605b0	0x000605b0	25 (0x19)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_GetInPlaceDecodeBufferSize	0x00000001800605c0	0x000605c0	26 (0x1a)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_GetNumSeekChunks	0x0000000180060600	0x00060600	27 (0x1b)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_GetSeekEntryPackedPos	0x0000000180060670	0x00060670	28 (0x1c)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_GetSeekTableMemorySizeNeeded	0x00000001800606e0	0x000606e0	29 (0x1d)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_GetZipLikeCompressionSettings	0x0000000180060d30	0x00060d30	30 (0x1e)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_MakeSeekChunkLen	0x0000000180060700	0x00060700	31 (0x1f)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZ_ThreadPhased_BlockDecoderMemorySizeNeeded	0x0000000180060e20	0x00060e20	32 (0x20)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZDecoder_Create	0x000000018005e0d0	0x0005e0d0	1 (0x1)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZDecoder_DecodeSome	0x000000018005e180	0x0005e180	2 (0x2)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZDecoder_Destroy	0x000000018005f060	0x0005f060	3 (0x3)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZDecoder_MakeValidCircularWindowSize	0x000000018005f070	0x0005f070	4 (0x4)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZDecoder_MemorySizeNeeded	0x000000018005f0f0	0x0005f0f0	5 (0x5)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleLZDecoder_Reset	0x000000018005f170	0x0005f170	6 (0x6)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1_CompressedBufferSizeNeeded	0x0000000180066ee0	0x00066ee0	46 (0x2e)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1_SelectDictionaryFromPackets	0x0000000180066ef0	0x00066ef0	47 (0x2f)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1_SelectDictionaryFromPackets_Trials	0x0000000180066f50	0x00066f50	48 (0x30)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1_Shared_SetWindow	0x0000000180067d60	0x00067d60	49 (0x31)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1_Shared_Size	0x0000000180067ea0	0x00067ea0	50 (0x32)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1TCP_Decode	0x0000000180065de0	0x00065de0	33 (0x21)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1TCP_Encode	0x0000000180065df0	0x00065df0	34 (0x22)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1TCP_State_InitAsCopy	0x0000000180065e00	0x00065e00	35 (0x23)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1TCP_State_Reset	0x0000000180065e60	0x00065e60	36 (0x24)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1TCP_State_Size	0x0000000180065e70	0x00065e70	37 (0x25)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1TCP_Train	0x0000000180065e80	0x00065e80	38 (0x26)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1UDP_Decode	0x0000000180066000	0x00066000	39 (0x27)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1UDP_Encode	0x0000000180066050	0x00066050	40 (0x28)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1UDP_State_Compact	0x0000000180066060	0x00066060	42 (0x2a)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1UDP_State_Size	0x0000000180066a60	0x00066a60	43 (0x2b)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1UDP_State_Uncompact	0x0000000180066a70	0x00066a70	44 (0x2c)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1UDP_StateCompacted_MaxSize	0x0000000180066a60	0x00066a60	41 (0x29)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodleNetwork1UDP_Train	0x0000000180066e50	0x00066e50	45 (0x2d)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodlePlugin_DisplayAssertion_Default	0x000000018005d1c0	0x0005d1c0	51 (0x33)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodlePlugin_Free_Default	0x000000018005d120	0x0005d120	52 (0x34)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodlePlugin_MallocAligned_Default	0x000000018005d0b0	0x0005d0b0	53 (0x35)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodlePlugin_Printf_Default	0x000000018005d150	0x0005d150	54 (0x36)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodlePlugins_SetAllocators	0x000000018005d230	0x0005d230	55 (0x37)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodlePlugins_SetAssertion	0x000000018005d240	0x0005d240	56 (0x38)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
		OodlePlugins_SetPrintf	0x000000018005d250	0x0005d250	57 (0x39)	oo2core_3_win64.dll	oo2core_3_win64.dll	Exported Function
	*/

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
