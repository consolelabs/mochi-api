package sliceutils

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
