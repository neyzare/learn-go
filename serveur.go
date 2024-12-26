package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/index", indexController)
	fmt.Println("le serveur est en cours sur le port : 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func indexController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "salut a tous")
}