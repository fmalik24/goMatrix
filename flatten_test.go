package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFlattenMatrix1(testHelper *testing.T) {

	falattened := getFlattenedMatrix([][]string{{"1", "2"}, {"3", "4"}, {"5", "6"}})
	if falattened != "1,2,3,4,5,6" {
		testHelper.Errorf("Not expecting this: %s", falattened)
	}
}

func TestGetFlattenMatrix2(testHelper *testing.T) {

	falattened := getFlattenedMatrix([][]string{{"1", "1", "2"}, {"3", "4", "5"}, {"6", "7", "8"}})
	if falattened != "1,1,2,3,4,5,6,7,8" {
		testHelper.Errorf("Not expecting this: %s", falattened)
	}
}

func TestFlatten(testHelper *testing.T) {

	csvFile, multipartWriter := createMultipartFormData(testHelper, []byte("1,2,3\n4,5,6\n7,8,9\n"))
	request, err := http.NewRequest("POST", "/flatten", &csvFile)
	if err != nil {
		testHelper.Fatal(err)
	}

	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(flatten)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		testHelper.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "1,2,3,4,5,6,7,8,9"
	if responseRecorder.Body.String() != expected {
		testHelper.Errorf("handler returned unexpected body: got %v want %v",
			responseRecorder.Body.String(), expected)
	}
}

func TestFlattenWrongFile(testHelper *testing.T) {

	csvFile, multipartWriter := createMultipartFormDataWithWrongFileName(testHelper, []byte("1,2\n4,5,6\n7,8,9\n"))
	request, err := http.NewRequest("POST", "/flatten", &csvFile)
	if err != nil {
		testHelper.Fatal(err)
	}

	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(flatten)

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

func TestFlattenWrongData(testHelper *testing.T) {

	csvFile, multipartWriter := createMultipartFormData(testHelper, []byte("1,2\n4,5,6\n7,8,9\n"))
	request, err := http.NewRequest("POST", "/flatten", &csvFile)
	if err != nil {
		testHelper.Fatal(err)
	}

	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(flatten)

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
