package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type vertice struct {
	valore item
	chiave string
	adj    []*vertice // insieme dei vertici adiacenti
}

type item struct {
	età   string
	hobby []string
}
type graph map[string]*vertice

func graphNew(n int) *graph {
	gr := make(graph)

	return &gr
}

func createNewVertex(it item,key string, adj []*vertice)*vertice{
	return &vertice{it,key,adj}

}

func readUser(gr *graph){
	scanner:=bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		txt:=strings.Split(scanner.Text(),";")
		following:=strings.Split(txt[3],",")
		adj:=make([]*vertice,0)
		for _,follower:= range following{
			if follower!=""{
				agg,ok:=(*gr)[follower]
				if !ok{
					(*gr)[follower]=createNewVertex(item{},follower,nil)
					agg = (*gr)[follower]
				}
				adj=append(adj,agg)
			}
		}
		it:=item{txt[1],strings.Split(txt[2], ",")}
		if txt[0] != "" {
			(*gr)[txt[0]] = createNewVertex(it, txt[0], adj)
			fmt.Println(*gr)
		}	
	}
}

func printHobbyFollowing(key string,gr graph){
	user:=gr[key]
	fmt.Println(user.valore.hobby)
	for _, v := range user.adj {
		fmt.Println(v.valore.hobby)
	}
}


func printHobbyFollower(key string,gr graph){
	user:=gr[key]
	fmt.Println(user.valore.hobby)
	for _, v := range user.adj {
		for _,i := range user.adj{
			if i==user{
				fmt.Println(v.valore.hobby)
			}

		}
	}
}

func main() {
	graph := graphNew(0)
	readUser(graph)
	for k,user:=range *graph{
		fmt.Println(k)
		fmt.Println(user)
		fmt.Println("----")

	}
	fmt.Println(graph)

}