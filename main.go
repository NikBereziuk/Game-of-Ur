package main

import (
	"fmt"
	"os"
)

type state struct {
	players    []player
	board      board
	currPlayer int
	gameOver   bool
}

func main() {

	state := state{
		currPlayer: 1,
	}

	fmt.Println("Game started.")

	state.board = instantiateBoard()

	players := make([]player, 0) //slice of players
	players = append(players, instantiatePlayer(1, "Nik"), instantiatePlayer(2, "Olmer"))
	state.players = players

	for state.gameOver == false {
		move(&state)
		fmt.Println("Move ended. Next player:", state.players[state.currPlayer-1].name)
		fmt.Println("Board state:", state.board)
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

	// precalculate moves concurrently, via channel
	precalc := make(chan possibleMoves)
	go precalcMoves(toss, st, precalc)

	isViableMove := false
	for isViableMove == false {
		pieceID, pieceErr := pickPiece()
		if pieceErr != nil {
			fmt.Println(pieceErr)
			os.Exit(888)
		}
		if pieceID == 999 { //tmp: force exit
			fmt.Println("force exit")
			(*st).gameOver = true
			os.Exit(999)
		}

		possibleMoves := <-precalc
		fmt.Printf("%+v", possibleMoves)

		//check if the picked piece can actually move
		pickedPiece := (*st).players[(*st).currPlayer-1].pieces[pieceID-1]
		for key, val := range possibleMoves {
			if key == pickedPiece && val.errMsg == "" {
				movePiece(pickedPiece, val.destCoord, st)
				isViableMove = true
				break
			}
		}
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
