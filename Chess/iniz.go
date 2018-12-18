package main

func inizPedoni(s scacchiera){
	for i:=0;i<8;i++{
		s[casella{1,i}]=pezzo{"♙",true}
		s[casella{6,i}]=pezzo{"♟️",false}

	}
}
func inizTorri(s scacchiera){
	s[casella{0,0}]=pezzo{"♖",true}
	s[casella{0,7}]=pezzo{"♖",true}
	s[casella{7,0}]=pezzo{"♜",false}
	s[casella{7,7}]=pezzo{"♜",false}
}
func inizCavalli(s scacchiera){
	s[casella{0,1}]=pezzo{"♘",true}
	s[casella{0,6}]=pezzo{"♘",true}
	s[casella{7,1}]=pezzo{"♞",false}
	s[casella{7,6}]=pezzo{"♞",false}
}
func inizAlfieri(s scacchiera){
	s[casella{0,2}]=pezzo{"♗",true}
	s[casella{0,5}]=pezzo{"♗",true}
	s[casella{7,2}]=pezzo{"♝",false}
	s[casella{7,5}]=pezzo{"♝",false}
}
func inizReRegina(s scacchiera){
	s[casella{0,3}]=pezzo{"♕",true}
	s[casella{0,4}]=pezzo{"♔",true}
	s[casella{7,3}]=pezzo{"♛",false}
	s[casella{7,4}]=pezzo{"♚",false}
}
func iniz(s scacchiera){
	inizPedoni(s)
	inizTorri(s)
	inizCavalli(s)
	inizAlfieri(s)
	inizReRegina(s)
}
