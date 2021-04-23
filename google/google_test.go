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

func TestDecompressString(t *testing.T) {
	tests := []struct {
		data, want string
	}{
		{"4[abc]2[e]6[o]p", "abcabcabcabceeoooooop"},
		{"10[a]", "aaaaaaaaaa"},
		{"3[abc]4[ab]c", "abcabcabcababababc"},
		{"2[3[a]b]", "aaabaaab"},
	}

	for i, test := range tests {
		if got, err := DecompressString(test.data); err != nil {
			t.Fatalf("test %d: %v", i, err)
		} else if got != test.want {
			t.Fatalf("test %d: want %v, got %v", i, test.want, got)
		}
	}
}
