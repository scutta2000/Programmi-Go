package main
import "fmt"
func main () {
    var n, i int
    fmt.Print("Quante volte ti devo salutare, capo ")
    fmt.Scan(&n)
    for i=0; i<n;i++{
        fmt.Println("Ciao")
    }
    fmt.Print("Premi enter per chiudere ")
    fmt.Scan(&n)
}
