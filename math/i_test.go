// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package math

import "testing"

func TestI(t *testing.T) {
	var a, b int = 0, 1

	if MaxI(a, b) != 1 {
		t.Error("Max fail")
	}

	if MaxI(b, a) != 1 {
		t.Error("Max fail")
	}

	if MinI(a, b) != 0 {
		t.Error("Min fail")
	}

	if MinI(b, a) != 0 {
		t.Error("Min fail")
	}
}
