package main
import "fmt"
func main(){
    var n, a int
    fmt.Print("Inserisci un numero: ")
    fmt.Scan(&n)
    t:=1
    for t>=1{
        fmt.Print("Inserisci un numero da controllare: ")
        fmt.Scan(&t)
        var i int
        for i=1;i<=t;i++{
            a=1
            // fmt.Println(i)
            //i alla n
            for j:=0;j<n;j++{
                a=a*i
                // fmt.Println("a= ",a)
            }
            if a==t{
                fmt.Println("il numero Inserito è uguale a",i,"alla",n)
                break
            }
            if i==t{
                fmt.Println("Il numero inserito non è la pontenza ennesima di alcun numero ")
        }
        }
    }
}
