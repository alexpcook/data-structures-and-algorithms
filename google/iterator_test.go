package google

import (
	"testing"
)

func TestIntIterator(t *testing.T) {
	data := []int{}
	_, err := NewIntIterator(data)
	if err == nil {
		t.Fatal("want error, got nil")
	}

	data = []int{1}
	itr, err := NewIntIterator(data)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; itr.HasNext(); i++ {
		got, err := itr.Next()
		if err != nil {
			t.Fatal(err)
		} else if got != data[i] {
			t.Fatalf("want %d, got %d", data[i], got)
		}
	}
	_, err = itr.Next()
	if err == nil {
		t.Fatal("want error, got nil")
	}

	data = []int{8, 3, 4}
	itr, err = NewIntIterator(data)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; itr.HasNext(); i++ {
		got, err := itr.Next()
		if err != nil {
			t.Fatal(err)
		} else if got != data[i] {
			t.Fatalf("want %d, got %d", data[i], got)
		}
	}
	_, err = itr.Next()
	if err == nil {
		t.Fatal("want error, got nil")
	}
}
