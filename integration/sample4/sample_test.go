package sample1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pepnova-9/err-handling-sample/controller"
	"github.com/pepnova-9/err-handling-sample/router"
)

const urlPrefix = "/sample4"

func TestSample_OK(t *testing.T) {
	router := router.Router()

	reqBody := controller.UpdateUserRequestBody{
		Name: "hoge",
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPut, urlPrefix+"/users/hoge", bytes.NewBuffer(jsonBody))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	fmt.Printf("Response Status: %v \n", rec.Result().Status)
	fmt.Printf("Response Body: %v \n", rec.Body.String())
}

func TestSample_Unauthorized(t *testing.T) {
	router := router.Router()

	reqBody := controller.UpdateUserRequestBody{
		Name: "hoge",
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPut, urlPrefix+"/users/unauthorized", bytes.NewBuffer(jsonBody))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	fmt.Printf("Response Status: %v \n", rec.Result().Status)
	fmt.Printf("Response Body: %v \n", rec.Body.String())
}

func TestSample_NotFound(t *testing.T) {
	router := router.Router()

	reqBody := controller.UpdateUserRequestBody{
		Name: "hoge",
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPut, urlPrefix+"/users/not_found", bytes.NewBuffer(jsonBody))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	fmt.Printf("Response Status: %v \n", rec.Result().Status)
	fmt.Printf("Response Body: %v \n", rec.Body.String())
}

func TestSample_Unknown(t *testing.T) {
	router := router.Router()

	reqBody := controller.UpdateUserRequestBody{
		Name: "hoge",
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPut, urlPrefix+"/users/unknown", bytes.NewBuffer(jsonBody))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	fmt.Printf("Response Status: %v \n", rec.Result().Status)
	fmt.Printf("Response Body: %v \n", rec.Body.String())
}

func TestSample_RepositoryPanic(t *testing.T) {
	router := router.Router()

	reqBody := controller.UpdateUserRequestBody{
		Name: "hoge",
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPut, urlPrefix+"/users/repository_panic", bytes.NewBuffer(jsonBody))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	fmt.Printf("Response Status: %v \n", rec.Result().Status)
	fmt.Printf("Response Body: %v \n", rec.Body.String())
}
