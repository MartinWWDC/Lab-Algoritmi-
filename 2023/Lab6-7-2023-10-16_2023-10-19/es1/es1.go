package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr:=make([]int,0)
	scanner:=bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		el,_:=strconv.Atoi(scanner.Text())
		if(el==0){
			break
		}
	
		h:=true
		for i:=0;i<len(arr);i++{
			if el<arr[i] {
				arr = append(arr[:i+1], arr[i:]...)
				arr[i] = el
				h=false
				break
			}
		}
		if h{
			arr=append(arr, el)
		}
	}
	fmt.Println(arr)

}