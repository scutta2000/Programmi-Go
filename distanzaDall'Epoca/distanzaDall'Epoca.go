package main

import "fmt"


type data struct{
            g int
            m int
            a int
}

func isBisestile(a int)bool{
    return a%4==0&&(a%100!=0||a%400==0)
}

func lungMese(m int, a int)int{
    switch m{
    case 11,4,6,9:
        return 30
	case 2:
        if isBisestile(a){
            return 29
        }else{
            return 28
        }
    default:
        return 31
    }
}
func lungAnno(a int)int{
    if isBisestile(a){
        return 366
    }else{
        return 365
    }
}

func distEpoca(d data)(s int){
    toggle:=1
    if d.a<1970{
        d.a=1970+(1970-d.a)
        toggle=-1
    }
    for anno:=1970;anno<d.a;anno++{
        s+=lungAnno(anno)
	}
    for mese:=1;mese<d.m;mese++{
        s+=lungMese(mese,d.a)*toggle
    }
    s+=(d.g-1)*toggle
    if d.a<1970{
        s*=-1
    }
    s*=toggle
    return
}

func main(){
    var g,m,a int
    var d data
    fmt.Print("Inserisci una data ")
    fmt.Scan(&g,&m,&a)
    d=data{g,m,a}
    fmt.Printf("La data inserita dista %d giorni dall'epoca",distEpoca(d))
}
