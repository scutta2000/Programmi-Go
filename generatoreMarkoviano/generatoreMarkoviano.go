package main
import (
    "fmt"
    "math/rand"
    "os"
    "bufio"
    // "strconv"
)

func follow (s string,m map[rune][]rune) {
    c1:=rune(s[0])
    for _,c:= range s[1:]{
        m[c1]=append(m[c1],c)
        c1=c
    }
}
// func stampa(m map [rune][]rune) (s string){
//     rand.Seed(5)
//     for k,v:=range m{
//         s+=fmt.Sprintf("%c:[",k)
//         for _,c:=range v{
//             s+=fmt.Sprintf("%c,",c)
//         }
//         s+="]"
//     }
//     return
// }
func genera(m map[rune][]rune, n int) (s string){
    c:='a'
    for i:=0;i<n;i++{
        // a:=rand.Int()%(len(m[c]))
        // s+=fmt.Sprintf("%c",m[c][a])
        // c=m[c][a]
        c=m[c][rand.Int()%len(m[c])]
        s+=fmt.Sprintf("%c",c)
    }
    return
}

func main(){
    m:=make(map [rune][]rune)
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        follow(scanner.Text(),m)
    }
    // follow("ciao come va",m)
    n:=100
    fmt.Println(genera(m,n))
}
