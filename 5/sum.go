package main

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAllTails(slices ...[]int) []int {
	var output []int

	for _, slice := range slices {
		if len(slice) == 0 {
			output = append(output, 0)
			continue
		}
		output = append(output, Sum(slice[1:]))
	}

	return output
}
