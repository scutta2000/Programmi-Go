package main

import "fmt"

func reverse (s string) (rs string) {
    for _,c:=range(s){
        rs=string(c)+rs
    }
    return

}


func main(){
    var s string
    fmt.Print("Inserisci la stringa da invertire: ")
    fmt.Scan(&s)
    fmt.Println(reverse(s))
}
