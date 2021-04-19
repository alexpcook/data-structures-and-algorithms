// Package nemo implements an algorithm to search for the string "nemo" in a slice of strings
package nemo

import (
	"fmt"
)

// FindNemo prints "found nemo" if names contains "nemo"
// O(n) --> linear time
func FindNemo(names []string) {
	for _, name := range names {
		if name == "nemo" {
			fmt.Println("found nemo")
		}
	}
}
