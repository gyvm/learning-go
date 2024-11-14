// 有向グラフと深さ優先探索
package main

import (
	"fmt"
)

const MaxVertexSize = 5
const StackSize = 1000

type Vertex struct {
	data    string
	visited bool
}

var top = -1
var stacks = make([]int, StackSize)

func push(element int) {
	if top >= MaxVertexSize-1 {
		fmt.Printf("Stack size exceeds MaxVertexSize (%d)\n", MaxVertexSize)
		return
	}

	top++
	stacks[top] = element
}

func pop() int {
	if isEmpty() {
		return -1
	}
	var data = stacks[top]
	top--
	return data
}

func peek() int {
	if isEmpty() {
		return -1
	}
	return stacks[top]
}

func isEmpty() bool {
	return top == -1
}

var size = 0
var vertices = make([]*Vertex, MaxVertexSize)

// adjacency = 隣接
var adjacencyMatrix [MaxVertexSize][MaxVertexSize]int

func addVertex(data string) {
	if size >= MaxVertexSize {
		fmt.Println("Reached MaxVertexSize")
		return
	}
	vertex := &Vertex{data: data, visited: false}
	vertices[size] = vertex
	size++
}

func addEdge(from int, to int) {
	if from < 0 || from >= size || to < 0 || to >= size {
		fmt.Println("Invalid edge index")
		return
	}
	adjacencyMatrix[from][to] = 1
}

func clearVisited() {
	for i := 0; i < size; i++ {
		vertices[i].visited = false
	}
}

func depthFirstSearch(start int) {
	if start < 0 || start >= size {
		fmt.Println("Invalid start vertex")
		return
	}

	vertices[start].visited = true
	fmt.Printf("%s", vertices[0].data)
	push(start)

	for !isEmpty() {
		row := peek()
		column := findAdjacencyUnVisitedVertex(row)
		if column == -1 {
			pop()
		} else {
			vertices[column].visited = true
			fmt.Printf("->%s", vertices[column].data)
			push(column)
		}
	}
	fmt.Println()
	clearVisited()
}

func findAdjacencyUnVisitedVertex(row int) int {
	for column := 0; column < size; column++ {
		if adjacencyMatrix[row][column] == 1 && !vertices[column].visited {
			return column
		}
	}
	return -1
}

func printGraph() {
	fmt.Println("Adjacency Matrix:")
	fmt.Print("  ")

	for i := 0; i < size; i++ {
		fmt.Printf("%s ", vertices[i].data)
	}
	fmt.Println()

	for i := 0; i < size; i++ {
		fmt.Printf("%s ", vertices[i].data)
		for j := 0; j < size; j++ {
			fmt.Printf("%d ", adjacencyMatrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	addVertex("A")
	addVertex("B")
	addVertex("C")
	addVertex("D")
	addVertex("E")

	addEdge(0, 1)
	addEdge(0, 2)
	addEdge(0, 3)
	addEdge(1, 2)
	addEdge(1, 3)
	addEdge(2, 3)
	addEdge(3, 4)

	printGraph()
	fmt.Println("\nDepth-first search traversal output:")
	depthFirstSearch(0)
}

// =>
// Adjacency Matrix:
//   A B C D E
// A 0 1 1 1 0
// B 0 0 1 1 0
// C 0 0 0 1 0
// D 0 0 0 0 1
// E 0 0 0 0 0
//
// Depth-first search traversal output:
// A->B->C->D->E
