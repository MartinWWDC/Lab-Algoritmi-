package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack interface{
	Pop() int 
	Push(i int)
}

type stackSlice []int

func newStack() stackSlice {
	return stackSlice([]int{})
}

// Effettua una Pop nella slice e ritorna il valore eliminato
func (s *stackSlice) Pop() int {
	n:=(*s)[len(*s)-1]
	*s=(*s)[:len(*s)-1]
	return n
}

// Effettua una Push nella slice
func (s *stackSlice) Push(i int) {
	*s = append(*s, i)
}


// 1.1. Valutazione di un’espressione in notazione postfissa
func valuta(espressione  string) int{
	var s Stack
	pila := newStack()
	s = &pila

	var tokens []string = strings.Split(espressione, " ")

	// leggi un token;
	for _, token := range tokens {
		n,err:=strconv.Atoi(token)
		// se il token è un numero
		if err==nil {
			s.Push(n)
		}

		// se il token è un operatore
		if isOperatore(token){
			// estrai due valori dalla pila;
			n1:=s.Pop()
			n2:=s.Pop()
			
			// applica ad essi l'operatore e inserisci il risultato nella pila;
			switch token {
			case "+":
				s.Push(n1+n2)
			case "-":
				s.Push(n2-n1)

			case "*":
				s.Push(n1*n2)

			case "/":
				s.Push(n2/n1)
			}		
		}
		

		
	}
	return s.Pop()
}


func isOperatore(token string)bool{
	return token=="+" ||token=="-" ||token=="*" ||token=="/"
}


func main() {
	// 1.1.
	var espressione1 string = "5 3 - 2 *"
	fmt.Println(valuta(espressione1))

	// 1.2
	
}