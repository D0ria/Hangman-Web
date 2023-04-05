package hangman

import (
	"bufio"
	"fmt"
	"os"
)

func PrintBourreau(life int) {
	// ouvre le fichier
	file, _ := os.Open("bourreau.txt") //ouverture bourreau.txt

	fileScanner := bufio.NewScanner(file) //flux bourreau.txt

	var hangmanPositions [10]string //stock chaque position
	var res string
	index := 9
	// lecture ligne par ligne
	for fileScanner.Scan() {
		if fileScanner.Text() != "" {
			res += fileScanner.Text() + "\n"
		} else {
			hangmanPositions[index] = res
			index--
			res = ""
		}
	}
	fmt.Println()
	fmt.Println(hangmanPositions[life])
}
