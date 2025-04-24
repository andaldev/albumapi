package api

import (
	"encoding/json"
	"log"
	"net/http"

	"albumapi/data"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

var all = data.GetAlbum()

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(all)
	})

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	log.Printf("Server has started on %s", s.addr)

	return server.ListenAndServe()
}
