package main

import (
	"github.com/zshearin/guitar-chords/chords"
	"github.com/zshearin/guitar-chords/print"
	
)

/*
type boardToPrint struct {
	chordPrint []string
}

type songLine struct {
	boardList []boardToPrint
}

type song struct {
	name string
	lines []songLine
}
*/

func main() {

	var aChord, dChord, gChord []string

	aChord = chords.GetChordToPrint("A")
	dChord = chords.GetChordToPrint("D")
	gChord = chords.GetChordToPrint("G")
	
	newBoard := print.CombineBoard(aChord, dChord)
	newBoard2 := print.CombineBoard(newBoard, gChord)
	newBoard3 := print.CombineBoard(newBoard2, dChord)
	print.PrintChords(newBoard3)

}


