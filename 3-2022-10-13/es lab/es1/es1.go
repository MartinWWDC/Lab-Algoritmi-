package main

import "fmt"

func main() {
	fmt.Println(multiply(1, 4))
}

func multiply(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	} else {
		return x + multiply(x, y-1)
	}
}
