package board


import(
	"fmt"
	"strconv"
)


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

func PrintLines(fretBoard []string) {

	for _, row := range fretBoard {
		fmt.Printf(row)
	}
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