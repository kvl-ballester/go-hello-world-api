package utils

import (
	"os"
	"path/filepath"
)

func GetProjectRoot() string {
	wd, err := os.Getwd() // Obtiene el directorio actual
	if err != nil {
		panic(err)
	}

	// Buscar el archivo go.mod desde el directorio actual hacia arriba
	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			return wd
		}

		parent := filepath.Dir(wd)
		if parent == wd {
			break // Llegamos a la raíz del sistema de archivos
		}
		wd = parent
	}

	return "" // No se encontró go.mod
}
