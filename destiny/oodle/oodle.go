package oodle

import "fmt"

// Decompress is a friendly wrapper around lzdecompress without the len of src/dst required
func Decompress(src []byte, dst []byte) int {
	return lzdecompress(src, dst)
}

func Decompress2(src []byte) []byte {
	dst := make([]byte, 262144)
	n := lzdecompress(src, dst)

	if n == 0 {
		panic(fmt.Sprintf("src(%d) n=0 :(", len(src), n))
	}

	return dst[:n]
}
