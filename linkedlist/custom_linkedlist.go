package linkedlist

import "fmt"

type LinkedList struct {
	head   *linkedListNode
	tail   *linkedListNode
	length int
}

type linkedListNode struct {
	value int
	next  *linkedListNode
}

// NewLinkedList creates a LinkedList with initial value head
// and an optional list of data to append after head.
// It has time complexity O(n) when it needs to loop over data, otherwise O(1).
func NewLinkedList(head int, data ...int) *LinkedList {
	linkedList := new(LinkedList)
	linkedList.length = len(data) + 1
	linkedList.head = &linkedListNode{head, nil}

	currentNode := linkedList.head
	for i := 0; i < len(data); i++ {
		currentNode.next = &linkedListNode{data[i], nil}
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
	newTail := &linkedListNode{data, nil}
	ll.tail.next = newTail
	ll.tail = newTail
	ll.length++
}

// Prepend adds data to the head of *LinkedList.
// It has time complexity O(1).
func (ll *LinkedList) Prepend(data int) {
	newHead := &linkedListNode{data, ll.head}
	ll.head = newHead
	ll.length++
}

// Insert adds data to *LinkedList at position index.
// It has time complexity O(n), except in the special cases of
// prepend (index == 0) and append (index == *LinkedList.length).
func (ll *LinkedList) Insert(data, index int) error {
	if index < 0 || index > ll.length {
		return fmt.Errorf("invalid index value %d", index)
	}

	if index == 0 {
		ll.Prepend(data)
		return nil
	} else if index == ll.length {
		ll.Append(data)
		return nil
	}

	var currNode *linkedListNode = ll.head.next
	var prevNode *linkedListNode = ll.head

	for i := 1; i < ll.length; i++ {
		if i == index {
			newNode := &linkedListNode{data, currNode}
			prevNode.next = newNode
			break
		}
		prevNode = currNode
		currNode = currNode.next
	}

	ll.length++

	return nil
}
