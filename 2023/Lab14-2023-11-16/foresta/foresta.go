package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type oggetto struct {
	nome string
	val  int	// non rilevante se l'indizio e' un'operazione
	dx   string
	op   rune // ' ' se numero
	sx   string
	tipo string // "num" se l'indizio è un numero, "op" se l'indizio è una operazione
}


func leggiInput() map[string]*oggetto {
	mapp:=make(map[string]*oggetto)
	scan:=bufio.NewScanner(os.Stdin)
	for scan.Scan(){
		arr:=strings.Split(scan.Text(),":")
		nome:=arr[0]
		sArr:=strings.Split(arr[1]," ")	
	
		if len(sArr)==2{
			n,_:=strconv.Atoi(sArr[0])
			mapp[nome]=&oggetto{nome,n,"",rune(0),"","num"}
		}else{
			mapp[nome]=&oggetto{nome,-1,sArr[1],[]rune(sArr[2])[0],sArr[3],"op"}
		}	
		fmt.Println(mapp[nome])

	}
	return mapp

}


type foresta struct{
	nodi map[string]*oggetto
}


func costruisciForesta(mappa map[string]*oggetto) foresta {
	return foresta{mappa}
}

func stampaAlbero(f foresta, name string) {

}

func sx(f foresta,n string) (string, bool) {
	oggetto, ok := f.nodi[n]
	if !ok {
	  return "", false
	}
	if oggetto.tipo == "num" {
	  return "", false
	}
	return oggetto.dx, true
  }

func dx(f foresta, n string) (string, bool) {
	oggetto, ok := f.nodi[n]
	if !ok {
	  return "", false
	}
	return oggetto.sx, true
}

func up(f foresta, n string) (string, bool) {
	oggetto, ok := f.nodi[n]
	if !ok {
	  return "", false
	}
	if oggetto.tipo == "num" {
	  return "", false
	}
	return string(oggetto.op), true
}

func calcolaPrezzo(f foresta, n string) int {
	return 0

}

// non modificare
func main() {
	mappa := leggiInput()
	f := costruisciForesta(mappa)
	stampaAlbero(f, os.Args[1])
}
