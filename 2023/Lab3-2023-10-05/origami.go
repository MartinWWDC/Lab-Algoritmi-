package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type dot struct { // dot or axis
	x, y int
}
type inst struct{
	axis string
	value int
}

func main() {
	scanner:= bufio.NewScanner(os.Stdin)
	dots:=make([]dot,0,1024)
	instr:=[]inst{}
	check:=true
	for scanner.Scan(){
		text:=scanner.Text()
		if text==""{
			check=false
			continue
		}
		if check{
			args := strings.Split(text, ",")
			x, _ := strconv.Atoi(args[0])
			y, _ := strconv.Atoi(args[1])
			dots = append(dots, dot{x, y})
			
		}else{
			args := strings.Split(text, "=")
			ch:=args[0][len(args[0])-1:]
			out,_:=strconv.Atoi(args[1])
			instr = append(instr, inst{ch,out})
			

		}
		
	}
	fmt.Println(dots)
	fmt.Println(instr)
	dt:=fold(dots,instr[1])
	fmt.Println(dt)
	printMatr(dots)


}

func printMatr(dots []dot){
	maxX:=0
	maxY:=0
	for i,d:=range dots{
		if i==0{
			maxX=d.x

			maxY=d.y
		}else{
			if maxX<dots[i].x{
				maxX=dots[i].x
			}
			if maxY<dots[i].y{
				maxY=dots[i].y
			}		
		}
	}
	matr:=make([][]string,maxX,maxY)
	for x:=0;x<maxX;x++{
		line:=[]string{}

		for y:=0;y<maxY;y++{
			c:=true
			for _,d:=range dots{
				if d.x==x && d.y==y{
					c=false
				}

			}
			if c{
				line=append(line,"-")

			}else{
				line=append(line,"#")

			}
			
			
		}
		matr=append(matr, line)
		fmt.Println(line)
	}
	
}

func fold(dots []dot,intr inst)[]dot{
	for i:=range dots{
		if intr.axis=="x" && dots[i].x>intr.value{
			fmt.Println("tac")
			dots[i].x+=2*(intr.value-dots[i].x)

		}
		if intr.axis=="y" && dots[i].y>intr.value{
			fmt.Println("tac")
			dots[i].y+=2*(intr.value-dots[i].y)

		}
	}
	return dots
	
}