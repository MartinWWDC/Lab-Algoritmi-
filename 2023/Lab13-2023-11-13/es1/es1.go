package main

import "fmt"

func main() {

	sonsMap := make(map[string][]string)
	sonsMap["Anna"] = []string{"Bruno", "Carlo", "Daniela"}
	sonsMap["Bruno"] = []string{"Enrico", "Francesco"}
	sonsMap["Gianni"] = []string{"Henry"}
	sonsMap["Francesco"] = []string{"Irene"}

	stampaSbordinati(sonsMap,"Anna")
	c:=quantiSenzaSubordinati(sonsMap)
	fmt.Println(c)

}

func stampaSbordinati(sonsMap map[string][]string, key string) {
	fmt.Println(sonsMap[key])
}

func quantiSenzaSubordinati(sonsMap map[string][]string)int{
	counter:=0
	for _,arr:= range sonsMap{
		for _,el:=range arr{
			_,check:=sonsMap[el]
			if !check{
				//fmt.Println(el)
				counter++
			}
		}				
	}
	return counter
}