package tree

// Binary search tree (assuming it's balanced)
// - Lookup O(log(n))
// - Insert O(log(n))
// - Delete O(log(n))
//
// If a binary search tree is not balanced (worst case)
// - Lookup O(n)
// - Insert O(n)
// - Delete O(n)
//
// Number of nodes = 2^h - 1 where h is the height of the tree
// Number of traversals to find a given node = log(n)
//
// Rules of binary search trees
// - Any node can have up to two children
// - The left child node is less than the current node
// - The right child node is greater than the current node

type node struct {
	value int
	left  *node
	right *node
}

// BinarySearchTree represents a BST.
type BinarySearchTree struct {
	root   *node
	height int
}

// NewBinarySearchTree creates a BST with the given root value and returns a pointer.
func NewBinarySearchTree(rootValue int) *BinarySearchTree {
	return &BinarySearchTree{
		&node{rootValue, nil, nil},
		1,
	}
}

// Insert adds the given value to the BST.
// It has time complexity O(log(n)) in the average case and O(n) in the worst case.
func (bst *BinarySearchTree) Insert(value int) {

}

// Lookup determines whether the given value is in the BST.
// It returns a non-nil error if the value is not present.
// It has time complexity O(log(n)) in the average case and O(n) in the worst case.
func (bst *BinarySearchTree) Lookup(value int) error {
	return nil
}

// Delete removes the given value from the BST.
// It returns a non-nil error if the value is not present.
// It has time complexity O(log(n)) in the average case and O(n) in the worst case.
func (bst *BinarySearchTree) Delete(value int) error {
	return nil
}
