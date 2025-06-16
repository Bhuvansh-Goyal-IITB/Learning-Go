package main

import (
	"slices"
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum of number of a slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("sum of empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(acc, x int) int {
			return acc * x
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})
	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(acc, x string) string {
			return acc + x
		}

		AssertEqual(t, Reduce([]string{"hello ", "world!"}, concatenate, ""), "hello world!")
	})
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})

		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})

	type Person struct {
		Name string
	}

	t.Run("find the best programmer", func(t *testing.T) {
		people := []Person{
			{Name: "Kent Beck"},
			{Name: "Martin Fowler"},
			{Name: "Chris James"},
		}

		king, found := Find(people, func(person Person) bool {
			return strings.Contains(person.Name, "Chris")
		})

		AssertTrue(t, found)
		AssertEqual(t, king, Person{Name: "Chris James"})
	})
}
