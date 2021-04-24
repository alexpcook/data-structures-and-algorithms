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

// Append adds data to the tail of *DoublyLinkedList.
// It has time complexity O(1).
func (dll *DoublyLinkedList) Append(data int) {
	newTail := new_doublynode(data, nil, dll.tail)
	dll.tail.next = newTail
	dll.tail = newTail
	dll.length++
}

// Prepend adds data to the head of *DoublyLinkedList.
// It has time complexity O(1).
func (dll *DoublyLinkedList) Prepend(data int) {
	newHead := new_doublynode(data, dll.head, nil)
	dll.head.previous = newHead
	dll.head = newHead
	dll.length++
}

// Insert adds data to *DoublyLinkedList at position index.
// It has time complexity O(n/2).
func (dll *DoublyLinkedList) Insert(data, index int) error {
	if index < 0 || index > dll.length {
		return fmt.Errorf("invalid index value %d", index)
	}

	prevNode := dll.get_node(index - 1)

	if prevNode == nil {
		dll.Prepend(data)
	} else if prevNode.next == nil {
		dll.Append(data)
	} else {
		nextNode := prevNode.next
		prevNode.next = new_doublynode(data, nextNode, prevNode)
		nextNode.previous = prevNode.next
		dll.length++
	}

	return nil
}

// Lookup returns the value at the specified index in *DoublyLinkedList.
// It has time complexity O(n/2).
func (dll *DoublyLinkedList) Lookup(index int) (int, error) {
	if index < 0 || index > dll.length-1 {
		return 0, fmt.Errorf("invalid index value %d", index)
	}
	return dll.get_node(index).value, nil
}

func (dll *DoublyLinkedList) get_node(index int) *doublynode {
	if index < 0 || index > dll.length-1 {
		return nil
	}

	var node *doublynode

	// This halves the time complexity of lookup, insert, and delete for
	// doubly-linked lists compared to singly-linked lists
	indexIsInFirstHalf := dll.length/2 > index
	if indexIsInFirstHalf {
		node = dll.head
		i := 0
		for i != index {
			node = node.next
			i++
		}
	} else {
		node = dll.tail
		i := dll.length - 1
		for i != index {
			node = node.previous
			i--
		}
	}

	return node
}
