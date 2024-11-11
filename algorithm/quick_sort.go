package main

import "fmt"

func quickSort(array []int, low int, high int) {
	if low > high {
		return
	}
	var i = low
	var j = high
	var threshold = array[low]

	for {
		if i >= j {
			break
		}
		for {
			if i >= j || array[j] <= threshold {
				break
			}
			j--
		}
		if i < j {
			array[i] = array[j]
			i++
		}
		for {
			if i >= j || array[i] > threshold {
				break
			}
			i++
		}
		if i < j {
			array[j] = array[i]
			j--
		}
	}

	array[i] = threshold
	quickSort(array, low, i-1)
	quickSort(array, i+1, high)
}

func qsort(array []int, length int) {
	if length > 0 {
		quickSort(array, 0, length-1)
	}
}

func main() {
	var scores = []int{50, 33, 99, 82, 73, 100, 2, 93}
	var length = len(scores)

	qsort(scores, length)

	for i := 0; i < length; i++ {
		if i == length-1 {
			fmt.Println(scores[i])
		} else {
			fmt.Printf("%d,", scores[i])
		}
	}
}

// => 2,33,50,73,82,93,99,100
