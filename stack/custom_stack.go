package stack

type node struct {
	value string
	next  *node
}

// Stack1 implements a stack data structure using linked lists.
// Use new(stack.Stack1) to create a new instance.
type Stack1 struct {
	top    *node
	bottom *node
	length int
}

// Peek gets the top entry of the stack.
// It has time complexity O(1).
func (s1 *Stack1) Peek() string {
	return ""
}

// Push adds an entry to the top of the stack.
// It has time complexity O(1).
func (s1 *Stack1) Push(entry string) {
}

// Pop removes the entry at the top of the stack and returns it.
// It has time complexity O(1).
func (s1 *Stack1) Pop() string {
	return ""
}
