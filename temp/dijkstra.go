package temp

import "container/heap"

const INF = 0x3f3f3f3f

type Node struct {
	to, c int
}

func NewNode(to, c int) Node {
	return Node{
		to: to,
		c:  c,
	}
}

type hp []Node

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].c < h[j].c }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(Node)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func (h *hp) Top() Node {
	a := *h
	return a[0]
}

func (h *hp) Empty() bool {
	return len(*h) == 0
}

func Dijkstra(edges [][]Node, n, s int, dis []int) {
	vis := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = INF
		vis[i] = 0
	}

	dis[s] = 0
	q := hp{}
	heap.Push(&q, NewNode(s, 0))
	for !q.Empty() {
		u := q.Top()
		heap.Pop(&q)
		if vis[u.to] == 1 {
			continue
		}
		vis[u.to] = 1
		for _, e := range edges[u.to] {
			if dis[e.to] > dis[u.to]+e.c {
				dis[e.to] = dis[u.to] + e.c
				heap.Push(&q, NewNode(e.to, dis[e.to]))
			}
		}
	}
}
