package main

import "fmt"

/*
Scrivete le funzioni specificate sotto. Per testarle potete utilizzare il main riportato qui, con gli opportuni
adattamenti
1 package main
2
3 import "fmt"
4
5 func main () {
6 const N = 10
7 parole := make([]string, N)
8
9 for i := 0; i < N; i ++ {
10 fmt . Scan (& parole [i ])
11 }
12
13 fmt .Println( nomeFunzione ( parole ))
14 }
1. Scrivere una funzione quanteConA(parole []string) int che, data una slice di stringhe, restituisca
quante stringhe contengono il carattere ‘a’.
2. Scrivere una funzione primaConA(parole []string) string che, data una slice di stringhe, restituisca la prima parola che contentiene il carattere ‘a?, o la stringa vuota.
3. Scrivere una funzione piuCorta(parole []string) int che, data una slice di stringhe, restituisca la
lunghezza della stringa più corta in termini di byte.
*/
func main() {
	const N = 10
	parole := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&parole[i])
	}
	fmt.Println(primaConA(parole))
}

func quanteConA(parole []string) int {
	num := 0
	for _, i := range parole {
		for _, g := range i {
			if g == 'a' || g == 'A' {
				fmt.Println("I:", i)
				num++
				break
			}
		}
	}
	return num
}

func primaConA(parole []string) string {
	for _, i := range parole {
		for _, g := range i {
			if g == 'a' || g == 'A' {
				return i
			}
		}
	}
	return ""

}
