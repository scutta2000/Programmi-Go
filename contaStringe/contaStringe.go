package main
import "fmt"



//controlla se una runa è presente in una stringa
func isInString(s string, r rune)bool{
    // fmt.Printf("r=%c \n",r)
    for _,c:=range(s){
        if c==r{
            return true
        }
    }
    return false
}


func main(){
    var input,soFar string
    fmt.Print("Inserisci la stringa ")
    fmt.Scan(&input)
    for _,c:=range(input){
        n:=0
        for _,r:=range(input){
            if r==c || r+32==c || r-32==c {
                n++
            }
        }
        if !isInString(soFar,c)&&!isInString(soFar,c+32)&&!isInString(soFar,c-32){
            fmt.Printf("%c : %d \n",c,n)
        }
        soFar+=string(c)
    }
}
