package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
 * Data da standard input una serie di interi positivi terminata da 0, stampare ’+’ ogni volta che il nuovo valore
è maggiore o uguale al precedente e ’-’ altrimenti.
*/

func main() {
	scanner:=bufio.NewScanner(os.Stdin)
	arr:=[]int{}
	signs:=[]string{}
	i:=0
	for scanner.Scan(){
		n,_:=strconv.Atoi(scanner.Text())
		if n==0{
			break
		}
		if i!=0{
			if arr[i-1]<n{
				signs = append(signs, "+")
			}else{
				signs = append(signs, "-")

			}
		}
		
		arr = append(arr, n)
		i++
	}
	fmt.Println(signs)
}