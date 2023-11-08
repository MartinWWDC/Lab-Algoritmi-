package main

import (
	"fmt"
	"os"
)

func main() {
	/* s è una variabile interfaccia
	s non ha ancora un tipo definito, può prendere il valore di un qualunque tipo
	che implementa l'iterfaccia stack */
	var s Stack
	fmt.Println("s è una variabile interfaccia Stack")

	if os.Args[1] == "list" {
		l := NewStackLinkedList()
		s = &l // s ha preso un valore di tipo *StackLinkedList
		fmt.Println("s ora ha un valore di tipo *StackLinkedList\n")
	}
	//var s stackLinkedList = NewStackLinkedList()

	if os.Args[1] == "slice" {
		a := NewStackSlice() // variabile di tipo stackSlice
		s = &a               // s ha preso un valore di tipo *StackSlice
		fmt.Println("s ora ha un valore di tipo *StackSlice\n")
		//var s stackSlice = NewStackSlice()
	}

	for i := 101; i < 105; i++ {
		s.Push(fmt.Sprint("<", i, ">")) // pila di stringhe
		//s.Push(i) // pila di interi
		fmt.Println(s)
		fmt.Println("top:", s.Top())
		fmt.Println("size", s.Size(), "\n")

	}

	for i := 0; i < 6; i++ {
		res := s.Pop()
		if res == nil {
			fmt.Println("pila vuota")
		}
		fmt.Println(res)
		fmt.Println(s)
		fmt.Println("la pila è vuota?", s.IsEmpty(), "\n")
	}

}

type Stack interface {
	Top() any
	Pop() any
	Push(any)
	IsEmpty() bool
	IsFull() bool
	Size() int
}

/*
/* implementazione del tipo stack mediante slice */

type stackSlice []any

func NewStackSlice() stackSlice {
	return stackSlice([]any{})
}

func (s *stackSlice) Top() any {
	return (*s)[len(*s)-1]
}

func (s *stackSlice) Pop() any {
	n := len(*s)
	if n == 0 {
		return nil
	} else {
		res := (*s)[n-1]
		*s = (*s)[:n-1]
		return res
	}
}

func (s *stackSlice) Push(x any) {
	*s = append(*s, x)
	fmt.Println("dopo push di", x, *s)
}

func (s *stackSlice) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stackSlice) IsFull() bool {
	return false
}

func (s *stackSlice) Size() int {
	return len(*s)
}

/* basta invocare la Sprint dopo aver convertito a slice */
func (s *stackSlice) String() string {
	return fmt.Sprint(*s)
}

/* implementazione del tipo stack mediante liste concatenate */

type stackLinkedList linkedList

func NewStackLinkedList() stackLinkedList {
	return stackLinkedList(linkedList{nil, 0})
}

func (s *stackLinkedList) Top() any {
	l := (*linkedList)(s)
	return l.head.item
}

func (s *stackLinkedList) Pop() any {
	l := (*linkedList)(s)
	res := l.removeFromHead()
	if res == nil {
		return nil
	} else {
		return res.item
	}
}

func (s *stackLinkedList) Push(x any) {
	l := (*linkedList)(s)
	l.addNewNode(x)
}

func (s *stackLinkedList) IsEmpty() bool {
	l := (*linkedList)(s)
	return l.head == nil
}

func (s *stackLinkedList) IsFull() bool {
	return false
}

func (s *stackLinkedList) Size() int {
	l := (*linkedList)(s)
	return l.num
}

func (s *stackLinkedList) String() string {
	return fmt.Sprint(linkedList(*s))
}

/* funzioni e metodi per gestire liste concatenate */

// nodo di lista concatenata di item (tipo generico ... any è un alias di interface{} )
type linkedList struct {
	head *listNode
	num  int
}

type listNode struct {
	item any
	next *listNode
}

func newNode(n any) *listNode {
	return &listNode{n, nil}
}

func (l *linkedList) addNewNode(n any) {
	node := newNode(n)
	node.next = l.head
	l.head = node
	l.num++
}

func (l *linkedList) removeFromHead() *listNode {
	if l.num == 0 {
		fmt.Println("lista vuota")
		return nil
	}

	fmt.Println("testa della lista:", l.head.item)
	l.num--
	node := l.head
	l.head = l.head.next
	//fmt.Println("sto per restituire", node)
	return node
}

func (l linkedList) String() string {
	var out string
	p := l.head
	for p != nil {
		out += fmt.Sprint(p.item)
		if p.next != nil {
			out += fmt.Sprint(" - ")
		}
		p = p.next
	}
	return out
}
