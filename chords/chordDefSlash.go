package chords

var chordDictSlash map[string]chord

func init() {

	chordDictSlash = map[string]chord{
		
		"G/B": chord{ 
				[]fingerPlacement{ 
					fingerPlacement{ finger: 1, string: 2, fret: 2, }, 
					fingerPlacement{ finger: 3, string: 5, fret: 3, },
					fingerPlacement{ finger: 4, string: 6, fret: 3, },
			}, [6]int{0,1,1,1,1,1},
		},
		
	}
}

