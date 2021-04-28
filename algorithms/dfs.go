package algorithms

// There are several ways to return the order of elements in depth-first search:
// 		101
//	33		105
// In order: 33, 101, 105 (returns entries in sequential order for a BST)
// Pre-order: 101, 33, 105 (returns parent before children, useful for recreating tree order)
// Post-order: 33, 105, 101 (returns children before parent)
//
// See tree.(*BinarySearchTree).DepthFirstSearchInOrder()
// See tree.(*BinarySearchTree).DepthFirstSearchPreOrder()
// tree.(*BinarySearchTree).DepthFirstSearchPostOrder()
