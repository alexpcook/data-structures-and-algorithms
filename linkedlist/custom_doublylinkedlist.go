package linkedlist

import "fmt"

type DoublyLinkedList struct {
	head   *doublynode
	tail   *doublynode
	length int
}

type doublynode struct {
	value    int
	next     *doublynode
	previous *doublynode
}

func new_doublynode(value int, next *doublynode, previous *doublynode) *doublynode {
	return &doublynode{value, next, previous}
}

// NewDoublyLinkedList creates a DoublyLinkedList with initial value head
// and an optional list of data to append after head.
// It has time complexity O(n) when it needs to loop over data, otherwise O(1).
func NewDoublyLinkedList(head int, data ...int) *DoublyLinkedList {
	dll := new(DoublyLinkedList)
	dll.length = len(data) + 1
	dll.head = new_doublynode(head, nil, nil)

	currentNode := dll.head
	for _, d := range data {
		currentNode.next = new_doublynode(d, nil, currentNode)
		currentNode = currentNode.next
	}
	dll.tail = currentNode

	return dll
}

// String implements the fmt.Stringer interface for DoublyLinkedList.
// It has time complexity O(n).
func (dll DoublyLinkedList) String() string {
	str := ""

	currentNode := dll.head
	for currentNode != nil {
		str += fmt.Sprintf("[%p] %v --> ", currentNode, currentNode.value)
		currentNode = currentNode.next
	}

	return str
}
