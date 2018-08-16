package graph

type Edge struct {
	From string
	To string
}

type Vertex struct {
	Visited bool
	Processed bool
}

func (v *Vertex) MarkVisited() {
	v.Visited = true
	v.Processed = false
}

func (v *Vertex) MarkProcessed() {
	v.Visited = false
	v.Processed = true
}

type Graph struct {
	Edges *[]Edge
	Vertices map[string]*Vertex
}

func InitializeGraph(edges *[]Edge) *Graph {
	var graph Graph
	graph.Edges = edges

	graph.Vertices = make(map[string]*Vertex)
	for _, edge := range *graph.Edges {
		graph.Vertices[edge.From] = &Vertex{}
		graph.Vertices[edge.To] = &Vertex{}
	}

	return &graph
}

