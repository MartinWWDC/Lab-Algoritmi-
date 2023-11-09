package main

import "fmt"

func main() {
	arr := []int{4, 3, 2, 1}
	SelectionSort(&arr)
	fmt.Println(arr)

}

func SelectionSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		minIndex := i
		for j := i; j < len(*arr); j++ {
			if (*arr)[j] < (*arr)[minIndex] {
				minIndex = j
			}
		}
		t := (*arr)[i]
		(*arr)[i] = (*arr)[minIndex]
		(*arr)[minIndex] = t
	}

}