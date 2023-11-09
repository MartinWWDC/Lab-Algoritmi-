package main

import (
	"fmt"
	"math/rand"
)

type bitreeNode struct {
	left * bitreeNode
	right * bitreeNode
	val int
}
type bitree struct {
	root * bitreeNode
}

func inorder ( node * bitreeNode ){
	if node == nil{
		return 
	}
	inorder(node.left)
	fmt.Println(node.val)
	inorder(node.right)


}
func preorder ( node * bitreeNode ){
	if node == nil{
		return 
	}
	fmt.Println(node.val)
	inorder(node.left)
	inorder(node.right)

}
func postorde ( node * bitreeNode ){
	if node == nil{
		return 
	}
	inorder(node.left)
	inorder(node.right)
	fmt.Println(node.val)

}

/*
1. genera a caso una sequenza di interi (di lunghezza massima fissata con una opportuna
macro) e la memorizza in una slice;
2. costruisce un albero binario a partire dalla slice (come descritto in seguito);
3. stampa l’albero nella reppresentazione a sommario (come descritta in seguito);
4. stampa i nodi dell’albero negli ordini determinati rispettivamente dalle visite in preordine,
inordine e postordine.

*/

/**
 * punto 1
 */

func generaSlice(l int)[]int{
	nums:=[]int{}
	for i := 0;i < l;i++{
		nums = append(nums, rand.Intn(100))
	}
	return nums

}

func newNode(i int)*bitreeNode{
	return &bitreeNode{nil,nil,i}
}
/**
 * es 1.2
 */
func stampaAlberoASommario ( node * bitreeNode , spaces int) {
	for i := 0; i < spaces ; i ++ {
	fmt.Print(" ")
	}
	fmt.Print("*")
	fmt.Println( node.val )
	if node.left != nil  {
	stampaAlberoASommario(node.left , spaces+1)
	}
	if node.right != nil{
		stampaAlberoASommario(node.right , spaces+1)
	}
}
	
func main() {
	//nums := generaSlice(8)
	//fmt.Println(nums)
	t := & bitree {nil}
	t. root = & bitreeNode {nil, nil, 50}
	t. root . left = newNode (20)
	t. root . right = newNode (80)
	t. root . left . left = newNode (10)
	t. root . left . left . right = newNode (15)
	//inorder(t.root)
	a :=  []int{69, 89, 28, 39, 66, 44, 12, 2, 71}
	t.root=arr2tree(a,0)
	stampaAlberoASommario(t.root,0)
}
func stampaAlbero ( node * bitreeNode ) {
	if node == nil {
	return
	}
	fmt .Print( node .val , "")
	if node . left == nil && node . right == nil {
	return
	}
	fmt .Print(" [")
	if node . left != nil {
	stampaAlbero ( node . left )
	} else {
	fmt .Print("-")
	}
	fmt .Print(", ")
	if node . right != nil {
	stampaAlbero ( node . right )
	} else {
	fmt .Print("-")
	}
	fmt .Print("]")
	}


/**
 es 1.4
*/ 
func arr2tree (a []int, i int) ( root * bitreeNode ) {
	if i >= len(a) {
	return nil
	}
	root = &bitreeNode{nil,nil,a[i]}
	root.left = arr2tree(a,2*i+1)
	root.right = arr2tree(a,2*i+2)
	return root
	}
	