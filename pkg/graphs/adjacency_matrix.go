package graphs

import "fmt"

type AdjacencyMatrix struct {
	size int
	m    [][]int
}

func NewAdjacencyMatrix(size int) *AdjacencyMatrix {
	am := &AdjacencyMatrix{
		size: size,
		m:    make([][]int, size),
	}
	for i := range am.m {
		am.m[i] = make([]int, size)
	}

	return am
}

func (am *AdjacencyMatrix) isIndexValid(i int) bool {
	return i >= 0 && i < am.size
}

func (am *AdjacencyMatrix) AddEdge(a, b int) error {
	if !(am.isIndexValid(a) && am.isIndexValid(b)) {
		return fmt.Errorf("AddEdge(a,b) index a or b out of bounds, size %v", am.size)
	}

	am.m[a][b] = 1
	am.m[b][a] = 1
	return nil
}

func (am *AdjacencyMatrix) RemoveEdge(a, b int) error {
	if !(am.isIndexValid(a) && am.isIndexValid(b)) {
		return fmt.Errorf("AddEdge(a,b) index a or b out of bounds, size %v", am.size)
	}

	am.m[a][b] = 0
	am.m[b][a] = 0
	return nil
}

func (am *AdjacencyMatrix) HasEdge(a, b int) (bool, error) {
	if !(am.isIndexValid(a) && am.isIndexValid(b)) {
		return false, fmt.Errorf("AddEdge(a,b) index a or b out of bounds, size %v", am.size)
	}

	return am.m[a][b] == 1, nil
}
