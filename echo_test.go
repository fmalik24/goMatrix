package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMatrix1(testHelper *testing.T) {
	matrix := getEcho([][]string{{"0", "1"}, {"2", "3"}, {"1", "2"}})
	if matrix != "0,1\n2,3\n1,2\n" {
		testHelper.Errorf("Not expecting this: %s", matrix)
	}
}

func TestEcho1(testHelper *testing.T) {

	// Arrange
	csvFile, multipartWriter := createMultipartFormData(testHelper, []byte("1,2,3\n4,5,6\n7,8,9\n"))

	//Act
	request, err := http.NewRequest("POST", "/echo", &csvFile)
	if err != nil {
		testHelper.Fatal(err)
	}
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(echo)
	handler.ServeHTTP(response, request)
	if status := response.Code; status != http.StatusOK {
		testHelper.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Assert.
	expected := "1,2,3\n4,5,6\n7,8,9\n"
	if response.Body.String() != expected {
		testHelper.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expected)
	}
}

func TestEchoWrongFile(testHelper *testing.T) {

	csvFile, multipartWriter := createMultipartFormDataWithWrongFileName(testHelper, []byte("1,2\n4,5,6\n7,8,9\n"))
	request, err := http.NewRequest("POST", "/echo", &csvFile)
	if err != nil {
		testHelper.Fatal(err)
	}

	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(echo)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		testHelper.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "We are unable to process your request. Can you try again with \nfile=@matrix.csv\n"
	if response.Body.String() != expected {
		testHelper.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expected)
	}
}

func TestEchoWrongData(testHelper *testing.T) {

	csvFile, multipartWriter := createMultipartFormData(testHelper, []byte("1,2\n4,5,6\n7,8,9\n"))
	request, err := http.NewRequest("POST", "/echo", &csvFile)
	if err != nil {
		testHelper.Fatal(err)
	}

	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(echo)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		testHelper.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "We are having a hard time reading the file. Can you make sure its a square and try again\n"
	if response.Body.String() != expected {
		testHelper.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expected)
	}
}
