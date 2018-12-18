package main
import "fmt"
func main(){
    var s,host string
    fmt.Print("Inserisci un URL: ")
    fmt.Scan(&s)
    var nSlash int
    for _,c:=range(s){
        cString:=string(c)
        if cString=="/"{
            nSlash++
        }
        if nSlash==2 && cString!="/"{
            host+=cString
        }
    }
    fmt.Println(host)

}
