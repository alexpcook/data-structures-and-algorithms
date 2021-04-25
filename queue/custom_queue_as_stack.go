package queue

import (
	"errors"

	"github.com/alexpcook/data-structures-and-algorithms/stack"
)

// QueueAsStack implements a queue as a stack.
type QueueAsStack struct {
	input  stack.Stack2
	output stack.Stack2
}

// String implements the fmt.Stringer interface on *QueueAsStack.
func (q *QueueAsStack) String() string {
	q.Peek()
	return q.output.String()
}

// Peek previews the next entry in the queue without dequeueing it.
// It has time complexity O(1) if the output stack already has data, else O(n) to transfer input to output.
func (q *QueueAsStack) Peek() (string, error) {
	if q.output.Length() == 0 {
		for next, err := q.input.Pop(); !errors.Is(err, stack.StackIsEmptyError{}); next, err = q.input.Pop() {
			q.output.Push(next)
		}
	}
	return q.output.Peek()
}

// Enqueue adds an entry to the queue.
// It has time complexity O(1).
func (q *QueueAsStack) Enqueue(entry string) {
	q.input.Push(entry)
}

// Dequeue removes an entry from the queue. It returns a non-nil error if the queue is empty.
// It has time complexity O(1) if the output stack already has data, else O(n) to transfer input to output.
func (q *QueueAsStack) Dequeue() (string, error) {
	q.Peek()
	return q.output.Pop()
}

// Length retrieves the number of entries in the queue.
func (q *QueueAsStack) Length() int {
	return q.input.Length() + q.output.Length()
}
