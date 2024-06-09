package routes

import (
	"github.com/gorilla/mux"

	"go-mongo/src/pkg/config"
	"go-mongo/src/pkg/controllers"
)

func RegisterUserRoutes(router *mux.Router) {
	uc := controllers.NewUserController(config.GetClient())
	router.HandleFunc("/user/", uc.CreateUser).Methods("POST")
	router.HandleFunc("/user/", uc.GetUsers).Methods("GET")
	router.HandleFunc("/user/{userID}", uc.GetUserByID).Methods("GET")
	router.HandleFunc("/user/{userID}", uc.UpdateUserByID).Methods("PUT")
	router.HandleFunc("/user/{userID}", uc.DeleteUserByID).Methods("DELETE")
}
