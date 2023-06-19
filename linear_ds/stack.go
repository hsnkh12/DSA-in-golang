package linear_ds

// Stack represents a stack data structure.
type Stack struct {
	top *Node
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

// Push adds a new node with the given value to the top of the stack.
func (s *Stack) Push(value int) {
	// Create a new node with the given value and set its next pointer to the current top
	var newNode *Node = &Node{Value: value, Next: s.top}
	// Set the new node as the new top of the stack
	s.top = newNode
}

// Pop removes the node from the top of the stack.
func (s *Stack) Pop() {
	// Check if the stack is empty
	if s.IsEmpty() {
		panic("Stack is already empty")
	}
	// Update the top pointer to remove the node from the top
	s.top = s.top.Next
}
