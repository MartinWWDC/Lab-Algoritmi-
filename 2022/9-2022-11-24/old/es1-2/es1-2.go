package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type item struct {
	nome    string
	cognome string
	età     int
}

type vertice struct {
	valore item
	chiave string
}

type grafo struct {
	vertici   map[string]*vertice
	adiacenti map[vertice][]*vertice
}

func main() {
	grafo := LeggiDaInput()
	printVertici(grafo)
}

func graphNew(n int) *grafo {
	vertici := make(map[string]*vertice)
	for i := 0; i < n; i++ {
		vertici[strconv.Itoa(i)] = &vertice{item{"", "", 0}, strconv.Itoa(i)}
	}
	return &grafo{vertici, nil}
}

func LeggiDaInput() *grafo {
	vertici := make(map[string]*vertice)
	adiacenti:=make(map[vertice][]*vertice)

	n, _ := strconv.Atoi(os.Args[1])
	scan := bufio.NewScanner(os.Stdin)
	i := 0
	for i <= n && scan.Scan() {
		arr := strings.Split(scan.Text(), ":")
		h, _ := strconv.Atoi(arr[2])
		item := item{arr[0], arr[1], h}
		vertici[genID(item)] = &vertice{item, genID(item)}
		i++
	}
	for scan.Scan() {
		arr := strings.Split(scan.Text(), " ")
		if _, ok := vertici[arr[0]]; ok {
			if _, ok := vertici[arr[1]]; ok {



	}
}
	return &grafo{vertici, nil}
}

func genID(i item) string {
	return string(i.nome[0]) + "" + string(i.cognome[0])
}
func printVertici(graph *grafo) {
	fmt.Println(&graph.vertici)

	for g, v := range graph.vertici {
		fmt.Println("id ", g)
		printVertice(*v)
	}
}
func printVertice(v vertice) {
	printItem(v.valore)
	fmt.Println("key:", v.chiave)
}

func printItem(i item) {
	fmt.Println("Item:")
	fmt.Println("Nome:", i.nome)
	fmt.Println("cognome:", i.cognome)
	fmt.Println("età:", i.età)

}
