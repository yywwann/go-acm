package meeting_rooms_ii

import (
	"container/heap"
	"sort"
)

type pair struct {
	start, end int
}
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func (h *hp) Top() pair {
	a := *h
	return a[0]
}

// 优先队列
// https://leetcode-cn.com/problems/meeting-rooms-ii/
func minMeetingRooms(intervals [][]int) int {
	meets := make([]pair, 0)
	for _, interval := range intervals {
		meets = append(meets, pair{
			interval[0],
			interval[1],
		})
	}

	sort.Slice(meets, func(i, j int) bool {
		return meets[i].start < meets[j].start
	})

	ans := 0
	h := hp{}
	for _, meet := range meets {
		for len(h) > 0 && h.Top().end <= meet.start {
			heap.Pop(&h)
		}
		heap.Push(&h, meet)
		if ans < len(h) {
			ans = len(h)
		}
	}

	return ans
}
