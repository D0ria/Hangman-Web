package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

func Accueil(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "accueil")
}

func Connexion(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "connexion")
}

func Jeu(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "jeu")
}

func Victoire(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "victoire")
}

func Defaite(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "defaite")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)

}

func main() {
	http.HandleFunc("/", Accueil)
	http.HandleFunc("/login", Connexion)
	http.HandleFunc("/game", Jeu)
	http.HandleFunc("/game/win", Victoire)
	http.HandleFunc("/game/lose", Defaite)

	fmt.Println("(http://localhost:8080) - Serveur démarré sur le port", port)
	http.ListenAndServe(port, nil)

}
