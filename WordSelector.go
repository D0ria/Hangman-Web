package hangman

import (
	"bufio"
	"os"
)

func WordSelector() string {
	file_words, _ := os.Open("words.txt")         //ouverture words.txt
	scanner_words := bufio.NewScanner(file_words) //flux words.txt

	var tab_words []string //tableau de mot
	for scanner_words.Scan() {
		tab_words = append(tab_words, scanner_words.Text())
	} //remplire le tableau de mot d'après words.txt
	to_found := tab_words[RandomNumber()] //mot à trouver
	return to_found
}
