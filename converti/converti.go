package main

import "fmt"

const SYMB = "0123456789ABCDEF" //set di simboli per num. fino a base 16

func main() {
    var n,b int
    fmt.Print("Inserisci un numero da convertire: ")
    fmt.Scan(&n)
    fmt.Print("Inserisci la base in cui converitre il numero: ")
    fmt.Scan(&b)
    if n==0{
        fmt.Println(0)
    }
    if b>16{
        fmt.Println("La base inserita è troppo alta, questo programma può convertire i numeri solo fino alla base 16")
    }else{
        var r int
        var totR string

        //restituisci il numero covertito ma in ordine invertito
        for n>0{
            r=n%b
            n=n/b
            totR+=string(SYMB[r])
        }
        // fmt.Println(totR,len(totR))
        var output string

        //inverte la totR
        for i:=len(totR)-1;i>=0;i--{
            // fmt.Println(i)
            output+=string(totR[i])
            // fmt.Println(output)
        }
        fmt.Println(output)
    }
}
