package google

// HasSumWithPair1 determines whether a pair of distinct elements in nums
// adds up to sum. It assumes nums is sorted.
// It has linear time complexity O(n).
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
// It has linear time complexity O(n).
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
