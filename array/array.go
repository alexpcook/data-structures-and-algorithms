package array

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
