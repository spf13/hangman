package main

import (
	"fmt"
	"strings"
)

func main() {
	// Welcome message
	fmt.Println("Welcome to Hangman!")
	fmt.Println("Can you guess the secret word before the hangman is complete?")
	fmt.Println()

	// Game settings
	secretWord := "programming" // Teachers can change this to any word
	wrongGuesses := 0
	maxWrongGuesses := 6

	// String to track what to display for the word
	displayWord := ""
	for i := 0; i < len(secretWord); i++ {
		displayWord = displayWord + "*"
	}

	// Keep track of all guesses
	allGuesses := ""

	// Game loop
	gameOver := false
	for gameOver == false {
		// 1. Display current word state
		fmt.Print("Current word: ")
		for i := 0; i < len(displayWord); i++ {
			fmt.Print(string(displayWord[i]) + " ")
		}
		fmt.Println()

		// 2. Show progress
		fmt.Printf("Wrong guesses: %d/%d\n", wrongGuesses, maxWrongGuesses)
		fmt.Printf("Letters you've guessed: %s\n", allGuesses)
		drawHangmanSwitch(wrongGuesses)

		// 3. Check win condition
		if !strings.Contains(displayWord, "*") {
			fmt.Println("Congratulations! You guessed the word!")
			gameOver = true
			break
		}

		// 4. Check lose condition
		if wrongGuesses >= maxWrongGuesses {
			fmt.Println("Sorry, you lost! The word was:", secretWord)
			gameOver = true
			break
		}

		// 5. Get player's guess
		fmt.Print("Guess a letter: ")
		var guess string
		fmt.Scan(&guess)

		// 6. Handle empty input
		if len(guess) == 0 {
			fmt.Println("Please enter a letter!")
			fmt.Println()
			continue
		}

		// 7. Take just the first character of input
		guessChar := string(guess[0])

		// 8. Check if letter was already guessed
		if strings.Contains(allGuesses, guessChar) {
			fmt.Println("You already guessed that letter!")
			fmt.Println()
			continue
		}

		// 9. Add to guesses
		allGuesses = allGuesses + guessChar + " "

		// 10. Check if guess is in the word and update display
		correctGuess := false
		newDisplay := ""

		for i := 0; i < len(secretWord); i++ {
			currentChar := string(secretWord[i])

			if currentChar == guessChar {
				// Add the guessed letter to the display
				newDisplay = newDisplay + currentChar
				correctGuess = true
			} else {
				// Keep whatever was already in that position
				newDisplay = newDisplay + string(displayWord[i])
			}
		}

		// Update the display word
		displayWord = newDisplay

		// 11. Record result of guess
		if correctGuess {
			fmt.Println("Good guess!")
		} else {
			fmt.Println("Wrong guess!")
			wrongGuesses++
		}

		fmt.Println()
	}

	// End game
	fmt.Println("Thanks for playing!")
}

func drawHangman(wrongGuesses int) {
	if wrongGuesses == 0 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	} else if wrongGuesses == 1 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	} else if wrongGuesses == 2 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println("  |   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	} else if wrongGuesses == 3 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println(" /|   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	} else if wrongGuesses == 4 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println(" /|\\  |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	} else if wrongGuesses == 5 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println(" /|\\  |")
		fmt.Println(" /    |")
		fmt.Println("      |")
		fmt.Println("=========")
	} else if wrongGuesses == 6 {
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println(" /|\\  |")
		fmt.Println(" / \\  |")
		fmt.Println("      |")
		fmt.Println("=========")
	}
}

func drawHangmanSwitch(wrongGuesses int) {
	switch wrongGuesses {
	case 0:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	case 1:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	case 2:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println("  |   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	case 3:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println(" /|   |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	case 4:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println(" /|\\  |")
		fmt.Println("      |")
		fmt.Println("      |")
		fmt.Println("=========")
	case 5:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println(" /|\\  |")
		fmt.Println(" /    |")
		fmt.Println("      |")
		fmt.Println("=========")
	case 6:
		fmt.Println("  +---+")
		fmt.Println("  |   |")
		fmt.Println("  O   |")
		fmt.Println(" /|\\  |")
		fmt.Println(" / \\  |")
		fmt.Println("      |")
		fmt.Println("=========")
	}
}
