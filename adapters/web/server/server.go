package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/leomoritz/fullcycle-arq-hexagonal-app/application"
)

type Webserver struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver(service application.ProductServiceInterface) *Webserver {
	return &Webserver{Service: service}
}

func (w Webserver) Serve() {
	router := mux.NewRouter()
	middleware := negroni.New(negroni.NewLogger())

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
