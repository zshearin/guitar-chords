package print

import(
	"fmt"
	"strings"
)



func PrintChords(fretBoard []string) {

	for _, row := range fretBoard {
		fmt.Printf(row)
	}
}


func CombineBoard(curLine []string, newChord []string) []string {

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