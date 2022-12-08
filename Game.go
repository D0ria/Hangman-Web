package hangman

type HangData struct {
	Life                  int
	Founded               []rune
	To_found              string
	To_found_RuneVersion  []rune
	Correct               bool
	Founded_StringVersion string
	Input_Letter          string
}

func Game(Hang *HangData) {

	Hang.Correct = false

	for j := 0; j < len(Hang.To_found); j++ { //lettre detecté, change également le mot révélé selon les lettres trouvées
		if Hang.Input_Letter == (string(Hang.To_found[j])) {
			Hang.Correct = true
			Hang.Founded[j] = (rune(Hang.Input_Letter[0]))
		}
	}

	if !Hang.Correct {
		Hang.Life--
	}

	Hang.Founded_StringVersion = SliceRuneToString(Hang.Founded)
}
