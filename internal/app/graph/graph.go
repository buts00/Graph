package graph

type Edge struct {
	Id          int
	Source      int `json:"Source"`
	Destination int `json:"Destination"`
	Weight      int `json:"Weight"`
}

type Graph struct {
	Edges []Edge
}
