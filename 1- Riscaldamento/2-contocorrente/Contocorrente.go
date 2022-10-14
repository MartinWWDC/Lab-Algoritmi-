package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Scrivere un programma che legge da riga di comando un intero che rappresenta il saldo di un conto corrente.
Il programma legge poi da standard input una serie di numeri interi che rappresentano spese da addebitare sul
conto e stampa il saldo finale. La lettura si interrompe quando il saldo è <=0.
NOTA: Se avete difficoltà a scrivere il programma, potete partire da qui:
http://parsons.problemsolving.io/puzzle/5bc3cba95a44449fb6212e3f5e588504

*/
func main() {
	saldo, _ := strconv.Atoi(os.Args[1])
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		movement, _ := strconv.Atoi(scanner.Text())
		saldo -= movement
		if saldo <= 0 {
			fmt.Println("Saldo insuffficente")

			break
		}
	}
}
