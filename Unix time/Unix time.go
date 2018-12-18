package main
import "fmt"
func main(){
	var t, anni, mesi, giorni, ore, minuti, secondi, sInAnni, sInMesi, sInGiorni, sInOre, sInMinuti int
	fmt.Print("Inserisci quanti secondi")
	fmt.Scan(&t)
	sInMinuti=60
	sInOre=sInMinuti*60
	sInGiorni=sInOre*24
	//condidero tutti i mesi di 30 giorni e di conseguenza gli anni di 360 giorni
	sInMesi=sInGiorni*30
	sInAnni=sInMesi*12

	t=t+1970*sInAnni


	anni=t/sInAnni
	fmt.Println("Anno:",anni)
	t=t%sInAnni
	
	mesi=t/sInMesi
	fmt.Println("Mese:",mesi+1)
	t=t%sInMesi
	
	giorni=t/sInGiorni
	fmt.Println("Giorno:",giorni+1)
	t=t%sInGiorni
	
	ore=t/sInOre
	fmt.Println("Ora:",ore)
	t=t%sInOre
	
	minuti=t/sInMinuti
	fmt.Println("Minuti:",minuti)
	t=t%sInMinuti
	
	secondi=t
	fmt.Println("Secondi:",secondi)
	
	
}