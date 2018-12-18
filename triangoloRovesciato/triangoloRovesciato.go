/*
Obbiettivo

* * *
 * *
  *
di altezza n

*/

package main
import "fmt"
import "time"
func main (){
    var n, a int
    a=-1
    r:=1
    fmt.Scan(&n)
    for {
        a=a*(-1)

        for ; (r!=n+1) && (r!=-1) ;r+=a{
            time.Sleep(1)
            for l:=0;l<r;l++{
                fmt.Print(" ")
            }
            for l:=0;l<n-r;l++{
                fmt.Print("* ")
            }
            fmt.Println()

        }
        if a==1{
            r--
        }else{
            r=1
        }
    }
}
