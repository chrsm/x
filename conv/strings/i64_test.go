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

var maxUInt64 = ^uint64(0)
var maxInt64 = int64(maxUInt64 >> 1)
var minInt64 = -(maxInt64 - 1)

func TestI64(t *testing.T) {
	numstr := "100"

	n := ToI64(numstr)
	if n != 100 {
		t.Error("Failed to get 100 from ToI64")
	}

	// we don't _really_ care if the string doesn't actually represent
	// a number. in that case, we just want a 0 back.
	invalid := "-_-f-ado039230_;lol_no<generics>"
	n = ToI64(invalid)
	if n != 0 {
		t.Errorf("Got %d from ToI64, expected 0", n)
	}
}

func TestCsvToI64(t *testing.T) {
	csvstr := fmt.Sprintf("%d,1,2,3,%d", minInt64, maxInt64)

	nslc := CsvToI64(csvstr)
	if nslc[0] != minInt64 || nslc[1] != 1 || nslc[2] != 2 || nslc[3] != 3 || nslc[4] != maxInt64 {
		t.Errorf("Failed to decode %s: got %v", csvstr, nslc)
	}

	invalid := "a,b,cDEf,4,5,6;7;8" // we'll expect some zeroes
	nslc = CsvToI64(invalid)
	if nslc[0] != 0 || nslc[1] != 0 || nslc[2] != 0 || nslc[3] != 4 || nslc[4] != 5 || nslc[5] != 0 {
		t.Errorf("Failed to decode %s: got %v", invalid, nslc)
	}
}

func TestStrsToI64(t *testing.T) {
	rstrs := _randBI64()
	strs := _toStrsI64(rstrs)

	nslc := StrsToI64(strs)

	for i := range nslc {
		if nslc[i] != rstrs[i] {
			t.Errorf("Expected nslc[%d] = randoms[%d]; got %v / %v", i, i, nslc[i], rstrs[i])
		}
	}

	t.Logf("%v == %v", strs, nslc)
}

func _randI64() int64 {
	m := math.Abs(math.Min(math.Pow(2, 32), math.Pow(2, 64)))

	return int64(rand.Intn(int(m)))
}

func _randBI64() []int64 {
	r := make([]int64, 10)

	for i := 0; i < len(r); i++ {
		r[i] = _randI64()
	}

	return r
}

func _toStrsI64(src []int64) []string {
	strs := make([]string, len(src))
	for i := range src {
		strs[i] = fmt.Sprintf("%d", src[i])
	}

	return strs
}

var I64_FAST int64

func BenchmarkFast_I64(b *testing.B) {
	var num int64
	// generate some random numbers
	x := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		x[i] = fmt.Sprintf("%d", _randI64())
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		num = ToI64_FAST(x[n*int(1000/b.N)])
	}

	I64_FAST = num
}

var I64_STD int64

func BenchmarkStdAtoi_I64(b *testing.B) {
	var num int64
	// generate some random numbers
	x := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		x[i] = fmt.Sprintf("%d", _randI64())
	}

	b.ResetTimer()
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		num = ToI64(x[n*int(1000/b.N)])
	}

	I64_STD = num
}