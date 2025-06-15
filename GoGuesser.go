// https://roadmap.sh/projects/number-guessing-game
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	theNumber := randInt(1, 100)
	// var guess int
	// var numOfGuess int
	var difficulty int

	fmt.Println("Welcome to GO Guessing game")
	fmt.Println("Guess a number between 1 and 100")
	fmt.Println("Select your difficulty")
	fmt.Println("1. Easy 10 guesses")
	fmt.Println("2. Medium 5 guesses")
	fmt.Println("3. Hard 3 guesses")
	_, err := fmt.Scan(&difficulty)

	// TODO: make it loop instead of a single game
	for {
		if err != nil {
			fmt.Println("Invalid input")
		} else if difficulty < 0 || difficulty > 3 {
			fmt.Println("Invalid number")
		} else if difficulty == 1 {
			guessGame(theNumber, 10)
			break
		} else if difficulty == 2 {
			guessGame(theNumber, 5)
			break
		} else if difficulty == 3 {
			guessGame(theNumber, 3)
			break
		}

	}
	// fmt.Print("Enter your guess: ")
	// _, err = fmt.Scan(&guess)

	// for range 10 {
	// 	if err != nil {
	// 		fmt.Println("Invalid input")
	// 	} else if guess < 0 || guess > 100 {
	// 		fmt.Printf("Invalid number, %d", guess)
	// 	} else {
	// 		if guess == theNumber {
	// 			fmt.Printf("You got it %d in %d", theNumber, numOfGuess)
	// 			break
	// 		} else {
	// 			fmt.Println("Incorrect")
	// 			numOfGuess += 1
	// 			fmt.Print("Enter your guess: ")
	// 			fmt.Scan(&guess)
	// 		}
	// 	}
	// }

	// fmt.Printf("unlucky the number was %d", theNumber)

}

func guessGame(theNumber, difficultyGuesses int) {
	var guess int
	var numOfGuess int

	fmt.Print("Enter your guess: ")
	_, err := fmt.Scan(&guess)
	for range difficultyGuesses {
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
				if numOfGuess == difficultyGuesses {
					break
				} else {
					fmt.Print("Enter your guess: ")
					fmt.Scan(&guess)
				}
			}

		}

	}
	fmt.Printf("unlucky the number was %d", theNumber)
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}
