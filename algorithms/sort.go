package algorithms

// BubbleSort sorts the data array in increasing order of integers.
// It has best case time complexity O(n) and average/worst case time complexity O(n^2).
// It has space complexity O(1).
func BubbleSort(data []int) {
	for i := len(data) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}
