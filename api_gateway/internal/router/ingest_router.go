package router

import (
	"net/http"

	"github.com/ShivangSrivastava/sentinel/api_gateway/internal/handler"
)

func registerIngestRoutes(mux *http.ServeMux, h handler.IngestHandler) {
	mux.HandleFunc("POST /create-parser", h.HandleIngest)
}
