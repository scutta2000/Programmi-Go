package main

import "fmt"

//tipi e funzioni base per implementare una scacchiera

// Colore alias per il colore dei pezzi
type Colore bool

//Tipo alias per il tipo dei pezzi
type Tipo byte

const (
	//BIANCO è un alias per false
	BIANCO = false
	//NERO è un alias per true
	NERO = true
)

//i simboli dei pezzi sono disposti nello stesso ordine in cui compaiono in tab Unicode ...
const (
	NULL = iota // 0
	RE
	REGINA
	TORRE
	ALFIERE
	CAVALLO
	PEDONE
)

type (
	//Pezzo = struttura per indicare i vari pezzi {Tipo, Colore}
	Pezzo struct {
		tipo   Tipo
		colore Colore
	}
	//Casella = struttura riga, colonna per indicare le caselle della Scacchiera
	Casella struct {
		riga byte //riga: 1..8
		col  rune //colonna: 'a'..'h'
	}
	//Scacchiera mappa che associa una casella ad un pezzo
	Scacchiera map[Casella]Pezzo
)

// StringPezzo restituisce la rapresentazione testuale (t) di un Pezzo degli scacchi
// se il pezzo non esiste restituisce ""
func StringPezzo(p Pezzo) (t string) {
	re := 0x2654 // "re bianco" in Unicode; i valori dei pezzi sono contigui (prima i bianchi poi i neri)
	if p.tipo >= RE && p.tipo <= PEDONE {
		if p.colore == NERO {
			re += 6
		}
		t = string(re - 1 + int(p.tipo)) //-1 perchè RE == 1
	}
	return
}

// verifica se le coordinate di una casella sono corrette
func check(r byte, c rune) bool {
	return r > 0 && r < 9 && c >= 'a' && c <= 'h'
}

//le operazioni di base sulla map sono ridefinite per comoditÃ

// ritorna il pezzo che si trova nella casella (r,c) della scacchiera s
// e true oppure false se la casella è piena o vuota (in quest'ultimo caso ritorna (Pezzo{},false))
func get(s Scacchiera, r byte, c rune) (p Pezzo, found bool) {
	p, found = s[Casella{r, c}]
	return
}

// dispone il pezzo p sulla casella (r,c) della scacchiera s
// non esegue alcun controllo
func put(s Scacchiera, p Pezzo, r byte, c rune) {
	s[Casella{r, c}] = p
}

// svuota la casella (r,c) della scacchiera s
// non esegue alcun controllo
func remove(s Scacchiera, r byte, c rune) {
	delete(s, Casella{r, c})
}

// Move se possibile, in base alla attuale disposizione dei pezzi, esegue la mossa (r1,b1) in (r2,b2)
// sulla scacchiera s
// ritorna (src, dest, true), dove src e dest sono i pezzi sulle caselle di partenza e di arrivo
// (quest'ultimo sarÃ  uguale al valre di default Pezzo{} se la casella di arrivo è vuota)
// ritorna (Pezzo{}, Pezzo{}, false) se la mossa non è valida per qualsiasi motivo
func Move(s Scacchiera, r1 byte, c1 rune, r2 byte, c2 rune) (src, dest Pezzo, ok bool) {
	if src, dest, ok = vaildMove(s, r1, c1, r2, c2); ok { // mossa valida
		put(s, src, r2, c2) // dispone il pezzo src sulla casella di destinazione
		remove(s, r1, c1)   // svuota la casella di partenza
	}
	return
}

// verifica se la mossa (r1,b1) in (r2,b2) sulla scacchiera s è valida
// [x]controlla che le coordinate delle caselle siano corrette, che la casella iniziale non sia vuota,
// [x]che la casella finale sia vuota o contenga un pezzo di colore diverso (attenzione alla regola del pedone!),
// [x]che la mossa sia valida per il pezzo da muovere, in base alle regole,...)
// se la mossa è valida ritorna (src, dest, true), dove src e dest sono i pezzi sulle caselle di
// partenza e di arrivo (quest'ultimo sarÃ  uguale a Pezzo{} se la casella di arrivo è vuota)
// ritorna (Pezzo{}, Pezzo{}, false) se la mossa non è valida
func vaildMove(s Scacchiera, r1 byte, c1 rune, r2 byte, c2 rune) (src, dest Pezzo, ok bool) {
	ok = true                            //ok è true di default e diventa false nelle if
	src, _ = get(s, r1, c1)              //pezzo sulla casella iniziale
	dest, destPopolata := get(s, r2, c2) // pezzo sulla casella finale
	var rMin, rMax byte
	var cMin, cMax rune
	if r1 <= r2 {
		rMin = r1
		rMax = r2
	} else {
		rMin = r2
		rMax = r1
	}
	if c1 <= c2 {
		cMin = c1
		cMax = c2
	} else {
		cMin = c2
		cMax = c1
	}
	if ok = check(r1, c1) && check(r2, c2); !ok { // le coordinate delle caselle sono ok
		ok = false
		// fmt.Println("CHECK FALLITO")
		return
	}
	if _, ok = get(s, r1, c1); !ok { //la casella iniziale non è vuota
		ok = false
		// fmt.Println("get FALLITO")
		return
	}
	if (src.colore == dest.colore) && destPopolata { //I colori tornano o la destinazione è vuota
		ok = false
		// fmt.Println("colore FALLITO")
		return
	}
	switch src.tipo {
	case TORRE:
		ok = torre(s, r1, c1, r2, c2, rMin, rMax, cMin, cMax)
	case ALFIERE:
		ok = alfiere(s, r1, c1, r2, c2, rMin, rMax, cMin, cMax)
	case RE:
		if rMax-rMin > 1 || cMax-cMin > 1 { //Controlla che non si sia mosso di più di una casella
			ok = false
		}
	case PEDONE:
		ok = pedone(s, r1, c1, r2, c2, rMin, rMax, cMin, cMax, destPopolata, src)
	case REGINA:
		ok = regina(s, r1, c1, r2, c2, rMin, rMax, cMin, cMax)
	case CAVALLO:
		if (rMax-rMin != 2 || cMax-cMin != 1) && (rMax-rMin != 1 || cMax-cMin != 2) {
			ok = false
		}
	}
	if ok {
		return
	}
	src = Pezzo{}
	dest = Pezzo{}
	return
}

//Funzione che restituisce se la mossa di una torre è legittima
func torre(s Scacchiera, r1 byte, c1 rune, r2 byte, c2 rune, rMin byte, rMax byte, cMin rune, cMax rune) (ok bool) {
	ok = true
	if r1 != r2 && c1 != c2 { //Controlla che sia una mossa che può fare la torre
		ok = false
	}
	// fmt.Println("TORRE IF",ok)
	for i := rMin + 1; i < rMax; i++ { //Controlla che tutte le righe in mezzo siano vuote
		if _, check := get(s, i, c1); check {
			ok = false
			// fmt.Println("TORRE IF",ok)
			break
		}
	}
	// fmt.Println("TORRE FOR",ok)
	for i := cMin + 1; i < cMax; i++ { //Controlla che tutte le colonne in mezzo siano vuote
		if _, check := get(s, r1, i); check {
			ok = false
			break
		}
	}
	// fmt.Println("TORRE FOR 2",ok)
	return
}

//Funzione che restituisce se la mossa di un alfiere è legittima
func alfiere(s Scacchiera, r1 byte, c1 rune, r2 byte, c2 rune, rMin byte, rMax byte, cMin rune, cMax rune) (ok bool) {
	ok = true
	if rMax-rMin != byte(cMax-cMin) { //Controllo movimento in diagonale
		fmt.Println(rMax-rMin, byte(cMax-cMin))
		ok = false
	}
	for ir, ic := rMin+1, cMin+1; ir < rMax-1; { //Controlla se tutte le diagonali sono vuote
		if _, check := get(s, ir, ic); check {
			ok = false
			break
		}
		ir++
		ic++
	}
	return
}

//Funzione che determina se la mossa di un pedone è legittima
func pedone(s Scacchiera, r1 byte, c1 rune, r2 byte, c2 rune, rMin byte, rMax byte, cMin rune, cMax rune, destPopolata bool, src Pezzo) (ok bool) {
	ok = true
	if destPopolata { //se sto mangiando
		if c1+1 != c2 && c1-1 != c2 { //devo spostarmi a destra o sinistra di 1
			// fmt.Println("DEVI SPOSTARTI A DESTRA O SINISTRA DI 1 ")
			ok = false
		}
		if src.colore == BIANCO {
			if r1+1 != r2 { //devo andare avanti di uno
				// fmt.Println("DEVI ANDARE AVANTI DI UNO")
				ok = false
			}
		} else {
			if r1-1 != r2 { //devo andare avanti di uno
				// fmt.Println("DEVI ANDARE AVANTI DI UNO")
				ok = false
			}
		}
	} else { //se non sto mangiando
		if c1 != c2 { //devo andare dritto
			// fmt.Println("DEVI ANDARE DRITTO")
			ok = false
		}
		if src.colore == BIANCO {
			if r1 == 2 { //Se è la prima mossa del pedone
				if r2-r1 > 2 { //Non posso muovermi di più di 2
					ok = false
				}
			} else { //Se non è la prima mossa del pedone
				if r2-r1 > 1 { //Non posso muovermi di più di 1
					ok = false
				}
			}
		} else {
			if r1 == 7 { //Se è la prima mossa del pedone
				if r1-r2 > 2 { //Non posso muovermi di più di 2
					ok = false
				}
			} else { //Se non è la prima mossa del pedone
				if r1-r2 > 1 { //Non posso muovermi di più di 1
					ok = false
				}
			}
		}
	}
	return ok
}

//Funzione che determina se la mossa di un regina è legittima
func regina(s Scacchiera, r1 byte, c1 rune, r2 byte, c2 rune, rMin byte, rMax byte, cMin rune, cMax rune) (ok bool) {
	ok = true
	if rMax-rMin == byte(cMax-cMin) { //diagonale
		for ir, ic := rMin+1, cMin+1; ir < rMax-1; { //Controlla se tutte le diagonali sono vuote
			if _, check := get(s, ir, ic); check {
				ok = false
				break
			}
			ir++
			ic++
		}
	} else if r1 == r2 || c1 == c2 { //dritto
		for i := cMin + 1; i < cMax; i++ { //Controlla che tutte le colonne in mezzo siano vuote
			if _, check := get(s, r1, i); check {
				ok = false
				break
			}
		}
	} else {
		ok = false
	}
	return
}
