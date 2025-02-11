package handler

import (
	"encoding/json"
	"fmt"
	"github.com/kvl-ballester/go-hello-world-api/internal/model"
	"github.com/kvl-ballester/go-hello-world-api/internal/service"
	"net/http"
	"time"
)

func MoviesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getMovies(w, r)
	case http.MethodPost:
		addMovie(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	director := r.URL.Query().Get("director")

	movies, err := service.GetMovies(title, director)
	if err != nil {
		http.Error(w, "Error obteniendo películas", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Error en datos de entrada", http.StatusBadRequest)
		return
	}

	if movie.Title == "" || movie.Director == "" {
		http.Error(w, "Faltan campos obligatorios", http.StatusBadRequest)
		return
	}

	// Convertir fecha de string a formato adecuado
	_, err := time.Parse("2006-01-02", movie.ReleaseDate)
	if err != nil {
		http.Error(w, "Formato de fecha inválido. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	err = service.AddMovie(movie)
	if err != nil {
		http.Error(w, "Error al guardar película", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Película añadida con éxito")

}
