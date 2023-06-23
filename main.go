package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pepnova-9/err-handling-sample/controller"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users/{userID}", controller.GetUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/users", controller.CreateUserHandler).Methods(http.MethodPut)
	r.Use(requestContextCheckMiddleware)
	return r
}

func main() {
	router := Router()
	http.Handle("/", router)

	server := &http.Server{
		Addr:    "localhost:8989",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
