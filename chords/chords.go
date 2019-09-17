package chords

import(
	"strings"
	"strconv"
)

func GetChordToPrint(chordName string) []string {
	chord := getChord(chordName)
	fretBoard := CreateEmptyBoard()
	newBoard := addFingerPlacements(chord, fretBoard)
	newBoard = addStringsToPlay(chord, newBoard)
	return newBoard
}

func getChord(chordName string) chord {
	return chordDict[chordName]
}

func addFingerPlacements(theChord chord, fretBoard []string) []string {

	for _, place := range theChord.fingerPlacements {
		fretBoard = AddFingerPlacement(fretBoard,place.string-1,place.fret-1,place.finger)
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


func CreateEmptyBoard() []string {

	var fretBoard [11]string
	
	fretBoard[0] = "===========\n"
	fretBoard[1] = "E A D G B e\n"
	fretBoard[2] = "===========\n"
	
	emptyLine := "| | | | | |\n"
	for i := 3; i < len(fretBoard); i++ {
		fretBoard[i] = emptyLine
	}
	newBoard := fretBoard[:]
	return newBoard

}

/*
fretBoard - current state of the board
*/
func AddFingerPlacement(fretBoard []string, x, y, val int) []string {

	curString := fretBoard[y+3]
	curStringBytes := []byte(curString)

	index1 := 2*x
	byteArray1 := curStringBytes[:index1]
	index2 := index1 + 1

	str := strconv.Itoa(val)
	byteArray2 := []byte(str)

	byteArray3 := curStringBytes[index2:]
	fullArray := append(byteArray1)

	for _, value := range byteArray2 {
		fullArray = append(fullArray, value)
	}
	for _, value := range byteArray3 {
		fullArray = append(fullArray, value)
	}
	fretBoard[y+3] = string(fullArray)
	return fretBoard
}