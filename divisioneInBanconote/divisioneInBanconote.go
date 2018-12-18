package main
import "fmt"
func main(){
	var tot, cento, cinquanta, venti, dieci, cinque, due, uno int
	fmt.Print("Inserisci la quantita di denaro ")
	fmt.Scan(&tot)
	cento= tot/100
	tot=tot%100
	cinquanta=tot/50
	tot=tot%50
	venti= tot/20
	tot=tot%20
	dieci= tot/10
	tot=tot%10
	cinque= tot/5
	tot=tot%5
	due=tot/2
	tot=tot%2
	uno=tot
	fmt.Println("devi usare", cento,"Banconote da cento,",cinquanta,"banconote da cinquanta",venti,"banconote da venti", dieci,"banconote da dieci",cinque,"banconote da cinque", due,"banconote da due", uno,"banconote da uno")
}
