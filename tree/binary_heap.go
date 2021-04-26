package tree

// Heaps are used when ordering is important, such as priority queues.
// Binary heaps contain at most two children and can be min at the top or max at the top
// Binary heaps are by nature balanced because they do left-to-right insertion
// - Lookup O(n), because there's no left/right relationship like in a binary tree
// - Insert O(log(n)), because value might have to bubble up to the top of the tree (O(1) in best case)
// - Delete O(log(n)), because tree might have to be rearranged depending on which value is deleted
