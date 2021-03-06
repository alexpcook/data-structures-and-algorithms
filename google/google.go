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

// DecompressString is the solution to sample problem (https://techdevguide.withgoogle.com/resources/former-interview-question-compression-and-decompression/).
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

// LakeVolume1 is the brute force solution to sample problem (https://techdevguide.withgoogle.com/resources/former-interview-question-volume-of-lakes/).
// It has time complexity O(n^2) due to nested loops, and space complexity O(1).
func LakeVolume1(heights []int) int {
	volume := 0

	for i, height := range heights {
		var maxLeft, maxRight int

		for j := i; j > -1; j-- {
			if heights[j] > maxLeft {
				maxLeft = heights[j]
			}
		}

		for j := i; j < len(heights); j++ {
			if heights[j] > maxRight {
				maxRight = heights[j]
			}
		}

		if maxLeft < maxRight {
			volume += maxLeft - height
		} else {
			volume += maxRight - height
		}
	}

	return volume
}

// LakeVolume2 is the dynamic programming approach to sample problem (https://techdevguide.withgoogle.com/resources/former-interview-question-volume-of-lakes/).
// It has time complexity O(n) and space complexity O(n).
func LakeVolume2(heights []int) int {
	max := 0
	maxLeft := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		if heights[i] > max {
			max = heights[i]
		}
		maxLeft[i] = max
	}

	max = 0
	maxRight := make([]int, len(heights))
	for i := len(heights) - 1; i > -1; i-- {
		if heights[i] > max {
			max = heights[i]
		}
		maxRight[i] = max
	}

	volume := 0
	for i, height := range heights {
		if maxLeft[i] < maxRight[i] {
			volume += maxLeft[i] - height
		} else {
			volume += maxRight[i] - height
		}
	}

	return volume
}
