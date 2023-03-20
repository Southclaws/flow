package dt

// Reduce applies fn to the collection, reducing the output to a single value.
func Reduce[T any, R any](collection []T, accumulator func(R, T) R, initial R) R {
	for _, item := range collection {
		initial = accumulator(initial, item)
	}

	return initial
}

// ReduceErr is the error handling version of Reduce.
func ReduceErr[T any, R any, E error](collection []T, accumulator func(R, T) (R, E), initial R) (R, error) {
	for _, item := range collection {
		var err error
		initial, err = accumulator(initial, item)
		if err != nil {
			return initial, err
		}
	}

	return initial, nil
}

// PrepReduce is the curried version of Reduce.
func PrepReduce[T any, R any](fn func(R, T) R, accumulator R) func([]T) R {
	return func(collection []T) R {
		for _, item := range collection {
			accumulator = fn(accumulator, item)
		}

		return accumulator
	}
}

// PrepReduceErr is the curried error handling version of Reduce.
func PrepReduceErr[T any, R any, E error](fn func(R, T) (R, E), accumulator R) func([]T) (R, error) {
	return func(collection []T) (R, error) {
		for _, item := range collection {
			var err error
			accumulator, err = fn(accumulator, item)
			if err != nil {
				return accumulator, err
			}
		}

		return accumulator, nil
	}
}
