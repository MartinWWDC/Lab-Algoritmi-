/* Dipendenti */

/* VERSIONE 1 */
type nodoAlbero struct {
	nome             string
	up               *NodoAlbero
	primoFiglio      *NodoAlbero
	prossimoFratello *NodoAlbero
}

/* VERSIONE 2 */
type Dipendente struct {
	nome        string
	supervisore *Dipendente
	subordinati []*Dipendente
}

/* VERSIONE 3 */
type azienda struct {
	supervisore map[string]string
	subordinati map[string][]string
}

/*
Nella versione 1 i nodi della foresta sono rappresentati esplicitamente in memoria
mediante il tipo "nodoForesta", definito come struttura con puntatori
- al padre ("up") al padre
- al primo figlio
- al prossimo fratello

I nomi scelti per il tipo e per i campi della struttura usano la terminologia degli alberi
e non fanno riferimento al problema dei dipendenti.

Usando la terminologia del problema si sarebbero potuti usare questi nomi:
nodoAlbero -> dipendente
up -> supervisore
primoFiglio -> primoSubordinato
prossimoFratello -> prossimoDipendente (ma questo nome non è altrettando chiaro di "prossimoFratello")

Non sono definiti tipi né per gli alberi né per la foresta.
Si possono aggiungere le definizioni di tipi opportuni,
oppure usare direttamente i nodi radice, da memorizzare opportunamente.

-----------------------

Anche nella versione 2 i nodi della foresta sono rappresentati esplicitamente in memoria;
questa volta però i nomi scelti per il tipo e per i campi della struttura fanno riferimento al problema dei dipendenti
e non alla terminologia tecnica degli alberi.

I nodi sono rappresentati dal tipo "Dipendente", definito come struttura con
- un puntatore ("supervisore") al padre
- un puntatore ("subordinati") alla slice dei figli (subordinati)

A differenza che nella versione 1, si evita di usare il doppio puntatore al primo figlio e al prossimo fratello e si sfrutta la possibilità di usare slice (dunque array di dimensione variabile).
Questa implementazione è senz'altro più comoda,
potrebbe però essere discutibile nel caso in cui l'elenco dei figli fosse molto lungo e variasse dinamicamente...

Usando la terminologia degli alberi si sarebbero potuti usare questi nomi:
Dipendente -> nodoAlbero
supervisore -> padre
subordinati -> figli

Anche qui, non sono definiti tipi né per gli alberi né per la foresta.
Se si volessero aggiungere le definizioni di tipi opportuni, sarebbe opportuno continuare a scegliere nomi che facciano riferimento al problema dei dipendenti, ad esempio "azienda" per la foresta.

----------------------------------

A differenza che nelle versioni precedenti, nella versione 3 è definito un tipo che rappresenta l'intera foresta: si tratta del tipo "foresta", definito come struttura costituita da due campi:
- "supervisore" è la mappa che associa ad ogni stringa (che identifica un dipendente) la stringa che identifica il supervisore
- "subordinati" è la mappa che associa ad ogni stringa (che identifica un dipendente) la slice delle stringhe che identificano i suoi subordinati

Diversamente dalle versioni 1 e 2, qui i nodi della foresta (che rappresentano i dipendenti) non sono rappresentati esplicitamente, ma il loro insieme (e dunque l'insieme dei nomi dei dipendenti) è identificato dall'insieme delle chiavi delle mappe "subordinati" e "supervisore". NB: attenzione alla gestione di radici e foglie...!

*/