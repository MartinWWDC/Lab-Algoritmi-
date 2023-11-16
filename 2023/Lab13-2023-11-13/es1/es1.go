package main

import "fmt"

func main() {

	sonsMap := make(map[string][]string)
	sonsMap["Anna"] = []string{"Bruno", "Carlo", "Daniela"}
	sonsMap["Bruno"] = []string{"Enrico", "Francesco"}
	sonsMap["Gianni"] = []string{"Henry"}
	sonsMap["Francesco"] = []string{"Irene"}

	stampaSbordinati(sonsMap,"Anna")

}

func stampaSbordinati(sonsMap map[string][]string, key string) {
	fmt.Println(sonsMap[key])
}