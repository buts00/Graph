package graph

import (
	"sort"
)

type Edge struct {
	Id          int
	Source      int
	Destination int
	Weight      int
}

type Graph struct {
	Edges []Edge
}

var (
	parent []int
)

func unionNodes(firstNode, secondNode int) {
	parentOfFirstNode := getParent(firstNode)
	parentOfSecondNode := getParent(secondNode)
	if parentOfFirstNode == parentOfSecondNode {
		return
	}
	parent[parentOfSecondNode] = parentOfFirstNode
}

func getParent(node int) int {
	if parent[node] == node {
		return node
	}
	parent[node] = getParent(parent[node])
	return parent[node]
}

func initParent(maxValue int) {
	parent = parent[:0]
	parent = append(parent, 0)
	for i := 1; i <= maxValue; i++ {
		parent = append(parent, i)
	}
}

func findMaxElement(graph Graph) int {
	mx := graph.Edges[0].Source
	for _, edge := range graph.Edges {
		if edge.Source > mx {
			mx = edge.Source
		}
		if edge.Destination > mx {
			mx = edge.Destination
		}
	}
	return mx
}

func MST(g Graph) []int {

	inMst := make([]int, 0)
	copyGraph := g
	copyGraph.Edges = nil
	copyGraph.Edges = append(copyGraph.Edges, g.Edges...)
	sort.Slice(copyGraph.Edges, func(i, j int) bool {
		return copyGraph.Edges[i].Weight < copyGraph.Edges[j].Weight
	})
	maxElement := findMaxElement(copyGraph)
	initParent(maxElement)

	for _, edge := range copyGraph.Edges {
		parentSource := getParent(edge.Source)
		parentDestination := getParent(edge.Destination)
		if parentSource != parentDestination {
			inMst = append(inMst, edge.Id)
			unionNodes(edge.Source, edge.Destination)
		}

	}

	return inMst

}
