package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMatrix(testHelper *testing.T) {

	// Arrange: A 2d string array is being initialized with test data
	testMatrix := [][]string{{"0", "1"}, {"2", "3"}, {"1", "2"}}

	// Act: The getEcho function is being called wiht test data
	matrix := getEcho(testMatrix)

	// Assert: The response is as per expectation
	if matrix != "0,1\n2,3\n1,2\n" {
		testHelper.Errorf("Not expecting this: %s", matrix)
	}
}

func TestEcho(testHelper *testing.T) {

	// Arrange:
	// The data neccessary to call the end point
	// csvFile created in memory with the given testMatrix
	testMatrix := []byte("1,2,3\n4,5,6\n7,8,9\n")
	csvFile, multipartWriter := createMultipartFormData(testHelper, testMatrix)

	// The function mapped to the url and the http action
	handlerFunction := http.HandlerFunc(echo)
	url := "/echo"
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

	// The response body is as per expectation
	expected := "1,2,3\n4,5,6\n7,8,9\n"
	if response.Body.String() != expected {
		testHelper.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expected)
	}
}

func TestEchoWrongFile(testHelper *testing.T) {

	// Arrange:
	// The data neccessary to call the end point
	// csvFile created in memory with the given testMatrix
	testMatrix := []byte("1,2,3\n4,5,6\n7,8,9\n")
	csvFile, multipartWriter := createMultipartFormDataWithWrongFileName(testHelper, testMatrix)

	// The function mapped to the url and the http action
	handlerFunction := http.HandlerFunc(echo)
	url := "/echo"
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

func TestEchoWrongData(testHelper *testing.T) {

	// Arrange:
	// The data neccessary to call the end point
	// csvFile created in memory with the given testMatrix
	testMatrix := []byte("1,2\n4,5,6\n7,8,9\n")
	csvFile, multipartWriter := createMultipartFormData(testHelper, testMatrix)

	// The function mapped to the url and the http action
	handlerFunction := http.HandlerFunc(echo)
	url := "/echo"
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
	// Check the response body is what we expect.
	expected := "we are having a hard time reading the file. can you make sure its a square and try again: \nrecord on line 2: wrong number of fields"
	if response.Body.String() != expected {
		testHelper.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expected)
	}
}
