package response

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func JSON[T any](w http.ResponseWriter, status int, message string, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := Response[T]{
		Status:  status,
		Message: message,
		Data:    data,
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, message, struct{}{})
}
