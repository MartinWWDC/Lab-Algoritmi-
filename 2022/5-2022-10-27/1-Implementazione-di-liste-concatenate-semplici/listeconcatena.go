package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type listNode struct {
	el   int
	next *listNode
}
type linkedList struct {
	len  int
	head *listNode
}

func main() {
	//scanner()
	//fmt.Println(replace("+    5"))
	input := ""
	list := newList()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
		input = replace(input)
		switch {
		case input[0] == '+':
			i, _ := strconv.Atoi(input[1:])
			//fmt.Println(i)
			check, _ := searchList(&list, i)
			//fmt.Println(check)
			if check == false {
				addNewNode(newNode(i), &list)
			}
		case input[0] == '-':
			i, _ := strconv.Atoi(input[1:])
			//fmt.Println(i)
			check, _ := searchList(&list, i)
			//fmt.Println(i)
			if check {
				removeNewNode(&list, i)
			}
		case input[0] == '?':
			i, _ := strconv.Atoi(input[1:])
			//fmt.Println(i)
			check, _ := searchList(&list, i)
			//fmt.Println(i)
			if check {
				fmt.Println("Appartiene:", i)
			} else {
				fmt.Println("Non Appartiene")

			}
		case input[0] == 'c':
			//i, _ := strconv.Atoi(input[1:])
			//fmt.Println(i)
			//check, _ := searchList(&list, i)
			//fmt.Println(i)
			fmt.Println("Len:", list.len)
		case input[0] == 'd':
			//i, _ := strconv.Atoi(input[1:])
			//fmt.Println(i)
			//check, _ := searchList(&list, i)
			//fmt.Println(i)
			list.head = nil
		case input[0] == 'f':
			break
		case input[0] == 'o':
			fmt.Println(reverseprintList(&list))

		}

	}

	//printList(&list)

}

func replace(str string) string {
	l := ""
	for _, i := range str {
		if i != ' ' {
			l += string(i)
		}
	}
	return l
}
func newNode(n int) listNode {
	return listNode{n, nil}
}
func printNode(node *listNode) {
	fmt.Println("val:", node.el)
	fmt.Println("next Node:", node.next)
}
func newList() linkedList {
	list := linkedList{}
	list.len = 0
	return list

}
func addNewNode(node listNode, list *linkedList) {
	if list.len == 0 {
		list.head = &node
		list.len++

	} else {
		max := list.len
		ptr := list.head
		for i := 0; i < max; i++ {
			if ptr.next == nil {
				ptr.next = &node
				list.len++
				return
			}
			ptr = ptr.next
		}

	}

}
func printList(list *linkedList) {
	fmt.Println("Print List")
	ptr := list.head
	for i := 0; i < list.len; i++ {
		printNode(ptr)
		ptr = ptr.next
	}
	fmt.Println("End Print List")
}
func searchList(list *linkedList, n int) (bool, listNode) {
	max := list.len
	ptr := list.head
	for i := 0; i < max; i++ {
		if ptr.el == n {
			//fmt.Println("Found!")
			//printNode(ptr)
			return true, *ptr
		}
		ptr = ptr.next

	}

	//fmt.Println("Not Found!")

	return false, listNode{}
}

func removeNewNode(list *linkedList, n int) {
	ptr := list.head
	var prev *listNode
	//fmt.Println("s", *ptr.next)
	for i := 0; i < list.len; i++ {
		//fmt.Println("s", ptr.el)
		if ptr.el == n {
			//fmt.Println("removing")
			list.len--
			switch {
			case i == 0:
				//fmt.Println("head")
				list.head = ptr.next
				//fmt.Println(list.head)

			case i == list.len:
				prev.next = nil
			default:
				prev.next = ptr.next

			}

			return
		}
		prev = ptr
		ptr = ptr.next
	}
}
func reverseprintList(list *linkedList) []int {
	arr := []int{}
	ptr := list.head
	for i := 0; i < list.len; i++ {
		//printNode(ptr)
		arr = append(arr, ptr.el)
		ptr = ptr.next
	}
	arr2 := []int{}
	for i := list.len - 1; i >= 0; i-- {
		arr2 = append(arr2, arr[i])
	}
	return arr2
}
