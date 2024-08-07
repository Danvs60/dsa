package graphs

import "errors"

type AdjacencyList struct {
	list map[int][]int
}

func (al *AdjacencyList) AddEdge(a, b int) error {
	if !(al.hasVertex(a) && al.hasVertex(b)) {
		return errors.New("either vertex does not exist")
	}
	if al.HasEdge(a, b) {
		return errors.New("already has edge")
	}

	al.list[a] = append(al.list[a], b)
	al.list[b] = append(al.list[b], a)
	return nil
}

func (al *AdjacencyList) RemoveEdge(a, b int) error {
	if !(al.hasVertex(a) && al.hasVertex(b)) {
		return errors.New("either vertex does not exist")
	}
	if !al.HasEdge(a, b) {
		return errors.New("no edge between these vertices")
	}

	al.list[a] = removeEdge(al.list[a], b)
	al.list[b] = removeEdge(al.list[b], a)
	return nil
}

func removeEdge(edges []int, vertex int) []int {
	for i, v := range edges {
		if v == vertex {
			return append(edges[:i], edges[i+1:]...)
		}
	}
	return edges
}

func (al *AdjacencyList) hasVertex(a int) bool {
	_, exists := al.list[a]
	return exists
}

func (al *AdjacencyList) HasEdge(a, b int) bool {
	for _, v := range al.list[a] {
		if v == b {
			return true
		}
	}

	return false
}

func (al *AdjacencyList) AddVertex(v int) error {
	if al.hasVertex(v) {
		return errors.New("vertex already exists")
	}
	al.list[v] = []int{}
	return nil
}

func (al *AdjacencyList) RemoveVertex(v int) error {
	if !al.hasVertex(v) {
		return errors.New("vertex does not exist")
	}
	delete(al.list, v)
	for _, edges := range al.list {
		for i, u := range edges {
			if u == v {
				edges = append(edges[:i], edges[i+1:]...)
				break
			}
		}
	}
	return nil
}
