package main

import (
	"math/rand"
	"time"
)

type dice [4]int

func rollDice() dice {
	var d dice

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := rand.Intn(9999)

	for i := 0; i < len(d); i++ {
		remaind := result % 10
		if remaind%2 == 0 {
			d[i] = 0
		} else {
			d[i] = 1
		}
		result = result / 10
	}

	return d
}
