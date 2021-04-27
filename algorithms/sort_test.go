package algorithms

import (
	"fmt"
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
		{[]int{4, 1, 3, 2, 1, 2, 9}, []int{1, 1, 2, 2, 3, 4, 9}},
	}
}

func RunSortFunctionTest(sortFunc func([]int), testCases []sortTest) error {
	for i, test := range testCases {
		if sortFunc(test.data); !reflect.DeepEqual(test.want, test.data) {
			return fmt.Errorf("test %d: want %v, got %v", i, test.want, test.data)
		}
	}
	return nil
}

func TestBubbleSort(t *testing.T) {
	err := RunSortFunctionTest(BubbleSort, sortTests())
	if err != nil {
		t.Fatal(err)
	}
}

func TestSelectionSort(t *testing.T) {
	err := RunSortFunctionTest(SelectionSort, sortTests())
	if err != nil {
		t.Fatal(err)
	}
}

func TestInsertionSort(t *testing.T) {
	err := RunSortFunctionTest(InsertionSort, sortTests())
	if err != nil {
		t.Fatal(err)
	}
}

func TestMergeSort(t *testing.T) {
	err := RunSortFunctionTest(MergeSort, sortTests())
	if err != nil {
		t.Fatal(err)
	}
}
