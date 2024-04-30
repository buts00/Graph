package Algorithms

import (
	"github.com/buts00/Graph/internal/app/graph"
	"sort"
)

type MST struct {
	parent []int
}

func NewMST() *MST {
	return &MST{}
}

func (m *MST) InitializeParent(maxValue int) {
	m.parent = make([]int, maxValue+1)
	for i := range m.parent {
		m.parent[i] = i
	}
}

func (m *MST) UnionNodes(firstNode, secondNode int) {
	parentOfFirstNode := m.GetParent(firstNode)
	parentOfSecondNode := m.GetParent(secondNode)
	if parentOfFirstNode == parentOfSecondNode {
		return
	}

	m.parent[parentOfSecondNode] = parentOfFirstNode
}

func (m *MST) GetParent(node int) int {
	if m.parent[node] == node {
		return node
	}
	m.parent[node] = m.GetParent(m.parent[node])
	return m.parent[node]
}

func (m *MST) FindMaxElement(graph graph.Graph) int {
	maxElement := graph.Edges[0].Source
	for _, edge := range graph.Edges {
		if edge.Source > maxElement {
			maxElement = edge.Source
		}
		if edge.Destination > maxElement {
			maxElement = edge.Destination
		}
	}
	return maxElement
}

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
		parentSource := m.GetParent(edge.Source)
		parentDestination := m.GetParent(edge.Destination)
		if parentSource != parentDestination {
			inMST = append(inMST, edge.Id)
			m.UnionNodes(parentSource, parentDestination)
		}
	}

	return inMST
}

func (m *MST) FindMST(g graph.Graph) []int {
	return m.ProcessGraphForMST(g)
}
