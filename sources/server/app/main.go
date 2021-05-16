package main

import (
	"github.com/c16a/microproxy/server/api/management"
	"github.com/c16a/microproxy/server/api/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	srv := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	router.HandleFunc("/", routes.HandleRoute).Methods(http.MethodGet, http.MethodPost)

	go setupManagementServer(router)

	log.Println("Starting reverse proxy")
	log.Fatalln(srv.ListenAndServe())
}

func setupManagementServer(proxyRouter *mux.Router) {
	mRouter := mux.NewRouter()

	srv := &http.Server{
		Addr:    ":7000",
		Handler: mRouter,
	}

	mRouter.HandleFunc("/routes", management.AddRoute).Methods(http.MethodPost)

	log.Println("Starting management server")
	log.Fatalln(srv.ListenAndServe())
}
