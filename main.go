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

func main() {

	state := state{
		currPlayer: 1,
	}

	fmt.Println("Test begin")

	// b := instantiateBoard()
	// fmt.Println(b)

	players := make([]player, 0) //slice of players
	players = append(players, instantiatePlayer(1, "Nik"), instantiatePlayer(2, "Olmer"))
	state.players = players
	fmt.Println(players)

	for state.gameOver == false {
		move(&state)
		fmt.Println("Move ended. Next player:", state.currPlayer)
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

	//TODO: calculate potential moves

	piece, pieceErr := pickPiece()
	if pieceErr != nil {
		fmt.Println(pieceErr)
	}
	if piece == 999 { //tmp: force exit
		fmt.Println("force exit")
		(*st).gameOver = true
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
		if err == nil {
			return pID, nil
		}
		fmt.Println("Not an integer. Try again")
		fmt.Print("Enter piece ID (int 1 - 7): -> ")
	}
	return 0, errors.New("scanner.Scan failed")
}
