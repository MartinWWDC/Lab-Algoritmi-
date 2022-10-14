package main

import (
	"fmt"
	"os"
)

/*
Si vuole scrivere un programma che legga una sequenza di stringhe e stampi, per ogni stringa, la sua prima
vocale (per semplicità assumiamo che le stringhe contengano solo lettere minuscole). Nel caso in cui una
stringa non contenga vocali il programma stampa “la parola non contiene vocali”.
*/
func main() {
	const N = 10
	numeri := os.Args[1:]
	for _, g := range numeri {
		for _, i := range g {
			if i == 'a' || i == 'e' || i == 'i' || i == 'o' || i == 'u' {
				fmt.Println(string(i))
				break
			}
		}
	}
}
