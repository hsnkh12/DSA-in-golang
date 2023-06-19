package non_linear_ds

import (
	"fmt"
)

type GNode struct {
	Vertex int
	Next   *GNode
}

type ADJlist struct {
	head *GNode
}

type Graph struct {
	NumVertices int
	lists       []*ADJlist
}

func CreateGraph(vertices int) *Graph {

	var newGraph *Graph = &Graph{NumVertices: vertices, lists: make([]*ADJlist, vertices)}

	for i := 0; i < vertices; i++ {

		newGraph.lists[i] = &ADJlist{head: nil}
	}

	return newGraph
}

func (g *Graph) AddEdge(src int, dest int) {

	g.DirectedAddEdge(src, dest)
	g.DirectedAddEdge(dest, src)

}

func (g *Graph) DirectedAddEdge(src int, dest int) {

	var newNode *GNode = &GNode{Vertex: dest}
	newNode.Next = g.lists[src].head
	g.lists[src].head = newNode

}

func (g *Graph) RemoveEdge(src int, dest int) {

	g.DrectedRemoveEdge(src, dest)
	g.DrectedRemoveEdge(dest, src)

}

func (g *Graph) DrectedRemoveEdge(src int, dest int) {

	var currentNode *GNode = g.lists[src].head
	prevNode := currentNode

	if currentNode == nil {
		panic("Adjc list is already empty")
	}

	if currentNode.Vertex == dest {
		g.lists[src].head = currentNode.Next
	}

	for currentNode != nil {

		if currentNode.Vertex == dest {
			prevNode.Next = currentNode.Next
			break
		}
		prevNode = currentNode
		currentNode = currentNode.Next
	}
}

func (g *Graph) PrintGraph() {
	for i := 0; i < g.NumVertices; i++ {
		fmt.Printf("Adjacency list for vertex %d: ", i)
		current := g.lists[i].head
		for current != nil {
			fmt.Printf("%d ", current.Vertex)
			current = current.Next
		}
		fmt.Println()
	}
}
