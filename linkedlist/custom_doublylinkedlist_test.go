package linkedlist

import (
	"fmt"
	"testing"
)

func TestNewDoublyLinkedList(t *testing.T) {
	dll := NewDoublyLinkedList(3)
	if length := dll.length; length != 1 {
		t.Fatalf("want length 1, got %d", length)
	}
	if dll.head != dll.tail {
		t.Fatalf("want head and tail to be equal, got head=%v and tail=%v", dll.head, dll.tail)
	}
	if head := dll.head.value; head != 3 {
		t.Fatalf("want head value 3, got %d", head)
	}
	if tail := dll.tail.value; tail != 3 {
		t.Fatalf("want tail value 3, got %d", tail)
	}
	if dll.head.next != nil || dll.head.previous != nil {
		t.Fatalf("want next and previous to be nil, got next=%v and previous=%v", dll.head.next, dll.head.previous)
	}

	dll = NewDoublyLinkedList(3, 4, 5)
	if length := dll.length; length != 3 {
		t.Fatalf("want length 3, got %d", length)
	}
	if head := dll.head.value; head != 3 {
		t.Fatalf("want head value 3, got %d", head)
	}
	if tail := dll.tail.value; tail != 5 {
		t.Fatalf("want tail value 5, got %d", tail)
	}
	if tailNext := dll.tail.next; tailNext != nil {
		t.Fatalf("want nil tail next, got %v", tailNext)
	}
	if headPrevious := dll.head.previous; headPrevious != nil {
		t.Fatalf("want nil head previous, got %v", headPrevious)
	}

	fmt.Println(dll)
}
