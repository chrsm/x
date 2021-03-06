// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package strings

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

var maxUint8 = ^uint8(0)

func TestU8(t *testing.T) {
	numstr := "100"

	n := ToU8(numstr)
	if n != 100 {
		t.Error("Failed to get 100 from ToU8")
	}

	// we don't _really_ care if the string doesn't actually represent
	// a number. in that case, we just want a 0 back.
	invalid := "-_-f-ado039230_;lol_no<generics>"
	n = ToU8(invalid)
	if n != 0 {
		t.Errorf("Got %d from ToU8, expected 0", n)
	}
}

func TestCsvToU8(t *testing.T) {
	csvstr := fmt.Sprintf("0,1,2,3,%d", maxUint8)

	nslc := CsvToU8(csvstr)
	if nslc[0] != 0 || nslc[1] != 1 || nslc[2] != 2 || nslc[3] != 3 || nslc[4] != maxUint8 {
		t.Errorf("Failed to decode %s: got %v", csvstr, nslc)
	}

	invalid := "a,b,cDEf,4,5,6;7;8" // we'll expect some zeroes
	nslc = CsvToU8(invalid)
	if nslc[0] != 0 || nslc[1] != 0 || nslc[2] != 0 || nslc[3] != 4 || nslc[4] != 5 || nslc[5] != 0 {
		t.Errorf("Failed to decode %s: got %v", invalid, nslc)
	}
}

func TestStrsToU8(t *testing.T) {
	rstrs := _randBU8()
	strs := _toStrsU8(rstrs)

	nslc := StrsToU8(strs)

	for i := range nslc {
		if nslc[i] != rstrs[i] {
			t.Errorf("Expected nslc[%d] = randoms[%d]; got %v / %v", i, i, nslc[i], rstrs[i])
		}
	}

	t.Logf("%v == %v", strs, nslc)
}

func _randU8() uint8 {
	m := math.Abs(math.Min(math.Pow(2, 32), math.Pow(2, 8)))

	return uint8(rand.Intn(int(m)))
}

func _randBU8() []uint8 {
	r := make([]uint8, 10)

	for i := 0; i < len(r); i++ {
		r[i] = _randU8()
	}

	return r
}

func _toStrsU8(src []uint8) []string {
	strs := make([]string, len(src))
	for i := range src {
		strs[i] = fmt.Sprintf("%d", src[i])
	}

	return strs
}

var U8_FAST uint8

func BenchmarkFast_U8(b *testing.B) {
	var num uint8
	// generate some random numbers
	x := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		x[i] = fmt.Sprintf("%d", _randU8())
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		num = ToU8_FAST(x[n*int(1000/b.N)])
	}

	U8_FAST = num
}

var U8_STD uint8

func BenchmarkStdAtoi_U8(b *testing.B) {
	var num uint8
	// generate some random numbers
	x := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		x[i] = fmt.Sprintf("%d", _randU8())
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		num = ToU8(x[n*int(1000/b.N)])
	}

	U8_STD = num
}
