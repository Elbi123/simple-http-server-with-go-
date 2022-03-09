package server

import (
	"github.com/Elbi123/learnmicro/router"
	"net/http"
)

func ServeAPI(endpoint string) error {
	router := router.EventRouter()

	return http.ListenAndServe(endpoint, router)
}
