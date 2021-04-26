// Implement solution to sample problem (https://techdevguide.withgoogle.com/resources/former-interview-question-flatten-iterators/).
package google

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() int
}

type node struct {
	value int
	next  *node
}

type IntIterator struct {
	head *node
}

// O(1)
func NewIntIterator(data []int) (*IntIterator, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("got no data")
	}
	iter := &IntIterator{&node{data[0], nil}}
	curNode := iter.head
	for i := 1; i < len(data); i++ {
		curNode.next = &node{data[i], nil}
		curNode = curNode.next
	}
	return iter, nil
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

type IntInterleaver struct {
	data  []IntIterator
	index int
}

// O(1)
func NewIntInterleaver(iterators ...IntIterator) (*IntInterleaver, error) {
	if len(iterators) == 0 {
		return nil, fmt.Errorf("got not iterators")
	}
	return &IntInterleaver{iterators, 0}, nil
}

// O(itl.data)
func (itl *IntInterleaver) HasNext() bool {
	for i, j := itl.index, 0; j < len(itl.data); i, j = (i+1)%len(itl.data), j+1 {
		if itl.data[i].HasNext() {
			return true
		}
	}
	return false
}

// O(itl.data)
func (itl *IntInterleaver) Next() (int, error) {
	for j := 0; j < len(itl.data); itl.index, j = (itl.index+1)%len(itl.data), j+1 {
		if itl.data[itl.index].HasNext() {
			defer func() {
				itl.index++
				itl.index %= len(itl.data)
			}()
			return itl.data[itl.index].Next()
		}
	}
	return 0, nil
}
