package algorithms

import (
	"reflect"
	"testing"
)

type sortTest struct {
	data, want []int
}

func sortTests() []sortTest {
	return []sortTest{
		{[]int{}, []int{}},
		{nil, nil},
		{[]int{1, 4, 2, 3}, []int{1, 2, 3, 4}},
		{[]int{8, 2, 7, 1}, []int{1, 2, 7, 8}},
		{[]int{3}, []int{3}},
		{[]int{5, 9, 2, 90, 1, -4, -100, -3000, 5, 3, 1, 8, 90}, []int{-3000, -100, -4, 1, 1, 2, 3, 5, 5, 8, 9, 90, 90}},
	}
}

func TestBubbleSort(t *testing.T) {
	tests := sortTests()
	for i, test := range tests {
		if BubbleSort(test.data); !reflect.DeepEqual(test.want, test.data) {
			t.Fatalf("test %d: want %v, got %v", i, test.want, test.data)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	tests := sortTests()
	for i, test := range tests {
		if SelectionSort(test.data); !reflect.DeepEqual(test.want, test.data) {
			t.Fatalf("test %d: want %v, got %v", i, test.want, test.data)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	tests := sortTests()
	for i, test := range tests {
		if InsertionSort(test.data); !reflect.DeepEqual(test.want, test.data) {
			t.Fatalf("test %d: want %v, got %v", i, test.want, test.data)
		}
	}
}
