# es 3

1. Modellate la situazione con un albero: cosa rappresentano i nodi dell’albero? cosa rappresenta la relazione padre/figlio?
   - i nodi dell'albero rappresentano i pianeti  menetre le relazioni rappresentano intorno a quale corpo celesteorbita il nodo figlio
2. Riformulate i problemi usando la terminologia degli alberi:
   a) Parte 1 - Cosa sono le orbite dirette? E le orbite indirette? Come può essere descritto, in termini di alberi, il numero di orbite indirette di un oggetto?
   b) Parte 2 - Come può essere descritta, in termini di alberi, La distanza tra gli oggetti
   attorno a cui orbitano YOU e SAN?
   1. due nodo sono detti in orbite diretti quando sono uniti da un arco. Mentre tutti i figli non diretti di un nodo sono dette le orbite indirette
   2. il numero di nodi padri rispetto al nodo di di cui si calcola la distanza che sono contrapposti  tra i due nodi di cui si sta calocolando la distanza,vengono sono i nodi le cui orbite si trovano in mezzo
3. Progettate una soluzione per calcolare:
   a) Parte 1 - il numero di orbite indirette
   b) Parte 2 - il numero di trasferimenti di orbita necessari
   Ragionare in termini di alberi vi aiuterà a impostare gli algoritmi risolutivi!
   1. dati due pianeti, uno di partenza e uno di arrivo,il numero di orbite indirette si può trovare facendo una ricerca ricorsiva sui padri fino al raggiungimento dell'elemento 
   2. dati due nodi per i quale cercare il path si  parte da entrambi i nodi e si cerca ricorsiva per ogni nodo padre il proprio padre finchè non si trova un nodo comune tra i due