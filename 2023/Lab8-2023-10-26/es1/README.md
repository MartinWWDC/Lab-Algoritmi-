# 1 Implementazione di liste concatenate semplici
Scrivete un programma con l’implementazione di una lista concatenata semplice, seguendo i
lucidi presentati a lezione. Definite i tipi listNode e linkedList e scrivete queste funzioni:
• newNode, che crea un nuovo nodo di lista;
• addNewNode, che inserisce un nuovo nodo in testa alla lista;
• printList, che stampa una lista;
• searchList, che cerca un elemento in una lista;
• removeItem, che cancella un item da una lista.
Per testare le vostre funzioni scrivete una funzione main che gestisca un insieme dinamico
(che variano nel tempo) di interi. Il main deve leggere da standard input una sequenza di istruzioni secondo il formato nella tabella qui sotto, dove n è un intero. I vari elementi sulla riga
sono separati da uno o più spazi. Quando una riga è letta, viene eseguita l’operazione associata;
le operazioni di stampa sono effettuate sullo standard output, e ogni operazione deve iniziare su
una nuova riga.
*Ultimo aggiornamento: 26 ottobre 2022 - 12:21:11
1
Istruzione Operazione
in input
+ n Se n non appartiene all’insieme lo inserisce, altrimenti non fa nulla
- n Se n appartiene all’insieme lo elimina, altrimenti non fa nulla
? n Stampa un messaggio che dichiara se n appartiene all’insieme
c Stampa il numero di elementi dell’insieme
p Stampa gli elementi dell’insieme (nell’ordine in cui compaiono nella lista)
o Stampa gli elementi dell’insieme nell’ordine inverso
d Cancella tutti gli elementi dell’insieme
f Termina l’esecuzione
Si assume che l’input sia inserito correttamente. Conviene scrivere le istruzioni di input in un
file in.txt ed eseguire il programma redirigendo lo standard input.