package exercise

// AnotherFunChallenge is the second exercise in Big O notation.
// It has time complexity O(4 + 7n), which can be simplified to O(n).
func AnotherFunChallenge(input []string) {
	a := 5  // O(1)
	b := 10 // O(1)
	c := 50 // O(1)

	for i := 0; i < len(input); i++ { // O(n)
		x := i + 1        // O(n)
		y := i + 2        // O(n)
		z := i + 3        // O(n)
		_, _, _ = x, y, z // Ignore
	}

	for j := 0; j < len(input); j++ { // O(n)
		p := j * 2  // O(n)
		q := j * 2  // O(n)
		_, _ = p, q // Ignore
	}

	whoAmI := "I don't know"     // O(1)
	_, _, _, _ = a, b, c, whoAmI // Ignore
}
