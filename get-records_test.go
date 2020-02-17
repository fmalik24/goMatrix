package main

import (
	"errors"
	"mime/multipart"
	"testing"
)

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
}




// func TestGetRecords(testHelper *testing.T) {

// 	var fileBuffer bytes.Buffer
// 	var err error
// 	multipartWriter := multipart.NewWriter(&fileBuffer)
// 	var fw io.Writer
// 	if fw, err = multipartWriter.CreateFormFile("file", "fileName.txt"); err != nil {
// 		testHelper.Errorf("Error creating writer: %v", err)
// 	}
// 	csvFileContents := []byte("1,2,3\n4,5,6\n7,8,9\n")

// 	body := new(bytes.Buffer)
// 	writer := multipart.NewWriter(body)
// 	part, err := writer.CreateFormFile("file", "ddd.txt")

// 	part.Write(csvFileContents)
// 	fw.Write(csvFileContents)

// 	multipartWriter.Close()

// 	values := map[string]io.Reader{
// 		"file":  io.Reader(bytes.NewReader(csvFileContents)),
// 		"other": strings.NewReader("hello world!"),
// 	}

// 	mat := io.Reader(bytes.NewReader(csvFileContents))

// 	mockRequestContext := &mockRequestContext{
// 		mockFormFile: func(fileName string) (multipart.File, *multipart.FileHeader, error) {
// 			// this time we want to mock our a failure scenario
// 			return io.Reader(bytes.NewReader(csvFileContents)), nil, nil
// 		},
// 	}

// 	// perform the function call we want to test
// 	_, err = getMatrixFromRequest(mockRequestContext)
// 	if err == nil {
// 		testHelper.Errorf("Not expecting this")
// 	}

// 	// ... check for panic
// }
