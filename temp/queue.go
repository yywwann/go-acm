package temp

type Queue struct {
	list []int
}

func NewQueue() Queue {
	return Queue{
		list: make([]int, 0),
	}
}

func (q *Queue) Push(i int) {
	q.list = append(q.list, i)
}

func (q *Queue) Pop() {
	q.list = q.list[1:len(q.list)]
}

func (q *Queue) Top() int {
	return q.list[0]
}

func (q *Queue) Empty() bool {
	return len(q.list) == 0
}
