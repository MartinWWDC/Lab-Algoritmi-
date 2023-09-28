package main

import "fmt"

var counter int

func main() {
	counter = 0
	arr := []int{1, 7, 2, 4, 21, 6, 8}
	fmt.Println(largest(arr))
	fmt.Println(counter)
}
func largest(numbers []int) int {
	n := len(numbers)
	if n == 1 {
		return numbers[0]
	}
	return max(numbers[n-1], largest(numbers[:n-1]))
}
func max(x int, y int) int {
	counter++
	if x > y {
		return x
	} else {
		return y
	}

}
