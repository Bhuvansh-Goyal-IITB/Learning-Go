package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "hola")
	})
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		stackOfInts := NewStack[int]()

		AssertTrue(t, stackOfInts.IsEmpty())

		stackOfInts.Push(123)
		AssertFalse(t, stackOfInts.IsEmpty())

		stackOfInts.Push(456)

		value, _ := stackOfInts.Pop()
		AssertEqual(t, value, 456)

		value, _ = stackOfInts.Pop()
		AssertEqual(t, value, 123)

		AssertTrue(t, stackOfInts.IsEmpty())

		stackOfInts.Push(1)
		stackOfInts.Push(2)

		firstValue, _ := stackOfInts.Pop()
		secondValue, _ := stackOfInts.Pop()

		AssertEqual(t, firstValue+secondValue, 3)
	})
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func AssertTrue(t testing.TB, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t testing.TB, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
