package main

import "fmt"

func main() {
	str := "+(4 8)*(4)+(6)"
	end := converti(str)
	fmt.Println(end)
	//arr := []rune{'4', '5'}
	//arr = remove(arr, len(arr)-2)
	//fmt.Println(string(arr))
}

func converti(espressione string) string {
	var pila []rune
	end := ""
	for g := range espressione {
		switch {
		case espressione[g] == '*' || espressione[g] == '+' || espressione[g] == '/' || espressione[g] == '-':
			pila = append(pila, rune(espressione[g]))
		case espressione[g] == ')':
			end += " " + string(pila[len(pila)-1])
			remove(pila, len(pila)-1)
		default:
			if espressione[g] != '(' && espressione[g] != ' ' {
				end += " " + string(espressione[g])
			}
		}
	}
	return end
}

func remove(slice []rune, s int) []rune {
	return append(slice[:s], slice[s+1:]...)
}
