package greeting

import "fmt"

// Boo prints the greeting "boo" n times.
// It has space complexity O(1) and time complexity O(n).
func Boo(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("boo")
	}
}

// Hi fills a slice of length n with greeting "hi".
// It has space complexity O(n) and time complexity O(n).
func Hi(n int) []string {
	slc := make([]string, n)
	for i := range slc {
		slc[i] = "hi"
	}
	return slc
}
