package main

import (
	"fmt"
	hangman "hangman/Internals"

	"github.com/fatih/color"
	//"github.com/fatih/color"
)

func main() {
	hangman.Letters = &[]string{}
	hangman.PressF11()
	hangman.ClearScreen()

	hangman.DisplayText(hangman.DisplayTextOptions{
		TextToPrint: "Choisit un mot à faire deviner:",
	})
	input := hangman.GetInput()
	InitializeHangman(input)
}

func InitializeHangman(text string) {
	hangman.ClearScreen()
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			hangman.HangmanChar = append(hangman.HangmanChar, " ")
			*hangman.Letters = append(*hangman.Letters, " ")
			fmt.Print(" ")
		} else {
			hangman.HangmanChar = append(hangman.HangmanChar, string(text[i]))
			*hangman.Letters = append(*hangman.Letters, "_")
			fmt.Print("_")
		}
	}
	hangman.Characters = &hangman.HangmanChar
	hangman.NewLine(1)
	getLetter()
}

func getLetter() {
	//fmt.Println(hangman.Characters)
	//fmt.Println(hangman.Letters)
	hangman.NewLine(1)
	hangman.DisplayText(hangman.DisplayTextOptions{
		TextToPrint: "Choisit une lettre",
	})
	input := hangman.GetInput()
	if len(input) != 1 {
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "Veuillez entrer une seule lettre",
			FgColor:     color.FgRed,
		})
		return
	} else {
		checkLetter(input)
	}
}

func checkLetter(letter string) {
	foundLetter := false
	for i := 0; i < len(*hangman.Characters); i++ {
		if letter == (*hangman.Characters)[i] {
			(*hangman.Letters)[i] = letter
			foundLetter = true
		}
	}
	if foundLetter {
		*hangman.PlayerLives -= 1
	}
	//fmt.Println(hangman.Letters)
	PrintHangman()
}

func PrintHangman() {
	hangman.ClearScreen()
	for i := 0; i < len(*hangman.Letters); i++ {
		fmt.Print(string((*hangman.Letters)[i]))
	}
	hangman.NewLine(1)
	for j := 0; j < len(*hangman.Letters); j++ {
		if (*hangman.Letters)[j] == "_" {
			break
		} else {
			if j == len(*hangman.Letters)-1 {
				hangman.DisplayText(hangman.DisplayTextOptions{
					TextToPrint: "Tu as gagné !",
					FgColor:     color.FgGreen,
				})
				return
			}
		}
	}
	getLetter()
}
