package linkedlist

import (
	"fmt"
	"reflect"
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

func TestAppend(t *testing.T) {
	tests := []struct {
		ll   *LinkedList
		data int
	}{
		{NewLinkedList(1), 2},
		{NewLinkedList(-4, -8, 19, 20), 45},
	}

	for i, test := range tests {
		oldTail := test.ll.tail
		oldLen := test.ll.length

		test.ll.Append(test.data)

		if oldTail.next != test.ll.tail {
			t.Fatalf("test %d: want old tail next to equal new tail, got next=%v and tail=%v", i, oldTail.next, test.ll.tail)
		}

		if test.ll.tail.value != test.data {
			t.Fatalf("test %d: want tail %d, got %d", i, test.data, test.ll.tail.value)
		}

		if oldLen+1 != test.ll.length {
			t.Fatalf("test %d: want length %d, got %d", i, oldLen+1, test.ll.length)
		}

		if tailPtr := test.ll.tail.next; tailPtr != nil {
			t.Fatalf("test %d: want tail next nil, got %v", i, tailPtr)
		}
	}
}

func TestPrepend(t *testing.T) {
	tests := []struct {
		ll   *LinkedList
		data int
	}{
		{NewLinkedList(1), 2},
		{NewLinkedList(-4, -8, 19, 20), 45},
	}

	for i, test := range tests {
		oldHead := test.ll.head
		oldLen := test.ll.length

		test.ll.Prepend(test.data)

		if oldHead != test.ll.head.next {
			t.Fatalf("test %d: want new head next to equal old head, got next=%v and head=%v", i, test.ll.head.next, oldHead)
		}

		if test.ll.head.value != test.data {
			t.Fatalf("test %d: want head %d, got %d", i, test.data, test.ll.head.value)
		}

		if oldLen+1 != test.ll.length {
			t.Fatalf("test %d: want length %d, got %d", i, oldLen+1, test.ll.length)
		}
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		ll      *LinkedList
		data    int
		index   int
		isError bool
		want    *LinkedList
	}{
		{NewLinkedList(3), 4, 0, false, NewLinkedList(4, 3)},
		{NewLinkedList(3), 4, 1, false, NewLinkedList(3, 4)},
		{NewLinkedList(3), 4, -1, true, nil},
		{NewLinkedList(3), 4, 2, true, nil},
		{NewLinkedList(4, 7, 1, -4), 5, 2, false, NewLinkedList(4, 7, 5, 1, -4)},
	}

	for i, test := range tests {
		err := test.ll.Insert(test.data, test.index)

		if test.isError {
			if err == nil {
				t.Fatalf("test %d: want error, got nil", i)
			}
			continue
		}

		if !reflect.DeepEqual(test.ll, test.want) {
			t.Fatalf("test %d: want %#v, got %#v", i, test.want, test.ll)
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		ll      *LinkedList
		index   int
		isError bool
		want    *LinkedList
	}{
		{NewLinkedList(3), 0, true, nil},
		{NewLinkedList(3, 4), -1, true, nil},
		{NewLinkedList(3, 4), 2, true, nil},
		{NewLinkedList(3, 4), 0, false, NewLinkedList(4)},
		{NewLinkedList(3, 4), 1, false, NewLinkedList(3)},
		{NewLinkedList(3, 4, 5), 1, false, NewLinkedList(3, 5)},
		{NewLinkedList(3, 4, 5, 6), 3, false, NewLinkedList(3, 4, 5)},
		{NewLinkedList(3, 4, 5, 6), 2, false, NewLinkedList(3, 4, 6)},
	}

	for i, test := range tests {
		err := test.ll.Delete(test.index)

		if test.isError {
			if err == nil {
				t.Fatalf("test %d: want error, got nil", i)
			}
			continue
		}

		if !reflect.DeepEqual(test.ll, test.want) {
			t.Fatalf("test %d: want %#v, got %#v", i, test.want, test.ll)
		}
	}
}
