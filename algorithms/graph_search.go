package algorithms

// DS needed: Graph, Queue
import (
	"ds/linear_ds"
	"ds/non_linear_ds"
	"fmt"
)

// Breadth first search
func BFS(graph non_linear_ds.Graph, start int) {

	// Create an array of boolean visited for each vertex, to keep track of visited vertices
	var visited = make([]bool, graph.NumVertices)

	// Initialize a new empty queue
	var queue = linear_ds.Queue{}

	// Set vertex start as true and enqueue it to the queue
	visited[start] = true
	queue.Enqueue(start)

	// If queue is not empty, keep searching
	for !queue.IsEmpty() {

		// Get the fist vertex of the queue, print it, and dequeue it
		var vertex int = queue.ShowFront().Value
		fmt.Print(vertex, " ")
		queue.Dequeue()

		// Get all adjacent nodes for the current vertex
		var adjacentNodes []*non_linear_ds.GNode = graph.GetAdjacentNodes(vertex)

		// Iterate through adjacent nodes for the vertex
		for i := 0; i < len(adjacentNodes); i++ {

			v := adjacentNodes[i].Vertex

			// If the vertex of the current node is not visited
			if !visited[v] {

				// Mark it as visited and then enqueue it to the queue
				visited[v] = true
				queue.Enqueue(v)

			}
		}

	}

}
