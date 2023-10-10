package mathutil

import (
	"golang.org/x/exp/constraints"
)

// IsEven returns true if the number is even.
func IsEven[T constraints.Integer](n T) bool {
	return (n % T(2)) == 0
}

// IsOdd returns true if the number is odd.
func IsOdd[T constraints.Integer](n T) bool {
	return (n % T(2)) != 0
}
