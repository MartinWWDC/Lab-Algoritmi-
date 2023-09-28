package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Nodi struct {
	N            int
	collegamenti []Galleria
}
type Galleria struct {
	part  *Nodi
	arr   *Nodi
	light int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	N, M, H, S := 0, 0, 0, 0
	var nod []Nodi
	for scanner.Scan() {
		if i == 0 {
			arr := strings.Split(scanner.Text(), " ")
			N, _ = strconv.Atoi(arr[0])
			M, _ = strconv.Atoi(arr[1])
			H, _ = strconv.Atoi(arr[2])
			S, _ = strconv.Atoi(arr[3])
			i++
			nod = GeneraNodi(N)

		} else {
			arr := (strings.Split(scanner.Text(), " "))
			A, _ := strconv.Atoi(arr[0])
			B, _ := strconv.Atoi(arr[1])
			L, _ := strconv.Atoi(arr[2])

			getNodo(A, nod).collegamenti = append(getNodo(A, nod).collegamenti, Galleria{getNodo(A, nod), getNodo(B, nod), L})
			getNodo(B, nod).collegamenti = append(getNodo(B, nod).collegamenti, Galleria{getNodo(B, nod), getNodo(A, nod), L})
		}
		if i == M {
			break
		}
	}
	fmt.Println("hi: ", N, H, S)
	fmt.Println(getNodo(H, nod))
	counter := FindPattern(H, S, nod)
	fmt.Println(counter)

}

func GeneraNodi(n int) (nodi []Nodi) {
	for i := 0; i <= n; i++ {
		nodi = append(nodi, Nodi{i, nil})
	}
	return
}

func getNodo(n int, nodi []Nodi) *Nodi {
	for h := range nodi {
		if nodi[h].N == n {
			return &nodi[h]
		}
	}
	return nil
}

func FindPattern(H int, S int, nod []Nodi) int {
	ptr := getNodo(H, nod)
	end := getNodo(S, nod)
	counter := 0
	for !Equals(ptr, end) {
		if counter > len(nod) {
			counter = -1
			break
		}
		fmt.Println("nope")
		ptr = checkLowLight(*ptr).arr
		counter++

	}

	return counter

}
func Equals(A *Nodi, B *Nodi) bool {
	if A.N == B.N {
		return true
	} else {
		return false
	}
}
func checkLowLight(A Nodi) Galleria {
	var lower Galleria
	for i, G := range A.collegamenti {
		if i == 0 {
			lower = G
		} else {
			if G.light < lower.light {
				lower = G
			}
		}
	}
	return lower
}
