package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	matr:=genMatr()
	//arr:=getApprox(0,1,matr)
	//fmt.Println(arr)
	toStringMatr(matr)
	steps,_:=strconv.Atoi(os.Args[1])
	for i:=0; i<=steps;i++{
		matr=loopStep(matr)
		toStringMatr(matr)

	}
	//matr1:=loopStep(matr)
	toStringMatr(matr)

}

func loopStep(matr [][]int)[][]int{
	for i:=range matr{
		for j:=range matr[i]{
			updateLights(i,j,&matr)

		} 
	}
	return matr
}

func updateLights(i int ,j int,matr *[][]int){
	(*matr)[i][j]++
	if (*matr)[i][j]>9{
		(*matr)[i][j]=0
		arr:=getApprox(i,j,(*matr))
		for _,el:=range arr{
			updateLights(el[0],el[1],matr)
		}

	}

}

func getApprox(i int,j int,matr [][]int)[][]int{
	mat:=make([][]int,0)
	if (i-1>=0){
		mat=append(mat,[]int{i-1,j})
		
	}
	if (i+1<len(matr)){
		mat=append(mat,[]int{i+1,j})
	}
	if (j-1>=0){
		mat=append(mat,[]int{i,j-1})
		
	}
	if (j+1<len(matr[i])){
		mat=append(mat,[]int{i,j+1})
		
	}
	return mat
}

func genMatr()[][]int {
	matr := make([][]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	for scanner.Scan() {
		cfrArr := strings.Split(scanner.Text(), "")
		intArr := make([]int, 0)
		for _, cfr := range cfrArr {
			n, _ := strconv.Atoi(cfr)
			intArr = append(intArr, n)
		}
		matr = append(matr, intArr)
		i++
	}
	return matr
}

func toStringMatr(matr [][]int) {
	for _, el := range matr {
		fmt.Println(el)
	}
	fmt.Println("")

}