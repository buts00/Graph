package algorithms

import (
	"container/heap"
	"github.com/buts00/Graph/internal/app/graph"
	"math"
)

const inf = math.MaxInt

type Pair struct {
	Weight int
	Node   int
}

type Dijkstra struct {
	SIZE        int
	Distance    []int
	Matrix      [][]Pair
	Predecessor []int
}

func NewDijkstra() *Dijkstra {
	return &Dijkstra{}
}

func (d *Dijkstra) FindMaxElement(graph graph.Graph) int {
	maxElement := 0
	for _, edge := range graph.Edges {
		maxElement = max(maxElement, edge.Source, edge.Destination)
	}
	return maxElement
}

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

func (d *Dijkstra) addEdges(g graph.Graph) {
	for i := 0; i < d.SIZE; i++ {
		_, src, dest, weight := g.AllValues(g.Edges[i])
		d.Matrix[src] = append(d.Matrix[src], Pair{weight, dest})
		d.Matrix[dest] = append(d.Matrix[dest], Pair{weight, src})
	}
}

func (d *Dijkstra) dijkstra(startPoint, destination int) {
	d.Distance[startPoint] = 0
	h := &intHeap{{0, startPoint}}
	heap.Init(h)
	for h.Len() > 0 {
		var curPair Pair = (*h)[0]
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

func (d *Dijkstra) FindDijkstra(startPoint, destination int, g graph.Graph) ([]graph.Edge, int) {
	maxElement := d.FindMaxElement(g)
	d.initDijkstra(maxElement, g)
	d.addEdges(g)
	d.dijkstra(startPoint, destination)

	// Reconstruct the shortest path from startPoint to destination as an array of edges
	path := make([]graph.Edge, 0)
	if d.Distance[destination] != inf {
		for at := destination; at != startPoint; at = d.Predecessor[at] {
			prev := d.Predecessor[at]
			// Find the weight of the edge
			weight := 0
			for _, pair := range d.Matrix[prev] {
				if pair.Node == at {
					weight = pair.Weight
					break
				}
			}
			path = append(path, graph.Edge{Source: prev, Destination: at, Weight: weight})
		}
	}

	return path, d.Distance[destination]
}
