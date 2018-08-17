package graph

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestDijkstra(t *testing.T) {
	var edges []Edge
	edges = append(edges,
		Edge{"1", "2", 7},
		Edge{"2", "1", 7},

		Edge{"2", "3", 10},
		Edge{"3", "2", 10},

		Edge{"1", "3", 9},
		Edge{"3", "1", 9},

		Edge{"2", "4", 15},
		Edge{"4", "2", 15},

		Edge{"1", "6", 14},
		Edge{"6", "1", 14},

		Edge{"3", "6", 2},
		Edge{"6", "3", 2},

		Edge{"3", "4", 11},
		Edge{"4", "3", 11},

		Edge{"6", "5", 9},
		Edge{"5", "6", 9},

		Edge{"4", "5", 6},
		Edge{"5", "4", 6})

	fmt.Println("Edges to process:", edges)

	expectedDistances := map[string]int{"1": 0, "2": 7, "3": 9, "4": 20, "5": 20, "6": 11}

	assert.Equal(t, expectedDistances, Dijkstra(&edges, "1"), "invalid path")
}
