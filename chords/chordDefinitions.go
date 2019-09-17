package chords

type fingerPlacement struct {
	finger, string, fret int	
}

type chord struct {
	fingerPlacements []fingerPlacement
	stringsToPlay [6]int
}

var chordDict map[string]chord

func init() {

	/*
	chordDict["A"] = chord{
		finger //so on so forth
	}
	*/

	chordDict = map[string]chord{
		"A": chord{
			[]fingerPlacement{
				fingerPlacement{
					finger: 1,
					string: 3,
					fret:	2,
				},
				fingerPlacement{
					finger: 2,
					string: 4,
					fret:	2,
				},
				fingerPlacement{
					finger: 3,
					string: 5,
					fret:	2,
				},
			},
			[6]int{0,1,1,1,1,1},
		},
		"C": chord{
			[]fingerPlacement{
				fingerPlacement{
					finger: 1,
					string: 5,
					fret:	1,
				},
				fingerPlacement{
					finger: 2,
					string: 3,
					fret:	2,
				},
				fingerPlacement{
					finger: 3,
					string: 2,
					fret:	3,
				},
			},
			[6]int{0,1,1,1,1,1},
		},
		"D": chord{
			[]fingerPlacement{
				fingerPlacement{
					finger: 1,
					string: 4,
					fret:	2,
				},
				fingerPlacement{
					finger: 2,
					string: 6,
					fret:	2,
				},
				fingerPlacement{
					finger: 3,
					string: 5,
					fret:	3,
				},
			},
			[6]int{0,0,1,1,1,1},
		},
		
		"G": chord{
			[]fingerPlacement{
				fingerPlacement{
					finger: 2,
					string: 1,
					fret:	3,
				},
				fingerPlacement{
					finger: 3,
					string: 5,
					fret:	3,
				},
				fingerPlacement{
					finger: 4,
					string: 6,
					fret:	3,
				},
			},
			[6]int{1,0,1,1,1,1},
		},
	}
	

}

