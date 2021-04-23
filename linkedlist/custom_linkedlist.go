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

func (ll LinkedList) String() string {
	str := ""

	currentNode := ll.head
	for currentNode != nil {
		str += fmt.Sprintf("[%p] %v --> ", currentNode, currentNode.value)
		currentNode = currentNode.next
	}

	return str
}
