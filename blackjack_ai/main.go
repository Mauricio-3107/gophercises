package main

import (
	"fmt"

	"github.com/Mauricio-3107/gophercises.git/blackjack_ai/blackjack"
)

func main() {
	opts := blackjack.Options{
		Hands: 2,
	}
	game := blackjack.New(opts)
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
