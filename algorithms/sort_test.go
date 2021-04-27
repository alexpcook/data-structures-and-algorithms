package algorithms

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input, want []int
	}{
		{[]int{}, []int{}},
		{nil, nil},
		{[]int{1, 4, 2, 3}, []int{1, 2, 3, 4}},
		{[]int{8, 2, 7, 1}, []int{1, 2, 7, 8}},
		{[]int{3}, []int{3}},
	}

	for i, test := range tests {
		if BubbleSort(test.input); !reflect.DeepEqual(test.want, test.input) {
			t.Fatalf("test %d: want %v, got %v", i, test.want, test.input)
		}
	}
}
