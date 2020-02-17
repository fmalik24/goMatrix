package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProductOfMatrixEnteries1(testHelper *testing.T) {

	product := getProductOfMatrixEnteries([][]string{{"1", "1", "2"}, {"3", "4", "5"}, {"6", "7", "8"}})
	if product != 40320 {
		testHelper.Errorf("Not expecting this: %d", product)
	}
}

func TestGetProductOfMatrixEnteriesWith0(testHelper *testing.T) {

	product := getProductOfMatrixEnteries([][]string{{"0", "1", "2"}, {"3", "4", "5"}, {"6", "7", "8"}})
	if product != 0 {
		testHelper.Errorf("Not expecting this: %d", product)
	}
}

func TestMultiply(testHelper *testing.T) {

	csvFile, multipartWriter := createMultipartFormData(testHelper, []byte("1,2,3\n4,5,6\n7,8,9\n"))
	request, err := http.NewRequest("GET", "/multiply", &csvFile)
	if err != nil {
		testHelper.Fatal(err)
	}

	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(productOfMatrixEnteries)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, request)

	if status := rr.Code; status != http.StatusOK {
		testHelper.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "362880"
	if rr.Body.String() != expected {
		testHelper.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestMulrongFile(testHelper *testing.T) {

	csvFile, multipartWriter := createMultipartFormDataWithWrongFileName(testHelper, []byte("1,2\n4,5,6\n7,8,9\n"))
	request, err := http.NewRequest("POST", "/product", &csvFile)
	if err != nil {
		testHelper.Fatal(err)
	}

	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(productOfMatrixEnteries)

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

func TestMulrongData(testHelper *testing.T) {

	csvFile, multipartWriter := createMultipartFormData(testHelper, []byte("1,2\n4,5,6\n7,8,9\n"))
	request, err := http.NewRequest("POST", "/product", &csvFile)
	if err != nil {
		testHelper.Fatal(err)
	}

	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(productOfMatrixEnteries)

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
