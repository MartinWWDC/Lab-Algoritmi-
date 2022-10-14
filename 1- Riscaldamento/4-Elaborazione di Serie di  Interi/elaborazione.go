package main

import "fmt"

/*
Scrivete le tre funzioni specificate sotto. Per testarle potete utilizzare il main riportato qui (con gli opportuni
adattamenti).
1 package main
2
3 import "fmt"
4
5 func main () {
6 const N = 10
7 numeri := make([]int, N)
8
9 for i := 0; i < N; i ++ {
10 fmt . Scan (& numeri [i ])
11 }
12
13 fmt .Println( nomeFunzione ( numeri ))
14 }
1. Scrivere una funzione stranoProdotto(numeri []int) int che, data come parametro una slice di
interi, trovi quelli che sono maggiori di 7 e multipli di 4 e ne restituisca il prodotto. Ad esempio, se la
slice contiene i numeri 12, 3, 4, 8, 9, 2, la funzione dovrà restituire il numero 96 (pari al prodotto di 12
per 8).
2. Scrivere una funzione pariDispari(numeri []int) che, data come parametro una slice di interi,
stampi, per ciascun numero, se è pari o dispari.
3. Scrivere una funzione minDispari(numeri []int) int che, data una slice di interi, restituisca il più
piccolo numero dispari (la slice può contenere sia numeri positivi che negativi)

*/
func main() {
	const N = 10
	numeri := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&numeri[i])
	}
	fmt.Println(numeri)
	fmt.Println(minDispari(numeri))
}

func stranoProdotto(numeri []int) int {
	res := 1
	for _, n := range numeri {

		if n > 7 && n%4 == 0 {
			fmt.Println(n)
			res *= n
		}
	}
	return res
}

func pariDispari(numeri []int) {
	for _, g := range numeri {
		if g%2 == 0 {
			fmt.Println("pari")
		} else {
			fmt.Println("dispari")

		}
	}
}

func minDispari(numeri []int) int {
	var min int
	for i := range numeri {
		if i == 0 {
			min = numeri[i]
		} else {
			if numeri[i] < min && numeri[i]%2 != 0 {
				min = numeri[i]
			}
		}
	}
	return min
}
