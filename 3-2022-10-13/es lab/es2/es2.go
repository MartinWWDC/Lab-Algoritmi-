package main

import "fmt"

func main() {
	arr := []int{1, 2, 5, 7, -2, 10, 9, 21, 3, 8}
	fmt.Println(largest(arr))
}
func largest(numbers []int) int {
	n := len(numbers)
	if n == 1 {
		return numbers[0]
	}
	return max(numbers[n-1], largest(numbers[:n-1]))
}
func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
