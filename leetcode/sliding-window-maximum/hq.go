package sliding_window_maximum

import "container/heap"

type Hq struct {
	num  []int
	list []int
}

func NewHq() Hq {
	return Hq{list: []int{}}
}

func (h *Hq) Len() int {
	return len(h.list)
}

func (h *Hq) Less(i, j int) bool {
	return h.num[h.list[i]] > h.num[h.list[j]]
	// return h.list[i] < h.list[j]
}

func (h *Hq) Swap(i, j int) {
	h.list[i], h.list[j] = h.list[j], h.list[i]
}

func (h *Hq) Push(v interface{}) {
	h.list = append(h.list, v.(int))
}

func (h *Hq) Pop() interface{} {
	n := len(h.list)
	x := h.list[n-1]
	h.list[n-1] = 0
	h.list = h.list[:n-1]
	return x
}

func (h *Hq) Top() int {
	return h.list[0]
}

func (h *Hq) Empty() bool {
	return len(h.list) == 0
}

func maxSlidingWindow1(nums []int, k int) []int {
	hq := NewHq()
	hq.num = nums
	ans := []int{}
	n := len(nums)

	for i := 0; i < k; i++ {
		heap.Push(&hq, i)
	}
	ans = append(ans, hq.num[hq.Top()])

	for i := k; i < n; i++ {
		for !hq.Empty() && hq.Top() <= i-k {
			heap.Pop(&hq)
		}
		heap.Push(&hq, i)
		ans = append(ans, hq.num[hq.Top()])
	}

	return ans
}
