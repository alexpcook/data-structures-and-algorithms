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
