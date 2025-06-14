package router

import (
	"github.com/gorilla/mux"
	"messaging/internal/handlers"
)

type HandlerRegistry struct {
	Message *handlers.MessageHandler
	Worker  *handlers.WorkerHandler
	// Sender    *handlers.SenderHandler
	// Contact *handlers.ContactHandler
	//Burada büyütülebilir bir yapının ibaresini bırakmak istedim.
	//Örneğin bir sender'a ait verileri istiyor olabiliriz.
}

func NewRouter(h *HandlerRegistry) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/messages", h.Message.GetAllSentMessages).Methods("GET")
	r.HandleFunc("/auto-sender/start", h.Worker.Start).Methods("POST")
	r.HandleFunc("/auto-sender/stop", h.Worker.Stop).Methods("POST")

	return r
}
