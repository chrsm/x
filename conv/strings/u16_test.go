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

var maxUint16 = ^uint16(0)

func TestU16(t *testing.T) {
	numstr := "100"

	n := ToU16(numstr)
	if n != 100 {
		t.Error("Failed to get 100 from ToU16")
	}

	// we don't _really_ care if the string doesn't actually represent
	// a number. in that case, we just want a 0 back.
	invalid := "-_-f-ado039230_;lol_no<generics>"
	n = ToU16(invalid)
	if n != 0 {
		t.Errorf("Got %d from ToU16, expected 0", n)
	}
}

func TestCsvToU16(t *testing.T) {
	csvstr := fmt.Sprintf("0,1,2,3,%d", maxUint16)

	nslc := CsvToU16(csvstr)
	if nslc[0] != 0 || nslc[1] != 1 || nslc[2] != 2 || nslc[3] != 3 || nslc[4] != maxUint16 {
		t.Errorf("Failed to decode %s: got %v", csvstr, nslc)
	}

	invalid := "a,b,cDEf,4,5,6;7;8" // we'll expect some zeroes
	nslc = CsvToU16(invalid)
	if nslc[0] != 0 || nslc[1] != 0 || nslc[2] != 0 || nslc[3] != 4 || nslc[4] != 5 || nslc[5] != 0 {
		t.Errorf("Failed to decode %s: got %v", invalid, nslc)
	}
}

func TestStrsToU16(t *testing.T) {
	rstrs := _randBU16()
	strs := _toStrsU16(rstrs)

	nslc := StrsToU16(strs)

	for i := range nslc {
		if nslc[i] != rstrs[i] {
			t.Errorf("Expected nslc[%d] = randoms[%d]; got %v / %v", i, i, nslc[i], rstrs[i])
		}
	}

	t.Logf("%v == %v", strs, nslc)
}

func _randU16() uint16 {
	m := math.Abs(math.Min(math.Pow(2, 32), math.Pow(2, 16)))

	return uint16(rand.Intn(int(m)))
}

func _randBU16() []uint16 {
	r := make([]uint16, 10)

	for i := 0; i < len(r); i++ {
		r[i] = _randU16()
	}

	return r
}

func _toStrsU16(src []uint16) []string {
	strs := make([]string, len(src))
	for i := range src {
		strs[i] = fmt.Sprintf("%d", src[i])
	}

	return strs
}

var U16_FAST uint16

func BenchmarkFast_U16(b *testing.B) {
	var num uint16
	// generate some random numbers
	x := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		x[i] = fmt.Sprintf("%d", _randU16())
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		num = ToU16_FAST(x[n*int(1000/b.N)])
	}

	U16_FAST = num
}

var U16_STD uint16

func BenchmarkStdAtoi_U16(b *testing.B) {
	var num uint16
	// generate some random numbers
	x := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		x[i] = fmt.Sprintf("%d", _randU16())
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		num = ToU16(x[n*int(1000/b.N)])
	}

	U16_STD = num
}