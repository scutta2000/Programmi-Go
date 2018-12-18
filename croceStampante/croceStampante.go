
/*

l'Obbiettivo è stampare

*   *
 * *
  *
 *  *
*    *

N per N

approccio stampante
    È meglio qualla cartesiana

*/


package main

import "fmt"
func main(){
    var n int
    fmt.Scan(&n)
    for r:=0;r<n;r++{
        if r<n/2{
            // fmt.Print("1")
            for i:=0;i<r;i++{
                fmt.Print(" ")
            }
            fmt.Print("*")
            for i:=0;i<n-2-2*r;i++{
                fmt.Print(" ")
            }
            fmt.Print("*")
        }
        if r==n/2 && n%2!=0{
            // fmt.Print("2")
            for i:=0;i<r;i++{
                fmt.Print(" ")
            }
            fmt.Print("*")
        }

        if r>n/2 || (r==n/2 && n%2==0){
            // fmt.Print("3")
            for i:=0;i<n-1-r;i++{
                fmt.Print(" ")
            }
            fmt.Print("*")
            for i:=0;i<n-2-2*(n-r-1);i++{
                fmt.Print(" ")
            }
            fmt.Print("*")

        }

        fmt.Println()
    }
}
