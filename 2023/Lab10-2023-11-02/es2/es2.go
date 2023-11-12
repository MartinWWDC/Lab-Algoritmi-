package main

import (
	"fmt"
	"strings"
)


type Stack interface{
	Pop() any
	Push(any)
	isIn(any) bool
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

func (s stackSlice)isIn(el any) bool {
	for _,i := range s{
		if el==i{
			return true	
		}
	}
	return false

}



func main() {
	//code:="<a> <b> </b> </c>"
	code:="<b> </b>"
	check(code)
}

func check(code string) {
	var s Stack
	pila := newStack()
	s = &pila

	tags := strings.Split(code, " ")
	for i,tag:= range tags {
		if tag[1]=='/'{
			el:=s.Pop()
			if el!=tag[2]{
				if !s.isIn(el){
					fmt.Println("erorre in posizione: ",i)
				}
			}
			
		}else{
			s.Push(tag[1])
		}
	}


}