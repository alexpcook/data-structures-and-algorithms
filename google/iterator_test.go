package google

import (
	"testing"
)

func TestIntIterator(t *testing.T) {
	data := []int{}
	itr := NewIntIterator(data)
	if itr.HasNext() {
		t.Fatal("want empty iterator")
	}
	_, err := itr.Next()
	if err == nil {
		t.Fatal("want error, got nil")
	}

	tests := [][]int{{1}, {8, 3, 4}}
	for _, test := range tests {
		itr = NewIntIterator(test)
		for i := 0; itr.HasNext(); i++ {
			got, err := itr.Next()
			if err != nil {
				t.Fatal(err)
			} else if got != test[i] {
				t.Fatalf("want %d, got %d", test[i], got)
			}
		}
		_, err = itr.Next()
		if err == nil {
			t.Fatal("want error, got nil")
		}
	}
}

func TestIntInterleaver(t *testing.T) {
	interleaver := NewIntInterleaver()
	if interleaver.HasNext() {
		t.Fatal("want empty interleaver")
	}

	if _, err := interleaver.Next(); err == nil {
		t.Fatal("want error, got nil")
	}

	emptyIter := NewIntIterator([]int{1})
	for emptyIter.HasNext() {
		if _, err := emptyIter.Next(); err != nil {
			t.Fatal(err)
		}
	}

	interleaver = NewIntInterleaver(*emptyIter)
	if interleaver.HasNext() {
		t.Fatal("want empty interleaver")
	}
	if _, err := interleaver.Next(); err == nil {
		t.Fatal("want error, got nil")
	}

	slices := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9},
	}
	iters := make([]IntIterator, 3)
	for i := 0; i < len(iters); i++ {
		iter := NewIntIterator(slices[i])
		iters[i] = *iter
	}

	interleaver = NewIntInterleaver(iters...)
	want := []int{1, 4, 6, 2, 5, 7, 3, 8, 9}

	for i := 0; interleaver.HasNext(); i++ {
		switch got, err := interleaver.Next(); {
		case err != nil:
			t.Fatal(err)
		case got != want[i]:
			t.Fatalf("want %d, got %d", want[i], got)
		}
	}

	if _, err := interleaver.Next(); err == nil {
		t.Fatal("want error, got nil")
	}
}
