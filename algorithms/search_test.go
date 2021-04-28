package algorithms

import (
	"fmt"
	"testing"
)

type searchTest struct {
	data    []int
	n       int
	want    int
	isError bool
}

func searchTestCases() []searchTest {
	return []searchTest{
		{nil, 3, 0, true},
		{[]int{}, 8, 0, true},
		{[]int{1, 2, 3, 5, 7, 9, 23}, 1, 0, false},
		{[]int{1, 2, 3, 5, 7, 9, 23}, 5, 3, false},
		{[]int{1, 2, 3, 5, 7, 9, 23}, -9, 0, true},
		{[]int{1, 2, 3, 5, 7, 9, 23}, 93, 0, true},
		{[]int{1, 2, 3, 5, 7, 9, 23}, 23, 6, false},
		{[]int{1, 2, 3, 5, 7, 9}, 1, 0, false},
		{[]int{1, 2, 3, 5, 7, 9}, 3, 2, false},
		{[]int{1, 2, 3, 5, 7, 9}, 0, 0, true},
		{[]int{1, 2, 3, 5, 7, 9}, 10, 0, true},
		{[]int{1, 2, 3, 5, 7, 9}, 9, 5, false},
		{[]int{1, 2, 3, 5, 7, 9}, 5, 3, false},
	}
}

func SearchTestRunner(searchFunc func([]int, int) (int, error), testCases []searchTest) error {
	for i, test := range testCases {
		got, err := searchFunc(test.data, test.n)
		if test.isError {
			if err == nil {
				return fmt.Errorf("test %d: want error, got nil", i)
			}
		} else if got != test.want {
			return fmt.Errorf("test %d: want %d, got %d", i, test.want, got)
		}
	}
	return nil
}

func TestLinearSearch(t *testing.T) {
	err := SearchTestRunner(LinearSearch, searchTestCases())
	if err != nil {
		t.Fatal(err)
	}
}

func TestBinarySearch(t *testing.T) {
	err := SearchTestRunner(BinarySearch, searchTestCases())
	if err != nil {
		t.Fatal(err)
	}
}
