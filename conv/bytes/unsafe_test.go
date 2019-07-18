package bytes

import (
	"math/rand"
	"testing"
	"time"
)

func TestBytesToU64(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := uint64(rnd.Int63())

	x := make([]byte, 8)
	x[0] = byte(n)
	x[1] = byte(n >> 8)
	x[2] = byte(n >> 16)
	x[3] = byte(n >> 24)
	x[4] = byte(n >> 32)
	x[5] = byte(n >> 40)
	x[6] = byte(n >> 48)
	x[7] = byte(n >> 56)

	ret := Btou64(x)
	if ret != n {
		t.Errorf("Couldn't cast byte slice to u64, expected %d, got %d", n, ret)
	}
}

func TestBytesToU32(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := uint32(rnd.Int31())

	x := make([]byte, 8)
	x[0] = byte(n)
	x[1] = byte(n >> 8)
	x[2] = byte(n >> 16)
	x[3] = byte(n >> 24)

	ret := Btou32(x)
	if ret != n {
		t.Errorf("Couldn't cast byte slice to u64, expected %d, got %d", n, ret)
	}
}

func TestU32ToBytes(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := uint32(rnd.Int31())

	expect := make([]byte, 4)
	expect[0] = byte(x)
	expect[1] = byte(x >> 8)
	expect[2] = byte(x >> 16)
	expect[3] = byte(x >> 24)

	bslice := U32tob(x)
	if len(bslice) != len(expect) {
		t.Errorf("Expected bslice length %d, got %d", len(expect), len(bslice))
	}

	for i := range expect {
		if expect[i] != bslice[i] {
			t.Errorf("Error at bslice[%d] - expected %X, got %X", i, expect[i], bslice[i])
		}
	}
}

func TestU64ToBytes(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := uint64(rnd.Int63())

	expect := make([]byte, 8)
	expect[0] = byte(x)
	expect[1] = byte(x >> 8)
	expect[2] = byte(x >> 16)
	expect[3] = byte(x >> 24)
	expect[4] = byte(x >> 32)
	expect[5] = byte(x >> 40)
	expect[6] = byte(x >> 48)
	expect[7] = byte(x >> 56)

	bslice := U64tob(x)
	if len(bslice) != len(expect) {
		t.Errorf("Expected bslice length %d, got %d", len(expect), len(bslice))
	}

	for i := range expect {
		if expect[i] != bslice[i] {
			t.Errorf("Error at bslice[%d] - expected %X, got %X", i, expect[i], bslice[i])
		}
	}
}

func TestStrToBytes(t *testing.T) {
	x := "this\nis\ta\xFFstring"
	expect := []byte(x)

	bslice := Strtob(x)
	if len(bslice) != len(expect) {
		t.Errorf("Expected bslice length %d, got %d", len(expect), len(bslice))
	}

	for i := range expect {
		if expect[i] != bslice[i] {
			t.Errorf("Error at bslice[%d] - expected %X, got %X", i, expect[i], bslice[i])
		}
	}
}

func TestBytesToStr(t *testing.T) {
	x := []byte("this\nis\ta\x0Fslice")
	expect := string(x) // copy

	str := Btostr(x)
	if len(str) != len(expect) {
		t.Errorf("Expected str length %d, got %d", len(expect), len(str))
	}

	for i := range expect {
		if expect[i] != str[i] {
			t.Errorf("Error at str[%d] - expected %X, got %X", i, expect[i], str[i])
		}
	}
}
