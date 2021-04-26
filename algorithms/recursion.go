package algorithms

// FactorialIterative calculates the factorial of n using an iterative approach.
// It has time complexity O(n) and space complexity O(1).
func FactorialIterative(n int) int {
	if n < 0 {
		return 0
	}
	nFactorial := 1
	for i := 2; i <= n; i++ {
		nFactorial *= i
	}
	return nFactorial
}

// FactorialRecursive calculates the factorial of n using a recursive approach.
// It has time complexity O(n) and space complexity O(n), due to recursion.
func FactorialRecursive(n int) int {
	switch {
	case n < 0:
		return 0
	case n < 2:
		return 1
	default:
		return n * FactorialRecursive(n-1)
	}
}

// FibonacciIterative returns the nth number of the Fibonacci sequence using an iterative approach.
// It has time complexity O(n) and space complexity O(1).
func FibonacciIterative(n int) int {
	current, next := 0, 1
	for i := 0; i < n-1; i++ {
		current, next = next, current+next
	}
	return current
}

// FibonacciRecursive returns the nth number of the Fibonacci sequence using a recursive approach.
// It has time complexity O(2^n) and space complexity O(n), due to recursion.
func FibonacciRecursive(n int) int {
	switch n {
	case 1:
		return 0
	case 2:
		return 1
	default:
		return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
	}
}
