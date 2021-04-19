// Package box implements a library to output the contents of a sample of boxes
package box

import "fmt"

// LogFirstBox prints the first box of the boxes slice.
// It has constant time complexity O(1).
func LogFirstBox(boxes []string) {
	fmt.Println(boxes[0])
}

// LogFirstTwoBoxes prints the first two boxes of the boxes slice.
// It has constant time complexity O(2), which can be simplified to O(1).
func LogFirstTwoBoxes(boxes []string) {
	fmt.Println(boxes[0])
	fmt.Println(boxes[1])
}

// LogAllPairsOfBoxes prints all pairs of boxes from the boxes slice.
// It has time complexity O(n^2).
func LogAllPairsOfBoxes(boxes []string) {
	for _, box1 := range boxes {
		for _, box2 := range boxes {
			fmt.Println(box1, box2)
		}
	}
}

// LogBoxesRecursive creates n loops.
// It has time complexity O(n!).
func LogBoxesRecursive(n int) {
	for i := 0; i < n; i++ { // O(n)
		LogBoxesRecursive(i - 1) // O(n-1), O(n-2), ...
	}
}
