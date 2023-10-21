package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var buffer string
	fmt.Scanln(&buffer)
	var fishs []int
	var splitBuffer = strings.Split(buffer, ",")
	for _, numString := range splitBuffer {
		initialTimer, _ := strconv.Atoi(numString)
		fishs=append(fishs, initialTimer)

	}
	lenR:=lantern(fishs,80)
	fmt.Println(lenR)

}

func lantern(fishs []int,day int) int{
	if day==0{
		return len(fishs)
	}
	newFishs:=[]int{}
	for _,timer := range fishs{
		if timer==0{
			newFishs = append(newFishs, 6,8)
		}else{
			newFishs = append(newFishs, timer-1)
		}
		
	}
	return lantern(newFishs,day-1)

}