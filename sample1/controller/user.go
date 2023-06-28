package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	usecase2 "github.com/pepnova-9/err-handling-sample/sample1/usecase"

	"github.com/gorilla/mux"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, err := usecase2.GetUserUsecase(r.Context(), vars["userID"])
	if err != nil {
		if errors.Is(r.Context().Err(), context.Canceled) {
			w.WriteHeader(499)
			return
		}

		switch {
		case errors.Is(err, usecase2.ErrUnauthorized):
			w.WriteHeader(http.StatusUnauthorized)
			return
		case errors.Is(err, usecase2.ErrUserNotFound):
			w.WriteHeader(http.StatusNotFound)
			return
		case errors.Is(err, usecase2.ErrUnexpectedError):
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("unexpected error handling"))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("default error handling"))
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type CreaetUserRequestBody struct {
	Name string `json:"name"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	requestBody := &CreaetUserRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := usecase2.CreateUserUsecase(r.Context(), usecase2.CreateUserInput{Name: requestBody.Name})
	if err != nil {
		if errors.Is(r.Context().Err(), context.Canceled) {
			w.WriteHeader(499)
			return
		}

		switch {
		case errors.Is(err, usecase2.ErrUnauthorized):
			w.WriteHeader(http.StatusUnauthorized)
			return
		case errors.Is(err, usecase2.ErrUserNotFound):
			w.WriteHeader(http.StatusNotFound)
			return
		case errors.Is(err, usecase2.ErrUnexpectedError):
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("unexpected error handling"))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("default error handling"))
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
