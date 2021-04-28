package tree

import "fmt"

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

// BSTValueExistsError represents an error when a duplicate value is
// attempted to be added to a BST.
type BSTValueExistsError int

// Error implements the error interface for BSTValueExistsError
func (val BSTValueExistsError) Error() string {
	return fmt.Sprintf("%d already exists in binary search tree", val)
}

// BSTValueNotExistsError represents an error when a non-existent value
// is attempted to be searched for or deleted from a BST.
type BSTValueNotExistsError int

// Error implements the error interface for BSTValueNotExistsError
func (val BSTValueNotExistsError) Error() string {
	return fmt.Sprintf("%d does not exist in binary search tree", val)
}

// BSTDeleteRootNodeError represents an error when only the root node
// exists and is attempted to be removed from a BST.
type BSTDeleteRootNodeError int

// Error implements the error interface for BSTDeleteRootNodeError
func (val BSTDeleteRootNodeError) Error() string {
	return fmt.Sprintf("only root node %d exists, so it cannot be removed", val)
}

// Insert adds the given value to the BST.
// It returns a non-nil error if the value is already present.
// It has time complexity O(log(n)) in the average case and O(n) in the worst case.
func (bst *BinarySearchTree) Insert(value int) error {
	var parentNode **node
	height := 1

	currentNode := bst.root
	for currentNode != nil {
		switch nodeValue := currentNode.value; {
		case value < nodeValue:
			parentNode = &currentNode.left
			currentNode = currentNode.left
		case value > nodeValue:
			parentNode = &currentNode.right
			currentNode = currentNode.right
		default:
			return BSTValueExistsError(value)
		}
		height++
	}

	*parentNode = &node{value, nil, nil}
	if height > bst.height {
		bst.height = height
	}

	return nil
}

// Lookup determines whether the given value is in the BST.
// It returns a non-nil error if the value is not present.
// It has time complexity O(log(n)) in the average case and O(n) in the worst case.
func (bst *BinarySearchTree) Lookup(value int) error {
	currentNode := bst.root
	for currentNode != nil {
		switch nodeValue := currentNode.value; {
		case value < nodeValue:
			currentNode = currentNode.left
		case value > nodeValue:
			currentNode = currentNode.right
		default:
			return nil
		}
	}

	return BSTValueNotExistsError(value)
}

// Delete removes the given value from the BST.
// It returns a non-nil error if the value is not present,
// or if only the root node exists and is attempted to be deleted.
// It has time complexity O(log(n)) in the average case and O(n) in the worst case.
func (bst *BinarySearchTree) Delete(value int) error {
	if value == bst.root.value && bst.root.left == nil && bst.root.right == nil {
		return BSTDeleteRootNodeError(bst.root.value)
	}

	var parentNode **node = &bst.root

	currentNode := bst.root
	for currentNode != nil {
		switch nodeValue := currentNode.value; {
		case value < nodeValue:
			parentNode = &currentNode.left
			currentNode = currentNode.left
		case value > nodeValue:
			parentNode = &currentNode.right
			currentNode = currentNode.right
		default:
			switch hasLeftChild, hasRightChild := currentNode.left != nil, currentNode.right != nil; {
			case hasLeftChild && hasRightChild:
				replaceLeft := currentNode.left
				currentNode = currentNode.right
				replaceRight := currentNode
				var parentNode2 **node = nil
				for currentNode.left != nil {
					parentNode2 = &currentNode.left
					currentNode = currentNode.left
				}
				*parentNode = currentNode
				currentNode.left = replaceLeft
				currentNode.right = replaceRight
				if parentNode2 != nil {
					*parentNode2 = nil
				}
			case hasLeftChild:
				*parentNode = currentNode.left
			case hasRightChild:
				*parentNode = currentNode.right
			default:
				*parentNode = nil
			}
			return nil
		}
	}

	return BSTValueNotExistsError(value)
}

// Height returns the depth of the BST.
func (bst *BinarySearchTree) Height() int {
	return bst.height
}

type bstQueue []*node

func (bstQ *bstQueue) Enqueue(node *node) {
	*bstQ = append(*bstQ, node)
}

func (bstQ *bstQueue) Dequeue() *node {
	if bstQ.IsEmpty() {
		return nil
	}
	val := (*bstQ)[0]
	*bstQ = (*bstQ)[1:]
	return val
}

func (bstQ *bstQueue) IsEmpty() bool {
	return len(*bstQ) == 0
}

func (bstQ *bstQueue) Next() *node {
	if bstQ.IsEmpty() {
		return nil
	}
	return (*bstQ)[0]
}

// BreadthFirstSearch traverses the nodes of the BST in breadth-first order.
// It has time complexity O(n), because we must visit all nodes, and space complexity O(n), because we potentially need to hold all nodes in memory.
func (bst *BinarySearchTree) BreadthFirstSearch() []int {
	result := make([]int, 0)

	queue := new(bstQueue)
	queue.Enqueue(bst.root)
	for !queue.IsEmpty() {
		currentNode := queue.Dequeue()
		if currentNode.left != nil {
			queue.Enqueue(currentNode.left)
		}
		if currentNode.right != nil {
			queue.Enqueue(currentNode.right)
		}
		result = append(result, currentNode.value)
	}

	return result
}

// BreadthFirstSearchRecursive traverses the nodes of the BST in breadth-first order using recursion.
// It has time complexity O(n), because we must visit all nodes, and space complexity O(n), because we potentially need to hold all nodes in memory.
func (bst *BinarySearchTree) BreadthFirstSearchRecursive() []int {
	if bst.root.left == nil && bst.root.right == nil { // Base recursion case is when the entire BST is a leaf
		return nil
	}

	result := make([]int, 0)
	result = append(result, bst.root.value)

	var leftResult, rightResult []int
	if bst.root.left != nil {
		result = append(result, bst.root.left.value)
		leftBST := &BinarySearchTree{root: bst.root.left}
		leftResult = leftBST.BreadthFirstSearchRecursive()
		if leftResult != nil {
			leftResult = leftResult[1:] // Don't double count
		}
	}
	if bst.root.right != nil {
		result = append(result, bst.root.right.value)
		rightBST := &BinarySearchTree{root: bst.root.right}
		rightResult = rightBST.BreadthFirstSearchRecursive()
		if rightResult != nil {
			rightResult = rightResult[1:] // Don't double count
		}
	}

	result = append(result, leftResult...)
	result = append(result, rightResult...)

	return result
}

// DepthFirstSearchInOrder traverses the nodes of the BST using in-order depth-first search.
// It has time complexity O(n), because we must visit all nodes, and space complexity O(h), where h is the height of the tree.
func (bst *BinarySearchTree) DepthFirstSearchInOrder() []int {
	result := make([]int, 0)

	// Append the left side of the BST
	if bst.root.left != nil {
		leftBST := &BinarySearchTree{root: bst.root.left}
		result = append(result, leftBST.DepthFirstSearchInOrder()...)
	}

	// Append the root
	result = append(result, bst.root.value)

	// Append the right side of the BST
	if bst.root.right != nil {
		rightBST := &BinarySearchTree{root: bst.root.right}
		result = append(result, rightBST.DepthFirstSearchInOrder()...)
	}

	return result
}

// DepthFirstSearchPreOrder traverses the nodes of the BST using pre-order depth-first search.
// It has time complexity O(n), because we must visit all nodes, and space complexity O(h), where h is the height of the tree.
func (bst *BinarySearchTree) DepthFirstSearchPreOrder() []int {
	result := make([]int, 0)

	// Append the root value
	result = append(result, bst.root.value)

	// Append the left side of the BST
	if bst.root.left != nil {
		leftBST := &BinarySearchTree{root: bst.root.left}
		result = append(result, leftBST.DepthFirstSearchPreOrder()...)
	}

	// Append the right side of the BST
	if bst.root.right != nil {
		rightBST := &BinarySearchTree{root: bst.root.right}
		result = append(result, rightBST.DepthFirstSearchPreOrder()...)
	}

	return result
}

// DepthFirstSearchPostOrder traverses the nodes of the BST using post-order depth-first search.
// It has time complexity O(n), because we must visit all nodes, and space complexity O(h), where h is the height of the tree.
func (bst *BinarySearchTree) DepthFirstSearchPostOrder() []int {
	result := make([]int, 0)

	// Append the left BST
	if bst.root.left != nil {
		leftBST := &BinarySearchTree{root: bst.root.left}
		result = append(result, leftBST.DepthFirstSearchPostOrder()...)
	}

	// Append the right BST
	if bst.root.right != nil {
		rightBST := &BinarySearchTree{root: bst.root.right}
		result = append(result, rightBST.DepthFirstSearchPostOrder()...)
	}

	// Append the root
	result = append(result, bst.root.value)

	return result
}

// IsValid determines whether the BST follows the expected rules:
// - The left subtree of a node contains only nodes with keys less than the node's key.
// - The right subtree of a node contains only nodes with keys greater than the node's key.
// - Both the left and right subtrees must also be binary search trees.
func (bst *BinarySearchTree) IsValid() bool {
	queue := new(bstQueue)
	queue.Enqueue(bst.root)

	for !queue.IsEmpty() {
		currentNode := queue.Dequeue()

		if currentNode.left != nil {
			if currentNode.value <= currentNode.left.value {
				return false
			}
			queue.Enqueue(currentNode.left)
		}

		if currentNode.right != nil {
			if currentNode.value >= currentNode.right.value {
				return false
			}
			queue.Enqueue(currentNode.right)
		}
	}

	return true
}
