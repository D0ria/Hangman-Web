package main

import (
	"fmt"
	"hangman"
	"net/http"
	"text/template"
)

func InitialiseStruct(Pts *hangman.HangData) { //Structure initiale
	Pts.Life = 10
	Pts.To_found = hangman.WordSelector()
	Pts.Founded = hangman.Founded(Pts.To_found)
	Pts.To_found_RuneVersion = hangman.StringToSliceRune(Pts.To_found)
	Pts.Correct = false
	Pts.Founded_StringVersion = hangman.SliceRuneToString(Pts.Founded)
	Pts.Input_Letter = ""
}

func Accueil(rw http.ResponseWriter, r *http.Request) { //Exécute la page d'accueil
	tmp, _ := template.ParseFiles("./templates/accueil.html")
	tmp.Execute(rw, r)
}

func Jeu(rw http.ResponseWriter, r *http.Request, Pts *hangman.HangData) { //Exécute la page de jeu
	tmp, _ := template.ParseFiles("./templates/jeu.html")
	tmp.Execute(rw, Pts)
}

func Defaite(rw http.ResponseWriter, r *http.Request) { //Exécute la page de défaite
	tmp, _ := template.ParseFiles("./templates/defaite.html")
	tmp.Execute(rw, r)
}

func Victoire(rw http.ResponseWriter, r *http.Request) { //Exécute la page de victoire
	tmp, _ := template.ParseFiles("./templates/victoire.html")
	tmp.Execute(rw, r)
}

func main() {

	fmt.Println("(http://localhost:8080) - Serveur démarré sur le port")

	Hang := &hangman.HangData{}
	InitialiseStruct(Hang) //Initialisation de la structure

	fs := http.FileServer(http.Dir("./static/")) //Liaison du css
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	rs := http.FileServer(http.Dir("./image/")) //Liaison des images
	http.Handle("/image/", http.StripPrefix("/image/", rs))

	ts := http.FileServer(http.Dir("./dafont/")) //Liaison de la police
	http.Handle("/dafont/", http.StripPrefix("/dafont/", ts))

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) { //requête page d'accueil
		Accueil(rw, r)
	})

	http.HandleFunc("/game", func(rw http.ResponseWriter, r *http.Request) { //requête page de jeu
		Jeu(rw, r, Hang)
	})

	http.HandleFunc("/hangman", func(rw http.ResponseWriter, r *http.Request) { //requête page inexistante qui calcule
		Hang.Input_Letter = r.FormValue("letter")
		hangman.Game(Hang)

		if Hang.Life == 0 { //boucle du jeu
			http.Redirect(rw, r, "/game/lose", http.StatusFound)
			InitialiseStruct(Hang)
		} else if Hang.Input_Letter == Hang.To_found {
			http.Redirect(rw, r, "/game/win", http.StatusFound)
			InitialiseStruct(Hang)
		} else if hangman.Equal(Hang.Founded, Hang.To_found_RuneVersion) {
			http.Redirect(rw, r, "/game/win", http.StatusFound)
			InitialiseStruct(Hang)
		} else if Hang.Life > 0 {
			http.Redirect(rw, r, "/game", http.StatusFound)
		}
	})

	http.HandleFunc("/game/win", func(rw http.ResponseWriter, r *http.Request) { //redirection page de victoire
		Victoire(rw, r)
	})

	http.HandleFunc("/game/lose", func(rw http.ResponseWriter, r *http.Request) { //redirection page de défaite
		Defaite(rw, r)
	})

	http.ListenAndServe(":8080", nil)
}
