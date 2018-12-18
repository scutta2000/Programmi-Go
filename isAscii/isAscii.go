package main
import "fmt"
func main(){
    var s string
    fmt.Print("Inserisci la stringa da controllare ")
    fmt.Scan(&s)
    for i,c:=range(s){
        if rune(s[i])!=c{
            fmt.Println("La stringa contiene almeno un carattere non ascii")
        }
        if i+1==len(s){
            fmt.Print("La stringa contiene solo caratteri ascii")
    }
    }
}
