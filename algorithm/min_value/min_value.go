package main

import "fmt"

func min(arrays []int, length int) int {
	var minIndex = 0
	for i := 1; i < length; i++ {
		if arrays[minIndex] > arrays[i] {
			minIndex = i
		}
	}
	return arrays[minIndex]
}

func main() {
	var score = []int{60, 80, 21, 92, 100, 4}
	var minValue = min(score, len(score))
	fmt.Printf("最小値は %d\n", minValue)
}
