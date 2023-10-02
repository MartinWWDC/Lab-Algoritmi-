package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner:=bufio.NewScanner(os.Stdin)
	check:=true
	g:=0
 	for scanner.Scan(){
		n,_ :=strconv.Atoi(scanner.Text())
		if(n==-1){
			break
		}
		if(n>100 && check){
			check=false
			g=n
		}
	}
	if(!check){
		fmt.Println(g)
	}else{
		fmt.Println("nessun numero maggiore di 100")
	}
}