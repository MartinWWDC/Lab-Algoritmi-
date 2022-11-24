package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

type Fold struct {
	axis   rune
	offset int
}

func main() {
	coordinates, folds := read_from_file("input2.txt")
	temp := make(map[Coordinate]bool, 0)

	for i := 0; i < len(folds); i++ {
		for j := 0; j < len(coordinates); j++ {
			if folds[i].axis == 'x' {
				if coordinates[j].x > folds[i].offset {
					coordinates[j].x = folds[i].offset - (coordinates[j].x - folds[i].offset)
				}
			} else {
				if coordinates[j].y > folds[i].offset {
					coordinates[j].y = folds[i].offset - (coordinates[j].y - folds[i].offset)
				}
			}
		}
	}

	for i := 0; i < len(coordinates); i++ {
		temp[coordinates[i]] = true
	}

	print_coordinates(temp)
	fmt.Print(len(temp))
	fmt.Print("\n\n")
}

func print_coordinates(map_coordinates map[Coordinate]bool) {
	var max_x, max_y int = 0, 0

	for value, _ := range map_coordinates {
		if value.x > max_x {
			max_x = value.x
		}

		if value.y > max_y {
			max_y = value.y
		}
	}

	for y := 0; y <= max_y; y++ {
		for x := 0; x <= max_x; x++ {
			_, exists := map_coordinates[Coordinate{x: x, y: y}]

			if exists {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func read_from_file(filename string) (coordinates []Coordinate, folds []Fold) {
	//Variabile che definisce quando leggere le coordinate dei punti e i dati sulle pieghe
	var read_coordinates bool = true
	//Variabili temporanee per la lettura dei dati
	var x, y, offset int
	var axis rune

	//Apro il file
	file, err := os.Open(filename)

	//Verifico se sono riuscito ad aprire il file
	if err == nil {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()

			//Quando trovo una linea vuota (corrispondente ad un a capo) inizio a leggere le pieghe
			if line == "" {
				read_coordinates = false
			}

			if read_coordinates {
				fmt.Sscanf(line, "%d,%d", &x, &y)
				coordinates = append(coordinates, Coordinate{x: x, y: y})

			} else {
				//Verifico che la linea non sia vuota
				if line != "" {
					split := strings.Split(line, " ")

					fmt.Sscanf(split[2], "%c=%d", &axis, &offset)
					folds = append(folds, Fold{axis: axis, offset: offset})
				}
			}
		}
	}

	return
}
