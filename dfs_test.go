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
		Edge{"e", "a"},
		Edge{"a", "b"},
		Edge{"a", "c"},
		Edge{"b", "d"})

	fmt.Println("graph edges: ", edges)

	assert.Equal(t, Dfs(&edges, "e"), []string{"e", "a", "b", "d", "c"}, "invalid path")
}
