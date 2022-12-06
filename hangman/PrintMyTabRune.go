package hangman

import (
	"fmt"
)

func PrintMyTabRune(tab_rune []rune, taille_tab int) { //print la decouverte du mot à chaque tour
	for i := 0; i < taille_tab; i++ {
		fmt.Print(string(tab_rune[i]))
	}
	fmt.Print("\n")
}