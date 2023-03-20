package dt

// Filter returns a new slice with only items where the predicate returns true.
func Filter[V any](collection []V, predicate func(V) bool) []V {
	result := []V{}

	for _, item := range collection {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// FilterErr is the error handling version of Filter.
func FilterErr[V any](collection []V, predicate func(V) (bool, error)) ([]V, error) {
	result := []V{}

	for _, item := range collection {
		ok, err := predicate(item)
		if err != nil {
			return nil, err
		}

		if ok {
			result = append(result, item)
		}
	}

	return result, nil
}

// PrepFilter is the curried version of Filter.
func PrepFilter[V any](predicate func(V) bool) func([]V) []V {
	return func(collection []V) []V {
		result := []V{}

		for _, item := range collection {
			if predicate(item) {
				result = append(result, item)
			}
		}

		return result
	}
}

// PrepFilterErr is the curried error handling version of Filter.
func PrepFilterErr[V any](predicate func(V) (bool, error)) func([]V) ([]V, error) {
	return func(collection []V) ([]V, error) {
		result := []V{}

		for _, item := range collection {
			ok, err := predicate(item)
			if err != nil {
				return nil, err
			}

			if ok {
				result = append(result, item)
			}
		}

		return result, nil
	}
}
