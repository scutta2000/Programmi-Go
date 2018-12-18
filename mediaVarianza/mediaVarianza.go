package main

import "fmt"

func media(num []float64) float64{
    var summ float64
    for _,n:= range(num){
        summ+=n
    }
    return summ / float64(len(num))
}

func varianza(num []float64) float64{
    var dist []float64
    dist=make([]float64,len(num))
    average:=media(num)
    for i,n:=range(num){
        dist[i]=(n-average)*(n-average)
    }
    return media(dist)
}

func main (){
    var num []float64
    fmt.Print("Inserisci una lista di numeri, Inserisci 0 per terminare ")
    for {
        var n float64
        fmt.Scan(&n)
        if n==0{
            break
        }
        num=append(num,n)
    }
    fmt.Printf("Media=%v Varianza=%v",media(num),varianza(num))
}
