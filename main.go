package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type state struct {
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

	// p1, pc1 := instantiatePlayer(1, "Nik")
	// fmt.Println(p1)
	// fmt.Println(pc1)

	for state.gameOver == false {
		move(&state)
		fmt.Println("Move ended")
	}
}

func move(st *state) {

	// throw dice
	diceRes := rollDice()
	fmt.Println("Dice roll result:", diceRes)
	//TODO: end move if dice toss result = 0

	//TODO: calculate potential moves

	piece, pieceErr := pickPiece()
	if pieceErr != nil {
		fmt.Println(pieceErr)
	}
	if piece == 999 { //tmp: force exit
		fmt.Println("force exit")
		(*st).gameOver = true
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
