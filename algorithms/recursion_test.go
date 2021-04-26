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
func fibonacciTests() (input, want []int) {
	input = []int{1, 2, 3, 4, 5, 6}
	want = []int{0, 1, 1, 2, 3, 5}
	if len(input) != len(want) {
		panic("input and want must be the same length")
	}
	return
}

func TestFibonacciIterative(t *testing.T) {
	input, want := fibonacciTests()
	for i := 0; i < len(input); i++ {
		if got := FibonacciIterative(input[i]); got != want[i] {
			t.Fatalf("want %d, got %d", want[i], got)
		}
	}
}

func TestFibonacciRecursive(t *testing.T) {
	input, want := fibonacciTests()
	for i := 0; i < len(input); i++ {
		if got := FibonacciRecursive(input[i]); got != want[i] {
			t.Fatalf("want %d, got %d", want[i], got)
		}
	}
}

func TestReverseStringRecursive(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"monday", "yadnom"},
		{"", ""},
		{"a", "a"},
		{"abc", "cba"},
	}

	for _, test := range tests {
		if got := ReverseStringRecursive(test.input); got != test.want {
			t.Fatalf("want %s, got %s", test.want, got)
		}
	}
}
