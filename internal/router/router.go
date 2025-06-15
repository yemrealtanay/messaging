package router

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
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
	r.HandleFunc("/start", h.Worker.Start).Methods("POST")
	r.HandleFunc("/stop", h.Worker.Stop).Methods("POST")
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	return r
}
