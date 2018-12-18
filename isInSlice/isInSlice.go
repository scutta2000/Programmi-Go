package main

import "fmt"

func isInSlice (n int, s []int) bool{
    for _,c:=range(s){
        if c==n{
            return true
        }
    }
    return false
}

func main(){
    var n int
    var s []int
    fmt.Print("Iserisci un intero: ")
    fmt.Scan(&n)
    fmt.Print("Inserisci gli interi che rientreranno nella slice, inserisci -1 per terminare ")
    for {
        var n int
        fmt.Scan(&n)
        if n==-1{
            break
        }
        s=append(s,n)
    }
    if isInSlice(n,s){
        fmt.Println("L'intero inserito è presente nella slice")
    }else{
        fmt.Println("L'intero inserito non è presente nella slice")
    }
}
