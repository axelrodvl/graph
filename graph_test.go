package graph

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestInitializeGraph(t *testing.T) {
	// G=({a, b, c, d, e}, {(e, a), (a, b), (a, c), (b, d)})

	var edges []Edge
	edges = append(edges,
		Edge{"e", "a", 0},
		Edge{"a", "b", 0},
		Edge{"a", "c", 0},
		Edge{"b", "d", 0})

	fmt.Println("Edges to process:", edges)

	graph := InitializeGraph(&edges)

	fmt.Println("Initialized graph: ", graph)

	assert.Equal(t, 5, len(graph.Vertices), "invalid vertices count")
	assert.Contains(t, graph.Vertices, "a", "graph not contains desired vertex")
	assert.Contains(t, graph.Vertices, "b", "graph not contains desired vertex")
	assert.Contains(t, graph.Vertices, "c", "graph not contains desired vertex")
	assert.Contains(t, graph.Vertices, "d", "graph not contains desired vertex")
	assert.Contains(t, graph.Vertices, "e", "graph not contains desired vertex")
}