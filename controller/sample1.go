package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pepnova-9/err-handling-sample/usecase"
)

type Sample1 struct{}

func (s *Sample1) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	sampleUsecase := &usecase.Sample1{}

	vars := mux.Vars(r)
	userID := vars["userID"]

	requestBody := &UpdateUserRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := sampleUsecase.UpdateUserUsecase(r.Context(), usecase.UpdateUserInput{UserID: userID, Name: requestBody.Name})
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
		}
		w.WriteHeader(http.StatusInternalServerError)
		// 500エラーの時はログにエラーを出力する
		log.Printf("%v", fmt.Errorf("failed to UpdateUserUsecase: %w", err))
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
