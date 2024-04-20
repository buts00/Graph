package graph

type Edge struct {
	Id          int
	Source      int `json:"Source" binding:"required"`
	Destination int `json:"Destination" binding:"required"`
	Weight      int `json:"Weight" binding:"required"`
}

type Graph struct {
	Edges []Edge
}

func (g *Graph) AllValues(e Edge) (int, int, int, int) {
	return e.Id, e.Source, e.Destination, e.Weight
}
