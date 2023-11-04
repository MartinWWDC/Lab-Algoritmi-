package main

import "fmt"

//2
type node struct{
	val int
	next *node
}

//3
type circNode struct{
	val int
	next *circNode
	prev *circNode
}

func main() {
	// Creazione di nodi per i dati
	node0 := &circNode{val: 0, next: nil, prev: nil}
	node1 := &circNode{val: 1, next: nil, prev: nil}
	node7 := &circNode{val: 7, next: nil, prev: nil}
	node3 := &circNode{val: 3, next: nil, prev: nil}
	nodeMinus2 := &circNode{val: -2, next: nil, prev: nil}

	// Collegamento dei nodi per creare una lista concatenata circolare.
	node0.next = node1
	node1.next = node7
	node7.next = node3
	node3.next = nodeMinus2
	nodeMinus2.next = node0

	node0.prev = nodeMinus2
	node1.prev = node0
	node7.prev = node1
	node3.prev = node7
	nodeMinus2.prev = node3

	printFromZero(node1)

}

func f(p *node, k int) int {
	var a int
	if p == nil {
		return 0
	}
	a = 1 + f(p.next, k)
	if a == k {
		fmt.Println(p.val)
	}
	return a
}



func printFromZero(head *circNode) {
	arr:=[]int{}
    if head == nil {
        return
    }

    current := head
    for {
		arr=append(arr,current.val)
        current = current.next
        if current == head {
            break // Hai completato un ciclo completo sulla lista.
        }
    }
	for i,el := range arr{
		if el==0{
			arr=append(arr[i:],arr[:i]...)
		}
	}
	fmt.Println(arr)
}
