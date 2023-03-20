package dt

// Map applies fn to each item in the list and returns a new list of equal size.
func Map[T any, R any](collection []T, fn func(T) R) []R {
	result := make([]R, len(collection))

	for i, item := range collection {
		result[i] = fn(item)
	}

	return result
}

// MapErr is the error handling version of Map.
func MapErr[T any, R any](collection []T, fn func(T) (R, error)) ([]R, error) {
	result := make([]R, len(collection))

	for i, item := range collection {
		var err error
		result[i], err = fn(item)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// PrepMap is the curried version of Map.
func PrepMap[T any, R any](fn func(T) R) func([]T) []R {
	return func(collection []T) []R {
		result := make([]R, len(collection))

		for i, item := range collection {
			result[i] = fn(item)
		}

		return result
	}
}

// PrepMapErr is the curried error handling version of Map.
func PrepMapErr[T any, R any](fn func(T) (R, error)) func([]T) ([]R, error) {
	return func(collection []T) ([]R, error) {
		result := make([]R, len(collection))

		for i, item := range collection {
			var err error
			result[i], err = fn(item)
			if err != nil {
				return nil, err
			}
		}

		return result, nil
	}
}
