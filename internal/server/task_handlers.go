package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type TaskResult struct {
	ID     int     `json:"id"`
	Result float64 `json:"result"`
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		task, found := GetNextTask()
		if !found {
			http.Error(w, "No task available", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"task": task})
	case http.MethodPost:
		var res TaskResult
		if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
			http.Error(w, "Invalid request body", http.StatusUnprocessableEntity)
			return
		}
		log.Printf("Received result for task %d: %f", res.ID, res.Result)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
