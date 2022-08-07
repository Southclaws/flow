package dt

type (
	MapIteratee[T any, R any]    func(T) R
	MapErrIteratee[T any, R any] func(T) (R, error)
)

// Map is stolen from samber/lo but has no `i` arg.
func Map[T any, R any](
	collection []T,
	iteratee MapIteratee[T, R],
) []R {
	result := make([]R, len(collection))

	for i, item := range collection {
		result[i] = iteratee(item)
	}

	return result
}

// MapErr is Map but handles errors.
func MapErr[T any, R any](
	collection []T,
	iteratee MapErrIteratee[T, R],
) ([]R, error) {
	result := make([]R, len(collection))

	for i, item := range collection {
		var err error
		result[i], err = iteratee(item)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}


func Filter[V any](
	collection []V,
	predicate func(V) bool,
) []V {
	result := []V{}

	for _, item := range collection {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

func FilterErr[V any](
	collection []V,
	predicate func(V) (bool, error),
) ([]V, error) {
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


// Reduce is stolen from samber/lo but improved with proper error handling and
// without the useless `i` argument.
//
// Most of the time you just want to Reduce(arr, DataTransformer) instead of
// constructing a whole new function and ignoring the `i int` parameter.
//
// Also, supports returning errors so you can break out of the iteration.
func Reduce[T any, R any](
	collection []T,
	accumulator ReduceIteratee[T, R],
	initial R,
) R {
	for _, item := range collection {
		initial = accumulator(initial, item)
	}

	return initial
}

func ReduceErr[T any, R any, E error](
	collection []T,
	accumulator ReduceErrIteratee[T, R, E],
	initial R,
) (R, error) {
	for _, item := range collection {
		var err error
		initial, err = accumulator(initial, item)
		if err != nil {
			return initial, err
		}
	}

	return initial, nil
}
