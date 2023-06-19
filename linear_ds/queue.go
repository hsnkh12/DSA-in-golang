package linear_ds

import (
	"fmt"
)

// Queue represents a FIFO (First In First Out) queue.
type Queue struct {
	front *Node
	rear  *Node
}

// Enqueue adds a new node with the given value to the rear of the queue.
func (q *Queue) Enqueue(value int) {
	// Create a new node
	var newNode *Node = &Node{Value: value, Next: nil}

	// If the queue is empty, make the new node both the front and rear
	if q.IsEmpty() {
		q.front = newNode
		q.rear = newNode
		return
	}

	// Update the links to add the new node to the rear
	q.rear.Next = newNode
	q.rear = newNode
}

// IsEmpty checks if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

// Dequeue removes the node from the front of the queue.
func (q *Queue) Dequeue() {
	// Check if the queue is empty
	if q.IsEmpty() {
		panic("Queue is empty")
	}

	// If there is only one node, update the front and rear and return
	if q.front == q.rear {
		q.front = nil
		q.rear = nil
		return
	}

	// Update the front to remove the node from the front
	q.front = q.front.Next
}

// ShowFront returns the node at the front of the queue.
func (q *Queue) ShowFront() *Node {
	// Check if the queue is empty
	if q.IsEmpty() {
		panic("Queue is empty")
	}
	return q.front
}

// Display prints the values of the nodes in the queue.
func (q *Queue) Display() {
	// Check if the queue is empty
	if q.IsEmpty() {
		panic("Queue is empty")
	}

	// Traverse the queue and print the values
	current := q.front
	for current != nil {
		fmt.Print(" ", current.Value, " ")
		current = current.Next
	}
	fmt.Println()
}
