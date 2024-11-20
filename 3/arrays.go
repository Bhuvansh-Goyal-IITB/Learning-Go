package main

import "fmt"

func main() {
	arr := [5]uint8{1, 2, 3, 4, 5}

	result := square(&arr)

	result[2] = 0

	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("result: %v\n", result)
}

func square(arr *[5]uint8) [5]uint8 {
	for i := range *arr {
		arr[i] = arr[i] * arr[i]
	}

	return *arr
}
