package router

import (
	"github.com/Elbi123/learnmicro/event"
	"github.com/gorilla/mux"
)


func EventRouter() *mux.Router {
	var eventsHandler event.EventServiceHandler

	nRouter := mux.NewRouter()

	eventsRouter := nRouter.PathPrefix("/events").Subrouter()
	/*eventsRouter.Methods("GET").Path("/{searchby}/{search}").HandleFunc(eventsHandler.FindEventHandler)
	eventsRouter.Methods("GET").Path("").HandleFunc(eventsHandler.AllEventHandler)
	eventsRouter.Methods("POST").Path("").HandleFunc(eventsHandler.AddNewEventHandler) */

	eventsRouter.HandleFunc("/", eventsHandler.FindEventHandler).Methods("GET")
	eventsRouter.HandleFunc("/{searchby}/{search}", eventsHandler.AllEventHandler).Methods("GET")
	eventsRouter.HandleFunc("/", eventsHandler.AddNewEventHandler).Methods("POST")

	return eventsRouter
}
