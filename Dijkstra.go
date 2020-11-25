// A Golang program for Dijkstra's single source shortest path algorithm.
// This is the program written for adjacency matrix representation of the graph
package main

import (
	"fmt"
	"math"
)

const (
	V int = 9
)

// Function to find the find the vertex with minimum distance value,
//from the set of vertices which are not yer included in tree of shortest path

func minimum_distance(distance [V]int, vertices [V]bool) int {

	// Initialize min value
	var min int = math.MaxInt64
	var min_index int = -1
	for i := 0; i < V; i++ {
		if vertices[i] == false && distance[i] <= min {
			min = distance[i]
			min_index = i
		}

	}
	return min_index
}

//  Function to print the final result

func final_result(distance [V]int, n int) {
	fmt.Println(" Distance of vertex from the source:")
	for i := 0; i < V; i++ {
		fmt.Println("Vertex ", i, "is at a distance of ", distance[i])
	}
}

// Function for the implementation of Dijkstra's single source shortest path
// algorithm for a graph represented using adjacency matrix

func dijkstra(graph [][]int, source int) {
	var distance [V]int  // Array for shortest distance from source to i
	var vertices [V]bool // Array for tracking visited vertices

	// Initializing all the distances to the vertices as INFINITE and Vertices[] as false
	for i := 0; i < V; i++ {
		distance[i] = math.MaxInt64
		vertices[i] = false
	}

	// Distance of source vertex from itself is always 0
	distance[source] = 0

	// Finding the shortest path to all the vertices
	for count := 0; count < V; count++ {

		// Considering the minimum distance vertex from the set of vertices not yet visited.
		// u is always the source in first iteration
		var u int = minimum_distance(distance, vertices)

		// Updating the visited vertex as processed
		vertices[u] = true

		// Update dist value of the adjacent vertices of the visited vertex.
		for r := 0; r < V; r++ {

			// Update distance[V] only if is not in vertices, there is an  edge from u to v,and total weight
			//  of path from source to V through u is smaller than current value of distance[V]
			if !vertices[r] && graph[u][r] != 0 &&
				distance[u] != math.MaxInt64 && distance[u]+graph[u][r] < distance[r] {
				distance[r] = distance[u] + graph[u][r]
			}
		}
	}
	// To print the distance array representing shortest paths from source to other vertices
	final_result(distance, V)
}

// creating the example graph using adjacent matrix representation

func main() {
	var graph = [][]int{
		{0, 4, 0, 0, 0, 0, 0, 8, 0},
		{4, 0, 8, 0, 0, 0, 0, 11, 0},
		{0, 8, 0, 7, 0, 4, 0, 0, 2},
		{0, 0, 7, 0, 9, 14, 0, 0, 0},
		{0, 0, 0, 9, 0, 10, 0, 0, 0},
		{0, 0, 4, 14, 10, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 1, 6},
		{8, 11, 0, 0, 0, 0, 1, 0, 7},
		{0, 0, 2, 0, 0, 0, 6, 7, 0}}
	dijkstra(graph, 0)
}
