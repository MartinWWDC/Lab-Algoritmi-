package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(valuta("4 8 + 4 * 6 +"))
	//fmt.Println(operator("-", []int{4, 5, 6}))
}

/**
 * leggi un token ( operatore o numero );
 * se il token è un numero
 * inseriscilo nella pila ;
 * se il token è un operatore
 * estrai due valori dalla pila ;
 * applica ad essi l’operatore;
 * inserisci il risultato nella pila;
 */

func valuta(exp string) int {
	fmt.Println(exp)
	g := 0
	pile := strings.Split(exp, " ")
	stack := []int{}

	for i := 0; i < len(pile); i++ {
		g, err := strconv.Atoi(pile[i])
		if err == nil {
			stack = append(stack, g)
		} else {
			break
		}
	}
	//fmt.Println(stack)
	//fmt.Println(len(stack), "<", len(exp))
	g = operator(pile[len(stack)], stack)
	if len(stack)+4 < len(exp) && len(stack) != 0 {
		str := strconv.Itoa(g)
		str = (str + " " + exp[len(stack)+3:])
		g = valuta(str)
		fmt.Println(g)

	}

	return g
}
func operator(operator string, ele []int) int {
	h := ele[0]
	for i := 1; i < len(ele); i++ {
		switch operator {
		case "+":
			h += ele[i]
		case "-":
			h -= ele[i]

		case "*":
			h *= ele[i]

		}
	}
	return h
}
