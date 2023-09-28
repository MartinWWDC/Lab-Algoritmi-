package main

import "fmt"

func main() {
	table := []int{-9, -1, 0, 13, 14, 12, 29, 31, 24, 36, 36, 44, 44, 8}
	x := 36
	check, i := algoX(table, x)
	fmt.Println(check, i)
	check, i = algoY(table, x)
	fmt.Println(check, i)
}

func algoX(table []int, x int) (bool, int) {
	//check if el from table  == x
	for i, el := range table {
		//fmt.Println(i)

		if el == x {
			return true, i
		}
	}

	return false, -1
}
func algoY(table []int, x int) (bool, int) {
	low, high := 0, len(table)-1
	for low <= high {
		mid := (low + high) / 2
		if table[mid] < x {
			fmt.Println(low, mid, high)
			low = mid + 1
		} else if table[mid] > x {
			high = mid - 1
		} else {
			return true, mid
		}
	}
	return false, -1
}

/*
3.
a)y
b)no
c)x
d)no
e)no (ne restituisce solo 1 (?))
f)x (cosa si intende per tutti? si intende che tutti gli elementi possono venir controllati o che  tutti gli elementi vengono controllati durante l'esecuzione?)
g)no
h)no
*/
/*
y è più eccicente  in quanto fa al massimo n/2 contronti mentre x ne fa al  masimo n
*/
/*
6.
si, basta aggiungere un array in cui vangano salvate le posizioni degli elementi e non fermare la ricerca appena si trova un elemento uguale

*/
/*
7.
no, non è facile in quando l'algormito  scarta parti dell'array
*/
