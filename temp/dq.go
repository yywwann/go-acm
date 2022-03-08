package temp

type Deque struct {
	list []int
}

func NewDeque() Deque {
	return Deque{
		list: make([]int, 0),
	}
}

func (q *Deque) PushBack(i int) {
	q.list = append(q.list, i)
}

func (q *Deque) PushFront(i int) {
	q.list = append([]int{i}, q.list...)
}

func (q *Deque) PopFront() {
	q.list = q.list[1:len(q.list)]
}

func (q *Deque) PopBack() {
	q.list = q.list[:len(q.list)-1]
}

func (q *Deque) Front() int {
	return q.list[0]
}

func (q *Deque) Back() int {
	return q.list[len(q.list)-1]
}

func (q *Deque) Empty() bool {
	return len(q.list) == 0
}
