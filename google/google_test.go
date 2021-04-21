package google

import (
	"sort"
	"testing"
)

type hasSumWithPairXTest struct {
	data []int
	sum  int
	want bool
}

func getHasSumWithPairXTestCases() []hasSumWithPairXTest {
	return []hasSumWithPairXTest{
		{[]int{8, 2, 5}, 7, true},
		{[]int{8, 2, 5}, 9, false},
		{nil, 7, false},
		{[]int{}, 10, false},
	}
}

func TestHasSumWithPair1(t *testing.T) {
	tests := getHasSumWithPairXTestCases()

	for _, test := range tests {
		sort.Slice(test.data, func(i, j int) bool {
			return test.data[i] < test.data[j]
		})

		if got := HasSumWithPair1(test.data, test.sum); got != test.want {
			t.Fatalf("want %v, got %v", test.want, got)
		}
	}
}

func TestHasSumWithPair2(t *testing.T) {
	tests := getHasSumWithPairXTestCases()

	for _, test := range tests {
		if got := HasSumWithPair2(test.data, test.sum); got != test.want {
			t.Fatalf("want %v, got %v", test.want, got)
		}
	}
}
