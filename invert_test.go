package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTransposedMatrix1(testHelper *testing.T) {

	trasnposedMatrix := getTransposedMatrix([][]string{{"0", "1", "2"}, {"3", "4", "5"}, {"6", "7", "8"}})
	if trasnposedMatrix != "0,3,6\n1,4,7\n2,5,8\n" {
		testHelper.Errorf("Not expecting this: %s", trasnposedMatrix)
	}
}

func TestGetTransposedMatrixError(testHelper *testing.T) {

	trasnposedMatrix := getTransposedMatrix([][]string{{"0", "1"}, {"3", "4", "3"}, {"6", "7"}})
	if trasnposedMatrix != "Invalid Entry: Row size is 3 which is not equal to column of size 2\n" {
		testHelper.Errorf("Not expecting this: %s", trasnposedMatrix)
	}
}

func TestInvert(testHelper *testing.T) {

	csvFile, multipartWriter := createMultipartFormData(testHelper, []byte("1,2,3\n4,5,6\n7,8,9\n"))
	request, err := http.NewRequest("POST", "/invert", &csvFile)
	if err != nil {
		testHelper.Fatal(err)
	}

	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(invert)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		testHelper.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "1,4,7\n2,5,8\n3,6,9\n"
	if response.Body.String() != expected {
		testHelper.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expected)
	}
}

func TestInvertBadData(testHelper *testing.T) {

	csvFile, multipartWriter := createMultipartFormData(testHelper, []byte("1,2\n4,5,6\n7,8,9\n"))
	request, err := http.NewRequest("POST", "/invert", &csvFile)
	if err != nil {
		testHelper.Fatal(err)
	}

	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(invert)

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

func TestInvertWrongFile(testHelper *testing.T) {

	csvFile, multipartWriter := createMultipartFormDataWithWrongFileName(testHelper, []byte("1,2\n4,5,6\n7,8,9\n"))
	request, err := http.NewRequest("POST", "/invert", &csvFile)
	if err != nil {
		testHelper.Fatal(err)
	}

	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(invert)

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
