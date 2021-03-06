package linkedlist

import "fmt"

type LinkedList struct {
	head   *node
	tail   *node
	length int
}

type node struct {
	value int
	next  *node
}

func new_node(value int, next *node) *node {
	return &node{value, next}
}

// NewLinkedList creates a LinkedList with initial value head
// and an optional list of data to append after head.
// It has time complexity O(n) when it needs to loop over data, otherwise O(1).
func NewLinkedList(head int, data ...int) *LinkedList {
	linkedList := new(LinkedList)
	linkedList.length = len(data) + 1
	linkedList.head = new_node(head, nil)

	currentNode := linkedList.head
	for _, d := range data {
		currentNode.next = new_node(d, nil)
		currentNode = currentNode.next
	}
	linkedList.tail = currentNode

	return linkedList
}

// String implements the fmt.Stringer interface for LinkedList.
// It has time complexity O(n).
func (ll LinkedList) String() string {
	str := ""

	currentNode := ll.head
	for currentNode != nil {
		str += fmt.Sprintf("[%p] %v --> ", currentNode, currentNode.value)
		currentNode = currentNode.next
	}

	return str
}

// Append adds data to the tail of *LinkedList.
// It has time complexity O(1).
func (ll *LinkedList) Append(data int) {
	newTail := new_node(data, nil)
	ll.tail.next = newTail
	ll.tail = newTail
	ll.length++
}

// Prepend adds data to the head of *LinkedList.
// It has time complexity O(1).
func (ll *LinkedList) Prepend(data int) {
	newHead := new_node(data, ll.head)
	ll.head = newHead
	ll.length++
}

// Insert adds data to *LinkedList at position index.
// It has time complexity O(n).
func (ll *LinkedList) Insert(data, index int) error {
	if index < 0 || index > ll.length {
		return fmt.Errorf("invalid index value %d", index)
	}

	prevNode := ll.get_node(index - 1)

	if prevNode == nil {
		ll.Prepend(data)
	} else if prevNode.next == nil {
		ll.Append(data)
	} else {
		prevNode.next = new_node(data, prevNode.next)
		ll.length++
	}

	return nil
}

// Delete removes the specified index from *LinkedList.
// It has time complexity O(n).
func (ll *LinkedList) Delete(index int) error {
	if index < 0 || index > ll.length-1 {
		return fmt.Errorf("invalid index value %d", index)
	}

	if ll.length == 1 {
		return fmt.Errorf("cannot delete when only one value remains")
	}

	prevNode := ll.get_node(index - 1)

	if prevNode == nil {
		ll.head = ll.head.next
	} else {
		prevNode.next = prevNode.next.next
	}

	ll.length--
	if ll.length == 1 {
		ll.tail = ll.head
	} else if ll.length == index {
		ll.tail = prevNode
	}

	return nil
}

// Lookup returns the value at the specified index in *LinkedList.
// It has time complexity O(n).
func (ll *LinkedList) Lookup(index int) (int, error) {
	if index < 0 || index > ll.length-1 {
		return 0, fmt.Errorf("invalid index value %d", index)
	}
	return ll.get_node(index).value, nil
}

func (ll *LinkedList) get_node(index int) *node {
	if index < 0 || index > ll.length-1 {
		return nil
	}

	node := ll.head
	i := 0
	for i != index {
		node = node.next
		i++
	}

	return node
}

// Reverse switches the order of the linked list. It modifies the pointer receiver.
// It has time and space complexity O(n).
func (ll *LinkedList) Reverse() {
	entries := make([]int, 0, ll.length)
	currentNode := ll.head
	for currentNode != nil { // Time complexity O(n)
		entries = append(entries, currentNode.value) // Space complexity O(n)
		currentNode = currentNode.next
	}

	llNew := NewLinkedList(entries[ll.length-1])
	for i := ll.length - 2; i > -1; i-- { // Time complexity O(n)
		llNew.Append(entries[i])
	}

	*ll = *llNew
}
