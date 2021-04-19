package exercise

func anotherFunction() {
	// Some arbitrary function
}

// FunChallenge is the first exercise in Big O notation.
// It has time complexity O(3 + 4n), which can be simplified to O(n).
func FunChallenge(input []string) int {
	a := 10    // O(1)
	a = 50 + 3 // O(1)

	for i := 0; i < len(input); i++ { // O(n)
		anotherFunction() // O(n)
		stranger := true  // O(n)
		_ = stranger      // Ignore this line, as it's only needed to satisfy the Go compiler
		a++               // O(n)
	}

	return a // O(1)
}
