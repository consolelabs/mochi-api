package maputils

func Keys[K comparable, V any](m map[K]V) (keys []K) {
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func Values[K comparable, V any](m map[K]V) (values []V) {
	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func HasKey[K comparable, V any](m map[K]V, key K) bool {
	if m == nil {
		return false
	}

	_, ok := m[key]
	return ok
}
