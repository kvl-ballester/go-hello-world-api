package model

type Movie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Director    string `json:"director"`
	ReleaseDate string `json:"release_date"`
	Rating      string `json:"rating"`
}
