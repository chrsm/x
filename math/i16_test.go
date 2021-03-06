// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package math

import "testing"

func TestI16(t *testing.T) {
	var a, b int16 = 0, 1

	if MaxI16(a, b) != 1 {
		t.Error("Max fail")
	}

	if MaxI16(b, a) != 1 {
		t.Error("Max fail")
	}

	if MinI16(a, b) != 0 {
		t.Error("Min fail")
	}

	if MinI16(b, a) != 0 {
		t.Error("Min fail")
	}
}
