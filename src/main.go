package main

import (
	"encoding/json"
	"net/http"

	"github.com/Robert076/tips-microservice/db"
	"github.com/Robert076/tips-microservice/utils"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if utils.IsMethodNotAllowed(w, r, http.MethodGet) {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		todos := db.GetTipsFromDatabase()
		json.NewEncoder(w).Encode(todos)
	})
}
