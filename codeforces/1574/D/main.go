package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Item interface {
	Compare(left Item) bool
}

type Heap interface {
	Push(i Item)
	Pop() Item
	Top() Item
	Len() int
	Update(old Item, new Item)
	Remove(item Item)
	Get(key interface{}) Item
}

type MinInt int

func (i MinInt) Compare(j Item) bool {
	return i < j.(MinInt)
}

type MaxInt int

func (i MaxInt) Compare(j Item) bool {
	return i > j.(MaxInt)
}

type item struct {
	item  Item
	index int
}

type heapImpl struct {
	p               *pq
	items           map[interface{}]*item
	key             func(Item) interface{}
	DisableIndexing bool
}

func (h *heapImpl) Push(i Item) {
	it := &item{i, 0}
	h.addItem(it)
	heap.Push(h.p, it)
}

func (h *heapImpl) Pop() Item {
	res := heap.Pop(h.p).(*item)
	h.removeItem(res.item)
	return res.item
}

func (h *heapImpl) Top() Item {
	if len(h.p.items) > 0 {
		return h.p.items[0].item
	}
	return nil
}

func (h *heapImpl) Len() int {
	return len(h.p.items)
}

func (h *heapImpl) Update(o Item, n Item) {
	old := h.items[h.key(o)]
	new := &item{n, old.index}
	h.removeItem(old.item)
	h.p.items[old.index] = new
	heap.Fix(h.p, old.index)
	h.addItem(new)
}

func (h *heapImpl) Remove(i Item) {
	index := h.items[h.key(i)].index
	h.removeItem(i)
	heap.Remove(h.p, index)
}

func (h *heapImpl) Get(key interface{}) Item {
	return h.items[key].item
}

func (h *heapImpl) removeItem(item Item) {
	if !h.DisableIndexing {
		delete(h.items, h.key(item))
	}
}

func (h *heapImpl) addItem(item *item) {
	if !h.DisableIndexing {
		h.items[h.key(item.item)] = item
	}
}

// A PriorityQueue implements heap.Interface and holds Items.
type pq struct {
	items []*item
}

func (pq pq) Len() int { return len(pq.items) }

func (pq pq) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq.items[i].item.Compare(pq.items[j].item)
}

func (pq pq) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].index = i
	pq.items[j].index = j
}

func (pq *pq) Push(x interface{}) {
	n := len(pq.items)
	item := x.(*item)
	item.index = n
	pq.items = append(pq.items, item)
}

func (pq *pq) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	pq.items = old[0 : n-1]
	return item
}

func NewHeapWithKey(getKey func(Item) interface{}) Heap {
	return &heapImpl{
		p: &pq{
			items: []*item{},
		},
		items:           make(map[interface{}]*item),
		key:             getKey,
		DisableIndexing: true,
	}
}

func NewHeap() Heap {
	return NewHeapWithKey(func(i Item) interface{} { return i })
}

type idx struct {
	value int
	arr   [10]int
}

func (left *idx) Compare(right Item) bool {
	return left.value > right.(*idx).value
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 8*1024*1024)

	var n int
	fmt.Fscanf(r, "%d\n", &n)

	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		var m int
		fmt.Fscanf(r, "%d", &m)
		arr[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscanf(r, "%d", &arr[i][j])
		}
		fmt.Fscanf(r, "\n")
	}

	var b int
	fmt.Fscanf(r, "%d\n", &b)

	seen := make(map[idx]bool)
	for i := 0; i < b; i++ {
		cur := idx{}
		for j := 0; j < n; j++ {
			fmt.Fscanf(r, "%d", &cur.arr[j])
			cur.arr[j]--
			cur.value += arr[j][cur.arr[j]]
		}
		fmt.Fscanf(r, "\n")
		seen[cur] = true
	}

	cur := idx{0, [10]int{}}
	for i := 0; i < n; i++ {
		cur.arr[i] = len(arr[i]) - 1
		cur.value += arr[i][cur.arr[i]]
	}

	h := NewHeap()
	h.Push(&cur)

	//alloc := [1001000]idx{}
	//a := 0

	used := make(map[idx]bool)

	for h.Len() > 0 {
		cur = *h.Pop().(*idx)
		if !seen[cur] {
			break
		}

		for i := 0; i < n; i++ {
			if cur.arr[i] == 0 {
				continue
			}
			next := cur
			next.arr = cur.arr
			next.value = cur.value - arr[i][cur.arr[i]] + arr[i][cur.arr[i]-1]
			next.arr[i]--

			if !used[next] {
				used[next] = true
				h.Push(&next)
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Print(cur.arr[i]+1, " ")
	}
	fmt.Println()
}
