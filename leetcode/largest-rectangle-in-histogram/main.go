package largest_rectangle_in_histogram

type Stack struct {
	stack []int
}

func NewStack() *Stack {
	return &Stack{
		stack: make([]int, 0),
	}
}

func (s *Stack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *Stack) Push(i int) {
	s.stack = append(s.stack, i)
}

func (s *Stack) Empty() bool {
	return len(s.stack) == 0
}

func largestRectangleArea(heights []int) int {
	n := len(heights)

	// left[i]  离i左边最近的比他高的pos
	// right[i] 离i右边最近的比他高的pos
	left, right := make([]int, n), make([]int, n)

	stack := NewStack()
	for i := 0; i < n; i++ {
		for !stack.Empty() && heights[stack.Top()] >= heights[i] {
			stack.Pop()
		}
		if stack.Empty() {
			left[i] = -1
		} else {
			left[i] = stack.Top()
		}
		stack.Push(i)
	}

	stack = NewStack()
	for i := n - 1; i >= 0; i-- {
		for !stack.Empty() && heights[stack.Top()] >= heights[i] {
			stack.Pop()
		}
		if stack.Empty() {
			right[i] = n
		} else {
			right[i] = stack.Top()
		}
		stack.Push(i)
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, heights[i]*(right[i]-left[i]-1))
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
