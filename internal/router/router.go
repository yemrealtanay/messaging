package router

import (
	"github.com/gorilla/mux"
	"messaging/internal/handlers"
)

type HandlerRegistry struct {
	Message *handlers.MessageHandler
	// Sender    *handlers.SenderHandler
	// Contact *handlers.ContactHandler
	//Burada büyütülebilir bir yapının ibaresini bırakmak istedim.
	//Örneğin bir sender'a ait verileri istiyor olabiliriz.
}

func NewRouter(h *HandlerRegistry) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/messages", h.Message.GetAllSentMessages).Methods("GET")

	return r
}
