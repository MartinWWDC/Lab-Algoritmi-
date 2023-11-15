# Es 1

Nella multinazionale Algoré il lavoro è organizzato in maniera gerarchica. Ogni dipendente è inquadrato in un certo livello di impiego. Tranne i dipendenti di massimo livello, ogni dipendente ha un supervisore, di cui è detto subordinato. Un dipendente può avere 0, 1, o più subordinati. Si considerino i seguenti compiti.

(a) Dato un certo dipendente, stampare l’elenco dei suoi subordinati.
(b) Contare quanti sono i dipendenti che non hanno alcun subordinato.
(c) Dato un certo dipendente, individuare chi è il suo supervisore.
(d) Dato un certo dipendente, stampare la lista dei dipendenti che si trovano sopra di lui gerarchicamente, partendo dal suo supervisore e risalendo la gerarchia fino a un dipendente di massimo livello.
(e) Stampare l’elenco di tutti i dipendenti –non importa l’ordine–, indicando per ciascuno chi è il suo supervisore (tranne nel caso di dipendenti di massimo livello).

## Es1.1

1. Modellate la situazione con una struttura dati opportuna:
   
   * descrivete come si possono rappresentare i dipendenti e le loro relazioni con la struttura dati scelta;
   
   * riformulate, nei termini della struttura dati scelta, ciascuno dei compiti enunciati sopra.

2. Descrivete come è opportuno implementare la struttura dati scelta.

3. Per ciascun compito, progettate e descrivete un algoritmo che consente di svolgere il compito, sfruttando le scelte di progettazione e implementazione fatte precedentemente.
   Gli algoritmi possono essere descritti a parole o in pseudocodice; può essere opportuno
   fare riferimento ad algoritmi noti.

---

Dentro di me sussuste un dibattito sulla corretta risoluzione di questo problema, procedo. La struttura dati da me scelta è l'albero e sono convito sia la struttura migliore per implementare una gerarchia aziendale. dove i nodi padre sono i sovrintendenti e i nodi figli sono i nodi subordinati.

Nasce ora un dubbio: come rappresentare l'albero? 

* nodicon lista di padri: in questo caso facente riferiemento che ogni ogni dipendente venga  rappresentato con la lista dei suoi superiori. 
  
  Questa implementazione comporterebbe un  ottimale implementazione delle funzionalità:
  
  * d: perchè permette una ricerca ricorsiva tra i responsabili
  
  * e

* nodi con lista dei figli:  dove ogni nodo ha associato la lista dei suoi sottoposti  questo permetterebbe un implementazione migliore per:
  
  * a
  
  * b
  
  * c
  
  


