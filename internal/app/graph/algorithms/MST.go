package algorithms

import (
	"github.com/buts00/Graph/internal/app/graph"
	"sort"
)

// MST represents the Minimum Spanning Tree algorithm.
type MST struct {
	parent []int
}

// NewMST initializes a new instance of MST.
func NewMST() *MST {
	return &MST{}
}

// InitializeParent initializes the parent array for Union-Find
func (m *MST) InitializeParent(maxValue int) {
	m.parent = make([]int, maxValue+1)
	for i := range m.parent {
		m.parent[i] = i
	}
}

// UnionNodes performs union operation in Union-Find.
func (m *MST) UnionNodes(firstNode, secondNode int) {
	parentOfFirstNode := m.GetParent(firstNode)
	parentOfSecondNode := m.GetParent(secondNode)
	if parentOfFirstNode == parentOfSecondNode {
		return
	}

	m.parent[parentOfSecondNode] = parentOfFirstNode
}

// GetParent finds the parent node in Union-Find.
func (m *MST) GetParent(node int) int {
	if m.parent[node] == node {
		return node
	}
	m.parent[node] = m.GetParent(m.parent[node])
	return m.parent[node]
}

// FindMaxElement FindMaxNode finds the maximum node ID in the graph.
func (m *MST) FindMaxElement(graph graph.Graph) int {
	maxElement := *graph.Edges[0].Source
	for _, edge := range graph.Edges {
		maxElement = max(maxElement, *edge.Source, *edge.Destination)
	}
	return maxElement
}

// ProcessGraphForMST processes the graph to find the MST edges.
func (m *MST) ProcessGraphForMST(g graph.Graph) []int {
	inMST := make([]int, 0)
	copyGraph := g
	copyGraph.Edges = nil
	copyGraph.Edges = append(copyGraph.Edges, g.Edges...)

	sort.Slice(copyGraph.Edges, func(i, j int) bool {
		return copyGraph.Edges[i].Weight < copyGraph.Edges[j].Weight
	})

	maxElement := m.FindMaxElement(copyGraph)
	m.InitializeParent(maxElement)

	for _, edge := range copyGraph.Edges {
		parentSource := m.GetParent(*edge.Source)
		parentDestination := m.GetParent(*edge.Destination)
		if parentSource != parentDestination {
			inMST = append(inMST, edge.Id)
			m.UnionNodes(parentSource, parentDestination)
		}
	}
	return inMST
}

// FindMST finds the Minimum Spanning Tree of the graph.
func (m *MST) FindMST(g graph.Graph) []int {
	return m.ProcessGraphForMST(g)
}
