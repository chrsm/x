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

var maxUInt16 = ^uint16(0)
var maxInt16 = int16(maxUInt16 >> 1)
var minInt16 = -(maxInt16 - 1)

func TestI16(t *testing.T) {
	numstr := "100"

	n := ToI16(numstr)
	if n != 100 {
		t.Error("Failed to get 100 from ToI16")
	}

	// we don't _really_ care if the string doesn't actually represent
	// a number. in that case, we just want a 0 back.
	invalid := "-_-f-ado039230_;lol_no<generics>"
	n = ToI16(invalid)
	if n != 0 {
		t.Errorf("Got %d from ToI16, expected 0", n)
	}
}

func TestCsvToI16(t *testing.T) {
	csvstr := fmt.Sprintf("%d,1,2,3,%d", minInt16, maxInt16)

	nslc := CsvToI16(csvstr)
	if nslc[0] != minInt16 || nslc[1] != 1 || nslc[2] != 2 || nslc[3] != 3 || nslc[4] != maxInt16 {
		t.Errorf("Failed to decode %s: got %v", csvstr, nslc)
	}

	invalid := "a,b,cDEf,4,5,6;7;8" // we'll expect some zeroes
	nslc = CsvToI16(invalid)
	if nslc[0] != 0 || nslc[1] != 0 || nslc[2] != 0 || nslc[3] != 4 || nslc[4] != 5 || nslc[5] != 0 {
		t.Errorf("Failed to decode %s: got %v", invalid, nslc)
	}
}

func TestStrsToI16(t *testing.T) {
	rstrs := _randBI16()
	strs := _toStrsI16(rstrs)

	nslc := StrsToI16(strs)

	for i := range nslc {
		if nslc[i] != rstrs[i] {
			t.Errorf("Expected nslc[%d] = randoms[%d]; got %v / %v", i, i, nslc[i], rstrs[i])
		}
	}

	t.Logf("%v == %v", strs, nslc)
}

func _randI16() int16 {
	m := math.Abs(math.Min(math.Pow(2, 32), math.Pow(2, 16)))

	return int16(rand.Intn(int(m)))
}

func _randBI16() []int16 {
	r := make([]int16, 10)

	for i := 0; i < len(r); i++ {
		r[i] = _randI16()
	}

	return r
}

func _toStrsI16(src []int16) []string {
	strs := make([]string, len(src))
	for i := range src {
		strs[i] = fmt.Sprintf("%d", src[i])
	}

	return strs
}

var I16_FAST int16

func BenchmarkFast_I16(b *testing.B) {
	var num int16
	// generate some random numbers
	x := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		x[i] = fmt.Sprintf("%d", _randI16())
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		num = ToI16_FAST(x[n*int(1000/b.N)])
	}

	I16_FAST = num
}

var I16_STD int16

func BenchmarkStdAtoi_I16(b *testing.B) {
	var num int16
	// generate some random numbers
	x := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		x[i] = fmt.Sprintf("%d", _randI16())
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		num = ToI16(x[n*int(1000/b.N)])
	}

	I16_STD = num
}
