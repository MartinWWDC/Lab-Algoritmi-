package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func contaOrbite(orbita string, orbite map[string]string) (tot int) {
	for {
		padre, ok := orbite[orbita]
		if ok {
			tot++
			orbita = padre
		} else {
			break
		}
	}
	return
}

func listaPadri(orbita string, orbite map[string]string) map[string]int {
	padri := make(map[string]int)
	var cammino int
	for {
		padre, ok := orbite[orbita]
		if ok {
			orbita = padre
			padri[orbita] = cammino
			cammino++
		} else {
			break
		}
	}
	return padri
}

func main() {
	orbite := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := strings.Split(scanner.Text(), ")")
		orbite[riga[1]] = riga[0]
	}
	var tot int
	for orbita, _ := range orbite {
		tot += contaOrbite(orbita, orbite)
	}
	//fmt.Println(tot)
	padriSAN := listaPadri("SAN", orbite)
	fmt.Println(padriSAN)
	var distanza, pos2 int
	inizio := "YOU"
	for {
		padre, ok := orbite[inizio]
		if ok {
			if pos1, ok := padriSAN[orbite[padre]]; ok {
				distanza = pos1 + pos2
				break
			} else {
				pos2++
				inizio = orbite[inizio]
			}

		} else {
			break
		}
	}
	fmt.Println(distanza + 1)
}
