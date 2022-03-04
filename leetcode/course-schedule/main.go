package course_schedule

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

func canFinish(numCourses int, prerequisites [][]int) bool {
	n := numCourses
	edges := make([][]int, n)
	indeg := make([]int, n)
	vis := make([]int, n)

	for _, prerequisite := range prerequisites {
		x := prerequisite[0]
		y := prerequisite[1]
		if edges[y] == nil {
			edges[y] = make([]int, 0)
		}

		edges[y] = append(edges[y], x)
		indeg[x] += 1
	}

	q := NewQueue()
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			q.Push(i)
		}
	}

	for !q.Empty() {
		cur := q.Top()
		q.Pop()

		if vis[cur] == 1 {
			continue
		}
		vis[cur] = 1
		for _, x := range edges[cur] {
			indeg[x] -= 1
			if indeg[x] == 0 {
				q.Push(x)
			}
		}
	}

	for _, v := range vis {
		if v == 0 {
			return false
		}
	}

	return true
}
