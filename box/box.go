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
