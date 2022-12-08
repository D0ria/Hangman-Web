package main

import (
	"fmt"
	"hangman"
	"net/http"
	"text/template"
)

func InitialiseStruct(Pts *hangman.HangData) {
	Pts.Life = 10
	Pts.To_found = hangman.WordSelector()
	Pts.Founded = hangman.Founded(Pts.To_found)
	Pts.To_found_RuneVersion = hangman.StringToSliceRune(Pts.To_found)
	Pts.Correct = false
	Pts.Founded_StringVersion = hangman.SliceRuneToString(Pts.Founded)
	Pts.Input_Letter = ""
}

func Accueil(rw http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("./templates/accueil.html")
	tmp.Execute(rw, r)
}

func Jeu(rw http.ResponseWriter, r *http.Request, Pts *hangman.HangData) {
	tmp, _ := template.ParseFiles("./templates/jeu.html")
	tmp.Execute(rw, Pts)
}

func Defaite(rw http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("./templates/defaite.html")
	tmp.Execute(rw, r)
}

func Victoire(rw http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("./templates/victoire.html")
	tmp.Execute(rw, r)
}

func main() {

	fmt.Println("(http://localhost:8080) - Serveur démarré sur le port")

	Hang := &hangman.HangData{}
	InitialiseStruct(Hang)
	fmt.Print(Hang.To_found)

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	rs := http.FileServer(http.Dir("./image/"))
	http.Handle("/image/", http.StripPrefix("/image/", rs))

	ts := http.FileServer(http.Dir("./dafont/"))
	http.Handle("/dafont/", http.StripPrefix("/dafont/", ts))

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		Accueil(rw, r)
	})

	http.HandleFunc("/game", func(rw http.ResponseWriter, r *http.Request) {
		Jeu(rw, r, Hang)
		// hangman.Game(Hang)
	})

	http.HandleFunc("/hangman", func(rw http.ResponseWriter, r *http.Request) {
		hangman.Game(Hang)

		if Hang.Life == 0 {
			http.Redirect(rw, r, "/game/lose", http.StatusFound)
			InitialiseStruct(Hang)
		} else if hangman.Equal(Hang.Founded, Hang.To_found_RuneVersion) {
			http.Redirect(rw, r, "/game/win", http.StatusFound)
			InitialiseStruct(Hang)
		} else if Hang.Life > 0 {
			http.Redirect(rw, r, "/game", http.StatusFound)
		}
	})

	http.HandleFunc("/game/win", func(rw http.ResponseWriter, r *http.Request) {
		Victoire(rw, r)
	})

	http.HandleFunc("/game/lose", func(rw http.ResponseWriter, r *http.Request) {
		Defaite(rw, r)
	})

	http.ListenAndServe(":8080", nil)
	//======

}
