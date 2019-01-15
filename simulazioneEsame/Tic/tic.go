package main

import (
	"fmt"
	"os"
	"strconv"
)

type Orario struct {
	secondi int
	minuti  int
	ore     int
}

func newOrario(s, m, o int) (bool, Orario) {
	if s < 0 || s > 59 {
		return false, Orario{0, 0, 0}
	}
	if m < 0 || m > 59 {
		return false, Orario{0, 0, 0}
	}
	if o < 0 || o > 23 {
		return false, Orario{0, 0, 0}
	}
	return true, Orario{s, m, o}
}
func tic(orario *Orario) {
	orario.secondi++
	if orario.secondi == 60 {
		orario.secondi = 0
		orario.minuti++
		if orario.minuti == 60 {
			orario.minuti = 0
			orario.ore++
			if orario.ore == 24 {
				orario.ore = 0
			}
		}
	}
}
func StringOrario(orario Orario) string {
	return fmt.Sprintf("%d:%d:%d", orario.secondi, orario.minuti, orario.ore)
}
func main() {
	var s, m, o, piuS int
	s, _ = strconv.Atoi(os.Args[1])
	m, _ = strconv.Atoi(os.Args[2])
	o, _ = strconv.Atoi(os.Args[3])
	piuS, _ = strconv.Atoi(os.Args[4])
	err, orario := newOrario(s, m, o)
	if !err {
		fmt.Println("parametri non validi")
		return
	}
	fmt.Println(StringOrario(orario))
	for i := 0; i < piuS; i++ {
		tic(&orario)
	}
	fmt.Println(StringOrario(orario))
}
