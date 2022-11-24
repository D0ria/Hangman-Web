package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// API routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "accueil.html")
	})
	http.HandleFunc("/idcard", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ok...")
	})

	port := ":8080"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified above
	log.Fatal(http.ListenAndServe(port, nil))
}
