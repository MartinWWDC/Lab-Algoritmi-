package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fatherMap:= make(map[string]string)

	scanner:=bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		txt:=strings.Split(scanner.Text(),")")
		fatherMap[txt[1]]= txt[0]
	
	}
	n:=countOrbits(fatherMap)
	fmt.Println(n)


}
func countOrbits(orbits map[string]string) int {
	n := 0
	for k, _ := range orbits {
		n += len(chain(k, orbits))
	}
	return n
}
func chain(k string, m map[string]string) []string {
	var chain []string
	for v, ok := m[k]; ok; v, ok = m[v] {
		chain = append(chain, v)
	}
	return chain
}