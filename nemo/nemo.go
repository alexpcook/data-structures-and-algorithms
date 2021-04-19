// Package nemo implements an algorithm to search for the string "nemo" in a slice of strings
package nemo

import (
	"fmt"
	"time"
)

// FindNemo prints "found nemo" if names contains "nemo"
func FindNemo(names []string) {
	t0 := time.Now()

	for _, name := range names {
		if name == "nemo" {
			fmt.Println("found nemo")
		}
	}

	fmt.Printf("FindNemo took %v\n", time.Since(t0))
}
