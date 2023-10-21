package main

import "fmt"

func sassi(h int) int {
	if h == 1 {
		return 1
	} else {
		return sassi(h-1) + (h * h)
	}
}

func main() {
	b := sassi(3)
	fmt.Println(b)
}