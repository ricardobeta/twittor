package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ricardobeta/twittor/middleW"
	"github.com/ricardobeta/twittor/routers"
	"github.com/rs/cors"
)

/*Manejadores seta el puerto y pone a escuchar al Servidor*/
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middleW.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
