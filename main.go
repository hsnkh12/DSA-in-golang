package main

import (
	"ds/linear_ds"
	"ds/non_linear_ds"
	"fmt"
)

func main() {

	var q linear_ds.Queue = linear_ds.Queue{}
	var ls linear_ds.SLinkedList = linear_ds.SLinkedList{}
	var dls linear_ds.DLinkedList = linear_ds.DLinkedList{}
	var g non_linear_ds.Graph = *non_linear_ds.CreateGraph(5)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(4, 2)
	g.AddEdge(4, 3)
	g.RemoveEdge(0, 1)
	q.Enqueue(1)
	ls.Push(1)
	dls.InsertFirst(1)
	g.PrintGraph()
	fmt.Println("")
}
