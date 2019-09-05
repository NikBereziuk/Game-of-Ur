package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type state struct {
	players    []player
	board      board
	currPlayer int
	gameOver   bool
}

type possibleMovesValues struct {
	destCoord coord
	errMsg    string
}

type possibleMoves map[piece]possibleMovesValues

func main() {

	state := state{
		currPlayer: 1,
	}

	fmt.Println("Test begin")

	state.board = instantiateBoard()

	players := make([]player, 0) //slice of players
	players = append(players, instantiatePlayer(1, "Nik"), instantiatePlayer(2, "Olmer"))
	state.players = players
	fmt.Println(players)

	for state.gameOver == false {
		move(&state)
		fmt.Println("Move ended. Next player:", state.players[state.currPlayer-1].name)
	}
}

func move(st *state) {

	// throw dice
	diceRes := rollDice()
	fmt.Println("Dice roll result:", diceRes)
	toss := diceRes.tossSum()
	//end move if dice toss result = 0
	if toss == 0 {
		(*st).currPlayer = nextPlayer((*st).currPlayer, (*st).players)
		return
	}

	possibleMoves := precalcMoves(toss, st)
	fmt.Println("Possible moves", "%+v\n", possibleMoves)

	pieceID, pieceErr := pickPiece()
	if pieceErr != nil {
		fmt.Println(pieceErr)
	}
	if pieceID == 999 { //tmp: force exit
		fmt.Println("force exit")
		(*st).gameOver = true
		os.Exit(999)
	}

	/* End of move: switch players.
	Exceptions:
	1) Piece was moved on a rosette
	2) Toss result: 4
	*/
	if toss != 4 {
		(*st).currPlayer = nextPlayer((*st).currPlayer, (*st).players)
	}

}

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

func precalcMoves(t int, st *state) (pm possibleMoves) {
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
	return pmMap
}
