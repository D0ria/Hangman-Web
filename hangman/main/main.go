package main

import (
	"bufio"
	"fmt"
	"hangman"
	"net/http"
	"os"
	"text/template"
)

func HandleHomePage(rw http.ResponseWriter, r *http.Request, str *hangman.HangData) {
	tmp, _ := template.ParseFiles(".html")
	tmp.Execute(rw, str)
}

func InitialiseStruct(Pts *hangman.HangData) {
	Pts.Life = 10
	Pts.To_found = hangman.WordSelector()
	Pts.Founded = hangman.Founded(Pts.To_found)
	Pts.To_found_RuneVersion = hangman.StringToSliceRune(Pts.To_found)
	Pts.Correct = false
	Pts.Founded_RuneVersion = hangman.SliceRuneToString(Pts.Founded)
}

func main() {
	//====== WEB SERVER
	fmt.Printf("Starting server at port 8080\n")

	Hang := &hangman.HangData{}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		InitialiseStruct(Hang)

		HandleHomePage(rw, r, Hang)
	})

	http.ListenAndServe(":8080", nil)
	//======

	InitialiseStruct(Hang)

	//fmt.Println(tab_words[nbr_words], "<--//===============TEST=imprime le mot du jeu
	//fmt.Println(to_found_RuneVersion, "<--//===============TEST=mot à trouver version rune

	// Fonction principal du jeu (loop)
	var want_play bool
	want_play = true
	respond := bufio.NewScanner(os.Stdin)
	InitialiseStruct(Hang)
	for {
		hangman.Game(Hang)
		fmt.Printf("Continuer à joue ? O/N")
		respond.Scan()
		respond := respond.Text()
		if respond == "O" {
			want_play = true
		} else if respond == "N" {
			want_play = false
		}
		if !want_play {
			break
		}
	}
}
