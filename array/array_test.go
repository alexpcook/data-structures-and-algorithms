package array

import "testing"

func TestSlicesHaveCommonItem(t *testing.T) {
	tests := []struct {
		s1, s2 []int
		want   bool
	}{
		{[]int{1, 2, 8, 3, 2}, []int{2, 9, 4}, true},
		{[]int{1, 2, 3}, []int{4, 5, 6}, false},
		{[]int{}, []int{4, 3, 2}, false},
		{[]int{4, 3, 2}, []int{}, false},
		{nil, nil, false},
	}

	for _, test := range tests {
		if got := SlicesHaveCommonItem(test.s1, test.s2); got != test.want {
			t.Fatalf("want %v, got %v", test.want, got)
		}
	}
}
