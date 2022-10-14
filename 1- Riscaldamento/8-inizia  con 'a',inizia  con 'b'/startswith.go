package main

import "fmt"

/*
Si vuole scrivere un programma che legga una sequenza di dieci stringhe e stampi il numero di stringhe che
cominciano con la lettera a e il numero di stringhe che cominciano con b.
Per calcolare il numero di stringhe che cominciano per a, usiamo il piano del conteggio; lo stesso piano serve
anche a calcolare il numero di stringhe che cominciano per b.

*/

func main() {
	const N = 10
	temp := ""
	worda, wordb := 0, 0
	for i := 0; i < N; i++ {
		fmt.Scan(&temp)
		if temp[0] == 'a' {
			//worda = append(worda, temp)
			worda++
		}
		if temp[0] == 'b' {
			//wordb = append(wordb, temp)
			wordb++

		}
	}
	fmt.Println(worda)
	fmt.Println(wordb)
}
