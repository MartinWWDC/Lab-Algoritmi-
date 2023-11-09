package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type listNode struct {
	item int
	next *listNode
}

type linkedList struct {
	head *listNode
	counter int
}

func main() {

	scanner:=bufio.NewScanner(os.Stdin)
	li:=newLinkedList()
	for scanner.Scan(){
		input:=strings.Split(scanner.Text()," ")
		switch input[0]{
			case "+":
				val,_:= strconv.Atoi(input[1])
				addNewNode(&li,val)
				li.counter++
			case "-":
				val,_:= strconv.Atoi(input[1])
				removeItem(&li,val)
			case "?":
				val,_:= strconv.Atoi(input[1])
				fmt.Println(isIn(li,val))
			case "c":
				fmt.Println(li.counter)
			case "p":
				printList(li)
			case "o":
				
			case "d":
				li=newLinkedList()
			case "f":
				break
		

		}
	}

}

func newLinkedList()linkedList{
	return linkedList{nil,0}

}

func newNode(val int) *listNode {
	return &listNode{val, nil}
}

func addNewNode(li *linkedList, val int) {
	node := newNode(val)
	node.next = li.head
	li.head = node

}

func printList(li linkedList) {
	fmt.Println("***")
	el := li.head

	for el != nil {
		fmt.Println(el.item)
		el = el.next

	}
	fmt.Println("***")

}

func searchList(li linkedList,val int)*listNode{
	el := li.head
	
	if el.item==val{
		return el
	}
	
	for el.item != val {
		el = el.next
		return el

	}

	return nil
}

func isIn(li linkedList, val int) bool {
    el := li.head

    for el != nil {
        if el.item == val {
            return true
        }
        el = el.next
    }

    return false
}

func removeItem(li *linkedList,val int){
	el := li.head
	for el.next !=nil {

		fmt.Println(el.item)

		if el.next.item==val{

			el.next=el.next.next
			break
		}
		el=el.next
	
	}
}


