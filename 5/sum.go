package main

func Sum(numbers []int) int {
	return Reduce(numbers, func(acc, x int) int {
		return acc + x
	}, 0)
}

func SumAllTails(slices ...[]int) []int {
	return Reduce(slices, func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		}
		return append(acc, Sum(x[1:]))
	}, []int{})
}
