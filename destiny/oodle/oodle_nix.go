// +build !windows

package oodle

import "log"

const CanDecompress = false

func lzdecompress(src []byte, dst []byte) int {
	log.Println("oodle/lzdecompress unimplemented on non-Windows platforms")

	return 0
}
