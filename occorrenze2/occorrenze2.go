/*Definire una funzione che, data una stringa s, ritorna una slice di un tipo occorrenza (una struct) formato da due campi:
simbolo (rune) e quanti (int). La slice conterrÃ  le occorrenze di tutti gli elementi della stringa, in ordine di apparizione.
Come per l'esercizio precedente, Definire un programma che stampa le occorrenze dei simboli (unicode) presenti negli argomenti passati da linea di comando.
PuÃ² essere utile definire una funzine stampaOccorrenza.*/

package main

import(
    "fmt"
    "os"
)
type occorrenza struct{
    simbolo rune
    quanti int
}

func isInSlice (o rune, occorrenze []occorrenza) (bool,int){
    for i,p:=range(occorrenze){
        if o==p.simbolo{
            return true, i
        }
    }
    return false, 0
}

func occorrenze(s string)(occorrenze []occorrenza){
    for _,c:=range(s){
        is,i:=isInSlice(c,occorrenze)
        if !is{
            // occorrenze=append(occorrenze,nuovo)
            occorrenze=append(occorrenze,occorrenza{c,1})
        }else{
            occorrenze[i].quanti ++
        }
    }
    return
}

func stampaOccorrenze(occorrenza []occorrenza){
    for i:=0;i<len(occorrenza);i++{
        // fmt.Printf("%c:%-3d",'a'+i,occorrenza[i])
        fmt.Printf("%c:%-3d",occorrenza[i].simbolo,occorrenza[i].quanti)
    }
}

func main(){
    for i:=1;i<len(os.Args);i++{
        stampaOccorrenze(occorrenze(os.Args[i]))
        // fmt.Print(occorrenze(os.Args[i]))
        fmt.Println()
    }
}
