package main

type pesceNode struct {
	name  int
	left  *pesceNode
	right *pesceNode
}
type pesceTree struct {
	root *pesceNode
}

func main() {
	var g pesceTree
	g.root = &pesceNode{0, &pesceNode{5, nil, nil}, &pesceNode{4, nil, nil}}

}
