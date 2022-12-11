package utils

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

func Result(w http.ResponseWriter, msg Message) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
