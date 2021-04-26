package graph

import (
	"reflect"
	"testing"
)

func TestAddVertex(t *testing.T) {
	tests := []struct {
		node    int
		isError bool
		nodes   int
		want    map[int][]int
	}{
		{1, false, 1, map[int][]int{1: {}}},
		{2, false, 2, map[int][]int{1: {}, 2: {}}},
		{7, false, 3, map[int][]int{1: {}, 2: {}, 7: {}}},
		{7, true, 3, map[int][]int{1: {}, 2: {}, 7: {}}},
		{1, true, 3, map[int][]int{1: {}, 2: {}, 7: {}}},
	}

	graph := &Graph{0, make(map[int][]int)}
	for i, test := range tests {
		err := graph.AddVertex(test.node)
		if test.isError {
			if err == nil {
				t.Fatalf("test %d: want error, got nil", i)
			}
		} else {
			if test.nodes != graph.numberOfNodes {
				t.Fatalf("test %d: want %d nodes, got %d nodes", i, test.nodes, graph.numberOfNodes)
			}
			if !reflect.DeepEqual(test.want, graph.adjacentList) {
				t.Fatalf("test %d: want %v graph, got %v graph", i, test.want, graph.adjacentList)
			}
		}
	}
}

func TestAddEdge(t *testing.T) {
	graph := &Graph{0, make(map[int][]int)}
	for _, datum := range []int{1, 2, 3} {
		err := graph.AddVertex(datum)
		if err != nil {
			t.Fatalf("error adding %d vertex to graph", datum)
		}
	}

	tests := []struct {
		node1, node2 int
		want1, want2 []int
		isError      bool
	}{
		{0, 1, nil, nil, true},
		{1, 0, nil, nil, true},
		{1, 2, []int{2}, []int{1}, false},
		{3, 1, []int{1}, []int{2, 3}, false},
	}

	for i, test := range tests {
		err := graph.AddEdge(test.node1, test.node2)
		if test.isError {
			if err == nil {
				t.Fatalf("test %d: want error, got nil", i)
			}
		} else {
			if !reflect.DeepEqual(test.want1, graph.adjacentList[test.node1]) {
				t.Fatalf("want node1 adjacent list %v, got node1 adjacent list %v", test.want1, graph.adjacentList[test.node1])
			}
			if !reflect.DeepEqual(test.want2, graph.adjacentList[test.node2]) {
				t.Fatalf("want node2 adjacent list %v, got node2 adjacent list %v", test.want2, graph.adjacentList[test.node2])
			}
		}
	}
}
