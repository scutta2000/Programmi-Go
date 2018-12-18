/*Definire una funzione che, data una stringa s, restituisce un array di 26 int contenenete le occorrenze
delle lettere dell'alfabeto latino presenti in s ('a' corrisponde all'indice 0, 'b' ad 1, etc.).
Non si faccia distinzione tra maiuscole e minuscole.
Definire quindi un programma che stampa le occorrenze delle lettere presenti negli argomenti passati da linea di comando.
PuÃ² essere utile definire una funzine stampaOccorrenze.*/

package main
import (
    "fmt"
    "os"
)
//Conta quante
func occorrenze(s string) (occorrenza [26]int) {
    var i rune
    for _,c:=range(s){
        if c>='a'&& c<='z' {
            i=c-'a'
        }else if c>='A'&& c<='Z'{
            i=c-'A'
        }
        occorrenza[i]++
    }
    return
}

func stampaOccorrenze(occorrenza [26]int){
    for i:=0;i<26;i++{
        fmt.Printf("%c:%-3d",'a'+i,occorrenza[i])
    }
}

func main(){
    for i:=1;i<len(os.Args);i++{
        stampaOccorrenze(occorrenze(os.Args[i]))
        fmt.Println()
    }
}
