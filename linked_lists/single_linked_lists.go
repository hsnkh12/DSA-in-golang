package linked_lists

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type SLinkedList struct {
	head *Node
}

func (ls *SLinkedList) Push(value int) {

	var newNode *Node = &Node{Value: value, Next: nil}
	var currentNode *Node = ls.head

	if ls.head == nil {
		ls.head = newNode
		return
	}

	for {

		if currentNode.Next == nil {
			currentNode.Next = newNode
			break
		}

		currentNode = currentNode.Next
	}
}

func (ls *SLinkedList) InsertFirst(value int) {

	var newNode *Node = &Node{Value: value, Next: nil}
	var currentNode *Node = ls.head

	if currentNode == nil {
		ls.head = newNode
	}

	newNode.Next = ls.head
	ls.head = newNode

}

func (ls *SLinkedList) Pop() {

	var currentNode *Node = ls.head

	if currentNode == nil {
		panic("List is already empty")
	}

	if currentNode.Next == nil {
		ls.head = nil
		return
	}

	for currentNode.Next.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = nil

}

func (ls *SLinkedList) Remove(index uint) {

	var currentNode *Node = ls.head

	if currentNode == nil {
		panic("Index out of range")
	}

	if index == 0 {
		ls.head = currentNode.Next
		return
	}

	var counter uint = 0
	var prevNode *Node = currentNode

	for currentNode != nil {

		if index == counter {
			prevNode.Next = currentNode.Next
			return
		}

		prevNode = currentNode
		currentNode = currentNode.Next
		counter++
	}

	panic("Index out of range")

}

func (ls *SLinkedList) PrintValues() {

	var currentNode *Node = ls.head

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
