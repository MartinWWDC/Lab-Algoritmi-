package main

import "fmt"

/**
 * es1
 * . Si tracci algoX (cioè: si simuli l’esecuzione tenendo traccia dello stato del programma)
 * per table indicata qui sotto e x pari a 14. Si elenchino in particolare tutti i valori della
 * variabile i durante l’esecuzione.
 *
 * table -9 -1 0 13 14 14 12 29 31 24 36 36 44 44 8
 *
 * RISPOSTA: 4
 */

/**
 * es2
 * Si tracci AlgoY (cioè: si simuli l’esecuzione tenendo traccia dello stato del programma)
 * per table indicata qui sotto e x pari a 14. Si scrivano in particolare i valori della variabili
 * low, mid e high subito dopo l’esecuzione della riga 6 (ogni volta che viene eseguita).
 * 		  0 1 2 3 4 5 6 7 8 9 10 11 12 13 14
 * table -9 -1 0 13 14 14 12 29 31 24 36 36 44 44 8
 */

/**
 * es3
 * Si indichi quali delle seguenti affermazioni sono vere per AlgoX e/o per AlgoY. Per ciascuna affermazione si usi una delle cinque opzioni seguenti: X (= l’affermazione è corretta
 * solo per AlgoX), Y (= l’affermazione è corretta solo per AlgoY), X & Y (= l’affermazione
 * è corretta sia per per AlgoX che per AlgoY), no (= l’affermazione non è corretta né per
 * AlgoX né per AlgoY), non so (= non so la risposta).
 * a) L’algoritmo esamina gli elementi partendo dall’indice minimo al massimo. X
 * b) L’algoritmo cerca l’elemento massimo di table. no
 * c) Come secondo valore, l’algoritmo restituisce sempre l’indice minore per l’elemento x, se questo è contenuto in table. X
 * d) L’algoritmo ordina gli elementi in table. no
 * e) L’algoritmo restituisce tutti gli indici in cui si trova l’elemento x. no
 * f) L’algoritmo esamina tutti gli elementi di table. no
 * g) L’algoritmo è corretto solo se table è ordinato. no
 * h) L’algoritmo alla fine restituisce sempre false, -1. no
 *
 */

func algoX ( table []int, x int) (bool, int) {
	for i , el := range table {
		if el == x {
			return true, i
		}
	}
	return false, -1
}

func algoY ( table []int, x int) (bool, int) {
	low , high := 0, len( table ) -1
		
	for low <= high {
		mid := ( low + high ) / 2
		if table [ mid ] < x {
			low = mid + 1
		} else if table [ mid ] > x {
			high = mid - 1
		} else {
			return true, mid
		}
		fmt.Println(low,high,mid)
	}
	return false, -1
}

func main() {

	//ES1
	table:= []int{-9,-1,0,13,14,14,12,29,31,24,36,36,44,44,8}

	x:=14
	b,i:=algoY(table,x)
	fmt.Println(b,i)
}