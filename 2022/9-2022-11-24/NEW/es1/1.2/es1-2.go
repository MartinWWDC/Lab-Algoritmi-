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
	hobbys  *[]string
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
	gr := LeggiInput()
	prinGraph(*gr)
	getHobby(*gr, "ugo")
}
func graphNew(n int) *grafo {
	var gr grafo
	verti := make(map[string]*vertice)
	adia := make(map[vertice][]*vertice)
	for i := 0; i < n; i++ {
		verti[strconv.Itoa(i)] = &vertice{item{"", "", 0, nil}, strconv.Itoa(i)}
		adia[*verti[strconv.Itoa(i)]] = nil
	}
	gr.vertici = verti
	gr.adiacenti = adia
	return &gr
}
func LeggiInput() *grafo {
	Scanner := bufio.NewScanner(os.Stdin)
	vert := make(map[string]*vertice)
	adiacenti := make(map[vertice][]*vertice)
	var vertici []vertice
	c := true
	for Scanner.Scan() {
		if Scanner.Text() == "-" {
			c = false
			for i := range vertici {
				//fmt.Println(vertici[i].chiave)
				vert[vertici[i].chiave] = &vertici[i]
			}
			continue
		}
		if c {
			arr := strings.Split(Scanner.Text(), ":")
			//fmt.Println(arr)
			et, _ := strconv.Atoi(arr[2])
			ar := strings.Split(arr[3], ",")
			in := item{arr[0], arr[1], et, &ar}
			vertici = append(vertici, vertice{in, arr[4]})
		} else {
			arr := strings.Split(Scanner.Text(), ":")
			//fmt.Println(arr[0])
			//fmt.Println("S:", vert[arr[0]])
			adiacenti[*vert[arr[0]]] = append(adiacenti[*vert[arr[0]]], vert[arr[1]])
			//adiacenti[*vert[arr[1]]] = append(adiacenti[*vert[arr[1]]], vert[arr[0]])
		}
	}

	//fmt.Println(vertici)
	//fmt.Println(vert)
	//fmt.Println(adiacenti)
	return &grafo{vert, adiacenti}
}

func prinGraph(graph grafo) {
	//ho un idea ma ho sonno
}
func getHobby(graph grafo, A string) (hobbys []string) {
	for _, k := range graph.vertici {
		if k.valore.nome == A {
			hobbys = *k.valore.hobbys
			break
		}

	}
	return
}
func getHobbyOf(graph grafo, A string) (hobbys []string) {
	var main *vertice
	for _, k := range graph.vertici {
		if k.valore.nome == A {
			main = k
			break
		}

	}
	for i, adia := range graph.adiacenti {
		if i == *main {
			for _, j := range adia {
				fmt.Println(j.valore.hobbys)
			}
		}
	}
	return
}
