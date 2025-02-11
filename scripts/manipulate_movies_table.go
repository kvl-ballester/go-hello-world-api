package main

import (
	"database/sql"
	"fmt"
	"github.com/kvl-ballester/go-hello-world-api/internal/model"
	"github.com/kvl-ballester/go-hello-world-api/utils"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func connectDB() (*sql.DB, error) {
	wd := utils.GetProjectRoot()
	db, err := sql.Open("sqlite3", wd+"/db/movies.sqlite")
	if err != nil {
		return nil, err
	}
	fmt.Println("Conectado a la base de datos")
	return db, nil
}

func dropTables(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE IF EXISTS movies")
	return err
}

func getMoviesTable(db *sql.DB) ([]model.Movie, error) {
	query := "SELECT * FROM movies"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	movies := []model.Movie{}
	for rows.Next() {
		var movie model.Movie
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Director, &movie.ReleaseDate, &movie.Rating); err != nil {
			return nil, err
		}

		movies = append(movies, movie)

	}

	return movies, nil
}

func insertMovies(db *sql.DB) error {
	movies := []model.Movie{
		{0, "Inception", "Christopher Nolan", "2010-07-16", "8.8"},
		{1, "The Godfather", "Francis Ford Coppola", "1972-03-24", "9.2"},
		{2, "Pulp Fiction", "Quentin Tarantino", "1994-10-14", "8.9"},
		{3, "The Dark Knight", "Christopher Nolan", "2008-07-18", "9.0"},
		{4, "Forrest Gump", "Robert Zemeckis", "1994-07-06", "8.8"},
		{5, "Interstellar", "Christopher Nolan", "2014-11-07", "8.6"},
		{6, "Fight Club", "David Fincher", "1999-10-15", "8.8"},
		{7, "The Matrix", "Lana Wachowski, Lilly Wachowski", "1999-03-31", "8.7"},
		{8, "Goodfellas", "Martin Scorsese", "1990-09-19", "8.7"},
		{9, "The Shawshank Redemption", "Frank Darabont", "1994-09-23", "9.3"},
	}

	tx, err := db.Begin() // Inicia una transacción
	if err != nil {
		log.Fatal("Error iniciando transacción:", err)
	}

	stmt, err := db.Prepare("INSERT INTO movies (ID, Title, Director, Release_Date, Rating) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Error preparando consulta:", err)
	}
	defer stmt.Close()

	for _, movie := range movies {
		_, err := stmt.Exec(movie.ID, movie.Title, movie.Director, movie.ReleaseDate, movie.Rating)
		if err != nil {
			log.Fatal("Error insertando película:", err)
		}
	}

	err = tx.Commit() // Confirma la transacción
	if err != nil {
		log.Fatal("Error confirmando transacción:", err)
	}

	log.Println("Se insertaron 10 películas de muestra correctamente.")
	return nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Println("Error al conectar a la base de datos:")
		panic(err)
	}

	//err = dropTables(db)
	//movies, err := getMoviesTable(db)
	err = insertMovies(db)
	if err != nil {
		fmt.Println("Error al eliminar la tabla:")
		panic(err)
	}
	movies, err := getMoviesTable(db)
	for _, movie := range movies {
		fmt.Println(movie)
	}

	defer db.Close()
}
