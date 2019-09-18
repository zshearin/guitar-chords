package chords

var chordDictMinor7 map[string]chord

func init() {

	chordDictMinor7 = map[string]chord{
		
		"Am7": chord{ 
				[]fingerPlacement{ 
					fingerPlacement{ finger: 1, string: 5, fret: 1, }, 
					fingerPlacement{ finger: 2, string: 3, fret: 2, },
					fingerPlacement{ finger: 4, string: 6, fret: 3, },
			}, [6]int{0,1,1,1,1,1},
		},
		
	}
}

