package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pepnova-9/err-handling-sample/controller"
)

func TestSample1_OK(t *testing.T) {
	router := Router()

	reqBody := controller.UpdateUserRequestBody{
		Name: "hoge",
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPut, "/sample1/users/hoge", bytes.NewBuffer(jsonBody))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	fmt.Printf("Response Status: %#v \n", rec.Result().Status)
	fmt.Printf("Response Body: %v \n", rec.Body.String())
}

func TestSample1_Unauthorized(t *testing.T) {
	router := Router()

	reqBody := controller.UpdateUserRequestBody{
		Name: "hoge",
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPut, "/sample1/users/unauthorized", bytes.NewBuffer(jsonBody))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	fmt.Printf("Response Status: %#v \n", rec.Result().Status)
	fmt.Printf("Response Body: %v \n", rec.Body.String())
}

func TestSample1_NotFound(t *testing.T) {
	router := Router()

	reqBody := controller.UpdateUserRequestBody{
		Name: "hoge",
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPut, "/sample1/users/not_found", bytes.NewBuffer(jsonBody))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	fmt.Printf("Response Status: %#v \n", rec.Result().Status)
	fmt.Printf("Response Body: %v \n", rec.Body.String())
}

func TestSample1_Unknown(t *testing.T) {
	router := Router()

	reqBody := controller.UpdateUserRequestBody{
		Name: "hoge",
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPut, "/sample1/users/unknown", bytes.NewBuffer(jsonBody))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	fmt.Printf("Response Status: %#v \n", rec.Result().Status)
	fmt.Printf("Response Body: %v \n", rec.Body.String())
}
