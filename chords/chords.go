package chords

import(
	"strings"
	"strconv"
)

type ChordPrint struct {
	Name string
	Value []string
}


func GetRowFromChordArray(chordNames []string) [][]string {
	var chordLine [][]string
	for _, val := range chordNames {

		curChord := GetChordToPrint(val)
		chordLine = append(chordLine, curChord.Value)

	}

	return chordLine
}

func GetChordToPrint(chordName string) ChordPrint {
	var chordToReturn ChordPrint
	chordToReturn.Name = chordName
	
	chord := getChord(chordName)
	newBoard := CreateEmptyBoard()
	newBoard = addFingerPlacements(chord, newBoard)
	newBoard = addStringsToPlay(chord, newBoard)
	newBoard = addNoteName(chordName, newBoard)
	
	chordToReturn.Value = newBoard
	return chordToReturn
}

func getChord(chordName string) chord {

	if  strings.Contains(chordName, "sus") {
		return chordDictSus[chordName]

	} else if strings.Contains(chordName, "/") {
		return chordDictSlash[chordName]
	} else if strings.Contains(chordName, "m7") {
		return chordDictMinor7[chordName]	
	} else {
		return chordDict[chordName]

	}
}

func addFingerPlacements(theChord chord, fretBoard []string) []string {

	for _, place := range theChord.fingerPlacements {
		fretBoard = addFingerPlacement(fretBoard,place.string-1,place.fret-1,place.finger)
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
func addFingerPlacement(fretBoard []string, x, y, val int) []string {

	curString := fretBoard[y+3]
	curStringBytes := []byte(curString)

	byteArrayBefore := curStringBytes[:(2*x)]
	byteArrayAfter := curStringBytes[(2*x+1):]

	str1 := string(byteArrayBefore)
	str2 := strconv.Itoa(val)
	str3 := string(byteArrayAfter)

	var sb strings.Builder
	sb.WriteString(str1)
	sb.WriteString(str2)
	sb.WriteString(str3)

	fretBoard[y+3] = sb.String()
	sb.Reset()
	return fretBoard
}

func addNoteName(noteName string, fretBoard []string) []string {

	var sb strings.Builder

	//Formatting based on length
	//===========	
	//===  G  ===
	//=== C2  ===
	//=== Am7 ===
	//===Dsus ===
	//===========	
	
	if (len(noteName) == 1) {
		sb.WriteString("===  ")
		sb.WriteString(noteName)
		sb.WriteString("  ===")
	} else if (len(noteName) == 2) {
		sb.WriteString("=== ")
		sb.WriteString(noteName)
		sb.WriteString("  ===")

	} else if (len(noteName) == 3) {
		sb.WriteString("=== ")
		sb.WriteString(noteName)
		sb.WriteString(" ===")

	} else if (len(noteName) == 4) {
		sb.WriteString("===")
		sb.WriteString(noteName)
		sb.WriteString(" ===")

	} else if (len(noteName) == 5) {
		sb.WriteString("===")
		sb.WriteString(noteName)
		sb.WriteString("===")
	
	} else {
		sb.WriteString("===========")
	}
	sb.WriteString("\n")
	fretBoard[0] = sb.String()
	sb.Reset()
	return fretBoard
}


