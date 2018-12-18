package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

var (
	nomeSeme   []string = []string{"♥️", "♦️", "♣️", "♠️"}
	nomeValore []string = []string{0: "A", 10: "J", 11: "Q", 12: "K"}
	tavolo     struct {
		mano []carta
		fish int
	}
	listaGiocatori []*giocatore
	listaMazzi     []*[]carta
)

func initialize() {
	for i := 1; i <= 9; i++ {
		nomeValore[i] = strconv.Itoa(i + 1)
	}
}

//valore da 0 a 51 di cui le prime 13 sono ♥️, poi ♦️ poi ♣️ poi ♠️
type (
	carta     int
	giocatore struct {
		nome string
		mano []carta
		fish int
	}
)

func formatCarta(c carta) string {
	return fmt.Sprintf("%s%s", nomeValore[c%13], nomeSeme[c/13])
}

func creaMazzo() []carta {
	res := make([]carta, 52)
	for i := 0; i < 52; i++ {
		res[i] = carta(i)
	}
	// for i,m:=range(listaMazzi){
	// 	if m==nil{
	// 		listaMazzi[i]=res
	// 		return res
	// 	}
	// }
	// listaMazzi=append(listaMazzi,res)
	return res
}

func stampaMazzo(m []carta) (s string) {
	for _, c := range m {
		s += fmt.Sprintf("%s", formatCarta(c))
	}
	return
}

func isPresent(m []carta, c carta) bool {
	for _, x := range m {
		if x == c {
			return true
		}
	}
	return false
}

func mescolaMazzo(m []carta) {
	for i := 0; i < len(m); i++ {
		j := i + rand.Int()%(len(m)-i)
		m[i], m[j] = m[j], m[i]
	}
}

func rimuoviCarta(i int, m []carta) {
	m = append(m[:i], m[i+1:]...)
}

func deal(g *giocatore, m []carta) {
	g.mano = make([]carta, 2)
	for i := 0; i < 2; i++ {
		g.mano[i] = m[0]
		rimuoviCarta(0, m)
	}
}

func flip(n int, m []carta) {
	for i := 0; i < n; i++ {
		tavolo.mano = append(tavolo.mano, m[0])
		rimuoviCarta(0, m)
	}
}
func nuovoGiocatore(nome string, fish int) giocatore {
	g := giocatore{nome: nome, fish: fish}
	// listaGiocatori=append(listaGiocatori,g)
	return g
}

func reset() {
	for i := 0; i < len(listaGiocatori); i++ {
		listaGiocatori[i].mano = nil
	}
	tavolo.mano = nil
	tavolo.fish = 0
	// for i:=0;i<len(listaMazzi);i++{
	// 	*listaMazzi[i]=creaMazzo()
	// 	mescolaMazzo(*listaMazzi[i])
	// }
	for _, m := range listaMazzi {
		*m = creaMazzo()
		mescolaMazzo(*m)

	}
}

func puntata(g *giocatore, n int) {
	tavolo.fish += n
	g.fish -= n
}

func main() {
	initialize()
	m := creaMazzo()
	listaMazzi = append(listaMazzi, &m)
	mescolaMazzo(m)

	pietro := nuovoGiocatore("Pietro", 500)
	ai := nuovoGiocatore("AI", 500)
	listaGiocatori = append(listaGiocatori, &pietro, &ai)

	for _, g := range listaGiocatori {
		deal(g, m)
	}

	fmt.Printf("%s la tua mano è: %s\nQuanto vuoi puntare? ", pietro.nome, stampaMazzo(pietro.mano))
	var n int
	fmt.Scan(&n)
	puntata(&pietro, n)
	puntata(&ai, n) //L'ai vede sempre
	fmt.Printf("L'ai vede, il piatto è di %d\n", tavolo.fish)

	//flop
	flip(3, m)
	fmt.Printf("Il flop è %s\nQuanto vuoi puntare? ", stampaMazzo(tavolo.mano))
	fmt.Scan(&n)
	puntata(&pietro, n)
	puntata(&ai, n) //L'ai vede sempre
	fmt.Printf("L'ai vede, il piatto è di %d\n", tavolo.fish)

	//turn
	flip(1, m)
	fmt.Printf("Il turn è %s\nQuanto vuoi puntare? ", stampaMazzo(tavolo.mano))
	fmt.Scan(&n)
	puntata(&pietro, n)
	puntata(&ai, n) //L'ai vede sempre
	fmt.Printf("L'ai vede, il piatto è di %d\n", tavolo.fish)

	//river
	flip(1, m)
	fmt.Printf("Il river è %s\nQuanto vuoi puntare? ", stampaMazzo(tavolo.mano))
	fmt.Scan(&n)
	puntata(&pietro, n)
	puntata(&ai, n) //L'ai vede sempre
	fmt.Printf("L'ai vede, il piatto è di %d\n", tavolo.fish)

	fmt.Printf("La tua mano è %s, il tavolo è %s e la mano dell'ai è %s\nChi ha vinto tu[0] o l'ai [1]? ", stampaMazzo(pietro.mano), stampaMazzo(tavolo.mano), stampaMazzo(ai.mano))
	var risultato int
	fmt.Scan(&risultato)
	switch risultato {
	case 0:
		pietro.fish += tavolo.fish
		fmt.Printf("Congratulazioni hai vinto %d fish, attualmente hai %d fish totali\n", tavolo.fish, pietro.fish)
	case 1:
		ai.fish += tavolo.fish
		fmt.Printf("Purtroppo hai perso, le tue fish attuali sono %d", pietro.fish)
	default:
		fmt.Println("ERRORE")
		return
	}
}
