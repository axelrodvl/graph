package graph

import (
    "testing"
    "fmt"
	"github.com/stretchr/testify/assert"
)

func TestDfs(t *testing.T) {
    // G=({a, b, c, d, e}, {(e, a), (a, b), (a, c), (b, d)})

	var edges []Edge
	edges = append(edges,
		Edge{"e", "a", 0},
		Edge{"a", "b", 0},
		Edge{"a", "c", 0},
		Edge{"b", "d", 0})

	fmt.Println("graph edges: ", edges)

	assert.Equal(t, []string{"e", "a", "b", "d", "c"}, Dfs(&edges, "e"), "invalid path")
}
