package array

import "fmt"

// SlicesHaveCommonItem determines whether the two input slices
// have at least one common element.
// It has linear time complexity O(n) and linear space complexity O(n).
func SlicesHaveCommonItem(s1, s2 []int) bool {
	s1Values := make(map[int]bool)

	for _, n1 := range s1 {
		s1Values[n1] = true
	}

	for _, n2 := range s2 {
		if _, sawInS1 := s1Values[n2]; sawInS1 {
			return true
		}
	}

	return false
}

// BasicArrayOperations contains examples of basic methods that can be performed on arrays and slices, such as
// Lookup O(1),
// Push O(1),
// Insert O(n),
// Delete O(n)
func BasicArrayOperations() {
	var letters []rune = []rune{'a', 'b', 'c', 'd'}
	fmt.Printf("base slice = %v %p\n", letters, letters)

	// Lookup O(1)
	fmt.Printf("lookup on base slice = %v %p\n", letters[0], letters)

	// Push O(1)
	// This operation has the space complexity cost of allocating a new underlying array
	// The address of letters changes here because of the newly allocated array
	letters = append(letters, 'e')
	fmt.Printf("push to base slice = %v %p\n", letters, letters)

	// Insert O(n)
	// The address of letters would change here if append hadn't been called earlier.
	letters = Insert(letters, 2, 'z')
	fmt.Printf("insert into letters = %v %p\n", letters, letters)

	// Delete
	// Since this operation is simply removing the last element, this should be O(1)
	letters = letters[:len(letters)-1]
	fmt.Printf("pop base slice = %v %p\n", letters, letters)

	// Removing an element at a specific index has time complexity O(n)
	letters = Delete(letters, 2)
	fmt.Printf("delete from letters = %v %p\n", letters, letters)
}

// Insert puts a rune at a specific position in a slice of runes.
// It has time complexity O(n).
func Insert(data []rune, index int, value rune) []rune {
	if elems := len(data); elems == 0 || elems == index {
		return append(data, value)
	}
	data = append(data[:index+1], data[index:]...)
	data[index] = value
	return data
}

// Delete removes a rune at a specific position in a slice of runes.
// It has time complexity O(n).
func Delete(data []rune, index int) []rune {
	return append(data[:index], data[index+1:]...)
}
