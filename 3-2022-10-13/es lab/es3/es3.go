package main

import "fmt"

func main() {
	fmt.Println(sassi(3))
}

func sassi(h int) int {
	if h == 0 {
		return 0
	} else {
		return h*h + sassi(h-1)
	}
}
