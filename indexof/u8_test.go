// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package indexof

import (
	"math/rand"
	"testing"
	"time"
)

func TestU8(t *testing.T) {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	expect := map[uint8]int{}

	var Uint8s []uint8

	c := 0
	for i := 0; i < 256; i++ {
		v := uint8(rnd.Int()) // l o l

		// skip if we've already done this one
		if _, ok := expect[v]; ok {
			continue
		}

		Uint8s = append(Uint8s, v)
		expect[v] = c
		c++
	}

	for i := range expect {
		got := U8(i, Uint8s)
		want := expect[i]

		if got != want {
			t.Errorf("Uint8(%d) - expected at position %d, got position %d", i, want, got)
		}
	}
}
