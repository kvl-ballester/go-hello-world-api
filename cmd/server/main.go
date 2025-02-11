package main

import (
	"fmt"
	"github.com/kvl-ballester/go-hello-world-api/internal/db"
	"github.com/kvl-ballester/go-hello-world-api/internal/handler"
	"log"
	"net/http"
)

func main() {
	db.InitDB()
	// Definir las rutas
	http.HandleFunc("/hello", handler.HelloHandler)

	http.HandleFunc("/date_html", handler.DateHtmlHandler)
	http.HandleFunc("/movies", handler.MoviesHandler)

	fmt.Println("API escuchando en el puerto 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
