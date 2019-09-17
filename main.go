package main

import (
	"github.com/zshearin/guitar-chords/chords"
	
	"strings"
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
	
	/*	
	chords.PrintChords(aChord)
	chords.PrintChords(dChord)
	*/

	newBoard := combineBoard(aChord, dChord)
	//chords.PrintChords(newBoard)
	newBoard2 := combineBoard(newBoard, gChord)
	chords.PrintChords(newBoard2)

}


func combineBoard(curLine []string, newChord []string) []string {

	var newBoard [11]string

	for i := 0; i < len(newBoard); i++ {

		string1 := curLine[i]
		string2 := newChord[i]

		string1 = strings.TrimSuffix(string1, "\n")
		string2 = strings.TrimSuffix(string2, "\n")
		
		var sb strings.Builder
		
		sb.WriteString(string1)
		sb.WriteString("     ")
		sb.WriteString(string2)
		sb.WriteString("\n")
		newBoard[i] = sb.String()
		sb.Reset()
	}

	return newBoard[:]

}