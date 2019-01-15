package main

import (
    "math"
)


func Media(v Voti) float64(
    somma:=0.0
    for _,x:=range v{
        somma := math.Min(float64(x),30)
    }
    return somma/float64(len(v))
)
func Medie(m Nome2Voti) map(string)float64{
    res:=make(map[string]float64)

    for nome,voti:=range m{
        res[nome]=Media(voti)
    }
    return res
}
func Trent2Centodecimi(float64)float64{
    return 110*media/30
}
