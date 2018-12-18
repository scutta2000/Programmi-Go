package main
import "fmt"



func main(){
    var s, sOut string
    fmt.Print("inserisci una stringa ")
    fmt.Scan(&s)
    for _,c:=range(s){
        flag:=true
        for _,r:=range(sOut){
            if c==r{
                flag=false
                break
            }
        }
        if flag{
            sOut+= string(c)
        }
    }
    fmt.Println(sOut)
}
