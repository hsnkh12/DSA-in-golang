package linear_ds

import (
	"fmt"
)

// DNode represents a node in a doubly linked list.
type DNode struct {
	Value int
	Prev  *DNode
	Next  *DNode
}

// DLinkedList represents a doubly linked list.
type DLinkedList struct {
	head *DNode
}

// Push adds a new node with the given value to the end of the list.
func (ls *DLinkedList) Push(value int) {
	// Create a new node
	var newNode *DNode = &DNode{Value: value, Prev: nil, Next: nil}

	// If the list is empty, make the new node the head and return
	if ls.head == nil {
		ls.head = newNode
		return
	}

	// Traverse the list to find the last node
	var currentNode *DNode = ls.head
	for {
		// If the current node is the last node, update the links and exit the loop
		if currentNode.Next == nil {
			newNode.Prev = currentNode
			currentNode.Next = newNode
			break
		}
		currentNode = currentNode.Next
	}
}

// InsertFirst adds a new node with the given value at the beginning of the list.
func (ls *DLinkedList) InsertFirst(value int) {
	// Create a new node
	var newNode *DNode = &DNode{Value: value, Next: nil, Prev: nil}

	// If the list is empty, make the new node the head
	if ls.head == nil {
		ls.head = newNode
	}

	// Update the links to insert the new node at the beginning
	newNode.Next = ls.head
	ls.head = newNode
}

// Remove deletes the node at the specified index from the list.
func (ls *DLinkedList) Remove(index uint) {
	// Start with the head node
	var currentNode *DNode = ls.head

	// Check if the list is empty
	if currentNode == nil {
		panic("Index out of range")
	}

	// If the index is 0, update the head and return
	if index == 0 {
		ls.head = currentNode.Next
		ls.head.Prev = nil
		return
	}

	// Initialize a counter to track the current index
	var counter uint = 0

	// Traverse the list to find the node at the specified index
	for currentNode != nil {
		// If the current index matches the specified index, update the links and exit the loop
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

// Reverse reverses the order of the nodes in the list.
func (ls *DLinkedList) Reverse() {
	// Start with the second node
	var currentNode *DNode = ls.head.Next

	// Check if the list is empty or has only one node
	if currentNode == nil {
		panic("List is already empty")
	}

	if currentNode.Next == nil {
		return
	}

	// Set the head's next pointer to nil
	ls.head.Next = nil

	// Reverse the links between nodes
	for {
		nextNode := currentNode.Next
		prevNode := currentNode.Prev
		currentNode.Prev = nextNode
		currentNode.Next = prevNode

		// If the next node is nil, update the head and exit the loop
		if nextNode == nil {
			ls.head = currentNode
			break
		}

		currentNode = nextNode
	}
}

// PrintValues prints the values of the nodes in the list.
func (ls *DLinkedList) PrintValues() {
	var currentNode *DNode = ls.head

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
