package main

import (
	"ds/algorithms"
	"ds/non_linear_ds"
	"fmt"
)

func main() {

	// var q linear_ds.Queue = linear_ds.Queue{}
	// var ls linear_ds.SLinkedList = linear_ds.SLinkedList{}
	// var dls linear_ds.DLinkedList = linear_ds.DLinkedList{}
	var g non_linear_ds.Graph = *non_linear_ds.CreateGraph(6)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 4)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)
	g.AddEdge(4, 5)

	g.PrintGraph()
	algorithms.BFS(g, 0)

	// q.Enqueue(1)
	// ls.Push(1)
	// dls.InsertFirst(1)
	fmt.Println("")
}
