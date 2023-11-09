package main

type listNode struct {
	item int
	next *listNode
}

type linkedListWithTail struct {
	head, tail *listNode
}

func newNode(val int) *listNode {
	return &listNode{val, nil}
}

//es3.1

func addNewNodeAtEnd(l *linkedListWithTail, val int) {
	if l.tail == nil {
		l.tail = newNode(val)
		l.head = l.tail
	} else {
		temp := l.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode(val)

	}
}

func main() {

}
