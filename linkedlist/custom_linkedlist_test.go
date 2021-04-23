package linkedlist

import (
	"fmt"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList(3)
	if length := ll.length; length != 1 {
		t.Fatalf("want length 1, got %d", length)
	}
	if ll.head != ll.tail {
		t.Fatalf("want head and tail to be equal, got head=%v and tail=%v", ll.head, ll.tail)
	}
	if head := ll.head.value; head != 3 {
		t.Fatalf("want head value 3, got %d", head)
	}
	if tail := ll.tail.value; tail != 3 {
		t.Fatalf("want tail value 3, got %d", tail)
	}

	ll = NewLinkedList(3, 4, 5)
	if length := ll.length; length != 3 {
		t.Fatalf("want length 3, got %d", length)
	}
	if head := ll.head.value; head != 3 {
		t.Fatalf("want head value 3, got %d", head)
	}
	if tail := ll.tail.value; tail != 5 {
		t.Fatalf("want tail value 5, got %d", tail)
	}
	if tailPtr := ll.tail.next; tailPtr != nil {
		t.Fatalf("want nil tail pointer, got %v", tailPtr)
	}

	fmt.Println(ll)
}
