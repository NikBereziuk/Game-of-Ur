package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type possibleMovesValues struct {
	destCoord coord
	errMsg    string
}

type possibleMoves map[piece]possibleMovesValues

func pickPiece() (int, error) {
	// ask player to pick a piece
	fmt.Print("Enter piece ID (int 1 - 7): -> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() == true {
		ans := scanner.Text()
		pID, err := strconv.Atoi(ans)
		if err == nil && ((pID > 0 && pID <= 7) || pID == 999) {
			return pID, nil
		} else if err == nil && (pID < 0 || pID > 7) {
			fmt.Println("Piece ID is out of bounds. Try again")
		} else if err != nil {
			fmt.Println("Not an integer. Try again")
		}
		fmt.Print("Enter piece ID (int 1 - 7): -> ")
	}
	return 0, errors.New("scanner.Scan failed")
}

func precalcMoves(t int, st *state, c chan possibleMoves) {
	pmMap := make(possibleMoves)

	//get path:
	path := (*st).players[(*st).currPlayer-1].path

	//loop through pieces:
	for _, pc := range (*st).players[(*st).currPlayer-1].pieces {
		newPathArrIdx := 0
		//Step 1. Get destination square
		if pc.coord == "" {
			newPathArrIdx = t - 1
		} else {
			for i, val := range path {
				if val == pc.coord {
					newPathArrIdx = i + t
				}
			}
		}
		if newPathArrIdx > len(path)-1 {
			pmMap[pc] = possibleMovesValues{errMsg: "You can't move this piece that far"}
			continue
		}
		newCoord := path[newPathArrIdx]
		//Step 2. See if the destination square is not occupied
		square := (*st).board[newCoord]
		if square.piece.id != 0 && square.piece.alliance == (*st).currPlayer {
			pmMap[pc] = possibleMovesValues{destCoord: newCoord, errMsg: "Destination square is occupied by your piece"}
		} else if (square.isWarzone == true && square.piece.alliance != (*st).currPlayer) || square.piece.id == 0 {
			pmMap[pc] = possibleMovesValues{destCoord: newCoord}
		}
	}
	c <- pmMap
}

func movePiece(pc piece, dest coord, st *state) {
	fmt.Println("Piece", pc.id, "is moving to", dest)

	square := (*st).board[dest]

	fmt.Printf("%v %+v", square, square)

	//Step 1. Remove enemy piece
	if square.isWarzone && square.piece.alliance != (*st).currPlayer {
		// TODO: move enemy piece
		// (*st).board[dest].piece = ""
	}
	//Step 2. Move piece
	// (*st).board[dest].piece = pc
}
