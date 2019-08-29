package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type state struct {
	currPlayer int
}

func main() {

	state := state{
		currPlayer: 1,
	}

	fmt.Println("Test begin")

	// b := instantiateBoard()
	// fmt.Println(b)

	p1, pc1 := instantiatePlayer(1, "Nik")
	fmt.Println(p1)
	fmt.Println(pc1)

	move(&state)

}

func move(st *state) {

	fmt.Println(st.currPlayer)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

	}

	// tossRes := 3 //tmp/

}
