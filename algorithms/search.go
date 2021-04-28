package algorithms

import "fmt"

// LinearSearch finds the first index of n in data. It returns a non-nil error if n is not in data.
// It has time complexity O(n) and space complexity O(1).
func LinearSearch(data []int, n int) (int, error) {
	for i, d := range data {
		if d == n {
			return i, nil
		}
	}
	return 0, fmt.Errorf("%d not found in %v", n, data)
}

// BinarySearch finds the index of n in data. It returns a non-nil error if n is not in data.
// It assumes data is sorted.
// It has time complexity O(log(n)) and space complexity O(log(n)).
func BinarySearch(data []int, n int) (int, error) {
	if len(data) == 0 {
		return 0, fmt.Errorf("%d not found in %v", n, data)
	} else if len(data) == 1 {
		if data[0] == n {
			return 0, nil
		}
		return 0, fmt.Errorf("%d not found in %v", n, data)
	}

	midpoint := len(data) / 2
	value := data[midpoint]

	var position int
	var err error
	switch {
	case value > n:
		position, err = BinarySearch(data[:midpoint], n)
	case value < n:
		position, err = BinarySearch(data[midpoint:], n)
		position += midpoint
	default:
		position = midpoint
	}

	return position, err
}
