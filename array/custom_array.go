package array

import "fmt"

// Array is a custom array type implemented from scratch.
type Array struct {
	Length int
	Data   []interface{}
}

// NewArray creates a new Array type populated with data and returns a pointer.
func NewArray(data ...interface{}) *Array {
	return &Array{len(data), data}
}

func (a *Array) String() string {
	return fmt.Sprintf("%v", a.Data)
}

// Get retrieves an item from the array by index.
func (a *Array) Get(index int) interface{} {
	return a.Data[index]
}

// Push adds a value to the end of the array, and returns the array's new length.
func (a *Array) Push(value interface{}) int {
	a.Data = append(a.Data, value)
	a.Length++
	return a.Length
}

// Pop removes the value at the end of the array and returns the removed value.
func (a *Array) Pop() interface{} {
	defer func() {
		a.Data = a.Data[:a.Length-1]
		a.Length--
	}()
	return a.Data[a.Length-1]
}

// Delete removes the value at the specified index from the array.
func (a *Array) Delete(index int) interface{} {
	defer func() {
		a.Data = append(a.Data[:index], a.Data[index+1:]...)
		a.Length--
	}()
	return a.Data[index]
}
