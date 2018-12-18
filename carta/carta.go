package main
import (
    "fmt"
    "strconv"
)


//dato un numero da 1 a 52 restituisce la carta corrispondente
func carta(x int) (string, string, bool) {
    if x<=0 ||x>52{
        return "","",false
    }
    var v, t string
     switch { //seme
     case x < 14:
     t = "cuori"
     case x < 27:
     t = "quadri"
     case x < 40:
     t = "fiori"
     default:
     t = "picche"
     }
     switch r := x % 13; r { // valore
     case 0:
     v = "K"
     case 1:
     v = "A"
     case 11:
     v = "J"
     case 12:
     v = "Q"
     default:
     v=strconv.Itoa(r)
    }
    return v,t, true
}

func valoreBJ(val string) (valnum int){
    switch val{
    case "A":
        valnum=1
    case "K","Q","J":
        valnum=10
    default:
        valnum,_=strconv.Atoi(val)
    }
    return
}


func main(){
    var x int
    fmt.Print("Inserisci un numero (0,52] ")
    fmt.Scan(&x)
    val,seme,ok:=carta(x)
    if !ok{
        fmt.Println("Valore non nel range")
        return
    }
    fmt.Printf("%s di %s \n",val,seme)
    fmt.Println("valore di black jack", valoreBJ(val))
}
