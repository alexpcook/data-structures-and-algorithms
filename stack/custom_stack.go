package stack

import "fmt"

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

func (s1 *Stack1) String() string {
	str := ""

	entry := s1.top
	for entry != nil {
		str += fmt.Sprintln(entry.value)
		entry = entry.next
	}

	return str
}

type StackIsEmptyError struct{}

func (s StackIsEmptyError) Error() string {
	return "stack is empty"
}

// Peek gets the top entry of the stack. It returns a non-nil error if the stack is empty.
// It has time complexity O(1).
func (s1 *Stack1) Peek() (string, error) {
	if s1.top == nil {
		return "", StackIsEmptyError{}
	}
	return s1.top.value, nil
}

// Push adds an entry to the top of the stack.
// It has time complexity O(1).
func (s1 *Stack1) Push(entry string) {
	s1.top = &node{entry, s1.top}
	if s1.length == 0 {
		s1.bottom = s1.top
	}
	s1.length++
}

// Pop removes the entry at the top of the stack and returns it. It returns a non-nil error if the stack is empty.
// It has time complexity O(1).
func (s1 *Stack1) Pop() (string, error) {
	if s1.top == nil {
		return "", StackIsEmptyError{}
	}

	poppedValue := s1.top.value
	s1.top = s1.top.next
	if s1.length == 1 {
		s1.bottom = nil
	}
	s1.length--

	return poppedValue, nil
}

// Stack2 implements a stack data structure using slices.
// Use new(stack.Stack2) to create a new instance.
type Stack2 struct {
	data   []string
	length int
}

func (s2 *Stack2) String() string {
	str := ""

	for i := s2.length - 1; i > -1; i-- {
		str += fmt.Sprintln(s2.data[i])
	}

	return str
}

// Peek gets the top entry of the stack. It returns a non-nil error if the stack is empty.
// It has time complexity O(1).
func (s2 *Stack2) Peek() (string, error) {
	if s2.data == nil {
		return "", StackIsEmptyError{}
	}
	return s2.data[s2.length-1], nil
}

// Push adds an entry to the top of the stack.
// It has time complexity O(1).
func (s2 *Stack2) Push(entry string) {
	if s2.data == nil {
		s2.data = make([]string, 0, 2) // Account for future growth with double the needed capacity
	}
	s2.data = append(s2.data, entry)
	s2.length++
}

// Pop removes the entry at the top of the stack and returns it. It returns a non-nil error if the stack is empty.
// It has time complexity O(1).
func (s2 *Stack2) Pop() (string, error) {
	if s2.data == nil {
		return "", StackIsEmptyError{}
	}

	poppedValue := s2.data[s2.length-1]
	if s2.length == 1 {
		s2.data = nil
	} else {
		s2.data = s2.data[:s2.length-1]
	}
	s2.length--

	return poppedValue, nil
}
