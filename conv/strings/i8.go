// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package strings

import (
	"strconv"
	"strings"
)

// ToI8 converts a string to a int8
// Use only when you don't care if the input is valid, as the error is thrown away.
// No efforts are made to handle overflow.
func ToI8(src string) int8 {
	num, _ := strconv.ParseInt(src, 10, 8)

	return int8(num)
}

// ToI8_FAST doesn't use the stdlib to convert numbers.
// It's "safe" if you don't expect negative integers, but really only
// saves 10-20 ns.
func ToI8_FAST(src string) int8 {
	var n int8

	for i := 0; i < len(src); i++ {
		n = n*10 + int8(src[i]-'0')
	}

	return n
}

// CsvToI8 converts a string to []int8, zero vals included.
func CsvToI8(src string) []int8 {
	if len(src) == 0 {
		return nil
	}

	split := strings.Split(src, ",")
	r := make([]int8, len(split))

	for i := 0; i < len(split); i++ {
		r[i] = ToI8(split[i])
	}

	return r
}

// StrsToI8 converts a []string to []int8.
func StrsToI8(src []string) []int8 {
	if len(src) == 0 {
		return nil
	}

	r := make([]int8, len(src))
	for i := range src {
		r[i] = ToI8(src[i])
	}

	return r
}