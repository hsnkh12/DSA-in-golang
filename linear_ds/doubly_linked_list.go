package linear_ds

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

func (ls *DLinkedList) InsertFirst(value int) {

	var newNode *DNode = &DNode{Value: value, Next: nil, Prev: nil}

	if ls.head == nil {
		ls.head = newNode
	}

	newNode.Next = ls.head
	ls.head = newNode

}

func (ls *DLinkedList) Remove(index uint) {

	var currentNode *DNode = ls.head

	if currentNode == nil {
		panic("Index out of range")
	}

	if index == 0 {
		ls.head = currentNode.Next
		ls.head.Prev = nil
		return
	}

	var counter uint = 0

	for currentNode != nil {

		if counter == index {

			prevNode := currentNode.Prev
			nextNode := currentNode.Next

			if prevNode != nil {
				prevNode.Next = nextNode
			}

			if nextNode != nil {
				nextNode.Prev = prevNode
			}

			return
		}

		currentNode = currentNode.Next
		counter++
	}

	panic("Index out of range")

}

func (ls *DLinkedList) Reverse() {

	var currentNode *DNode = ls.head.Next

	if currentNode == nil {
		panic("List is already empty")
	}

	if currentNode.Next == nil {
		return
	}

	ls.head.Next = nil

	for {

		nextNode := currentNode.Next
		prevNode := currentNode.Prev
		currentNode.Prev = nextNode
		currentNode.Next = prevNode

		if nextNode == nil {
			ls.head = currentNode
			break
		}

		currentNode = nextNode

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
