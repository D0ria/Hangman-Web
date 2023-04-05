package hangman

func StringToSliceRune(word_to_found string) []rune { //transforme un string en un tableau de rune
	var result []rune
	for i := 0; i < len(word_to_found); i++ {
		result = append(result, (rune(word_to_found[i])))
	}
	return result
}