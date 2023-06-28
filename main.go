package main

import (
	"net/http"

	"github.com/pepnova-9/err-handling-sample/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	sample1 := controller.Sample1{}
	r.HandleFunc("/sample1/users/{userID}", sample1.UpdateUserHandler).Methods(http.MethodPut)
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
