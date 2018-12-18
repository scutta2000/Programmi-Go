
package main
import (
    "fmt"
    "math"
)
func main(){
    var n float64
    for i:=2;;i++{

        // controllo se i è primo
        var d int
        d=2
        for ;d<i;d++{
            if i%d==0{
                break
            }
        }
        //entra in questa if solo se i è primo
        if d==i{
            n++
        }

        //Converto le variabili in float64p
        var iFloat, nFloat float64
        iFloat=float64(i)
        nFloat=float64(n)

        //in teoria all infinito dovrebbe fare 1
        fmt.Println((nFloat*math.Log(iFloat))/iFloat)



    }
}
