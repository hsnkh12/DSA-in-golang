package linear_ds

import (
	"fmt"
)

// Node represents a node in a singly linked list.
type Node struct {
	Value int
	Next  *Node
}

// SLinkedList represents a singly linked list.
type SLinkedList struct {
	head *Node
}

// Push adds a new node with the given value to the end of the list.
func (ls *SLinkedList) Push(value int) {
	// Create a new node
	var newNode *Node = &Node{Value: value, Next: nil}

	// If the list is empty, make the new node the head and return
	if ls.head == nil {
		ls.head = newNode
		return
	}

	// Traverse the list to find the last node
	var currentNode *Node = ls.head
	for {
		// If the current node is the last node, update the links and exit the loop
		if currentNode.Next == nil {
			currentNode.Next = newNode
			break
		}

		currentNode = currentNode.Next
	}
}

// InsertFirst adds a new node with the given value at the beginning of the list.
func (ls *SLinkedList) InsertFirst(value int) {
	// Create a new node
	var newNode *Node = &Node{Value: value, Next: nil}

	// If the list is empty, make the new node the head
	if ls.head == nil {
		ls.head = newNode
	}

	// Update the links to insert the new node at the beginning
	newNode.Next = ls.head
	ls.head = newNode
}

// Pop removes the last node from the list.
func (ls *SLinkedList) Pop() {
	var currentNode *Node = ls.head

	// Check if the list is empty
	if currentNode == nil {
		panic("List is already empty")
	}

	// If the list has only one node, update the head and return
	if currentNode.Next == nil {
		ls.head = nil
		return
	}

	// Traverse the list to find the second-to-last node
	for currentNode.Next.Next != nil {
		currentNode = currentNode.Next
	}

	// Remove the last node by updating the links
	currentNode.Next = nil
}

// Remove deletes the node at the specified index from the list.
func (ls *SLinkedList) Remove(index uint) {
	var currentNode *Node = ls.head

	// Check if the list is empty
	if currentNode == nil {
		panic("Index out of range")
	}

	// If the index is 0, update the head and return
	if index == 0 {
		ls.head = currentNode.Next
		return
	}

	var counter uint = 0
	var prevNode *Node = currentNode

	// Traverse the list to find the node at the specified index
	for currentNode != nil {
		// If the current index matches the specified index, update the links and exit the loop
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

// PrintValues prints the values of the nodes in the list.
func (ls *SLinkedList) PrintValues() {
	var currentNode *Node = ls.head

	// Check if the list is empty
	if currentNode == nil {
		fmt.Print("")
		return
	}

	// Traverse the list and print the values
	for {
		fmt.Print(" ", currentNode.Value, " ")
		if currentNode.Next == nil {
			break
		}

		currentNode = currentNode.Next
	}
	fmt.Println()
}

// Reverse reverses the order of the nodes in the list.
func (ls *SLinkedList) Reverse() {
	// Check if the list is empty
	if ls.head == nil {
		panic("List is already empty")
	}

	// If the list has only one node, no need to reverse
	if ls.head.Next == nil {
		return
	}

	var currentNode *Node = ls.head.Next
	var prevNode *Node = ls.head
	ls.head.Next = nil

	// Reverse the links between nodes
	for {
		nextNode := currentNode.Next
		currentNode.Next = prevNode
		prevNode = currentNode

		if nextNode == nil {
			ls.head = currentNode
			break
		}

		currentNode = nextNode
	}
}
