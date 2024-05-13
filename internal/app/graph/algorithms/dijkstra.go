package algorithms

import (
	"container/heap"
	"github.com/buts00/Graph/internal/app/graph"
	"math"
	"sort"
)

const inf = math.MaxInt

type Pair struct {
	Weight int
	Node   int
}

type Dijkstra struct {
	SIZE     int
	Distance []int
	Matrix   [][]Pair
	Parent   []int
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
	for i := range d.Distance {
		d.Distance[i] = math.MaxInt
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
			break // Якщо досягли пункту призначення, виходимо з циклу
		}
		for _, pair := range d.Matrix[curNode] {
			weight := pair.Weight
			to := pair.Node
			if d.Distance[to] > d.Distance[curNode]+weight {
				d.Distance[to] = d.Distance[curNode] + weight
				heap.Push(h, Pair{Weight: d.Distance[to], Node: to})
			}
		}
	}
	path := []Pair{}
	current := destination
	for current != startPoint {
		path = append([]Pair{{Weight: d.Distance[current], Node: current}}, path...)
		current = d.Parent[current]
	}
	path = append([]Pair{{Weight: 0, Node: startPoint}}, path...)

}

func (d *Dijkstra) FindDijkstra(startPoint, destination int, g graph.Graph) []Pair {
	maxElement := d.FindMaxElement(g)
	d.initDijkstra(maxElement, g)
	d.addEdges(g)
	d.dijkstra(startPoint, destination)
	distance := make([]Pair, 0)
	for i := range d.Distance {
		if d.Distance[i] != inf {
			distance = append(distance, Pair{Node: i, Weight: d.Distance[i]})
		} else if len(d.Matrix[i]) > 0 {
			distance = append(distance, Pair{Node: i, Weight: 1e9})
		}
	}
	sort.Slice(distance, func(i, j int) bool {
		return distance[i].Weight < distance[j].Weight
	})
	return distance

}
