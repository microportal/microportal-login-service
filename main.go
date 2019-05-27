package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"microportal-resource-service/controller"
	"net/http"
	"os"
)

var (
	lc         controller.LoginController
	port       string
	pathPrefix string
)

func init() {
	_ = gotenv.Load()
	port = os.Getenv("PORT")
	pathPrefix = os.Getenv("PATH_PREFIX")
	lc = controller.LoginController{}
	lc.Init()
}

func main() {
	if port == "" {
		port = "8080"
	}
	if pathPrefix == "" {
		pathPrefix = "/login-service"
	}
	addr := fmt.Sprint(":", port)

	router := mux.NewRouter().PathPrefix(pathPrefix).Subrouter()

	router.HandleFunc("/login", lc.Login).Methods(http.MethodPost)
	router.HandleFunc("/token", lc.ValidateToken).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(addr, router))
}
