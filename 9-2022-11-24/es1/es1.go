package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type graph struct {
	n         int // numero di vertici
	adiacenti []*linkedList
}

type linkedList struct {
	value       int
	connections []*linkedList
}

func main() {
	h := LeggiDaInput()
	printGraph(h)
	fmt.Println(checkConnections(h, 2, 1))
}

func nuovoGrafo(n int) *graph {
	var connesctions []*linkedList
	for i := 0; i < n; i++ {
		connesctions = append(connesctions, &linkedList{i, nil})
	}
	//fmt.Println("conn: ", connesctions)
	return &graph{n, connesctions}
}

func LeggiDaInput() *graph {
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
func findElement(graph *graph, n int) *linkedList {
	for _, g := range graph.adiacenti {
		if g.value == n {
			//printNode(g)
			return g
		}
	}
	return nil

}
func printNode(nodo *linkedList) {
	fmt.Println("Valore:", nodo.value)
	fmt.Print("conneionei: ")
	for _, g := range nodo.connections {
		fmt.Print(g.value, " ")
	}
	fmt.Println("")
}

func printGraph(graph *graph) {
	fmt.Println("grandezza:", graph.n)
	for _, g := range graph.adiacenti {
		printNode(g)
	}
}
func checkConnections(graph *graph, x int, y int) bool {
	g := findElement(graph, x)
	for _, h := range g.connections {
		if h.value == y {
			return true
		}
	}

	return false

}
