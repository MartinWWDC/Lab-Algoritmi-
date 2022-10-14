package main

import (
	"fmt"
	"os"
)

/*
1. Scrivete un programma che legga una riga di caratteri terminata da a-capo, poi legga un’altra lettera e
stampi quante volte la lettera compare nella prima riga.
2. Scrivete un programma che legga una riga di caratteri terminata da a-capo e che visualizzi un istogramma con una barra per ogni carattere dell’alfabeto, lunga quanto il numero delle sue occorrenze. Il
programma non deve visualizzare le barre delle lettere che non compaiono e non deve fare distinzione
fra maiuscole e minuscole (a tal fine potete usare le funzioni dichiarate nel file ctype.h)
3. Due parole costituiscono un anagramma se l’una si ottiene dall’altra permutando le lettere (es: attore,
teatro). Scrivete un programma che legga due parole e verifichi se sono anagrammi.
Suggerimento: potete sfruttare il programma scritto per l’esercizio precedente.
*/
func main() {
	es2()
}

func es1() {
	var lett string
	j := 0
	frase := os.Args[1:]
	fmt.Scan(&lett)
	for _, i := range frase {
		for _, h := range i {
			if string(h) == lett {
				j++
			}
		}

	}
	fmt.Println(j)
}

func es2() {
	mapp := make(map[rune]int)
	frase := os.Args[1:]
	for _, i := range frase {
		for _, h := range i {
			mapp[h]++
		}

	}
	fmt.Println(mapp)
	printMapp(mapp)
}
func printMapp(mapp map[rune]int) {
	for key, n := range mapp {
		fmt.Print(string(key), ":")
		for i := 0; i < n; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
