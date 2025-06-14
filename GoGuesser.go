// https://roadmap.sh/projects/number-guessing-game
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	theNumber := randInt(1, 100)
	var guess int
	var numOfGuess int

	fmt.Println("Welcome to GO Guessing game")
	fmt.Println("Guess a number between 1 and 100")
	fmt.Println("10 Guesses")

	fmt.Print("Enter your guess: ")
	_, err := fmt.Scan(&guess)

	for i := 0; i < 10; i++ {
		if err != nil {
			fmt.Println("Invalid input")
		} else if guess < 0 || guess > 100 {
			fmt.Printf("Invalid number, %d", guess)
		} else {
			if guess == theNumber {
				fmt.Printf("You got it %d in %d", theNumber, numOfGuess)
				break
			} else {
				fmt.Println("Incorrect")
				numOfGuess += 1
				fmt.Print("Enter your guess: ")
				fmt.Scan(&guess)
			}
		}

	}

}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}
