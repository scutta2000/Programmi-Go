// Implementa una calcolatrice che esegue ripetutamente operazioni la cui sintassi è:
// <operatore> <operando>
// dove:
// <operatore> è uno fra +,-,*,\, ed S (s),
// <operando> è un valore float.
// La calcolatrice visualizza e salva in un registro il risultato di ogni operazione aritmetica.
// Il valore corrente del registro è il primo operando (implicito) di una operazione.
// L'operazione
// S x
// assegna al registro il nuovo valore x.
// Se viene digitato un operatore diverso viene stampato il messaggio "operatore sconosciuto".
// Se il divisore è 0.0 viene stampato il messaggio "divisione per zero" (in entrambi i casi il programma procede).
// Il comando Q (q) termina il programma.

package main

import "fmt"

func main() {
	fmt.Println("comandi: <operatore> <valore> (S <valore> per resettare); Q per uscire")
	var valore, memoria float64
	var operatore string
	for operatore!="Q"&&operatore!="q"{
		fmt.Println(memoria)
		fmt.Scan(&operatore,&valore)
		switch operatore{
		case "+":
			memoria+=valore
		case "-":
			memoria-=valore
		case "*":
			memoria*=valore
		case "/":
			if valore == 0{
				fmt.Println("Divisione per zero ")
				break
			}
			memoria/=valore
		case "S","s":
			memoria=valore
		case "Q","q":
			// in questo caso non fa niente è il programma esce nel controllo del for
		default:
			fmt.Println("Il comando inserito non è corretto")
		}
	}
}
