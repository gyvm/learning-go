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
	top++
	stacks[top] = element
}

func pop() int {
	if top == -1 {
		return -1
	}
	var data = stacks[top]
	top--
	return data
}

func peek() int {
	if top == -1 {
		return -1
	}
	return stacks[top]
}

func isEmpty() bool {
	if top == -1 {
		return true
	}
	return false
}

var size = 0
var vertices = make([]*Vertex, MaxVertexSize)

// adjacency = 隣接
var adjacencyMatrix [MaxVertexSize][MaxVertexSize]int

func addVertex(data string) {
	var vertex Vertex
	vertex.data = data
	vertex.visited = false
	vertices[size] = &vertex
	size++
}

func addEdge(from int, to int) {
	adjacencyMatrix[from][to] = 1
}

func clear() {
	for i := 0; i < size; i++ {
		vertices[i].visited = false
	}
}

func depthFirstSearch() {
	vertices[0].visited = true
	fmt.Printf("%s", vertices[0].data)
	push(0)
	for {
		if isEmpty() {
			break
		}
		var row = peek()
		var column = findAdjacencyUnVisitedVertex(row)
		if column == -1 {
			pop()
		} else {
			vertices[column].visited = true
			fmt.Printf("->%s", vertices[column].data)
			push(column)
		}
	}
	fmt.Printf("\n")
	clear()
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
	fmt.Printf("Two-demensional array travesal vertex edge and adjacent array : \n ")
	for i := 0; i < MaxVertexSize; i++ {
		fmt.Printf("%s", vertices[i].data)
	}
	fmt.Printf("\n")

	for i := 0; i < MaxVertexSize; i++ {
		fmt.Printf("%s", vertices[i].data)
		for j := 0; j < MaxVertexSize; j++ {
			fmt.Printf("%d", adjacencyMatrix[i][j])
		}
		fmt.Printf("\n")
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
	fmt.Printf("\nDepth-first search traversal output : \n")
	depthFirstSearch()
}
