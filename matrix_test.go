package main

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMatrix(testHelper *testing.T) {
	matrix, _ := getEcho([][]string{{"0", "1"}, {"2", "3"}, {"1", "2"}})
	if matrix != "0,1\n2,3\n1,2\n" {
		testHelper.Errorf("Not expecting this: %s", matrix)
	}
}

func TestGetTransposedMatrix(testHelper *testing.T) {

	trasnposedMatrix, _ := getTransposedMatrix([][]string{{"0", "1", "2"}, {"3", "4", "5"}, {"6", "7", "8"}})
	if trasnposedMatrix != "0,3,6\n1,4,7\n2,5,8\n" {
		testHelper.Errorf("Not expecting this: %s", trasnposedMatrix)
	}
}

func TestGetSumOfMatrixEnteries(testHelper *testing.T) {

	sum, _ := getSumOfMatrixEnteries([][]string{{"0", "1", "2"}, {"3", "4", "5"}, {"6", "7", "8"}})
	if sum != "36" {
		testHelper.Errorf("Not expecting this: %s", sum)
	}
}

func TestGetProductOfMatrixEnteries(testHelper *testing.T) {

	product, _ := getProductOfMatrixEnteries([][]string{{"1", "1", "2"}, {"3", "4", "5"}, {"6", "7", "8"}})
	if product != "40320" {
		testHelper.Errorf("Not expecting this: %s", product)
	}
}

func TestGetFlattenMatrix(testHelper *testing.T) {

	falattened, _ := getFlattenedMatrix([][]string{{"1", "1", "2"}, {"3", "4", "5"}, {"6", "7", "8"}})
	if falattened != "1,1,2,3,4,5,6,7,8" {
		testHelper.Errorf("Not expecting this: %s", falattened)
	}
}

func TestEcho(testHelper *testing.T) {

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

func TestInvert(testHelper *testing.T) {

	csvFile, multipartWriter := createMultipartFormData(testHelper)
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

func TestMultiply(testHelper *testing.T) {

	csvFile, multipartWriter := createMultipartFormData(testHelper)
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

func TestSum(testHelper *testing.T) {

	csvFile, multipartWriter := createMultipartFormData(testHelper)

	response := getFunctionFromServer(csvFile, testHelper, multipartWriter)

	// request, err := http.NewRequest("POST", "/sum", &csvFile)
	// if err != nil {
	// 	testHelper.Fatal(err)
	// }

	// request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	// response := httptest.NewRecorder()
	// handler := http.HandlerFunc(sumOfMatrixEnteries)

	// // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// // directly and pass in our Request and ResponseRecorder.
	// handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		testHelper.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "45"
	if response.Body.String() != expected {
		testHelper.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expected)
	}
}

func getFunctionFromServer(csvFile bytes.Buffer, testHelper *testing.T, multipartWriter *multipart.Writer, response *ResponseRecorder) {
	request, err := http.NewRequest("POST", "/sum", &csvFile)
	if err != nil {
		testHelper.Fatal(err)
	}

	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	response = httptest.NewRecorder()
	handler := http.HandlerFunc(sumOfMatrixEnteries)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(response, request)
}

func TestFlatten(testHelper *testing.T) {

	csvFile, multipartWriter := createMultipartFormData(testHelper)
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

type mockRequestContext struct {
	mockFormFile func(fileName string) (multipart.File, *multipart.FileHeader, error)
}

// 2
func (m *mockRequestContext) FormFile(filename string) (multipart.File, *multipart.FileHeader, error) {
	if m.mockFormFile != nil {
		return nil, nil, errors.New("failed to increment")
	}
	return nil, nil, nil
}

func TestGetRecordsExceptionMessage(testHelper *testing.T) {
	mockRequestContext := &mockRequestContext{
		mockFormFile: func(fileName string) (multipart.File, *multipart.FileHeader, error) {
			// this time we want to mock our a failure scenario
			return nil, nil, errors.New("Failed to ")
		},
	}

	// perform the function call we want to test
	_, err := getMatrixFromRequest(mockRequestContext)
	if err == nil {
		testHelper.Errorf("Not expecting this")
	}

	// ... check for panic
}

// func TestGetRecords(testHelper *testing.T) {

// 	data, w := createMultipartFormData(testHelper)

// 	mockRequestContext := &mockRequestContext{
// 		mockFormFile: func(fileName string) (multipart.File, *multipart.FileHeader, error) {
// 			// this time we want to mock our a failure scenario
// 			return &data, nil, nil
// 		},
// 	}

// 	// perform the function call we want to test
// 	_, err := getMatrixFromRequest(mockRequestContext)
// 	if err == nil {
// 		testHelper.Errorf("Not expecting this")
// 	}

// 	// ... check for panic
// }
