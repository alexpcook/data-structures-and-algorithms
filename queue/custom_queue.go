package queue

import "fmt"

type node struct {
	value rune
	next  *node
}

// Queue implements a queue data structure using a linked list.
// Create an instance using new(queue.Queue).
type Queue struct {
	first  *node
	last   *node
	length int
}

type QueueIsEmptyError struct{}

func (q QueueIsEmptyError) Error() string {
	return "queue is empty"
}

// String implements the fmt.Stringer interface for *Queue.
func (q *Queue) String() string {
	str := ""
	entry := q.first
	for entry != nil {
		str += fmt.Sprintf("%c, ", entry.value)
		entry = entry.next
	}
	return str
}

// Peek examines the next entry to be retrieved from the queue. It returns a non-nil error if the queue is empty.
// It has time complexity O(1).
func (q *Queue) Peek() (rune, error) {
	if q.length == 0 {
		return 0, QueueIsEmptyError{}
	}
	return q.first.value, nil
}

// Enqueue adds an entry to the queue.
// It has time complexity O(1).
func (q *Queue) Enqueue(entry rune) {
	newEntry := &node{entry, nil}
	if q.last != nil {
		q.last.next = newEntry
	}
	q.last = newEntry
	if q.length == 0 {
		q.first = q.last
	}
	q.length++
}

// Dequeue retrieves the next entry from the queue. It returns a non-nil error if the queue is empty.
// It has time complexity O(1).
func (q *Queue) Dequeue() (rune, error) {
	if q.length == 0 {
		return 0, QueueIsEmptyError{}
	}
	nextEntry := q.first.value
	q.first = q.first.next
	if q.length == 1 {
		q.last = nil
	}
	q.length--
	return nextEntry, nil
}
