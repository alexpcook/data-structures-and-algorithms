package graph

import "fmt"

// Graph represents a graph using the number of nodes and an adjacent list.
type Graph struct {
	numberOfNodes int
	adjacentList  map[int][]int
}

// AddVertex adds a vertex with value node to the graph. It returns a non-nil error
// if the value node already exists in the graph. It has constant time complexity O(1).
func (g *Graph) AddVertex(node int) error {
	if _, exists := g.adjacentList[node]; exists {
		return fmt.Errorf("%d already exists in graph", node)
	}
	g.adjacentList[node] = make([]int, 0)
	g.numberOfNodes++
	return nil
}

// AddEdge adds a connection between node1 and node2 to the graph. It returns a non-nil error
// if either node does not exist in the graph. It has constant time complexity O(1).
func (g *Graph) AddEdge(node1, node2 int) error {
	node1Connections, node1Exists := g.adjacentList[node1]
	node2Connections, node2Exists := g.adjacentList[node2]
	if !node1Exists || !node2Exists {
		return fmt.Errorf("both nodes must exist in order to connect them")
	}
	g.adjacentList[node1] = append(node1Connections, node2)
	g.adjacentList[node2] = append(node2Connections, node1)
	return nil
}
