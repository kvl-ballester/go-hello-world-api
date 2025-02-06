package service

import "time"

// GetCurrentDate devuelve la fecha y hora actual en un formato legible.
func GetCurrentDate() string {
	return time.Now().Format("02/01/2006 15:04:05")
}
