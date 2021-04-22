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
// It has time complexity O(1), assuming that the _hash function is extremely fast and there are no collisions.
// If there are collisions, it has time complexity O(n).
func (ht *HashTable) Get(key string) string {
	node := ht.data[ht._hash(key)]
	for i := 0; i < len(node)-1; i += 2 {
		if node[i] == key {
			return node[i+1] // If no collisions, we always get here upon first iteration (i == 0), so time complexity is O(1).
		}
	}
	return ""
}

// Set sets a key/value pair of the hash table.
// It has time complexity O(1), assuming that the _hash function is extremely fast and there are no collisions.
// If there are collisions, it has time complexity O(n).
func (ht *HashTable) Set(key, value string) {
	hash := ht._hash(key)
	if ht.data[hash] == nil {
		ht.data[hash] = make([]string, 2)
		ht.data[hash][0] = key
		ht.data[hash][1] = value
	} else {
		for i := 0; i < len(ht.data[hash])-1; i += 2 {
			if ht.data[hash][i] == key {
				ht.data[hash][i+1] = value
				return // If no collisions, we always get here upon first iteration (i == 0), so time complexity is O(1).
			}
		}
		ht.data[hash] = append(ht.data[hash], key, value) // This is the O(n) case when collisions occur.
	}
}

// Delete removes a key/value pair of the hash table.
// It has time complexity O(1), assuming that the _hash function is extremely fast and there are no collisions.
// If there are collisions, it has time complexity O(n).
func (ht *HashTable) Delete(key string) {
	hash := ht._hash(key)
	for i := 0; i < len(ht.data[hash])-1; i += 2 {
		if ht.data[hash][i] == key {
			ht.data[hash] = append(ht.data[hash][:i], ht.data[hash][i+2:]...)
			return // If no collisions, we always get here upon first iteration (i == 0), so time complexity is O(1).
		}
	}
}

// String implements the fmt.Stringer interface on *HashTable
func (ht *HashTable) String() string {
	str := "["
	for _, hash := range ht.data {
		for i := 0; i < len(hash)-1; i += 2 {
			str += fmt.Sprintf(" %s:%s", hash[i], hash[i+1])
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
