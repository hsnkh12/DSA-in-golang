package linear_ds

type Stack struct {
	top *Node
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Push(value int) {

	var newNode *Node = &Node{Value: value, Next: s.top}
	s.top = newNode
}

func (s *Stack) Pop() {

	if s.IsEmpty() {
		panic("Stack is already empty")
	}
	s.top = s.top.Next
}
