package handler

import (
	"github.com/kvl-ballester/go-hello-world-api/internal/service"
	"html/template"
	"net/http"
)

func DateHtmlHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/date.html"))
	date := service.GetCurrentDate()

	data := struct {
		Now string
	}{
		Now: date,
	}

	tmpl.Execute(w, data)

}
