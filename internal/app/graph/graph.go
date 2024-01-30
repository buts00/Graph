package graph

type Edge struct {
	Source      int
	Destination int
	Weight      int
}

type Graph struct {
	Edges *[]Edge
}
