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

// InsertionSort sorts the data slice in increasing order of integers.
// It has best case time complexity O(n) and average/worst case time complexity O(n^2).
// It has space complexity O(1).
func InsertionSort(data []int) {
	for i := 0; i < len(data)-1; i++ {
		if data[i+1] < data[i] {
			for j := 0; j < i+1; j++ {
				if data[j] > data[i+1] {
					data[j], data[i+1] = data[i+1], data[j]
				}
			}
		}
	}
}

// MergeSort sorts the data slice in increasing order of integers.
// It has time complexity O(n*log(n)) and space complexity O(n).
func MergeSort(data []int) {
	length := len(data)
	if length < 2 {
		return
	}

	midpoint := length / 2
	left := make([]int, len(data[:midpoint]))
	right := make([]int, len(data[midpoint:]))
	copy(left, data[:midpoint])
	copy(right, data[midpoint:])

	MergeSort(left)
	MergeSort(right)

	i := 0
	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(left) && rightIndex < len(right) {
		leftValue, rightValue := left[leftIndex], right[rightIndex]
		if leftValue < rightValue {
			data[i] = leftValue
			leftIndex++
		} else {
			data[i] = rightValue
			rightIndex++
		}
		i++
	}

	for leftIndex < len(left) {
		data[i] = left[leftIndex]
		leftIndex++
		i++
	}

	for rightIndex < len(right) {
		data[i] = right[rightIndex]
		rightIndex++
		i++
	}
}
