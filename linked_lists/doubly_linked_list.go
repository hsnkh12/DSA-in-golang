package linked_lists

import (
	"fmt"
)

type DNode struct {
	Value int
	Prev  *DNode
	Next  *DNode
}

type DLinkedList struct {
	head *DNode
}

func (ls *DLinkedList) Push(value int) {

	var newNode *DNode = &DNode{Value: value, Prev: nil, Next: nil}

	if ls.head == nil {
		ls.head = newNode
		return
	}

	var currentNode *DNode = ls.head

	for {
		if currentNode.Next == nil {
			newNode.Prev = currentNode
			currentNode.Next = newNode
			break
		}
		currentNode = currentNode.Next
	}
}

func (ls *DLinkedList) PrintValues() {

	var currentNode *DNode = ls.head

	if currentNode == nil {
		fmt.Print("")
		return
	}

	for {

		fmt.Print(" ", currentNode.Value, " ")
		if currentNode.Next == nil {
			break
		}

		currentNode = currentNode.Next
	}
	fmt.Println()

}
