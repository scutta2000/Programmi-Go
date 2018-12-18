/*

l'Obbiettivo è stampare

*   *
 * *
  *
 *  *
*    *

N per N

approccio piano cartesiano
*/


package main
import "fmt"
func main(){
    var n int
    fmt.Scan(&n)
    for y:=0;y<n;y++{
        for x:=0;x<n;x++{
            if x==y || y==n-1-x{
                fmt.Print("*")
            }else{
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}
