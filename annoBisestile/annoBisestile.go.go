package main
import "fmt"
func main(){
  var anno int
  fmt.Print("Inserisci l'anno da controllare ")
  fmt.Scan(&anno)
  if anno%4==0 && anno%100!=0{
    fmt.Println("L'anno inserito è bisestile")
  }else if anno%400==0{
    fmt.Println("L'anno inserito è bisestile")
  }else{
    fmt.Println("L'anno inserito non è bisestile")
  }
  
}