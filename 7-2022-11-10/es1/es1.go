package main

import (
	"fmt"
)

const MAX_SLICE = 2 //Max Size of slice

type bitreeNode struct {
	left  *bitreeNode
	right *bitreeNode
	val   int
}
type bitree struct {
	root *bitreeNode
}

func main() {
	
	arr := []int{5, 8, 2}

	root := addElement(arr, 0)

	stampaAlberoASommario(root, 0)
	stampaAlbero(root)

}
func newNode(val int) *bitreeNode {
	return &bitreeNode{nil, nil, val}
}
func printNode(node *bitreeNode) {
	fmt.Println("val:", node.val)
	fmt.Println("left:", node.left)
	fmt.Println("right:", node.right)
}
func stampaAlbero(node *bitreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.val, "")
	if node.left == nil && node.right == nil {
		return
	}
	fmt.Print(" [")
	if node.left != nil {
		stampaAlbero(node.left)
	} else {
		fmt.Print("-")
	}
	fmt.Print(", ")
	if node.right != nil {
		stampaAlbero(node.right)
	} else {
		fmt.Print("-")
	}
	fmt.Print("]")
}

func stampaAlberoASommario(node *bitreeNode, spaces int) {
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
	fmt.Print("*")
	fmt.Println(node.val)
	if node.right != nil || node.left != nil {

		if node.left != nil {
			stampaAlberoASommario(node.left, spaces+1)

		} else {
			for i := 0; i < spaces; i++ {
				fmt.Print(" ")
			}
			fmt.Println(" *")
		}

		if node.right != nil {
			stampaAlberoASommario(node.right, spaces+1)
		} else {
			for i := 0; i < spaces; i++ {
				fmt.Print(" ")
			}
			fmt.Println(" *")
		}
	}

}
func addElement(arr []int, i int) *bitreeNode {
	var node bitreeNode
	if i >= len(arr) {
		return nil
	}
	node.val = arr[i]
	node.left = addElement(arr, 2*i+1)
	//fmt.Println(node.left.val)
	node.right = addElement(arr, 2*i+2)
	//fmt.Println(node.right.val)

	return &node

}
