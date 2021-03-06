// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package strings

import (
	"strconv"
	"strings"
)

// ToU32 converts a string to a uint32
// Use only when you don't care if the input is valid, as the error is thrown away.
// No efforts are made to handle overflow.
func ToU32(src string) uint32 {
	num, _ := strconv.ParseUint(src, 10, 32)

	return uint32(num)
}

// ToU32_FAST doesn't use the stdlib to convert numbers.
// It's "safe" if you don't expect negative integers, but really only
// saves 10-20 ns.
func ToU32_FAST(src string) uint32 {
	var n uint32

	for i := 0; i < len(src); i++ {
		n = n*10 + uint32(src[i]-'0')
	}

	return n
}

// CsvToU32 converts a string to []uint32, zero vals included.
func CsvToU32(src string) []uint32 {
	if len(src) == 0 {
		return nil
	}

	split := strings.Split(src, ",")
	r := make([]uint32, len(split))

	for i := 0; i < len(split); i++ {
		r[i] = ToU32(split[i])
	}

	return r
}

// StrsToU32 converts a []string to []uint32.
func StrsToU32(src []string) []uint32 {
	if len(src) == 0 {
		return nil
	}

	r := make([]uint32, len(src))
	for i := range src {
		r[i] = ToU32(src[i])
	}

	return r
}
