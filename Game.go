package hangman

import (
	"bufio"
	"fmt"
	"os"
)

type HangData struct {
	Life                 int
	Founded              []rune
	To_found             string
	To_found_RuneVersion []rune
	Correct              bool
	Founded_RuneVersion  string
}

func Game(Pts *HangData) {

	//flux input
	input := bufio.NewScanner(os.Stdin)

	fmt.Println("Good Luck, you have 10 attempts.")
	for {
		Pts.Correct = false
		PrintMyTabRune(Pts.Founded, len(Pts.To_found))
		fmt.Println("Your choose: ")
		input.Scan()
		input_text := input.Text()

		for j := 0; j < len(Pts.To_found); j++ { //lettre detector, change egalement le mot révélé selon les lettres trouver
			if input_text == (string(Pts.To_found[j])) {
				Pts.Correct = true
				Pts.Founded[j] = (rune(input_text[0]))
			}
		}

		if !Pts.Correct {
			Pts.Life--
			PrintBourreau(Pts.Life)
			if Pts.Life >= 1 { //pour évité de print après le mot
				fmt.Print("Not present in the word, ", Pts.Life, " attempts remaining", "\n")
			}

		}

		if Pts.Life <= 0 {
			fmt.Println("YOU DEATH ! TES PARENTS NE T'AIMENT PAS !")
			break
		}

		if Equal(Pts.To_found_RuneVersion, Pts.Founded) {
			fmt.Println("YOU WIN ! TES PARENT T'AIMENT !")
			break
		}
	}
}
