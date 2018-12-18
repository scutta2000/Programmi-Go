package main

import "fmt"

func main() {
	var x1a, x2a, y1a, y2a, x1b, x2b, y1b, y2b float64
	fmt.Print("Inserisci le coordinate del angolo in basso a sinistra del primo rettangolo ")
	fmt.Scan(&x1a, &y1a)
	fmt.Print("Inserisci le coordinate dell'angolo in altro a destra ")
	fmt.Scan(&x2a, &y2a)

	fmt.Print("Inserisci le coordinate del angolo in basso a sinistra del secondo rettangolo ")
	fmt.Scan(&x1b, &y1b)
	fmt.Print("Inserisci le coordinate dell'angolo in altro a destra ")
	fmt.Scan(&x2b, &y2b)

	//calcolo areee
	var areaA, areaB float64
	areaA = (x1a - x2a) * (y1a - y2a)
	areaB = (x1b - x2b) * (y1b - y2b)
	fmt.Println("Le aree valgono rispettivamente: ", areaA, areaB)

	//confronto aree
	if areaA > areaB {
		fmt.Println("La prima area è più grande della seconda")
	} else if areaA == areaB {
		fmt.Println("Le due aree sono uguali")
	} else {
		fmt.Println("La seconda area è più grande della seconda")
	}

	//sovrapposizione

	/*
		1 è l'angolo in basso a sinistra
		2 è l'angolo in alto a destra
		3 è l'angolo in basso a destra
		4 è l'angolo in alto a sinistra
	*/
	var b1Compreso, b2Compreso, b3Compreso, b4Compreso, a1Compreso, a2Compreso, a3Compreso, a4Compreso bool

	b1Compreso = (x1b > x1a && x1b < x2a) && (y1b > y1a && y1b < y2a)
	b2Compreso = (x2b > x1a && x2b < x2a) && (y2b > y1a && y2b < y2a)
	b3Compreso = (x2b > x1a && x2b < x2a) && (y1b > y1a && y1b < y2a)
  b4Compreso = (x1b > x1a && x1b < x2a) && (y2b > y1a && y2b < y2a)
	a1Compreso = (x1a > x1b && x1a < x2b) && (y1a > y1b && y1a < y2b)
	a2Compreso = (x2a > x1b && x2a < x2b) && (y2a > y1b && y2a < y2b)
  a3Compreso = (x2a > x1b && x2a < x2b) && (y1a > y1b && y1a < y2b)
	a4Compreso = (x1a > x1b && x1a < x2b) && (y2a > y1b && y2a < y2b)

	if b1Compreso || b2Compreso || b3Compreso || b4Compreso || a1Compreso || a2Compreso {
		if b1Compreso && b2Compreso && b3Compreso && b4Compreso {
			fmt.Println("Il secondo triangolo è contenuto del primo")
			fmt.Println("Quindi l'area dell'intersezzione è uguale all'area del secondo: ", areaB)
		} else if a1Compreso && a2Compreso && a3Compreso && a4Compreso {
			fmt.Println("Il primo triangolo è contenuto del secondo")
			fmt.Println("Quindi l'area dell'intersezzione è uguale all'area del primo: ", areaA)
		} else {
			fmt.Println("I due triangoli sono sovrapposti ma uno non è contenuto nell'altro")
			if a1Compreso && !a2Compreso && !a3Compreso && !a4Compreso {
				fmt.Println("E l'area compresa è ", (x1a-x2b)*(y1a-y2b))
			} else if !a1Compreso && a2Compreso && !a3Compreso && !a4Compreso {
				fmt.Println("E l'area compresa è ", (x2a-x1b)*(y2a-y1b))
			} else if !a1Compreso && !a2Compreso && a3Compreso && !a4Compreso {
				fmt.Println("E l'area compresa è ", (x2a-x1b)*(y2b-y1a))
			} else if !a1Compreso && !a2Compreso && !a3Compreso && a4Compreso {
				fmt.Println("E l'area compresa è ", (x2b-x1a)*(y2a-y1b))
			}else if b1Compreso && b4Compreso{
        fmt.Println("E l'area compresa è ",(x2a-x1b)*(y2b-y1b))
      }else if b2Compreso && b3Compreso{
        fmt.Println("E l'area compresa è ",(x2b-x1a)*(y2b-y1b))
      }

		}
	} else {
		fmt.Print("I due rettangoli non sono sovrapposti")
	}
}
