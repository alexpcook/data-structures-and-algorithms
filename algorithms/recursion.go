package algorithms

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
