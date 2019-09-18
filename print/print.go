package print

import(
	"fmt"
//	"strings"
//	"strconv"
	"github.com/zshearin/guitar-chords/format"
)


func CombineAndPrintLine(boards [][]string) []string {
	newBoard := format.CombineChordPrints(boards)
	PrintLine(newBoard)
	return newBoard
}

func PrintLine(fretBoard []string) {

	for _, row := range fretBoard {
		fmt.Printf(row)
	}
	fmt.Printf("\n")
}
