package Dijkstra

type intHeap []Pair

func (h intHeap) Len() int               { return len(h) }
func (h intHeap) Less(i int, j int) bool { return h[i].Weight < h[j].Weight }
func (h intHeap) Swap(i int, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x any)            { *h = append(*h, x.(Pair)) }

func (h *intHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
