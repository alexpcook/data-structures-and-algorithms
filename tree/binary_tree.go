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
