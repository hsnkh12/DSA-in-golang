package main

import (
	"ds/linear_ds"
	"fmt"
)

func main() {

	var q linear_ds.Queue = linear_ds.Queue{}
	var ls linear_ds.SLinkedList = linear_ds.SLinkedList{}
	var dls linear_ds.DLinkedList = linear_ds.DLinkedList{}

	q.Enqueue(1)
	ls.Push(1)
	dls.InsertFirst(1)

	fmt.Println("")
}
