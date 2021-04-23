package google

import (
	"strconv"
	"strings"
)

// HasSumWithPair1 determines whether a pair of distinct elements in nums
// adds up to sum. It assumes nums is sorted.
// It has linear time complexity O(n) and constant space complexity O(1).
func HasSumWithPair1(nums []int, sum int) bool {
	low := 0
	high := len(nums) - 1

	for high > low {
		tmp := nums[low] + nums[high]
		if tmp > sum {
			high--
		} else if tmp < sum {
			low++
		} else {
			return true
		}
	}

	return false
}

// HasSumWithPair2 determines whether a pair of distinct elements in nums
// adds up to sum. It does not assume that nums is sorted.
// It has linear time complexity O(n) and linear space complexity O(n).
func HasSumWithPair2(nums []int, sum int) bool {
	pastNumsComplement := make(map[int]int)

	for _, n := range nums {
		if _, present := pastNumsComplement[n]; present {
			return true
		}
		pastNumsComplement[sum-n] = 0
	}

	return false
}

// DecompressString is the solution to sample problem  https://techdevguide.withgoogle.com/resources/former-interview-question-compression-and-decompression/#!
// It has linear time complexity O(n).
func DecompressString(data string) (string, error) {
	decompressedData := ""
	multiplier := ""
	i := 0

	for i < len(data) && data[i] != ']' { // O(n)
		b := data[i]
		if b >= 97 && b <= 122 {
			decompressedData += string(b)
			i++
		} else if b >= 48 && b <= 57 {
			multiplier += string(b)
			i++
		} else if b == '[' {
			factor, err := strconv.Atoi(multiplier)
			if err != nil {
				return "", err
			}

			subDecompressedData, err := DecompressString(data[i+1:])
			if err != nil {
				return "", err
			}

			decompressedData += strings.Repeat(subDecompressedData, factor)
			multiplier = ""
			i += len(subDecompressedData) + 2
		}
	}

	return decompressedData, nil
}
