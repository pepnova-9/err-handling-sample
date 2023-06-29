package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/pepnova-9/err-handling-sample/errwrapper"

	"github.com/gorilla/mux"
	"github.com/pepnova-9/err-handling-sample/usecase"
)

type Sample4 struct{}

func (s *Sample4) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	sampleUsecase := &usecase.Sample4{}

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
		log.Printf("%v", err)
		// StackErrorを採用するならここで出力
		if stackErr := (*errwrapper.StackError)(nil); errors.As(err, &stackErr) {
			fmt.Println("stack:", string(stackErr.Stack))
		}
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
