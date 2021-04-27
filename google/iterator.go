// Implement solution to sample problem (https://techdevguide.withgoogle.com/resources/former-interview-question-flatten-iterators/).
package google

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() int
}

type intIteratorNode struct {
	value int
	next  *intIteratorNode
}

type IntIterator struct {
	head *intIteratorNode
}

// O(n), where n is the length of data
func NewIntIterator(data []int) *IntIterator {
	if len(data) == 0 {
		return new(IntIterator)
	}

	iter := &IntIterator{&intIteratorNode{data[0], nil}}
	curNode := iter.head

	for i := 1; i < len(data); i++ {
		curNode.next = &intIteratorNode{data[i], nil}
		curNode = curNode.next
	}

	return iter
}

// O(n), where n is the number of entries in the linked list
func (itr *IntIterator) String() string {
	str := "["
	currentNode := itr.head
	for currentNode != nil {
		str += fmt.Sprintf("%d-->", currentNode.value)
		currentNode = currentNode.next
	}
	return str + "]"
}

// O(1)
func (itr *IntIterator) HasNext() bool {
	return itr.head != nil
}

// O(1)
func (itr *IntIterator) Next() (int, error) {
	if itr.head == nil {
		return 0, fmt.Errorf("empty iterator")
	}
	defer func() { itr.head = itr.head.next }()
	return itr.head.value, nil
}

type intInterleaverEntry struct {
	iterator IntIterator
	next     *intInterleaverEntry
}

type IntInterleaver struct {
	first, last *intInterleaverEntry
	length      int
}

// O(n), where n is the number of iterators passed to the constructor
func NewIntInterleaver(iterators ...IntIterator) *IntInterleaver {
	interleaver := new(IntInterleaver)

	for _, iter := range iterators {
		if iter.HasNext() {
			entry := &intInterleaverEntry{iter, nil}
			if interleaver.last != nil {
				interleaver.last.next = entry
			}
			interleaver.last = entry
			if interleaver.length == 0 {
				interleaver.first = interleaver.last
			}
			interleaver.length++
		}
	}

	return interleaver
}

// O(n), where n is the number of iterators in the queue
// func (itl *IntInterleaver) String() string {
// 	str := ""
// 	currentEntry := itl.first
// 	for currentEntry != nil {
// 		str += fmt.Sprint(currentEntry.iterator)
// 		currentEntry = currentEntry.next
// 	}
// 	return str
// }

// O(1)
func (itl *IntInterleaver) HasNext() bool {
	return itl.length > 0
}

// O(1)
func (itl *IntInterleaver) Next() (int, error) {
	if itl.length == 0 {
		return 0, fmt.Errorf("empty interleaver")
	}

	entry := itl.first
	itl.first = itl.first.next
	itl.length--

	defer func() {
		if entry.iterator.HasNext() {
			if itl.last != nil {
				itl.last.next = entry
			}
			itl.last = entry
			if itl.length == 0 {
				itl.first = itl.last
			}
			itl.length++
		}
	}()

	return entry.iterator.Next()
}
