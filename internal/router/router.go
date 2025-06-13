package router

import (
	"github.com/gorilla/mux"
	"messaging/internal/handler/message"
)

func NewRouter(messageHandler *message.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/messages", messageHandler.GetAllSentMessages).Methods("GET")

	return r
}
