package chords

var chordDictSus map[string]chord

func init() {

	chordDictSus = map[string]chord{
		
		"Gsus": chord{ 
				[]fingerPlacement{ 
					fingerPlacement{ finger: 1, string: 5, fret: 1, }, 
					fingerPlacement{ finger: 2, string: 1, fret: 3, },
					fingerPlacement{ finger: 4, string: 6, fret: 3, },
			}, [6]int{0,1,1,1,1,1},
		},

		
		"Dsus": chord{ 
			[]fingerPlacement{ 
					fingerPlacement{ finger: 1, string: 4, fret: 2, }, 
					fingerPlacement{ finger: 3, string: 5, fret: 3, },
					fingerPlacement{ finger: 4, string: 6, fret: 3, },
			}, [6]int{0,0,1,1,1,1},
		},
	}
}

