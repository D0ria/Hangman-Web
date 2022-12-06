package hangman

import (
	"bufio"
	"fmt"
	"os"
)

func PrintBourreau(life int) {
	// open the file
	file, _ := os.Open("bourreau.txt") //ouverture bourreau.txt

	fileScanner := bufio.NewScanner(file) //flux bourreau.txt

	var hangmanPositions [10]string //stocke chaque position
	var res string
	index := 9
	// read line by line
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

	// for , v := range hangmanPositions {
	//     fmt.Println(v)
	// }
	fmt.Println(hangmanPositions[life])
}
