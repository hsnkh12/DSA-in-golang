package non_linear_ds

import (
	"fmt"
)

// GNode represents a node in the adjacency list.
type GNode struct {
	Vertex int
	Next   *GNode
}

// ADJlist represents the adjacency list for a vertex.
type ADJlist struct {
	Head *GNode
}

// Graph represents a graph data structure.
type Graph struct {
	NumVertices int
	Lists       []*ADJlist
}

// CreateGraph creates a new graph with the specified number of vertices.
func CreateGraph(vertices int) *Graph {
	// Create a new graph with the specified number of vertices
	var newGraph *Graph = &Graph{NumVertices: vertices, Lists: make([]*ADJlist, vertices)}

	// Initialize the adjacency lists for each vertex
	for i := 0; i < vertices; i++ {
		newGraph.Lists[i] = &ADJlist{Head: nil}
	}

	return newGraph
}

// AddEdge adds an edge between two vertices in the graph.
func (g *Graph) AddEdge(src int, dest int) {
	// Add an edge from source to destination
	g.DirectedAddEdge(src, dest)
	// Add an edge from destination to source (undirected graph)
	g.DirectedAddEdge(dest, src)
}

// DirectedAddEdge adds a directed edge from source to destination.
func (g *Graph) DirectedAddEdge(src int, dest int) {
	// Create a new node for the destination vertex
	var newNode *GNode = &GNode{Vertex: dest}
	// Make the new node point to the current head of the adjacency list for the source vertex
	newNode.Next = g.Lists[src].Head
	// Set the new node as the new head of the adjacency list for the source vertex
	g.Lists[src].Head = newNode
}

// GetAdjacentNodes returns the list of adjacent nodes for a given vertex.
func (g *Graph) GetAdjacentNodes(src int) []*GNode {
	var currentNode *GNode = g.Lists[src].Head
	var adjacentNodes []*GNode

	// Traverse the adjacency list for the source vertex and collect the adjacent nodes
	for currentNode != nil {
		adjacentNodes = append(adjacentNodes, currentNode)
		currentNode = currentNode.Next
	}

	return adjacentNodes
}

// RemoveEdge removes an edge between two vertices in the graph.
func (g *Graph) RemoveEdge(src int, dest int) {
	// Remove the edge from source to destination
	g.DirectedRemoveEdge(src, dest)
	// Remove the edge from destination to source (undirected graph)
	g.DirectedRemoveEdge(dest, src)
}

// DirectedRemoveEdge removes a directed edge from source to destination.
func (g *Graph) DirectedRemoveEdge(src int, dest int) {
	var currentNode *GNode = g.Lists[src].Head
	prevNode := currentNode

	if currentNode == nil {
		panic("Adjacency list is already empty")
	}

	// Check if the head node is the one to be removed
	if currentNode.Vertex == dest {
		g.Lists[src].Head = currentNode.Next
	}

	// Traverse the adjacency list and find the node to be removed
	for currentNode != nil {
		if currentNode.Vertex == dest {
			prevNode.Next = currentNode.Next
			break
		}
		prevNode = currentNode
		currentNode = currentNode.Next
	}
}

// PrintGraph prints the adjacency list representation of the graph.
func (g *Graph) PrintGraph() {
	for i := 0; i < g.NumVertices; i++ {
		fmt.Printf("Adjacency list for vertex %d: ", i)
		current := g.Lists[i].Head
		for current != nil {
			fmt.Printf("%d ", current.Vertex)
			current = current.Next
		}
		fmt.Println()
	}
}
