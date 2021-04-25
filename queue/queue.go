package queue

// Queues operate on a first-in, first-out (FIFO) paradigm
// - Lookup O(n) (not a typical operation)
// - Enqueue O(1) (similar to push, adds an entry to the queue)
// - Dequeue O(1) (similar to pop, removes the entry that was added first to the queue)
// - Peek O(1) (view the first item to come out)

// Queues can be implemented using either arrays or linked lists
// Linked lists would be best because you can enqueue/dequeue entries to the head/tail
// Arrays would require shifting entries, which has O(n) time complexity, so they are NOT a good choice
