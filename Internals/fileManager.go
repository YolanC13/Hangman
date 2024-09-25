package hangman

import (
	"bufio"
	"fmt"
	"os"
)

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
