package main

import (
	"fmt"
	"github.com/zshearin/guitar-chords/chords"
	"github.com/zshearin/guitar-chords/print"
	"github.com/zshearin/guitar-chords/format"
	
)


func main() {

	//Get repeated part of verse
	chordsVerseRepeat := []string{"G", "Gsus"}
	board := getBoardForChordsWithRepeat(chordsVerseRepeat, 4)

	//Get non repeated part of verse	
	chordsVerseNonRepeat := []string{"C2", "G/B", "Am7", "Dsus"}
	board2 := getBoardForChords(chordsVerseNonRepeat)
	
	//Combine repeated and non repeated parts
	verse := format.CombineTwoBoards(board,board2, true)
	//Get chorus
	chordsChorus := []string{"C2","G","Am7","Dsus"}
	chorus := getBoardForChordsWithRepeat(chordsChorus, 2)
	
	chordsBridgeLine1 := []string{"C2", "G", "Am7", "G"}

	chordsBridgeLine1SingleMeasure := []string{"Dsus"}
	singleMeasure := getBoardForChords(chordsBridgeLine1SingleMeasure)

	chordsBridgeLine2 := []string{"C2", "G", "Dsus", "Dsus"}

	bridgeLine1 := getBoardForChords(chordsBridgeLine1)
	bridgeLine2 := getBoardForChords(chordsBridgeLine2)

	bridgeLine1 = format.CombineTwoBoards(bridgeLine1, singleMeasure, false)

	fmt.Println("Verse")
	print.PrintLine(verse)
	fmt.Println("\n\nChorus")
	print.PrintLine(chorus)
	fmt.Println("\n\nBridge")
	print.PrintLine(bridgeLine1)
	print.PrintLine(bridgeLine2)
}	

func getBoardForChords(chordNames []string) []string {
	return getBoardForChordsWithRepeat(chordNames, 1)
}

func getBoardForChordsWithRepeat(chordNames []string, val int) []string {

	listOfChordPrints := chords.GetRowFromChordArray(chordNames)
	aggregatedChordPrints := format.CombineChordPrints(listOfChordPrints)

	if !(val == 1) {
		aggregatedChordPrints = format.AddRepeatLine(aggregatedChordPrints, val)
	}

	return aggregatedChordPrints
}

