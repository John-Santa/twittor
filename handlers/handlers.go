package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/John-Santa/twittor/middlewares"
	"github.com/John-Santa/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Routes() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.CheckDB(routers.Registro)).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
