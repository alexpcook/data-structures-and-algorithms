package linkedlist

import (
	"fmt"
	"reflect"
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

func TestDoublyAppend(t *testing.T) {
	tests := []struct {
		dll  *DoublyLinkedList
		data int
		want *DoublyLinkedList
	}{
		{NewDoublyLinkedList(1), 2, NewDoublyLinkedList(1, 2)},
		{NewDoublyLinkedList(-4, -8, 19, 20), 45, NewDoublyLinkedList(-4, -8, 19, 20, 45)},
	}

	for i, test := range tests {
		test.dll.Append(test.data)
		if !reflect.DeepEqual(test.dll, test.want) {
			t.Fatalf("test %d: want %#v, got %#v", i, test.want, test.dll)
		}
	}
}

func TestDoublyPrepend(t *testing.T) {
	tests := []struct {
		dll  *DoublyLinkedList
		data int
		want *DoublyLinkedList
	}{
		{NewDoublyLinkedList(1), 2, NewDoublyLinkedList(2, 1)},
		{NewDoublyLinkedList(-4, -8, 19, 20), 45, NewDoublyLinkedList(45, -4, -8, 19, 20)},
	}

	for i, test := range tests {
		test.dll.Prepend(test.data)
		if !reflect.DeepEqual(test.dll, test.want) {
			t.Fatalf("test %d: want %#v, got %#v", i, test.want, test.dll)
		}
	}
}

func TestDoublyLookup(t *testing.T) {
	tests := []struct {
		dll         *DoublyLinkedList
		index, want int
		isError     bool
	}{
		{NewDoublyLinkedList(1, 2, 3), -1, 0, true},
		{NewDoublyLinkedList(1, 2, 3), 3, 0, true},
		{NewDoublyLinkedList(1, 2, 3), 0, 1, false},
		{NewDoublyLinkedList(1, 2, 3), 1, 2, false},
		{NewDoublyLinkedList(1, 2, 3), 2, 3, false},
		{NewDoublyLinkedList(10), 0, 10, false},
	}

	for i, test := range tests {
		got, err := test.dll.Lookup(test.index)

		if test.isError {
			if err == nil {
				t.Fatalf("test %d: want error, got nil", i)
			}
			continue
		}

		if got != test.want {
			t.Fatalf("test %d: want %d, got %d", i, test.want, got)
		}
	}
}
