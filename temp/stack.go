package temp

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
