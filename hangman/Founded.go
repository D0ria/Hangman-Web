package hangman

import (

)

func Founded(to_found string) []rune{
	var founded []rune                   
	for i := 0; i < len(to_found); i++ { //initialisation de founded pour qu'il devient un miroir de "_"
		founded = append(founded, '_')
	}
	return founded
}