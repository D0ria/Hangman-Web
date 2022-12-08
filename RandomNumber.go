package hangman

import (
	"math/rand"
	"time"
)

func RandomNumber() int { //Lettre indice
	rand.Seed(time.Now().UnixNano()) //seed
	nbr_words := rand.Intn(85) + 1   //Random number for pick word
	return nbr_words
}
