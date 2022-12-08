package main

import (
	"fmt"
	"hangman"
	"html/template"
	"net/http"

	hangman "hangman/hangman"
)

const port = ":8080"

func Accueil(rw http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("./templates/accueil.html")
	tmp.Execute(rw, r)
}

func Jeu(rw http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("./templates/jeu.html")
	tmp.Execute(rw, r)
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

	func InitialiseStruct(Pts *hangman.HangData) {
		Pts.Life = 10
		Pts.To_found = hangman.WordSelector()
		Pts.Founded = hangman.Founded(Pts.To_found)
		Pts.To_found_RuneVersion = hangman.StringToSliceRune(Pts.To_found)
		Pts.Correct = false
		Pts.Founded_RuneVersion = hangman.SliceRuneToString(Pts.Founded)
	}
	
	HangPts := HangData{10,"","","","",""}
	InitialiseStruct(Pts)

	fmt.Println("(http://localhost:8080) - Serveur démarré sur le port", port)

	// var data = Page{"Hangman-Web ", Founded_RuneVersion, tabURL[Life], Life, string(To_found_RuneVersion), string(Founded_RuneVersion), Correct, To_found  , Founded} //actualisation de la variable data

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		Accueil(rw, r)
	})

	http.HandleFunc("/game", func(rw http.ResponseWriter, r *http.Request) {
		Jeu(rw, r)
	})

	http.HandleFunc("/game/win", func(rw http.ResponseWriter, r *http.Request) {
		Victoire(rw, r)
	})

	http.HandleFunc("/game/lose", func(rw http.ResponseWriter, r *http.Request) {
		Defaite(rw, r)
	})

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	rs := http.FileServer(http.Dir("./image/"))
	http.Handle("/image/", http.StripPrefix("/image/", rs))

	ts := http.FileServer(http.Dir("./dafont/"))
	http.Handle("/dafont/", http.StripPrefix("/dafont/", ts))

	http.HandleFunc("/hangman", func(rw http.ResponseWriter, r *http.Request) {
		Pts.InputLetter = r.FormValue("letter")
		Equal(Pts)
		Founded(Equal(Pts), Pts)
		htp.Redirect(rw, r, "/", http.StatusFound)
	})

	http.ListenAndServe(port, nil)
}
