package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	arr := []int{}
	i := 0
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		arr = append(arr, n)

		if i != 0 {
			if n >= arr[i-1] {
				fmt.Println("+")
			} else {
				fmt.Println("-")
			}
		}
		i++

	}

}
