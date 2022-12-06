package main

import (
	"fmt"
	"html/template"
	"net/http"

	hangman "hangman/hangman"

)

type HangData struct {
	Life                 int
	Founded              []rune
	To_found             string
	To_found_RuneVersion []rune
	Correct              bool
	Founded_RuneVersion  string
	TabURL 				 []string
}

const port = ":8080"

func InitialiseStruct(Pts *hangman.HangData) {
	Pts.Life = 10
	Pts.To_found = hangman.WordSelector()
	Pts.Founded = hangman.Founded(Pts.To_found)
	Pts.To_found_RuneVersion = hangman.StringToSliceRune(Pts.To_found)
	Pts.Correct = false
	Pts.Founded_RuneVersion = hangman.SliceRuneToString(Pts.Founded)
	Pts.TabUrl = [
	"",
	"./position/pictures/pendu_10.png",
	"./position/pictures/pendu_9.png",
	"./position/pictures/pendu_8.png",
	"./position/pictures/pendu_7.png",
	"./position/pictures/pendu_6.png",
	"./position/pictures/pendu_5.png",
	"./position/pictures/pendu_4.png",
	"./position/pictures/pendu_3.png",
	"./position/pictures/pendu_2.png",
	"./position/pictures/pendu_1.png"
	]
}

func Website() {
	if string(To_found_RuneVersion) == string(Founded_RuneVersion) {
		Correct = true
	}
	if game.Lettre != "" && game.Lettre != " " && !same {
		To_found   = append(To_found  , game.Lettre)
		Founded_RuneVersion += game.Lettre
		Founded_RuneVersion += ", "
	}
	data = Page{"Hangman-Web ", Founded_RuneVersion, tabURL[Life], Life, string(To_found_RuneVersion), string(Founded_RuneVersion), Correct, To_found  , Founded} // actualisation de data
	tmpl.ExecuteTemplate(w, "index", data) //execution de la template "index" avec les données
}


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

func Restart() {
	Life                 = 10
	Founded              = hangman.Founded(Pts.To_found)
	To_found             = hangman.WordSelector()
	To_found_RuneVersion = hangman.StringToSliceRune(Pts.To_found)
	Correct              = false
	Founded_RuneVersion  = hangman.SliceRuneToString(Pts.Founded)
}

func main() {
	fmt.Println("(http://localhost:8080) - Serveur démarré sur le port", port)

	// var data = Page{"Hangman-Web ", Founded_RuneVersion, tabURL[Life], Life, string(To_found_RuneVersion), string(Founded_RuneVersion), Correct, To_found  , Founded} //actualisation de la variable data
	var data HangData
	Pts := &data

	InitialiseStruct(Pts)
	fmt.pri

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

	http.ListenAndServe(port, nil)
}
