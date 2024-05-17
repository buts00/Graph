package algorithms

import (
	"container/heap"
	"github.com/buts00/Graph/internal/app/graph"
	"math"
)

// inf represents infinity
const inf = math.MaxInt

// Pair represents a weighted edge in the graph
type Pair struct {
	Weight int // Weight of the edge
	Node   int // Destination node of the edge
}

// Dijkstra struct represents the Dijkstra algorithm
type Dijkstra struct {
	SIZE        int      // Number of edges in the graph
	Distance    []int    // Array to store the minimum distance from the source to each node
	Matrix      [][]Pair // Adjacency list representation of the graph
	Predecessor []int    // Array to store the predecessor of each node in the shortest path
}

// NewDijkstra initializes a new instance of Dijkstra
func NewDijkstra() *Dijkstra {
	return &Dijkstra{}
}

// FindMaxElement finds the maximum element in the graph
func (d *Dijkstra) FindMaxElement(graph graph.Graph) int {
	maxElement := 0
	for _, edge := range graph.Edges {
		maxElement = max(maxElement, *edge.Source, *edge.Destination)
	}
	return maxElement
}

// initDijkstra initializes the Dijkstra struct with required size
func (d *Dijkstra) initDijkstra(maxElement int, g graph.Graph) {
	d.SIZE = len(g.Edges)
	d.Matrix = make([][]Pair, maxElement+1)
	d.Distance = make([]int, maxElement+1)
	d.Predecessor = make([]int, maxElement+1)
	for i := range d.Distance {
		d.Distance[i] = math.MaxInt
		d.Predecessor[i] = -1
	}
}

// addEdges adds edges to the adjacency list
func (d *Dijkstra) addEdges(g graph.Graph) {
	for i := 0; i < d.SIZE; i++ {
		_, src, dest, weight := g.AllValues(g.Edges[i])
		if src == dest {
			continue
		}
		d.Matrix[src] = append(d.Matrix[src], Pair{weight, dest})
		d.Matrix[dest] = append(d.Matrix[dest], Pair{weight, src})
	}

}

// dijkstra runs the Dijkstra algorithm
func (d *Dijkstra) dijkstra(startPoint, destination int) {
	d.Distance[startPoint] = 0
	h := &intHeap{{0, startPoint}}
	heap.Init(h)
	for h.Len() > 0 {
		curPair := (*h)[0]
		curNode := curPair.Node
		heap.Pop(h)

		if curNode == destination {
			break
		}

		for _, pair := range d.Matrix[curNode] {
			weight := pair.Weight
			to := pair.Node
			if d.Distance[to] > d.Distance[curNode]+weight {
				d.Distance[to] = d.Distance[curNode] + weight
				d.Predecessor[to] = curNode
				heap.Push(h, Pair{Weight: d.Distance[to], Node: to})
			}
		}
	}
}

// FindDijkstra finds the shortest path using Dijkstra algorithm
func (d *Dijkstra) FindDijkstra(startPoint, destination int, g graph.Graph) ([]graph.Edge, int) {
	maxElement := d.FindMaxElement(g)
	d.initDijkstra(maxElement, g)
	d.addEdges(g)
	d.dijkstra(startPoint, destination)

	path := make([]graph.Edge, 0)
	if d.Distance[destination] == inf {
		return path, -1
	}

	for at := destination; at != startPoint; at = d.Predecessor[at] {
		prev := d.Predecessor[at]
		weight := 0
		for _, pair := range d.Matrix[prev] {
			if pair.Node == at {
				weight = pair.Weight
				break
			}
		}
		path = append(path, graph.Edge{Source: &prev, Destination: &at, Weight: weight})
	}

	return path, d.Distance[destination]
}
