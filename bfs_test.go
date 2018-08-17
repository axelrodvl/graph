package graph

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestBfs(t *testing.T) {
	var edges []Edge
	edges = append(edges,
		Edge{"1", "2", 0},
		Edge{"2", "5", 0},
		Edge{"5", "9", 0},
		Edge{"1", "3", 0},
		Edge{"3", "6", 0},
		Edge{"3", "7", 0},
		Edge{"6", "10", 0},
		Edge{"1", "4", 0},
		Edge{"4", "8", 0})

	fmt.Println("graph edges:", edges)

	assert.Equal(t, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, Bfs(&edges, "1"), "invalid path")
	assert.Equal(t, []string{"3", "6", "7", "10"}, Bfs(&edges, "3"), "invalid path")
	assert.Equal(t, []string{"2", "5", "9"}, Bfs(&edges, "2"), "invalid path")
}