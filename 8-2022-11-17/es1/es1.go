package main

import (
	"fmt"
)

func main() {

	mappa := make(map[string]string)
	mappa["bruno"] = "anna"
	mappa["carlo"] = "anna"
	mappa["daniela"] = "anna"
	mappa["anna"] = "giacomo"
	mappa["giacomo"] = ""

	//aggiungiElemento(mappa, "s:i")
	fmt.Println(mappa)
	//stampaSubordinati(mappa, "anna")
	//quantiSenzaSubordinati(mappa)
	//supervisore(mappa, "daniela")
	//stampaImpiegatiSopra(mappa, "bruno")
	stampaImpiegatiConSupervisore(mappa)
}

func stampaSubordinati(mappa map[string]string, s string) {
	for i, j := range mappa {
		if j == s {
			fmt.Println(i)
		}
	}
}

func quantiSenzaSubordinati(mappa map[string]string) {
	for i, j := range mappa {
		if j == "" {
			fmt.Println(i)
		}
	}
}
func supervisore(mappa map[string]string, s string) {
	if _, j := mappa[s]; j {
		fmt.Println(mappa[s])
	}

}
func stampaImpiegatiSopra(mappa map[string]string, s string) {
	if mappa[s] != "" {
		fmt.Println(mappa[s])
		stampaImpiegatiSopra(mappa, mappa[s])
	}
}

func stampaImpiegatiConSupervisore(mappa map[string]string) {
	for i, j := range mappa {
		if j != "" {
			fmt.Println("dimpendente: ", i, " con supervisore:", j)
		}
	}
}
