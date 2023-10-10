package sliceutil

import (
	"math/rand"
)

// Filter returns a new slice containing only elements where the function "f" returns true
func Filter[T any](slice []T, f func(T) bool) []T {
	var n []T

	for i := range slice {
		if f(slice[i]) {
			n = append(n, slice[i])
		}
	}

	return n
}

// Map (apply-to-all) returns a new slice containing elements of the output from function "f" of each element of slice
func Map[T any](slice []T, f func(T) T) []T {
	var n []T

	for i := range slice {
		n = append(n, f(slice[i]))
	}

	return n
}

// All returns true if all elements in slice satisfy the predicate "f"
func All[T any](slice []T, f func(T) bool) bool {
	for i := range slice {
		if !f(slice[i]) {
			return false
		}
	}

	return true
}

// Any returns true if any elements in slice satisfy the predicate "f"
func Any[T any](slice []T, f func(T) bool) bool {
	for i := range slice {
		if f(slice[i]) {
			return true
		}
	}

	return false
}

// PointerSlice returns a slice of pointers that point to the items in the original slice ([]T ==> []*T)
//
// This is convenient for many AWS SDK methods which require []*T
func PointerSlice[T any](slice []T) []*T {
	v := make([]*T, len(slice))

	for i := range v {
		v[i] = &slice[i]
	}

	return v
}

// Shuffle randomizes the order of the elements in the slice.
func Shuffle[T any](slice []T) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
