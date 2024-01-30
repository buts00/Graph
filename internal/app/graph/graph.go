package graph

type Node struct {
	Value int
}

type Edge struct {
	Source      int
	Destination int
	Weight      int
}

type Graph struct {
	Nodes *[]Node
	Edges *[]Edge
}
