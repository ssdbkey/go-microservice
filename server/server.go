// Package server contains all the logic regarding the http server, CORS, and route registering.
package server

import (
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/ssdbkey/go-microservice-test/handlers"
)

// NewServer creates and return a http.Server struct, sets timeouts, register routes and optionally adds CORS handling.
func NewServer(h *handlers.Handler) *http.Server {

	router := httprouter.New()

	// get path
	router.GET("/path", h.Path())

	srvr := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return srvr
}
