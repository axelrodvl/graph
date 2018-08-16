package graph

import (
	"fmt"
	"errors"
)

func Dfs(edges *[]Edge, firstVertex string) []string {
	var path []string

	graph := InitializeGraph(edges)
	fmt.Println("graph vertices: ", graph.Vertices)

	if e := getPath(graph, &path, firstVertex); e == nil {
		fmt.Println("graph traversed successfully, no cycles found")
	} else {
		fmt.Println(e)
	}

	fmt.Println(path)

	return path
}

func getPath(graph *Graph, path *[]string, vertexName string) error {
	*path = append(*path, vertexName)
	fmt.Println("processing vertex :", vertexName)

	vertex := graph.Vertices[vertexName]

	if vertex.Visited && !vertex.Processed {
		return errors.New("cycle found")
	}

	vertex.MarkVisited()

	for _, edge := range *graph.Edges {
		if edge.From == vertexName {
			getPath(graph, path, edge.To)
		}
	}

	vertex.MarkProcessed()

	return nil
}
