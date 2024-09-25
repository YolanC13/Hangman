package main

import (
	"bufio"
	"fmt"
	hangman "hangman/Internals"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {
	hangman.Letters = &[]string{}
	//hangman.PressF11()
	hangman.ClearScreen()
	MainMenuDisplay()
}

func InitializeHangman(text string) {
	hangman.ClearScreen()
	HangmanAsciiPrinter(*hangman.PlayerLives)
	hangman.NewLine(2)
	fmt.Print("Mot à deviner : ")
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
				hangman.GetInput()
				MainMenuDisplay()
				return
			}
		}
	}
	getLetter(true)
}

func InitializeVariables(text string) {
	*hangman.PlayerLives = 9
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

func getLetter(x bool) {

	//DEBUG
	//fmt.Println(hangman.Characters)
	//fmt.Println(hangman.Letters)

	if x {
		hangman.NewLine(1)
		hangman.BoxStrings([]string{"Choisit une lettre"})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "/!\\ Tu as " + fmt.Sprint(*hangman.PlayerLives) + " vies /!\\",
			FgColor:     color.FgRed,
		})
	}
	input := hangman.GetInput()
	if input == "" || !IsLetter(input) {
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "Veuillez entrer seulement des lettres",
			FgColor:     color.FgRed,
		})
		getLetter(false)
	} else if len(input) != 1 {
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "Vous avez tenté de rentrer plus d'une lettre vous perdez 2 vies",
			FgColor:     color.FgRed,
		})
		*hangman.PlayerLives -= 2
		getLetter(false)
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
	fmt.Print("Mot à deviner : ")
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
					hangman.GetInput()
					MainMenuDisplay()
					return
				}
			}
		}
		getLetter(true)
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
		hangman.GetInput()
		MainMenuDisplay()
	}
}

func MainMenuDisplay() {
	hangman.ClearScreen()
	MainMenuIntro()
	hangman.NewLine(3)
	hangman.DisplayText(hangman.DisplayTextOptions{
		TextToPrint: "1. Jouer",
		IsCentered:  true,
	})
	hangman.DisplayText(hangman.DisplayTextOptions{
		TextToPrint: "2. Ajouter un mot",
		IsCentered:  true,
	})
	hangman.DisplayText(hangman.DisplayTextOptions{
		TextToPrint: "3. Quitter",
		IsCentered:  true,
	})
	hangman.NewLine(2)
	input := hangman.GetInput()
	switch input {
	case "1":
		words := hangman.LoadFile()
		if len(words) > 0 {
			hangman.Letters = &[]string{}
			hangman.HangmanChar = []string{}
			InitializeVariables(words[rand.Intn(len(words))])
		} else {
			fmt.Println("Pas de mots trouvés dans le fichier words.txt")
		}
	case "2":
		AddWord()
	case "3":
		hangman.ClearScreen()
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "Merci d'avoir joué!",
			IsCentered:  true,
		})
		time.Sleep(2 * time.Second)
		hangman.ClearScreen()
		return
	default:
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: "Veuillez entrer un chiffre entre 1 et 3",
			FgColor:     color.FgRed,
		})
		time.Sleep(2 * time.Second)
		MainMenuDisplay()
	}
}

func MainMenuIntro() {
	width, _ := hangman.SizeTest()
	fmt.Println(width)
	for i := width; i > width/2; i -= 2 {
		time.Sleep(40_000_000 * time.Nanosecond)
		hangman.ClearScreen()
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: strings.Repeat(" ", i/2) + "             ___    ___    _  _     ___    _   _  ",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: strings.Repeat(" ", i/2) + "    o O O   | _ \\  | __|  | \\| |   |   \\  | | | | ",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: strings.Repeat(" ", i/2) + "   o        |  _/  | _|   | .` |   | |) | | |_| | ",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: strings.Repeat(" ", i/2) + "  TS__[O]  _|_|_   |___|  |_|\\_|   |___/   \\___/  ",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: strings.Repeat(" ", i/2) + " {======||\"\"\"\"\"\"||\"\"\"\"\"\"||\"\"\"\"\"\"|_|\"\"\"\"\"\"||\"\"\"\"\"|",
		})
		hangman.DisplayText(hangman.DisplayTextOptions{
			TextToPrint: strings.Repeat(" ", i/2) + "./o--000' `-0-0-' `-0-0-' `-0-0-' `-0-0-' `-0-0-'",
		})
	}
}

func AddWord() {
	//MERCI CHAT GPT
	hangman.ClearScreen()

	file, err := os.OpenFile("words.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	r := bufio.NewReader(file)
	words := []string{}
	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			//fmt.Printf("ReadLine: %q\n", line)
			words = append(words, string(line))
		}
		if err != nil {
			break
		}
	}
	fmt.Println("écrivez le mot à ajouter à la liste:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	word := scanner.Text()
	words = append(words, word+"\n")

	w := bufio.NewWriter(file)
	for _, word := range words {
		fmt.Fprintln(w, word)
	}
	w.Flush()

	hangman.DisplayText(hangman.DisplayTextOptions{
		TextToPrint: "Mot ajouté à la liste",
		FgColor:     color.FgGreen,
	})
	hangman.NewLine(1)
	fmt.Println("Appuyez pour continuer")
	hangman.GetInput()
	MainMenuDisplay()
}
