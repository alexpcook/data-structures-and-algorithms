package graph

// Graphs are made up of nodes (vertexes)
// Nodes are connected using edges
// Graphs are great for representing relationships
// Graphs are a more general version of trees, which are a more general version of linked lists
//
// Graphs can be directed (unidirectional, e.g. one-way street) or undirected (bidirectional, e.g. highway)
// Graph edges can be weighted or unweighted
// Graphs can be cyclic or acyclic

// How to think about building graphs
// - An edge list shows the connections between nodes
func EdgeList() {
	// 0 is connected to 2
	// 2 is connected to 3
	// 2 is connected to 1
	// 1 is connected to 3
	graph := [][]int{{0, 2}, {2, 3}, {2, 1}, {1, 3}}
	_ = graph
}

// - An adjacent list contains each node and its neighbors
func AdjacentList() {
	// 0 is adjacent to 2
	// 1 is adjacent to 2, 3
	// 2 is adjacent to 0, 1, 3
	// 3 is adjacent to 1, 2
	graph := [][]int{{2}, {2, 3}, {0, 1, 3}, {1, 2}}
	_ = graph
}

// - An adjacent matrix lists whether each node is connected to all others
func AdjacentMatrix() {
	graph := [][]int{
		{0, 0, 1, 0}, // 0 is connected to 2
		{0, 0, 1, 1}, // 1 is connected to 2, 3
		{1, 1, 0, 1}, // 2 is connected to 0, 1, 3
		{0, 1, 1, 0}, // 3 is connected to 1, 2
	}
	_ = graph
}
