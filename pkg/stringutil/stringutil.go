package stringutil

import (
	"slices"
)

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse[T ~string](s T) T {
	r := []rune(s)

	slices.Reverse(r)

	return T(r)
}
