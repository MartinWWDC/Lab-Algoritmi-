# 3 Lista con puntatore a tail

Considerate ora una lista con concatenata di interi, dotata di due riferimenti al primo e all’ultimo
elemento della lista, come descritto dal tipo linkedListwithTail nella porzione di codice qui
sotto. Quando la lista è vuota, sia head che tail sono nil. La funzione newNode alloca lo
spazio per un nuovo nodo e ne inizializza il valore con l’argomento passato

```go
type listNode struct {
 item int
 next * listNode
}

type linkedListWithTail struct {
 head , tail * listNode
}

func newNode ( val int) * listNode {
 return & listNode {val , nil}
}
```

## 3.1 Inserimento alla fine

Considerate la seguente funzione (incompleta) che aggiunge un elemento e in fondo alla lista:

```go
func addNewNodeAtEnd (l * linkedListWithTail , val int) {
 if l. tail == nil {
  l. tail = newNode ( val )
  l. head = l. tail
 } else {
 // MISSING CODE
 }
}
```

Quale dei seguenti frammenti di codice completa la funzione addAtEnd? Spiegate quali
problemi si riscontrano scegliendo ciascuna delle altre opzioni.
A)

```go
l. tail . next = val

l. tail = val
```


B)

```go
temp := newNode ( val )
l. tail . next = temp
```


C) 

```go
temp := newNode ( val )
l. tail = temp
```


D) 

```go
l. tail . next = newNode ( val )
l. tail = l. tail . next
```

E) 

```go
temp := l. head
for temp . next != nil {
temp = temp . next
}
temp . next = newNode ( val )
```

## 3.2 Confronto tra lista semplice e lista con tail

Nella lista concatenata linkedListWithTail abbiamo i riferimenti all’inizio e alla fine della lista. Confrontate questa implementazione con quella di una lista semplice che non ha un riferimento alla fine della lista, cioè definita come segue:




```go
type linkedList struct {
  head * listNode
}
```

Per quali delle seguenti operazioni si ha un tempo migliore con linkedListWithTail piuttosto
che con linkedList? Scegliete tutte le opzioni corrette e giustificate la risposta.
A) restituisci 1 se la lista contiene un dato elemento
B) cancella l’ultimo elemento della lista
C) aggiungi un elemento all’inizio della lista
D) aggiungi un elemento alla fine della lista
