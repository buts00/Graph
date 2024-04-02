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

func (g *Graph) AllValues(e Edge) (int, int, int, int) {
	return e.Id, e.Source, e.Destination, e.Weight
}
