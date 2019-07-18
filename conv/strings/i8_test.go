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

var maxUInt8 = ^uint8(0)
var maxInt8 = int8(maxUInt8 >> 1)
var minInt8 = -(maxInt8 - 1)

func TestI8(t *testing.T) {
	numstr := "100"

	n := ToI8(numstr)
	if n != 100 {
		t.Error("Failed to get 100 from ToI8")
	}

	// we don't _really_ care if the string doesn't actually represent
	// a number. in that case, we just want a 0 back.
	invalid := "-_-f-ado039230_;lol_no<generics>"
	n = ToI8(invalid)
	if n != 0 {
		t.Errorf("Got %d from ToI8, expected 0", n)
	}
}

func TestCsvToI8(t *testing.T) {
	csvstr := fmt.Sprintf("%d,1,2,3,%d", minInt8, maxInt8)

	nslc := CsvToI8(csvstr)
	if nslc[0] != minInt8 || nslc[1] != 1 || nslc[2] != 2 || nslc[3] != 3 || nslc[4] != maxInt8 {
		t.Errorf("Failed to decode %s: got %v", csvstr, nslc)
	}

	invalid := "a,b,cDEf,4,5,6;7;8" // we'll expect some zeroes
	nslc = CsvToI8(invalid)
	if nslc[0] != 0 || nslc[1] != 0 || nslc[2] != 0 || nslc[3] != 4 || nslc[4] != 5 || nslc[5] != 0 {
		t.Errorf("Failed to decode %s: got %v", invalid, nslc)
	}
}

func TestStrsToI8(t *testing.T) {
	rstrs := _randBI8()
	strs := _toStrsI8(rstrs)

	nslc := StrsToI8(strs)

	for i := range nslc {
		if nslc[i] != rstrs[i] {
			t.Errorf("Expected nslc[%d] = randoms[%d]; got %v / %v", i, i, nslc[i], rstrs[i])
		}
	}

	t.Logf("%v == %v", strs, nslc)
}

func _randI8() int8 {
	m := math.Abs(math.Min(math.Pow(2, 32), math.Pow(2, 8)))

	return int8(rand.Intn(int(m)))
}

func _randBI8() []int8 {
	r := make([]int8, 10)

	for i := 0; i < len(r); i++ {
		r[i] = _randI8()
	}

	return r
}

func _toStrsI8(src []int8) []string {
	strs := make([]string, len(src))
	for i := range src {
		strs[i] = fmt.Sprintf("%d", src[i])
	}

	return strs
}

var I8_FAST int8

func BenchmarkFast_I8(b *testing.B) {
	var num int8
	// generate some random numbers
	x := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		x[i] = fmt.Sprintf("%d", _randI8())
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		num = ToI8_FAST(x[n*int(1000/b.N)])
	}

	I8_FAST = num
}

var I8_STD int8

func BenchmarkStdAtoi_I8(b *testing.B) {
	var num int8
	// generate some random numbers
	x := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		x[i] = fmt.Sprintf("%d", _randI8())
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		num = ToI8(x[n*int(1000/b.N)])
	}

	I8_STD = num
}
