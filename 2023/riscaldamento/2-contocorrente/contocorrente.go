package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	saldo,_ := strconv.Atoi(os.Args[1])
	scanner:=bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		spesa,_:=strconv.Atoi(scanner.Text())
		saldo-=spesa
		fmt.Println(saldo)
		if(saldo<=0){
			break
		}
	}
}