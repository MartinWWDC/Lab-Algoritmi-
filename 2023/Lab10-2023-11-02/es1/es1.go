package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack interface{
	Pop() any
	Push(any)
}

type stackSlice []any

func newStack() stackSlice {
	return stackSlice([]any{})
}

// Effettua una Pop nella slice e ritorna il valore eliminato
func (s *stackSlice) Pop() any {
	n:=(*s)[len(*s)-1]
	*s=(*s)[:len(*s)-1]
	return n
}

// Effettua una Push nella slice
func (s *stackSlice) Push(i any) {
	*s = append(*s, i)
}


// 1.1. Valutazione di un’espressione in notazione postfissa
func valuta(espressione  string) any{
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
				s.Push(n1.(int)+n2.(int))
			case "-":
				s.Push(n2.(int)-n1.(int))

			case "*":
				s.Push(n1.(int)*n2.(int))

			case "/":
				s.Push(n2.(int)/n1.(int))
			}		
		}
		

		
	}
	return s.Pop()
}


func isOperatore(token string)bool{
	return token=="+" ||token=="-" ||token=="*" ||token=="/"
}

// 1.2. Conversione da notazione infissa a notazione postfissa



func main() {
	// 1.1.
	var espressione1 string = "5 3 - 2 *"
	fmt.Println(valuta(espressione1))

	// 1.2
	// Puppatina al bomberazzo??	
}