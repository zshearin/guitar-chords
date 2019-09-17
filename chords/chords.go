package chords

import(
	"github.com/zshearin/guitar-chords/chords/board"
	"strings"
)

/*
func PrintChordSequence(chordNames []string) {
	var currentBoard []string
	


	for _, chordName := range chordNames {
		newBoard := GetBoard(chordName)
		
		newBoard = board.CombineBoards(currentBoard,newBoard)
		currentBoard = newBoard
	}

	board.PrintBoard(currentBoard)

}
*/

/*
func PrintChord(chordName string) {

	newBoard := GetBoard(chordName)

	board.PrintBoard(newBoard)

}
*/
func getChord(chordName string) chord {
	return chordDict[chordName]
}

func addFingerPlacements(theChord chord, fretBoard []string) []string {

	for _, place := range theChord.fingerPlacements {
		fretBoard = board.AddFingerPlacement(fretBoard,place.string-1,place.fret-1,place.finger)
	}
	return fretBoard

}

func addStringsToPlay(theChord chord, fretBoard []string) []string {
	var sb strings.Builder
	for i, val := range theChord.stringsToPlay {
		
		if (val == 1) {
			sb.WriteString("O")
		} else {
			sb.WriteString("X")
		}

		if (i != 5) {
			sb.WriteString("=")
		}
	}
	sb.WriteString("\n")

	fretBoard[2] = sb.String()
	sb.Reset()
	return fretBoard

}
func GetChordToPrint(chordName string) []string {
	chord := getChord(chordName)
	fretBoard := board.CreateEmptyBoard()
	newBoard := addFingerPlacements(chord, fretBoard)
	newBoard = addStringsToPlay(chord, newBoard)
	return newBoard
}


