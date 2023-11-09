package main

import "fmt"

func multiply(x, y int) int {
	if y == 1 {
		return x
	} else {
		return x + multiply(x, y-1)
	}
}

func main() {

	fmt.Println(multiply(5, 5))
}