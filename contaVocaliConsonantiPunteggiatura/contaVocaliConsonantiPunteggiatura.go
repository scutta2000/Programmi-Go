package main
import (
    "fmt"
    "unicode"
    )

func contaVocaliConsonantiPuteggiatura(s string) (int, int, int){
    var cc, cv, cp int
    for _,r:=range(s){
        switch{
        case r=='a'||r=='e'||r=='i'||r=='o'||r=='u':
            cv++
        case unicode.IsPunct(r):
            cp++
        case unicode.IsLetter(r):
            cc++
        }
    }
    return cv, cc, cp
}


func main(){
    var s string
    fmt.Print("inserisci una stringa: ")
    fmt.Scan(&s)
    cv,cc,cp:= contaVocaliConsonantiPuteggiatura(s)
    fmt.Printf("vocali: %d, consonanti: %d, punteggiatura:%d",cv,cc,cp)
}
