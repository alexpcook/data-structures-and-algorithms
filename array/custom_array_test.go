package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCustomArray(t *testing.T) {
	data := make([]interface{}, 4)
	data[0] = 1
	data[1] = 6 + 4i
	data[2] = "testing"
	data[3] = 'r'

	ary := NewArray(data...)
	if ary.Length != 4 {
		t.Fatalf("expected array length of 4, got %d", ary.Length)
	}

	for i, d := range data {
		if item := ary.Get(i); item != d {
			t.Fatalf("expected element of %v, got %v", d, item)
		}
	}

	newItem := []string{"a", "b", "cde"}
	newLength := ary.Push(newItem)
	if newLength != 5 {
		t.Fatalf("expected array length of 5, got %d", ary.Length)
	}

	removedItem := ary.Pop()
	if !reflect.DeepEqual(newItem, removedItem) {
		t.Fatalf("expected removed item to be %v, got %v", newItem, removedItem)
	}
	if ary.Length != 4 {
		t.Fatalf("expected array length of 4, got %d", ary.Length)
	}

	deletedItem := ary.Delete(1)
	if deletedItem != 6+4i {
		t.Fatalf("expected to remove %v, but removed %v\n", 6+4i, deletedItem)
	}
	if ary.Length != 3 {
		t.Fatalf("expected array length of 3, got %d", ary.Length)
	}

	fmt.Println(ary)
}
