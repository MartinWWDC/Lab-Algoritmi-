package main

import "fmt"

/**
 * Scrivere una funzione stranoProdotto(numeri []int) int che, data come parametro una slice di
interi, trovi quelli che sono maggiori di 7 e multipli di 4 e ne restituisca il prodotto. Ad esempio, se la
slice contiene i numeri 12, 3, 4, 8, 9, 2, la funzione dovrà restituire il numero 96 (pari al prodotto di 12
per 8)
*/
func stranoProdotto(numeri []int)int{
	n:=1
	for i := 0;i < len(numeri);i++{
		if numeri[i] >7 && numeri[i]%4==0{
			n*=numeri[i]
		}
	}
	return n

}

/**
 * Scrivere una funzione pariDispari(numeri []int) che, data come parametro una slice di interi, stampi, 
 * per ciascun numero, se è pari o dispari
 */

func pariDispari(numeri []int){
	for  r := range numeri{
		if numeri[r]%2==0{
			fmt.Println("pari")
		}else{
			fmt.Println("dispari")

		}
	}

}
/**
 * Scrivere una funzione minDispari(numeri []int) int che, data una slice di interi, restituisca il più
 * piccolo numero dispari (la slice può contenere sia numeri positivi che negativi).
 */

func minDispari(numeri []int) int{
	min:=0
	for r := range numeri{
		if r==0{
			min=numeri[r]
		}else if min>numeri[r]{
			min=numeri[r]
		}
	}
	return min
}
func main () {
	const N = 10
	numeri := make([]int, N)

	for i := 0; i < N; i ++ {
	fmt.Scan(& numeri [i ])
	}

	//fmt.Println(stranoProdotto(numeri))
	fmt.Println(minDispari(numeri))
	//pariDispari(numeri)
} 

