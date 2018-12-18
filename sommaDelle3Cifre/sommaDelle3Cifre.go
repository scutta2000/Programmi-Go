package main
import "fmt"
func main(){
  var a, a1, a2, a3 int
  fmt.Print("Inserisci un numero di 3 cifre ")
  fmt.Scan(&a)
  a1=a/100
  a-=a1*100
  fmt.Println(a)
  a2=a/10
  a-=a2*10
  a3=a
  a=a1+a2+a3
  if a<10{
    fmt.Println("La somma delle cifre del numero è minore di 10")
  }else if a>10{
    fmt.Println("La somma delle cifre del numero è maggiore di 10")
  }else{
    fmt.Println("La somma delle cifre del numero è uguale a 10")
  }

}