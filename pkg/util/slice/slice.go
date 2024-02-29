package sliceutils

import (
	"reflect"
	"sort"

	"golang.org/x/exp/constraints"
)

// check if the given slice contains the given elemement
func Contains[T comparable](s []T, elem T) bool {
	m := make(map[T]bool)
	for _, k := range s {
		m[k] = true
	}

	return m[elem]
}

// returns a new slice containing the given slice's elements, which pass the provided callback function ('callbackFn' returns true)
func Filter[T any](s []T, callbackFn func(elem T) bool) (result []T) {
	for _, item := range s {
		if !callbackFn(item) {
			continue
		}
		result = append(result, item)
	}

	return
}

// returns a new slice containing the given slice's elements after being transformed by the provided callback function
func Map[A, B any](input []A, callbackFn func(a A) B) (result []B) {
	for _, a := range input {
		b := callbackFn(a)
		result = append(result, b)
	}

	return
}

// returns true if at lease one element of the input array passes the implementation of provided callback function
// otherwise return false
func Some[T any](s []T, callbackFn func(elem T) bool) bool {
	for _, item := range s {
		if callbackFn(item) {
			return true
		}
	}

	return false
}

func IndexOf[T comparable](s []T, elem T) int {
	for i, item := range s {
		if item == elem {
			return i
		}
	}

	return -1
}

func FindDuplications[T comparable](s []T) (duplications []T) {
	var idx int
	duplications = Filter(s, func(e T) bool {
		cond := idx != IndexOf(s, e)
		idx++
		return cond
	})

	return
}

// Check if 2 slices have the same elements without considering their order
// Supported types: int, float, string
func Equal[T constraints.Ordered](a []T, b []T) bool {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return reflect.DeepEqual(a, b)
}

// returns the first element of the given slice, which pass the provided callback function ('callbackFn' returns true)
func Find[T any](s []T, callbackFn func(elem T) bool) *T {
	for _, item := range s {
		if callbackFn(item) {
			return &item
		}
	}

	return nil
}
