package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-mongo/src/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting Server at PORT 9010")
	err := http.ListenAndServe(":9010", r)
	if err != nil {
		log.Fatal(err)
	}
}
