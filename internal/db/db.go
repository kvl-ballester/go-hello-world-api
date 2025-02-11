package db

import (
	"database/sql"
	"github.com/kvl-ballester/go-hello-world-api/utils"
	_ "github.com/mattn/go-sqlite3" // Driver para SQLite
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	wd := utils.GetProjectRoot()
	DB, err = sql.Open("sqlite3", wd+"/db/movies.sqlite") // Base de datos SQLite en un archivo local
	if err != nil {
		log.Fatal("Error al abrir la base de datos:", err)
	}

	// Crear la tabla si no existe
	createTable := `
	CREATE TABLE IF NOT EXISTS movies (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		director TEXT NOT NULL,
		release_date TEXT NOT NULL,
		rating REAL NOT NULL
	);`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal("Error creando la tabla:", err)
	}

	log.Println("Base de datos inicializada correctamente")
}
