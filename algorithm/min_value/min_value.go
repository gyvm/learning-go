package main

import (
	"fmt"
	"slices"
)

func findMin(arrays []int, length int) int {
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
	var minValue = findMin(score, len(score))
	fmt.Printf("最小値は%d\n", minValue)

	// 組み込み関数
	// https://tip.golang.org/ref/spec#Min_and_max
	fmt.Println(slices.Min(score)) // sliceの場合
	fmt.Println(min(40, 23, 12, 43))
	fmt.Println(min("c", "a", "b")) // アルファベット順
}

//=> 最小値は4
//4
//12
//a
