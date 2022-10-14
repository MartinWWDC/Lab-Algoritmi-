package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Scrivere un programma supera100.go che legge da standard input una sequenza di interi positivi terminata
da -1 e stampa il primo numero che supera 100, se presente; altrimenti stampa “nessun numero maggiore di
100”.
NOTA: Se avete difficoltà a scrivere il programma, potete partire da qui:
http://parsons.problemsolving.io/puzzle/e334652c5f7f4a82ad21565a01e7cbb7

*/

func main() {
	h := true
	arr := os.Args[1:]
	for i := range arr {
		g, _ := strconv.Atoi(arr[i])
		if g > 100 {
			fmt.Println(g)
			h = false
			break

		}

	}
	if h {
		fmt.Println("nessun numero maggiore di 100")
	}

}
