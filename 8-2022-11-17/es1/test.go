package main

import "fmt"

type g struct {
	h int
}

func main() {
	var n [5]int
	var h g
	add(n, 5)

	fmt.Println(n)

	check(h)
	fmt.Println(h)

}

func add(n [5]int, g int) {
	n[0] = g

}
func check(h g) {
	h.h = 7

}
