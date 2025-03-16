package main

import (
	"fmt"
	"strings"
)

func main() {
	// --- SETUP ---
	// Welcome message
	fmt.Println("Welcome to Hangman!")
	fmt.Println("Can you guess the secret word before the hangman is complete?")
	fmt.Println()

	// Game settings
	secretWord := "programming"
	secretWord = strings.ToLower(secretWord)
	guessedLetters := ""
	wrongGuesses := 0
	maxWrongGuesses := 6

	// --- GAME LOOP ---
	for {
		// 1. Display current word state
		fmt.Print("Current word: ")
		wordComplete := true
		for _, char := range secretWord {
			if strings.ContainsRune(guessedLetters, char) {
				fmt.Print(string(char) + " ")
			} else {
				fmt.Print("_ ")
				wordComplete = false
			}
		}
		fmt.Println()

		// 2. Show progress
		fmt.Printf("Wrong guesses: %d/%d\n", wrongGuesses, maxWrongGuesses)
		drawHangman(wrongGuesses)

		// 3. Check win/lose conditions
		if wordComplete {
			fmt.Println("Congratulations! You guessed the word!")
			break
		}

		if wrongGuesses >= maxWrongGuesses {
			fmt.Println("Sorry, you lost! The word was:", secretWord)
			break
		}

		// 4. Get player's guess
		fmt.Print("Guess a letter: ")
		var guess string
		fmt.Scan(&guess)

		// Make guess lowercase
		guess = strings.ToLower(guess)

		// 5. Validate guess
		if len(guess) == 0 {
			fmt.Println("Please enter a letter!")
			fmt.Println()
			continue
		}

		guessChar := rune(guess[0])

		if guessChar < 'a' || guessChar > 'z' {
			fmt.Println("Please enter a valid letter!")
			fmt.Println()
			continue
		}

		if strings.ContainsRune(guessedLetters, guessChar) {
			fmt.Println("You already guessed that letter!")
			fmt.Println()
			continue
		}

		// 6. Record and process guess
		guessedLetters += string(guessChar)

		if strings.ContainsRune(secretWord, guessChar) {
			fmt.Println("Good guess!")
		} else {
			fmt.Println("Wrong guess!")
			wrongGuesses++
		}

		fmt.Println()
	}

	// --- END GAME ---
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
