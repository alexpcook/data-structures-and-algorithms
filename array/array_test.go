package array

import (
	"reflect"
	"testing"
)

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

func TestBasicArrayOperations(t *testing.T) {
	BasicArrayOperations()
}

func TestInsert(t *testing.T) {
	tests := []struct {
		data  []rune
		index int
		value rune
		want  []rune
	}{
		{[]rune{'a', 'b', 'c'}, 0, 'z', []rune{'z', 'a', 'b', 'c'}},
		{[]rune{'a', 'b', 'c'}, 1, 'z', []rune{'a', 'z', 'b', 'c'}},
		{[]rune{'a', 'b', 'c'}, 2, 'z', []rune{'a', 'b', 'z', 'c'}},
		{[]rune{'a', 'b', 'c'}, 3, 'z', []rune{'a', 'b', 'c', 'z'}},
		{[]rune{}, 5, 'z', []rune{'z'}},
		{nil, 5, 'z', []rune{'z'}},
	}

	for _, test := range tests {
		if got := Insert(test.data, test.index, test.value); !reflect.DeepEqual(got, test.want) {
			t.Fatalf("want %v, got %v", test.want, got)
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		data  []rune
		index int
		want  []rune
	}{
		{[]rune{'a', 'b', 'c'}, 0, []rune{'b', 'c'}},
		{[]rune{'a', 'b', 'c'}, 1, []rune{'a', 'c'}},
		{[]rune{'a', 'b', 'c'}, 2, []rune{'a', 'b'}},
	}

	for _, test := range tests {
		if got := Delete(test.data, test.index); !reflect.DeepEqual(test.want, got) {
			t.Fatalf("want %v, got %v", test.want, got)
		}
	}
}

func TestMergeSortedSlices(t *testing.T) {
	tests := []struct {
		s1, s2, want []int
	}{
		{[]int{}, []int{}, []int{}},
		{nil, nil, []int{}},
		{nil, []int{1, 57, 398}, []int{1, 57, 398}},
		{[]int{4, 8, 19, 89}, []int{}, []int{4, 8, 19, 89}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{4, 5, 6}, []int{1, 2, 3}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{14}, []int{1, 13, 45, 78}, []int{1, 13, 14, 45, 78}},
		{[]int{67, 89}, []int{13, 54, 67, 68, 87, 90}, []int{13, 54, 67, 67, 68, 87, 89, 90}},
		{[]int{-7, -1, 0, 4}, []int{-17, -2, 1, 5}, []int{-17, -7, -2, -1, 0, 1, 4, 5}},
	}

	for i, test := range tests {
		if got := MergeSortedSlices(test.s1, test.s2); !reflect.DeepEqual(got, test.want) {
			t.Fatalf("test %d: want %v, got %v", i, test.want, got)
		}
	}
}
