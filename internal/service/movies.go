package service

import (
	"fmt"
	"github.com/kvl-ballester/go-hello-world-api/internal/db"
	"github.com/kvl-ballester/go-hello-world-api/internal/model"
)

func GetMovies(title, director string) ([]model.Movie, error) {
	query := "SELECT * FROM movies WHERE 1=1"

	var args []interface{}

	if title != "" {
		query += " AND Title LIKE ?"
		args = append(args, "%"+title+"%")
	}
	if director != "" {
		query += " AND Director LIKE ?"
		args = append(args, "%"+director+"%")
	}

	fmt.Println(query)

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var movies []model.Movie
	for rows.Next() {
		var m model.Movie
		if err := rows.Scan(&m.ID, &m.Title, &m.Director, &m.ReleaseDate, &m.Rating); err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, nil
}

func AddMovie(m model.Movie) error {
	query := fmt.Sprintf(
		"INSERT INTO movies (Title, Director, ReleaseDate, Rating) VALUES ('%s', '%s', '%s', '%s')",
		m.Title,
		m.Director,
		m.ReleaseDate,
		m.Rating)

	_, err := db.DB.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
