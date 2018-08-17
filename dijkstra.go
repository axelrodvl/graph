package graph

const MaxInt = int(^uint(0) >> 1)
const UNDEFINED = "UNDEFINED"

func Dijkstra(edges *[]Edge, firstVertex string) map[string]int {
	graph := InitializeGraph(edges)
	distance := initializeDistance(graph, firstVertex)

	for {
		nextVertex := UNDEFINED

		for vertexName, vertex := range graph.Vertices {
			if !vertex.Visited {
				if nextVertex == UNDEFINED {
					nextVertex = vertexName
				} else {
					if distance[vertexName] < distance[nextVertex] {
						nextVertex = vertexName
					}
				}
			}
		}

		if nextVertex != UNDEFINED {
			graph.processDistance(distance, nextVertex)
		} else {
			break
		}
	}

	return distance
}

func initializeDistance(graph *Graph, firstVertex string) map[string]int {
	var distance = make(map[string]int)
	for vertexName := range graph.Vertices {
		distance[vertexName] = MaxInt
	}
	distance[firstVertex] = 0
	
	return distance
}

func (graph *Graph) processDistance(distance map[string]int, vertex string) {
	for _, edge := range *graph.Edges {
		if edge.From == vertex && !graph.Vertices[edge.To].Visited {
			if distanceFromCurrentVertex := distance[vertex] + edge.Length; distanceFromCurrentVertex < distance[edge.To] {
				distance[edge.To] = distanceFromCurrentVertex
			}
		}
	}

	graph.Vertices[vertex].Visited = true
}