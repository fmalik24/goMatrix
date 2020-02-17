package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTransposedMatrix(testHelper *testing.T) {

	// Arrage:
	testData := [][]string{{"0", "1", "2"}, {"3", "4", "5"}, {"6", "7", "8"}}

	// Act:
	trasnposedMatrix := getTransposedMatrix(testData)

	// Assert:
	if trasnposedMatrix != "0,3,6\n1,4,7\n2,5,8\n" {
		testHelper.Errorf("Not expecting this: %s", trasnposedMatrix)
	}
}

func TestGetTransposedMatrixError(testHelper *testing.T) {

	// Arrange:
	testData := [][]string{{"0", "1"}, {"3", "4", "3"}, {"6", "7"}}

	// Act:
	trasnposedMatrix := getTransposedMatrix(testData)

	// Assert:
	if trasnposedMatrix != "Invalid Entry: Row size is 3 which is not equal to column of size 2\n" {
		testHelper.Errorf("Not expecting this: %s", trasnposedMatrix)
	}
}

func TestInvert(testHelper *testing.T) {

	// Arrange:
	// The data neccessary to call the end point
	// csvFile created in memory with the given testMatrix
	testMatrix := []byte("1,2,3\n4,5,6\n7,8,9\n")
	csvFile, multipartWriter := createMultipartFormData(testHelper, testMatrix)

	// The function mapped to the url and the http action
	handlerFunction := http.HandlerFunc(invert)
	url := "/Invert"
	httpVerb := "POST"

	// Setup the request
	request, err := http.NewRequest(httpVerb, url, &csvFile)
	if err != nil {
		testHelper.Fatal(err)
	}

	// Setup the Content-Type to be of MultipartFomData
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// Setup the response recorder
	response := httptest.NewRecorder()

	//Act:
	// Trigger HTTP request with the given data
	handlerFunction.ServeHTTP(response, request)

	// Assert:
	// The status code is as per expectation

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

	// Arrange:
	// The data neccessary to call the end point
	// csvFile created in memory with the given testMatrix
	testMatrix := []byte("1,2,3\n4,5,6\n7,8\n")
	csvFile, multipartWriter := createMultipartFormData(testHelper, testMatrix)

	// The function mapped to the url and the http action
	handlerFunction := http.HandlerFunc(invert)
	url := "/Invert"
	httpVerb := "POST"

	// Setup the request
	request, err := http.NewRequest(httpVerb, url, &csvFile)
	if err != nil {
		testHelper.Fatal(err)
	}

	// Setup the Content-Type to be of MultipartFomData
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// Setup the response recorder
	response := httptest.NewRecorder()

	//Act:
	// Trigger HTTP request with the given data
	handlerFunction.ServeHTTP(response, request)

	// Assert:
	// The status code is as per expectation

	if status := response.Code; status != http.StatusOK {
		testHelper.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "we are having a hard time reading the file. can you make sure its a square and try again: \nrecord on line 3: wrong number of fields"
	if response.Body.String() != expected {
		testHelper.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expected)
	}
}

func TestInvertWrongFile(testHelper *testing.T) {

	// Arrange:
	// The data neccessary to call the end point
	// csvFile created in memory with the given testMatrix
	testMatrix := []byte("1,2,3\n4,5,6\n7,8,9\n")
	csvFile, multipartWriter := createMultipartFormDataWithWrongFileName(testHelper, testMatrix)

	// The function mapped to the url and the http action
	handlerFunction := http.HandlerFunc(invert)
	url := "/Invert"
	httpVerb := "POST"

	// Setup the request
	request, err := http.NewRequest(httpVerb, url, &csvFile)
	if err != nil {
		testHelper.Fatal(err)
	}

	// Setup the Content-Type to be of MultipartFomData
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// Setup the response recorder
	response := httptest.NewRecorder()

	//Act:
	// Trigger HTTP request with the given data
	handlerFunction.ServeHTTP(response, request)

	// Assert:
	// The status code is as per expectation
	if status := response.Code; status != http.StatusOK {
		testHelper.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "we are unable to process your request. can you try again with \nfile=@matrix.csv"
	if response.Body.String() != expected {
		testHelper.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expected)
	}
}
