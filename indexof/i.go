// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package indexof

func I(v int, src []int) int {
	for i := range src {
		if src[i] == v {
			return i
		}
	}

	return -1
}