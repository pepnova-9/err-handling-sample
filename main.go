package main

import (
	"net/http"

	"github.com/pepnova-9/err-handling-sample/router"
)

func main() {
	router := router.Router()
	http.Handle("/", router)

	server := &http.Server{
		Addr:    "localhost:8989",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}

}
