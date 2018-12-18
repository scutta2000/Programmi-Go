/*Definire una funzione che, dato un array di int, ne modifica gli elementi facendo in modo che il valore
in posizione i sia pari alla somma dei valori nelle posizioni 0,..,i-1.
Definire quindi un programma che applichi la funzione una sequenza di valori passati dalla linea di comando.*/
package main
import (
    "fmt"
    "os"
    "strconv"
)

func accumula (a []int){
    for i:=1;i<len(a);i++{
        a[i]+=a[i-1]
    }
}
func main(){
    var a []int
    for i:=1;i<len(os.Args);i++{
        b,err:=strconv.Atoi(os.Args[i])
        if err !=nil{
            return
        }
        a=append(a,b)
    }
    fmt.Println("a prima",a)
    accumula(a)
    fmt.Println("a dopo ",a)
}
