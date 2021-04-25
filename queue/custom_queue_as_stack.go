package queue

import (
	"errors"

	"github.com/alexpcook/data-structures-and-algorithms/stack"
)

// QueueAsStack implements a queue as a stack.
type QueueAsStack struct {
	data   stack.Stack1
	length int
}

// String implements the fmt.Stringer interface on *QueueAsStack.
func (q *QueueAsStack) String() string {
	return q.data.String()
}

// Peek previews the next entry in the queue without dequeueing it.
// It has time complexity O(1).
func (q *QueueAsStack) Peek() (string, error) {
	return q.data.Peek()
}

// Enqueue adds an entry to the queue.
// It has time complexity O(n) because it has to reverse the order of the underlying stack to preserve
// the FIFO behavior of a queue.
func (q *QueueAsStack) Enqueue(entry string) {
	for next, err := q.data.Pop(); !errors.Is(err, stack.StackIsEmptyError{}); next, err = q.data.Pop() {
		defer func(nextEntry string) {
			q.data.Push(nextEntry)
		}(next)
	}
	q.data.Push(entry)
	q.length++
}

// Dequeue removes an entry from the queue. It returns a non-nil error if the queue is empty.
// It has time complexity O(1).
func (q *QueueAsStack) Dequeue() (string, error) {
	val, err := q.data.Pop()
	if err == nil {
		q.length--
	}
	return val, err
}
