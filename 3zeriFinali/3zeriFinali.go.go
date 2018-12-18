package main
import "fmt"
func main(){
  var a int
  fmt.Print("Scrivi il numero da controllare ")
  fmt.Scan(&a)
  if a%1000==0{
    fmt.Println("Il numero termina con tre zeri")
  }else{
    fmt.Println("Il numero non termina con tre zeri")
  }
}