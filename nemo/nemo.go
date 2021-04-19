// Package nemo implements an algorithm to search for the string "nemo" in a slice of strings
package nemo

import (
	"fmt"
)

// FindNemo prints "found nemo" each time it encounters name "nemo".
// It has linear time complexity O(n).
func FindNemo(names []string) {
	for _, name := range names {
		if name == "nemo" {
			fmt.Println("found nemo")
			break
		}
	}
}
