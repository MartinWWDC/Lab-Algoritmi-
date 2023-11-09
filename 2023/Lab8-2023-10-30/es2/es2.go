package main

import "fmt"

type circNode struct {
	val  int
	next *circNode
	prev *circNode
}

func f(p *circNode, k int) int {
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

func f2(head *circNode,k int)int {
	a:=-1
	c:=true
    if head == nil {
        return 1
    }

    current := head
    for {
        current = current.next
		if !c{
			a++
		}
		if current.val==0 && c{
			c=false
			head=current
		}else if current == head {
			break
		}
		if a==k{
			fmt.Println(current.val)
		}
		
    }

	
	return a
}
func f3(head *circNode,k int)int {
	a:=0
    if head == nil {
        return 1
    }

    current := head
    for {
        current = current.prev
		
		if a==k{
			fmt.Println(current.val)
			break
		}
		a++
		
    }

//  T = 1 + n(1) = O(n)
	

	return k
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

	//printFromZero(node1)
	f3(node0,1)

}