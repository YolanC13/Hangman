package main

import (
	"bufio"
	"fmt"
	hangman "hangman/Internals"
	"math/rand"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	hangman.Letters = &[]string{}
	//hangman.PressF11()
	hangman.ClearScreen()
	words := LoadFile()
	//fmt.Println("Words found in the file:", len(words))
	//fmt.Println(words)
	if len(words) > 0 {
		InitializeVariables(words[rand.Intn(len(words))])
	} else {
		fmt.Println("No words found in the file.")
		// Return an empty slice if the loop completes without errors
	}
}

func InitializeHangman(text string) {
	hangman.ClearScreen()
	HangmanAsciiPrinter(*hangman.PlayerLives)
	hangman.NewLine(2)
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

func InitializeVariables(text string) {
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			hangman.HangmanChar = append(hangman.HangmanChar, " ")
			*hangman.Letters = append(*hangman.Letters, " ")
		} else {
			hangman.HangmanChar = append(hangman.HangmanChar, strings.ToLower(string(text[i])))
			*hangman.Letters = append(*hangman.Letters, "_")
		}
	}

	//AJOUTE DES LETTRES ALEATOIREMENT
	if len(*hangman.Letters) > 9 {
		for i := 0; i < 2; i++ {
			x := rand.Intn(len(*hangman.Characters))
			if (*hangman.Letters)[x] == "_" {
				(*hangman.Letters)[x] = (*hangman.Characters)[x]
			} else {
				i--
			}
		}
	} else if len(*hangman.Letters) > 5 {
		x := rand.Intn(len(*hangman.Characters))
		(*hangman.Letters)[x] = (*hangman.Characters)[x]
	}
	InitializeHangman(text)
}

func getLetter() {
	//fmt.Println(hangman.Characters)
	//fmt.Println(hangman.Letters)

	hangman.NewLine(1)
	hangman.DisplayText(hangman.DisplayTextOptions{
		TextToPrint: "Choisit une lettre",
	})
	input := hangman.GetInput()
	if input == "" || !IsLetter(input) {
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "Veuillez entrer seulement des lettres",
			FgColor:     color.FgRed,
		})
		getLetter()
	} else if len(input) != 1 {
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "Veuillez entrer une seule lettre",
			FgColor:     color.FgRed,
		})
		*hangman.PlayerLives -= 2
		getLetter()
	} else {
		checkLetter(strings.ToLower(input))
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
	if !foundLetter {
		*hangman.PlayerLives -= 1
	}
	//fmt.Println(hangman.Letters)
	PrintHangman()
}

func PrintHangman() {
	hangman.ClearScreen()
	HangmanAsciiPrinter(*hangman.PlayerLives)
	hangman.NewLine(2)
	if *hangman.PlayerLives > 0 {
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
}

func IsLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && r != ' ' {
			return false
		}
	}
	return true
}

func HangmanAsciiPrinter(lives int) {
	switch lives {
	case 9:
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "        ",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "        ",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "        ",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "        ",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "        ",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "        ",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "        ",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "===========",
		})
	case 8:
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       +",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "===========",
		})
	case 7:
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " +-----+",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "===========",
		})
	case 6:
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " +-----+",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " |     |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "===========",
		})
	case 5:
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " +-----+",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " |     |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " O     |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "===========",
		})
	case 4:
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " +-----+",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " |     |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " O     |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " |     |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "===========",
		})
	case 3:
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " +-----+",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " |     |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " O     |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "/|     |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "===========",
		})
	case 2:
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " +-----+",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " |     |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " O     |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "/|\\    |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "===========",
		})

	case 1:
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " +-----+",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " |     |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " O     |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "/|\\    |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "/      |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "===========",
		})
	case 0:
		hangman.ClearScreen()
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " +-----+",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " |     |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: " O     |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "/|\\    |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "/ \\    |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "       |",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "===========",
		})
		hangman.NewLine(2)
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "Tu as perdu !",
			FgColor:     color.FgRed,
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "LE MOT ÉTAIT : " + strings.Join(*hangman.Characters, ""),
		})
	}
}

func LoadFile() []string {
	//Section 1
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	r := bufio.NewReader(file)
	words := []string{}
	// Section 2
	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			//fmt.Printf("ReadLine: %q\n", line)
			words = append(words, string(line))
		}
		if err != nil {
			return words
		}
	}
}
