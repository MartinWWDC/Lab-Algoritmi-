package main

import (
	"fmt"
	"os"
)

/*
Vogliamo leggere una sequenza di N interi (almeno 3), che descrivono il consumo di elettricità in N giorni, e
stampare i giorni in cui il consumo è stato inferiore rispetto sia al giorno prima che al giorno dopo. I giorni
sono numerati a partire da 1.
*/
func main() {
	consumi := os.Args[1:]
	for i := range consumi {
		if i != 0 && i != len(consumi)-1 && consumi[i] < consumi[i-1] && consumi[i] < consumi[i+1] {
			fmt.Println("pos:", i+1, "consumo: ", consumi[i])
		}
	}

}
