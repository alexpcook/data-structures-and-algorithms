package queue

import (
	"errors"
	"fmt"
	"testing"
)

func TestQueueString(t *testing.T) {
	q := new(Queue)
	q.Enqueue('a')
	q.Enqueue('i')
	fmt.Println(q)
}

func TestQueueIsEmptyError(t *testing.T) {
	err := QueueIsEmptyError{}
	got := err.Error()
	want := "queue is empty"
	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestQueuePeek(t *testing.T) {
	q := new(Queue)
	val, err := q.Peek()
	if !errors.Is(err, QueueIsEmptyError{}) {
		t.Fatalf("want QueueIsEmptyError, got %T = %v", err, err)
	}
	if val != 0 {
		t.Fatalf("want %v, got %v", 0, val)
	}

	q.Enqueue('t')
	val, err = q.Peek()
	if err != nil {
		t.Fatalf("want nil error, got %T = %v", err, err)
	}
	if val != 't' {
		t.Fatalf("want %c, got %c", 't', val)
	}
	if q.length != 1 {
		t.Fatalf("want queue length 1, got %d", q.length)
	}
}

func TestQueueEnqueue(t *testing.T) {
	q := new(Queue)
	q.Enqueue('r')

	if q.length != 1 {
		t.Fatalf("want queue depth 1, got %d", q.length)
	}
	if want := (node{'r', nil}); *q.first != want || *q.last != want {
		t.Fatalf("want queue first == last == %#v, got first == %#v and last == %#v", want, *q.first, *q.last)
	}

	q.Enqueue('a')

	if q.length != 2 {
		t.Fatalf("want queue depth 2, got %d", q.length)
	}

	last := node{'a', nil}
	first := node{'r', &last}

	if q.first.value != first.value {
		t.Fatalf("want queue first value == %#v, got first value == %#v", first.value, q.first.value)
	}
	if *q.first.next != last {
		t.Fatalf("want queue first next == %#v, got first next == %#v", last, *q.first.next)
	}
	if *q.last != last {
		t.Fatalf("want queue last == %#v, got last == %#v", last, *q.last)
	}
}

func TestQueueDequeue(t *testing.T) {
	q := new(Queue)
	val, err := q.Dequeue()
	if !errors.Is(err, QueueIsEmptyError{}) {
		t.Fatalf("want QueueIsEmptyError, got %T = %v", err, err)
	}
	if val != 0 {
		t.Fatalf("want %v, got %v", 0, val)
	}

	q.Enqueue('a')
	q.Enqueue('b')

	val, err = q.Dequeue()
	if err != nil {
		t.Fatalf("want nil error, got %T = %v", err, err)
	}
	if val != 'a' {
		t.Fatalf("want %c, got %c", 'a', val)
	}
	if q.length != 1 {
		t.Fatalf("want queue length 1, got %d", q.length)
	}

	val, err = q.Dequeue()
	if err != nil {
		t.Fatalf("want nil error, got %T = %v", err, err)
	}
	if val != 'b' {
		t.Fatalf("want %c, got %c", 'b', val)
	}
	if q.length != 0 {
		t.Fatalf("want queue length 0, got %d", q.length)
	}
	if q.first != nil || q.last != nil {
		t.Fatalf("want queue first and queue last == nil, got first == %#v and last == %#v", q.first, q.last)
	}
}
