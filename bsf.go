package graph

func Bfs(edges *[]Edge, firstVertex string) []string {
	var path []string

	var queue []string = make([]string, 1)
	queue[0] = firstVertex

	for len(queue) > 0 {
		current := queue[0]
		path = append(path, current)
		queue = queue[1:]

		for _, edge := range *edges {
			if edge.From == current {
				queue = append(queue, edge.To)
			}
		}
	}

	return path
}
