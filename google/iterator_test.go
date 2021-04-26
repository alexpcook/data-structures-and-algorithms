package google

import "testing"

func TestIntInterleaver(t *testing.T) {
	slc1 := []int{1, 2, 3}
	slc2 := []int{4, 5}
	slc3 := []int{6, 7, 8, 9}

	iter1, err := NewIntIterator(slc1)
	if err != nil {
		t.Fatal(err)
	}

	iter2, err := NewIntIterator(slc2)
	if err != nil {
		t.Fatal(err)
	}
	iter3, err := NewIntIterator(slc3)
	if err != nil {
		t.Fatal(err)
	}

	interleaver, err := NewIntInterleaver(*iter1, *iter2, *iter3)
	if err != nil {
		t.Fatal(err)
	}

	want := []int{1, 4, 6, 2, 5, 7, 3, 8, 9}
	i := 0
	for interleaver.HasNext() {
		if got := interleaver.Next(); got != want[i] {
			t.Fatalf("want %d, got %d", want[i], got)
		}
		i++
	}
}
