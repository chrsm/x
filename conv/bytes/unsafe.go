package bytes

import "unsafe"

// Btou64 converts a byte slice to a uint64, unsafely.
func Btou64(src []byte) uint64 {
	return *(*uint64)(unsafe.Pointer(&src[0]))
}

// U64tob converts a uint64 to a byte slice, unsafely.
func U64tob(src uint64) []byte {
	return (*(*[8]byte)(unsafe.Pointer(&src)))[:]
}

// Btou32 converts a byte slice to a uint32, unsafely.
func Btou32(src []byte) uint32 {
	return *(*uint32)(unsafe.Pointer(&src[0]))
}

// U32tob converts a uint32 to a byte slice, unsafely.
func U32tob(src uint32) []byte {
	return (*(*[4]byte)(unsafe.Pointer(&src)))[:]
}

// Btostr converts a byte slice to a string, unsafely.
func Btostr(src []byte) string {
	return *(*string)(unsafe.Pointer(&src))
}

// Strtob converts a string to a byte slice, unsafely.
func Strtob(src string) []byte {
	return *(*[]byte)(unsafe.Pointer(&src))
}
