package main
import (
    "fmt"
    "bufio"
    "os"
)

func rigaDiAsterischi(n int){
    for i:=0;i<n;i++{
        fmt.Print("*")
    }
    fmt.Println()
}
func stampaIncorniciata (msg string){
    n:=0
    for range(msg){
        n++
    }
    rigaDiAsterischi(n+4)
    fmt.Printf("* %s *\n",msg)
    rigaDiAsterischi(n+4)
}
func main(){
    scanner:=bufio.NewScanner(os.Stdin)
    fmt.Println("Inserisci delle stringhe da incorniciare")
    for scanner.Scan(){
        riga:=scanner.Text()
        if riga !=""{
            stampaIncorniciata(riga)
        }
    }
}
