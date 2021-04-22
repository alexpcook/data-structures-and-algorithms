package hashtable

import "fmt"

// Time complexity of hash table operations
// - Insert O(1)
// - Lookup O(1)
// - Delete O(1)
// - Search O(1)

// BasicHashTableOperations illustrates some basic hash table (map) operations.
// Ignoring the problem of collisions, operations on hash tables generally
// have constant time complexity O(1).
func BasicHashTableOperations() {
	phonebook := map[string]string{
		"John Doe":  "123-456-7890",
		"Sally Doe": "098-765-4321",
	}

	// Insert O(1)
	phonebook["Fred Doe"] = "123-234-3456"

	// Lookup O(1)
	fmt.Println(phonebook["John Doe"])

	// Delete O(1)
	delete(phonebook, "Fred Doe")
	fmt.Println(phonebook)

	// Search O(1) is the same time complexity as lookup because
	// no iteration is required to find the value given a key
	fmt.Println(phonebook["Sally Doe"])
}
