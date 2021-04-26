package algorithms

import "testing"

func factorialTests() (input, want []int) {
	input = []int{-4, 0, 1, 2, 3, 4, 5}
	want = []int{0, 1, 1, 2, 6, 24, 120}
	if len(input) != len(want) {
		panic("input and want must be the same length")
	}
	return
}

func TestFactorialIterative(t *testing.T) {
	input, want := factorialTests()
	for i := 0; i < len(input); i++ {
		if got := FactorialIterative(input[i]); got != want[i] {
			t.Fatalf("want %d, got %d", want[i], got)
		}
	}
}

func TestFactorialRecursive(t *testing.T) {
	input, want := factorialTests()
	for i := 0; i < len(input); i++ {
		if got := FactorialRecursive(input[i]); got != want[i] {
			t.Fatalf("want %d, got %d", want[i], got)
		}
	}
}
