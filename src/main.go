package main

import (
	"encoding/json"
	"net/http"
)

func getTodosFromDatabase() {

}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if IsMethodNotAllowed(w, r, http.MethodGet) {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		todos, _ := GetTodosFromDatabase()
		json.NewEncoder(w).Encode(todos)
	})
}
