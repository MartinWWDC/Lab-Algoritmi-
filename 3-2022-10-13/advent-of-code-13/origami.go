package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
https://adventofcode.com/2021/day/13
--- Day 13: Transparent Origami ---
You reach another volcanically active part of the cave. It would be nice if you could do some kind of thermal imaging so you could tell ahead of time which caves are too hot to safely enter.

Fortunately, the submarine seems to be equipped with a thermal camera! When you activate it, you are greeted with:

Congratulations on your purchase! To activate this infrared thermal imaging
camera system, please enter the code found on page 1 of the manual.
Apparently, the Elves have never used this feature. To your surprise, you manage to find the manual; as you go to open it, page 1 falls out. It's a large sheet of transparent paper! The transparent paper is marked with random dots and includes instructions on how to fold it up (your puzzle input).
*/
type coord struct {
	x int
	y int
}
type fold struct {
	ty   string
	size int
}

func main() {
	coords := []coord{}
	folds := []fold{}
	maxh := 0
	maxl := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text()[0] != 'f' {
			xy := strings.Split(s.Text(), ",")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			//fmt.Println("x:", x, ">", maxh)
			if x > maxh {
				maxh = x
			}
			//fmt.Println("y:", y, ">", maxl)
			if y > maxl {
				maxl = y
			}
			coords = append(coords, coord{x, y})
		} else {
			ax := strings.Split(s.Text(), "=")
			size := ax[0][len(ax[0])-1]
			//fmt.Println(string(size), ax[1])
			line, _ := strconv.Atoi(ax[1])
			folds = append(folds, fold{string(size), line})
		}

	}

	//fmt.Println(len(coords))
	//fmt.Println(maxl, maxh)
	//printMatrix(maxl, maxh, coords)
	fmt.Println()
	for _, fold := range folds {
		maxl, maxh, coords = foldMatrix(fold.ty, fold.size, coords)

	}
	g := printMatrix(maxl, maxh, coords)
	fmt.Println(g)
	//_, _, coords = foldMatrix("y", 7, coords)
	//foldMatrix("x", 5, coords)
}
func printMatrix(maxl int, maxh int, coords []coord) int {
	check := true
	counter := 0
	for i := 0; i <= maxl; i++ {
		for j := 0; j <= maxh; j++ {
			for _, coord := range coords {
				if coord.x == j && coord.y == i {
					fmt.Print("#")
					counter++
					check = false
					break
				}
			}
			if check {
				fmt.Print(".")

			} else {
				check = true
			}
		}
		fmt.Println()
	}
	return counter
}
func foldMatrix(ax string, size int, coords []coord) (int, int, []coord) {
	newcoords := []coord{}
	maxl, maxh := findthebiggers(coords)
	if ax == "x" {
		for _, coor := range coords {
			if coor.x < size+1 {
				newcoords = append(newcoords, coor)

			} else {
				x := (coor.x - maxh) * -1
				newcoords = append(newcoords, coord{x, coor.y})

			}
		}

	} else {
		for _, coor := range coords {
			if coor.y < size+1 {
				newcoords = append(newcoords, coor)

			} else {
				y := (coor.y - maxl) * -1
				newcoords = append(newcoords, coord{coor.x, y})

			}
		}
	}
	//maxl, maxh = findthebiggers(newcoords)
	//fmt.Println(maxl, maxh)
	if ax == "x" {
		maxh = maxh - size - 1

	} else {
		maxl = maxl - size - 1

	}
	//printMatrix(maxl, maxh, newcoords)

	return maxl, maxh, newcoords
}
func findthebiggers(coords []coord) (int, int) {
	maxh, maxl := 0, 0
	for _, coord := range coords {
		if coord.x > maxh {
			maxh = coord.x
		}
		if coord.y > maxl {
			maxl = coord.y
		}
	}
	return maxl, maxh
}
