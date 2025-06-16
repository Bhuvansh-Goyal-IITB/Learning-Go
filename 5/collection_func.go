package main

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	result := initialValue
	for _, val := range collection {
		result = f(result, val)
	}
	return result
}

func Find[T any](collection []T, predicate func(T) bool) (value T, found bool) {
	for _, v := range collection {
		if predicate(v) {
			return v, true
		}
	}
	return
}
