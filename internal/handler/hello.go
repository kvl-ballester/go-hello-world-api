package handler

import (
	"fmt"
	"github.com/kvl-ballester/go-hello-world-api/internal/service"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	message := service.GetHelloMessage()
	fmt.Fprintf(w, message)
}
