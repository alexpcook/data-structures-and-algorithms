package hashtable

import (
	"fmt"
)

// HashTable implements a custom hash table type with string key/value pairs.
type HashTable struct {
	size int
	data [][]string
}

// NewHashTable creates a hash table with a given number of entries and returns a pointer.
func NewHashTable(size int) *HashTable {
	return &HashTable{size, make([][]string, size)}
}

// Get returns a value from the hash table given a key.
// It has time complexity O(1), assuming that the _hash function is extremely fast.
func (ht *HashTable) Get(key string) string {
	node := ht.data[ht._hash(key)]
	if node == nil {
		return ""
	}
	return node[1]
}

// Set sets a key/value pair of the hash table.
// It has time complexity O(1), assuming that the _hash function is extremely fast.
func (ht *HashTable) Set(key, value string) {
	hash := ht._hash(key)
	if ht.data[hash] == nil {
		ht.data[hash] = make([]string, 2)
	}
	ht.data[hash][0] = key
	ht.data[hash][1] = value
}

// Delete removes a key/value pair of the hash table.
// It has time complexity O(1), assuming that the _hash function is extremely fast.
func (ht *HashTable) Delete(key string) {
	ht.data[ht._hash(key)] = nil
}

// String implements the fmt.Stringer interface on *HashTable
func (ht *HashTable) String() string {
	str := "["
	for _, kv := range ht.data {
		if kv != nil {
			str += fmt.Sprintf(" %s: %s", kv[0], kv[1])
		}
	}
	return str + " ]"
}

// _hash returns a unique position in the underlying data array to store a given key.
func (ht *HashTable) _hash(key string) int {
	hash := 0
	for i, r := range key {
		hash = (hash + int(r)*i) % ht.size
	}
	return hash
}
