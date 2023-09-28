package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := input()
	fmt.Println(n)
	fmt.Println(mergeSort(n))
	//fmt.Println(concat(n[0:1], n[4:6]))

}

func input() (arr []int) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		n, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, n)

		//	if n == 0 {
		//		break
		//	}

	}
	return
}

func insertionsort(items []int) []int {
	//var n = len(items)
	for i := range items {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
	return items
}

func selectionSort(arr []int) []int {
	for n := range arr {

		for i := n; i < len(arr); i++ {
			fmt.Println(arr)
			if arr[n] > arr[i] {
				g := arr[i]
				arr[i] = arr[n]
				arr[n] = g
			}
		}
	}
	return arr
}

func mergeSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}
	sub1 := mergeSort(arr[:len(arr)/2])
	sub2 := mergeSort(arr[len(arr)/2:])
	return merge(sub1, sub2)
}
func merge(arr1 []int, arr2 []int) (final []int) {
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] > arr2[j] {
			final = append(final, arr2[j])
			j++
		} else {
			final = append(final, arr1[i])
			i++
		}
	}
	for ; i < len(arr1); i++ {
		final = append(final, arr1[i])
	}
	for ; j < len(arr2); j++ {
		final = append(final, arr2[j])
	}
	return
}
