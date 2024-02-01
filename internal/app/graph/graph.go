package graph

type Edge struct {
	Id          int
	Source      int
	Destination int
	Weight      int
}

type Graph struct {
	Edges []Edge
}
