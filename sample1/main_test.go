package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pepnova-9/err-handling-sample/sample1/controller"
)

func TestGetUserEndpoint_OK(t *testing.T) {
	router := Router()

	req := httptest.NewRequest("GET", "/users/hoge", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	fmt.Printf("Response Status: %#v \n", rec.Result().Status)
	fmt.Printf("Response Body: %v \n", rec.Body.String())

}

func TestGetUserEndpoint_Unauthorized(t *testing.T) {
	router := Router()

	req := httptest.NewRequest("GET", "/users/unauthorized", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	fmt.Printf("Response Status: %v \n", rec.Result().Status)
	fmt.Printf("Response Body: %v \n", rec.Body.String())

}

func TestGetUserEndpoint_NotFound(t *testing.T) {
	router := Router()

	req := httptest.NewRequest("GET", "/users/not_found", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	fmt.Printf("Response Status: %v \n", rec.Result().Status)
	fmt.Printf("Response Body: %v \n", rec.Body.String())

}

func TestGetUserEndpoint_Unknown(t *testing.T) {
	router := Router()

	req := httptest.NewRequest("GET", "/users/unknown", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	fmt.Printf("Response Status: %v \n", rec.Result().Status)
	fmt.Printf("Response Body: %v \n", rec.Body.String())

}

func TestUpdateUserEndpoint_OK(t *testing.T) {
	router := Router()

	reqBody := controller.CreaetUserRequestBody{
		Name: "hoge",
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/users/hoge", bytes.NewBuffer(jsonBody))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	fmt.Printf("Response Status: %#v \n", rec.Result().Status)
	fmt.Printf("Response Body: %v \n", rec.Body.String())

}
