package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMatrix1(testHelper *testing.T) {
	matrix, _ := getEcho([][]string{{"0", "1"}, {"2", "3"}, {"1", "2"}})
	if matrix != "0,1\n2,3\n1,2\n" {
		testHelper.Errorf("Not expecting this: %s", matrix)
	}
}

func TestEcho1(testHelper *testing.T) {

	csvFile, multipartWriter := createMultipartFormData(testHelper)
	request, err := http.NewRequest("POST", "/echo", &csvFile)
	if err != nil {
		testHelper.Fatal(err)
	}

	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	//Recording the resposne
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
	expected := "1,2,3\n4,5,6\n7,8,9\n"
	if response.Body.String() != expected {
		testHelper.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expected)
	}
}

func createMultipartFormData(testHelper *testing.T) (bytes.Buffer, *multipart.Writer) {
	var fileBuffer bytes.Buffer
	var err error
	multipartWriter := multipart.NewWriter(&fileBuffer)
	var fw io.Writer
	if fw, err = multipartWriter.CreateFormFile("file", "fileName.txt"); err != nil {
		testHelper.Errorf("Error creating writer: %v", err)
	}
	csvFileContents := []byte("1,2,3\n4,5,6\n7,8,9\n")
	fw.Write(csvFileContents)

	multipartWriter.Close()
	return fileBuffer, multipartWriter
}
