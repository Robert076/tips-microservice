package main

import (
	"encoding/json"
	"log"
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
		tips := db.GetTipsFromDatabase()

		if err := json.NewEncoder(w).Encode(tips); err != nil {
			http.Error(w, "Failed to encode tips", http.StatusInternalServerError)
			return
		}
	})
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
