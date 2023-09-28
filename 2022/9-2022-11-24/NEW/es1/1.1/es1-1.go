package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type grafo struct {
	n         int
	adiacenti []*linkedList
}
type linkedList struct {
	value       int
	connections []*linkedList
}

func main() {
	//gr := nuovoGrafo(5)
	//printGraph(*gr)
	r := LeggiDaInput()
	printGraph(*r)
}

func nuovoGrafo(n int) *grafo {
	var arrLinked []*linkedList
	for i := 0; i < n; i++ {
		arrLinked = append(arrLinked, &linkedList{i, nil})
	}
	return &grafo{n, arrLinked}
}

func printGraph(graph grafo) {
	fmt.Println("fin")
	//fmt.Println(graph)
	fmt.Println("N nodi:", graph.n)
	for i := 0; i < graph.n; i++ {
		printNodo(*graph.adiacenti[i])
	}
}
func printNodo(node linkedList) {
	fmt.Println("Nodo:", node.value)
	fmt.Println("Connections: ", node.connections)
}

func LeggiDaInput() *grafo {
	n, _ := strconv.Atoi(os.Args[1])
	graph := nuovoGrafo(n)
	scan := bufio.NewScanner(os.Stdin)
	i := 0
	for i <= n && scan.Scan() {
		txt := scan.Text()
		arr := strings.Split(txt, " ")
		fmt.Println("arr ", i, ":", arr)
		var arrG []*linkedList
		for _, h := range arr {
			n, _ := strconv.Atoi(h)
			arrG = append(arrG, findElement(graph, n))
		}
		graph.adiacenti[i].connections = arrG
		i++
	}
	return graph
}
func findElement(graph *grafo, n int) *linkedList {
	for _, g := range graph.adiacenti {
		if g.value == n {
			//printNode(g)
			return g
		}
	}
	return nil

}
func checkConnections(graph *grafo, x int, y int) bool {
	g := findElement(graph, x)
	for _, h := range g.connections {
		if h.value == y {
			return true
		}
	}

	return false

}
