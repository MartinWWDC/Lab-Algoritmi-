package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "<a> <b> </b> <c> <d> </d>"
	fmt.Println(isBalanced(input))
}

func isBalanced(list string) bool {
	var pila []string
	tags := strings.Split(list, " ")
	h := true
	for g := range tags {
		//fmt.Println(tags[g])
		//fmt.Println(tags[g][0:2])
		//fmt.Println(pila)

		if tags[g][0:2] == "</" {
			//fmt.Println("ss:")

			if pila[len(pila)-1][1] == tags[g][2] {
				pila = remove(pila, len(pila)-1)
			} else {
				h = false
				fmt.Println("errore in pos:", g+1)
			}
			//fmt.Println(tags[g])

		} else {
			pila = append(pila, tags[g])
		}
	}
	if h {
		return len(pila) == 0

	}
	return h
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
