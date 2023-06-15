package main

import (
	"ds/linked_lists"
)

func main() {

	var ls linked_lists.SLinkedList = linked_lists.SLinkedList{}

	ls.Push(1)
	ls.Push(2)
	ls.Push(4)
	ls.Reverse()
	ls.PrintValues()
}
