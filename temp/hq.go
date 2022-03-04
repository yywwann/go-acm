package temp

type Hq struct {
	list []int
}

func NewHq() Hq {
	return Hq{list: []int{}}
}

func (h *Hq) Len() int {
	return len(h.list)
}

func (h *Hq) Less(i, j int) bool {
	return h.list[i] < h.list[j]
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
