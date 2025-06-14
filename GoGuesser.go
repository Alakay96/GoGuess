// https://roadmap.sh/projects/number-guessing-game
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randInt(1, 100)

	fmt.Print("Welcome to GO Guessing game")
	fmt.Print("Guess a number between 1 and 100")
	fmt.Print("10 Guesses")

}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}
