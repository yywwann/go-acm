package sliding_window_maximum

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

func maxSlidingWindow2(nums []int, k int) []int {
	dq := NewDeque() // 存下标, nums从小到大
	ans := []int{}
	n := len(nums)

	for i := 0; i < n; i++ {
		//
		for !dq.Empty() && nums[i] >= nums[dq.Front()] {
			dq.PopFront()
		}
		dq.PushFront(i)
		//
		if i >= k-1 {
			for !dq.Empty() && dq.Back() < i-k+1 {
				dq.PopBack()
			}
			ans = append(ans, nums[dq.Back()])
		}
	}

	return ans
}

func maxSlidingWindow3(nums []int, k int) []int {
	dq := NewDeque() // 存下标, nums从大到小
	ans := []int{}
	n := len(nums)

	for i := 0; i < n; i++ {
		//
		for !dq.Empty() && nums[i] >= nums[dq.Back()] {
			dq.PopBack()
		}
		dq.PushBack(i)
		//
		if i >= k-1 {
			for !dq.Empty() && dq.Front() < i-k+1 {
				dq.PopFront()
			}
			ans = append(ans, nums[dq.Front()])
		}
	}

	return ans
}
