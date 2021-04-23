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
