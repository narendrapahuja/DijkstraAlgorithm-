// Project Group Members :
// Shravya Donthisaram, Levi Barrantes, Zhiyan Yu, Dylan Jezo,  Drew Cavanaugh , and  Narendra Pahuja
// Date Modified : 1st December, 2020

// Dijkstra's algorithm is an algorithm for finding the shortest paths between nodes in a graph.
// Below is the Golang program for Dijkstra's single source shortest path algorithm
// with input as a graph in adjacent matrix representation and
// output as the shortest distance from source to every vertex and the path.
// This is the program written for adjacency matrix representation of the graph

package main

// Packages to format and use the maximum intial distances for the vertices
import (
	"fmt"
	"math"
)

// Constant declared to represent the number of vertices in the graph
// defining the number of vertices
const (
	V int = 9
)

// Function to find the find the vertex with minimum distance value,
// from the set of vertices which are not yer included in tree of shortest path
// Input parameters are distances and vertices
// Returns the Minimum distance to a vertex
func findMinimumDistance(distance [V]int, vertices [V]bool) int {

	// Initialize min value
	var min int = math.MaxInt64
	var minIndex int = -1
	for i := 0; i < V; i++ {
		if vertices[i] == false && distance[i] <= min {
			min = distance[i]
			minIndex = i
		}

	}
	return minIndex
}

// Function to print the final result
// Input parameters are distances and the path
// Prints the Vertex, distance and the path

func displayResult(distance [V]int, path [V]int) {
	fmt.Println(" Distance of vertex from the source:")
	for i := 0; i < V; i++ {
		fmt.Println("\n\nVertex ", i, "is at a distance of ", distance[i], "\n Shortest Path is")
		printPath(path, i)
	}

}

// Function to print the shortest path
// Input parameters are path and the vertex
// Prints the path

func printPath(path [V]int, n int) {
	// Base case if n is source
	if path[n] == -1 {
		fmt.Printf("%d %s", n, " ")
		return
	}
	printPath(path, path[n])
	fmt.Printf("%d %s", n, " ")

}

// Function for the implementation of Dijkstra's single source shortest path
// algorithm for a graph represented using adjacency matrix
// Input is the graph in the form of adjacent matrix and source vertex
// Output

func dijkstra(graph [][]int, source int) {
	var distance [V]int  // Array for shortest distance from source to i
	var vertices [V]bool // Array for tracking visited vertices
	var path [V]int      // Array for storing vertices following the shortest path

	// Initializing path to source as -1, all the distances to the vertices as INFINITE and Vertices[] as false
	for i := 0; i < V; i++ {
		path[0] = -1
		distance[i] = math.MaxInt64
		vertices[i] = false
	}

	// Distance of source vertex from itself is always 0
	distance[source] = 0

	// Finding the shortest path to all the vertices
	for count := 0; count < V; count++ {

		// Considering the minimum distance vertex from the set of vertices not yet visited.
		// u is always the source in first iteration
		var u int = findMinimumDistance(distance, vertices)

		// Updating the visited vertex as processed
		vertices[u] = true

		// Update dist value of the adjacent vertices of the visited vertex.
		for r := 0; r < V; r++ {

			// Update distance[V] only if is not in vertices, there is an  edge from u to v,and total weight
			//  of path from source to V through u is smaller than current value of distance[V]
			if !vertices[r] && graph[u][r] != 0 &&
				distance[u] != math.MaxInt64 && distance[u]+graph[u][r] < distance[r] {
				distance[r] = distance[u] + graph[u][r]
				path[r] = u
			}

		}
	}
	// To print the distance array representing shortest paths and distance from source to other vertices
	displayResult(distance, path)
}

// Creating the example graph using adjacent matrix representation

func main() {
	var graph = [][]int{{0, 4, 0, 0, 0, 0, 0, 8, 0},
		{3, 0, 7, 0, 0, 0, 0, 10, 0},
		{0, 7, 0, 6, 0, 3, 0, 0, 1},
		{0, 0, 6, 0, 8, 13, 0, 0, 0},
		{0, 0, 0, 8, 0, 9, 0, 0, 0},
		{0, 0, 3, 0, 10, 0, 1, 0, 0},
		{0, 0, 0, 13, 0, 1, 0, 2, 5},
		{8, 10, 0, 0, 0, 0, 2, 0, 6},
		{0, 0, 1, 0, 0, 0, 5, 6, 0}}

	dijkstra(graph, 0)
}
