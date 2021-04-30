package algorithms

import (
	"sync"
	"time"
)

// Dynamic programming uses caching to store the results of calculations for future use
// Here's a simple example
type cache struct {
	data  map[int]int
	mutex sync.Mutex
}

// O(1)
func (c *cache) entryExists(n int) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, exists := c.data[n]
	return exists
}

// O(1)
func (c *cache) getEntry(n int) int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.data[n]
}

// O(1)
func (c *cache) addEntry(n, result int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[n] = result
}

// AddTo80 simulates an expensive add 80 operation.
func AddTo80(n int) int {
	time.Sleep(time.Second * 4) // Simulate an expensive calculation
	return n + 80
}

// MemorizedAddTo80 returns the cached value of the add 80 operation if it exists, else it computes add 80.
func MemorizedAddTo80() func(int) int {
	var addTo80Cache cache = cache{make(map[int]int), sync.Mutex{}}

	return func(n int) int {
		if addTo80Cache.entryExists(n) {
			return addTo80Cache.getEntry(n)
		}
		result := AddTo80(n)
		addTo80Cache.addEntry(n, result)
		return result
	}
}

// FibonacciRecursiveMemorized uses recursion and dynamic programming to implement an O(n) time complexity
// Fibonacci sequence.
func FibonacciRecursiveMemorized() func(int) int {
	fibCache := make(map[int]int)

	var fibMemFunc func(int) int
	fibMemFunc = func(n int) int {
		if cachedVal, exists := fibCache[n]; exists {
			return cachedVal
		}
		if n < 3 {
			return n - 1
		}
		fibN := fibMemFunc(n-1) + fibMemFunc(n-2)
		fibCache[n] = fibN
		return fibN
	}

	return fibMemFunc
}
