// +build windows

package oodle

import (
	"bytes"
	"strings"
	"testing"
)

func TestCompressDecompress(t *testing.T) {
	src := []byte(strings.Repeat("a", 5000))

	compressed := lzcompress(src)
	t.Logf("comp len: %d", len(compressed))

	// decomp
	buf := make([]byte, 5000)
	n := lzdecompress(compressed, buf)
	t.Logf("decomp len: %d", n)

	if !bytes.Equal(buf, src) {
		t.Fatalf("buf != src")
	}

	// should fail
	x := make([]byte, 20)
	x[1] = 0x02
	dst := make([]byte, 30)
	dst[29] = 0x99
	n = lzdecompress(x, dst)
	if n != 0 {
		t.Errorf("invalid data passed decomp?")
	}
}
