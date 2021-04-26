// Implement solution to sample problem (https://techdevguide.withgoogle.com/resources/former-interview-question-flatten-iterators/).
package google

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() int
}

type IntIterator struct {
	data  []int
	index int
}

// O(1)
func NewIntIterator(data []int) (*IntIterator, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("got no data")
	}
	return &IntIterator{data, 0}, nil
}

// O(1)
func (itr *IntIterator) HasNext() bool {
	return itr.index < len(itr.data)
}

// O(1)
func (itr *IntIterator) Next() int {
	if itr.index > len(itr.data)-1 {
		return 0
	}
	defer func() {
		itr.index++
	}()
	return itr.data[itr.index]
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
func (itl *IntInterleaver) Next() int {
	for j := 0; j < len(itl.data); itl.index, j = (itl.index+1)%len(itl.data), j+1 {
		if itl.data[itl.index].HasNext() {
			defer func() {
				itl.index++
				itl.index %= len(itl.data)
			}()
			return itl.data[itl.index].Next()
		}
	}
	return 0
}
