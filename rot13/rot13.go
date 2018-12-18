package main
import "fmt"
func main(){
    var sIn,sOut string
    fmt.Print("Iserisci una stringa: ")
    fmt.Scan(&sIn)
    for _,c:=range(sIn){
        // non funziona perchè il mondo è triste: c=(c)%26+91
        c=(c-84)%26+97
        sOut+=string(c)
    }
    fmt.Println(sOut)
}
