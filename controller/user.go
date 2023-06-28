package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/pepnova-9/err-handling-sample/usecase"

	"github.com/gorilla/mux"
)

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user, err := usecase.GetUserUsecase(r.Context(), vars["userID"])
	if err != nil {
		if errors.Is(r.Context().Err(), context.Canceled) {
			w.WriteHeader(499)
			return
		}

		switch {
		case errors.Is(err, usecase.ErrUnauthorized):
			w.WriteHeader(http.StatusUnauthorized)
			return
		case errors.Is(err, usecase.ErrUserNotFound):
			w.WriteHeader(http.StatusNotFound)
			return
		case errors.Is(err, usecase.ErrUnexpectedError):
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

	user, err := usecase.CreateUserUsecase(r.Context(), usecase.CreateUserInput{Name: requestBody.Name})
	if err != nil {
		if errors.Is(r.Context().Err(), context.Canceled) {
			w.WriteHeader(499)
			return
		}

		switch {
		case errors.Is(err, usecase.ErrUnauthorized):
			w.WriteHeader(http.StatusUnauthorized)
			return
		case errors.Is(err, usecase.ErrUserNotFound):
			w.WriteHeader(http.StatusNotFound)
			return
		case errors.Is(err, usecase.ErrUnexpectedError):
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
