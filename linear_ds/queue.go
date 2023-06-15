package linear_ds

import (
	"fmt"
)

// FIFO (First In First Out)
type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) Enqueue(value int) {

	var newNode *Node = &Node{Value: value, Next: nil}

	if q.IsEmpty() {
		q.front = newNode
		q.rear = newNode
		return
	}

	q.rear.Next = newNode
	q.rear = newNode
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

func (q *Queue) Dequeue() {

	if q.IsEmpty() {
		panic("Queue is empty")
	}

	if q.front == q.rear {
		q.front = nil
		q.rear = nil
		return
	}

	q.front = q.front.Next

}

func (q *Queue) ShowFront() *Node {

	if q.IsEmpty() {
		panic("Queue is empty")
	}

	fmt.Println(q.front)
	return q.front

}

func (q *Queue) Display() {

	if q.IsEmpty() {
		panic("Queue is empty")
	}

	current := q.front
	for current != nil {
		fmt.Print(" ", current.Value, " ")
		current = current.Next
	}
	fmt.Println()
}
