package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

func Accueil(rw http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("./templates/accueil.html")
	tmp.Execute(rw, r)
}

func Connexion(rw http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("./templates/connexion.html")
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
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		Accueil(rw, r)
	})

	http.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {
		Connexion(rw, r)
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

	fmt.Println("(http://localhost:8080) - Serveur démarré sur le port", port)
	http.ListenAndServe(port, nil)

}
