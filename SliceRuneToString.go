package hangman

func SliceRuneToString(founded []rune) string {
	var result string
	for i := 0; i < len(founded); i++ {
		result += (string(founded[i]))
	}
	return result
}
