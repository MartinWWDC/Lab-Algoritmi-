package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//partOne()
}

func partOne(){
	matr:=genMatr()
	//toStringMatr(matr)

	arr:=getLowPoints(matr)
	riskL:=getRiskLevel(arr)
	fmt.Println(riskL)
}

func genMatr() [][]int {
	matr:=make([][]int,0)
	scanner:=bufio.NewScanner(os.Stdin)
	i:=0
	for scanner.Scan(){
		cfrArr:=strings.Split(scanner.Text(),"")
		intArr:=make([]int,0)
		for _,cfr:=range cfrArr{
			n,_:=strconv.Atoi(cfr)
			intArr = append(intArr, n)
		}
		matr=append(matr, intArr)
		i++
	}
	return matr
}

func toStringMatr(matr [][]int){
	for _,el:= range matr{
		fmt.Println(el)
	}
}

func getLowPoints(matr [][]int)[]int{
	lowP:=make([]int,0)
	for i,_:=range matr{
		for j,_:= range matr[i]{
			check:=true
			if (i-1>=0){
				if !(matr[i][j] < matr[i-1][j]){
					check=false
				}
			}
			if (i+1<len(matr)){
				if !(matr[i][j] < matr[i+1][j]){
					check=false
				}
			}
			if (j-1>=0){
				if !(matr[i][j] < matr[i][j-1]){
					check=false
				}
			}
			if (j+1<len(matr[i])){
				if !(matr[i][j] < matr[i][j+1]){
					check=false
				}
			}
			if check{
				lowP=append(lowP,matr[i][j])
			}

		}
	}
	return lowP
}

func getRiskLevel(arr []int)int{
	n:=0
	for _,i := range arr{
		n+=i+1
	}
	return n
}