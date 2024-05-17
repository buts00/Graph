package graph

// Edge represents an edge in a graph.
type Edge struct {
	Id          int
	Source      *int `json:"Source" binding:"required"`
	Destination *int `json:"Destination" binding:"required"`
	Weight      int  `json:"Weight" binding:"required"`
}

// Graph represents a graph with a collection of edges.
type Graph struct {
	Edges []Edge `json:"edges" binding:"required"`
}

// AllValues returns all the values of an Edge.
func (g *Graph) AllValues(e Edge) (int, int, int, int) {
	return e.Id, *e.Source, *e.Destination, e.Weight
}
