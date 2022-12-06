package hangman


func Equal(a, b []rune) bool { //fonction comparation de tableau de rune(condition de victoire)
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}