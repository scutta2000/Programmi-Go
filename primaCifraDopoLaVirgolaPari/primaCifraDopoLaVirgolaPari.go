package main
import "fmt"
func main(){
  var a float64
  var b int
  fmt.Print("inserisci il numero ")
  fmt.Scan(&a)
  a*=10
  b=int(a)
  b=b%10
  if b%2==0{
    fmt.Println("La prima cira dopo la virgola del numero è pari ")
  }else{
    fmt.Println("La prima cifra dopo la virgola del numero è dispari")
  }
}