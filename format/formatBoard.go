package format

import(
	"strings"
	"strconv"
)

const chordEndMeasure = "|  "

/*
outer loop - lines to print

inner loop - iterates through and concatenates each line, then assigns to newBoard
*/
func CombineChordPrints(boards [][]string) []string {

	var newBoard [11]string

	//for each line in the new board
	for lineNum := 0; lineNum < len(newBoard); lineNum++ {
		
		
		var sb strings.Builder
		sb.Reset()
		//for each 		
		for boardNum := 0; boardNum < len(boards); boardNum++ {

			//j - the board 
			curString := boards[boardNum][lineNum]
			curString = strings.TrimSuffix(curString, "\n")

			sb.WriteString(curString)
			sb.WriteString("  ")
			sb.WriteString(chordEndMeasure)
		}
		sb.WriteString("\n")
		newBoard[lineNum] = sb.String()
		sb.Reset()
	}
	return newBoard[:]
}

/*


/  === C2  ===      ===  G  ===      === Am7 ===      ===Dsus === 2x \
|  E A D G B e      E A D G B e      E A D G B e      E A D G B e    |
|  X=O=O=O=O=O      O=X=O=O=O=O      X=O=O=O=O=O      X=X=O=O=O=O    |
|  | | | | | |      | | | | | |      | | | | 1 |      | | | | | |    |
|  | | 1 | | |      | | | | | |      | | 2 | | |      | | | 1 | |    |
|  | 2 | | 3 4      2 | | | 3 4      | | | | | 4      | | | | 3 4    |
|  | | | | | |      | | | | | |      | | | | | |      | | | | | |    |
|  | | | | | |      | | | | | |      | | | | | |      | | | | | |    |
|  | | | | | |      | | | | | |      | | | | | |      | | | | | |    |
|  | | | | | |      | | | | | |      | | | | | |      | | | | | |    |
\  | | | | | |      | | | | | |      | | | | | |      | | | | | |    /


*/

func AddRepeatLine(fretBoard []string, repeatTimes int) []string {


	for i, val := range fretBoard {
		var sb strings.Builder
		val = strings.TrimSuffix(val,"\n")
		str1 := val[0:len(val)-5]
		if (i == 0) {
			sb.WriteString("/  ")
			sb.WriteString(str1)
			sb.WriteString(" ")
			sb.WriteString(" \\")
			sb.WriteString(strconv.Itoa(repeatTimes))
			sb.WriteString("x")
			} else if (i ==10) {
			sb.WriteString("\\  ")
			sb.WriteString(str1)
			sb.WriteString("  /  ")			
		} else {
			sb.WriteString("|  ")
			sb.WriteString(str1)
			sb.WriteString("  ")
			sb.WriteString(chordEndMeasure)
		}
		sb.WriteString("   \n")
		fretBoard[i] = sb.String()
		sb.Reset()
	}
	return fretBoard
}




func CombineTwoBoards(board1 []string, board2 []string, measureSeparator bool) []string {

	var sb strings.Builder
	var newBoard [11]string

	for i := range board1 {
		str1 := board1[i]
		str2 := board2[i]
		str1 = strings.TrimSuffix(str1, "\n")
		str1 = strings.TrimSuffix(str1, "   ")
		if (measureSeparator) {
			sb.WriteString(str1)
			sb.WriteString(" ")
			sb.WriteString(chordEndMeasure)
		} else {

			str1 = strings.TrimSuffix(str1, chordEndMeasure)
			sb.WriteString(str1)

		}
		sb.WriteString(" ")
		sb.WriteString(str2)
 
		newBoard[i] = sb.String()
		sb.Reset()
	}

	return newBoard[:]
}