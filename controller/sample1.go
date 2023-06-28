package controller

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pepnova-9/err-handling-sample/usecase"
)

type Sample1 struct{}

type UpdateUserRequestBody struct {
	Name string `json:"name"`
}

func (s *Sample1) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	sample1 := &usecase.Sample1{}

	vars := mux.Vars(r)
	userID := vars["userID"]

	requestBody := &UpdateUserRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := sample1.UpdateUserUsecase(r.Context(), usecase.UpdateUserInput{UserID: userID, Name: requestBody.Name})
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
