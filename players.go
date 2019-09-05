package main

type piece struct {
	id        int
	coord     coord
	endedGame bool
}

type player struct {
	alliance int
	name     string
	path     []coord
	pieces   [7]piece
}

func instantiatePlayer(a int, n string) player {
	p := player{
		alliance: a,
		name:     n,
		path:     getPass(a),
	}

	var pc [7]piece
	for i := 0; i < len(pc); i++ {
		pc[i] = piece{
			id: i + 1,
		}
	}
	p.pieces = pc

	return p
}

func getPass(a int) []coord {
	p := make([]coord, 0)
	// hardcoded paths:
	if a == 1 {
		p = append(p, "A4", "A3", "A2", "A1", "B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8", "A8", "A7", "A6")
	} else if a == 2 {
		p = append(p, "C4", "C3", "C2", "C1", "B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8", "C8", "C7", "C6")
	}
	return p
}

func nextPlayer(currPlayerIdx int, arr []player) int {
	for i, v := range arr {
		if v.alliance == currPlayerIdx {
			if i == len(arr)-1 {
				return arr[0].alliance
			} 
			return arr[i+1].alliance
		} 
	}
	return -1
}
