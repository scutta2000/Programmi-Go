package main
import "fmt"

type pazziente struct{
    nome string
    cognome string
    medicine somministrazione
}
type farmaco string
type somministrazione map[farmaco]int
 var dosaggioMassimo=somministrazione{"tachipirina":20, "paracetamolo":8, "oki":14, "propoli":100}

func killer(listaPazzienti[]pazziente)(ret string){
    for _,p:=range listaPazzienti{//per ogni pazziente
        for k,v:=range dosaggioMassimo{
            if p.medicine[k]>v{
                ret+=fmt.Sprintf(" %s",p.nome)
            }

        }
    }
    if ret == ""{
        ret="Non è morto nessuno per overdose"
    }else{
        ret="Sono morti di overdose: "+ret
    }
    return
}

func main(){
    febbre:=somministrazione{"tachipirina":15,"paracetamolo":17}
    p1:=pazziente{"Marco","Rossi",febbre}
    p2:=pazziente{"Giovanni","Certore",somministrazione{"oki":19,"propoli":25}}
    listaPazzienti:=[]pazziente{p1,p2}
    fmt.Println(killer(listaPazzienti))
}
