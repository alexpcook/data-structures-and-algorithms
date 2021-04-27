package algorithms

// BubbleSort sorts the data slice in increasing order of integers.
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

// SelectionSort sorts the data slice in increasing order of integers.
// It has time complexity O(n^2) and space complexity O(1).
func SelectionSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		smallestIndex := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[smallestIndex] {
				smallestIndex = j
			}
		}
		data[i], data[smallestIndex] = data[smallestIndex], data[i]
	}
}
