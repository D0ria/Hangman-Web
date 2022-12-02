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

func Game(Hang *HangData) {

	//flux input
	input := bufio.NewScanner(os.Stdin)

	fmt.Println("Good Luck, you have 10 attempts.")
	for {
		Hang.Correct = false
		PrintMyTabRune(Hang.Founded, len(Hang.To_found))
		fmt.Println("Your choose: ")
		input.Scan()
		input_text := input.Text()

		for j := 0; j < len(Hang.To_found); j++ { //lettre detector, change egalement le mot révélé selon les lettres trouver
			if input_text == (string(Hang.To_found[j])) {
				Hang.Correct = true
				Hang.Founded[j] = (rune(input_text[0]))
			}
		}

		if !Hang.Correct {
			Hang.Life--
			PrintBourreau(Hang.Life)
			if Hang.Life >= 1 { //pour évité de print après le mot
				fmt.Print("Not present in the word, ", Hang.Life, " attempts remaining", "\n")
			}

		}

		if Hang.Life <= 0 {
			fmt.Println("YOU DEATH ! TES PARENTS NE T'AIMENT PAS !")
			break
		}

		if Equal(Hang.To_found_RuneVersion, Hang.Founded) {
			fmt.Println("YOU WIN ! TES PARENT T'AIMENT !")
			break
		}
	}
}
