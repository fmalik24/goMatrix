package main

import (
	"errors"
	"mime/multipart"
	"testing"
)

type mockRequestContext struct {
	mockFormFile func(fileName string) (multipart.File, *multipart.FileHeader, error)
}

func (m *mockRequestContext) FormFile(filename string) (multipart.File, *multipart.FileHeader, error) {
	if m.mockFormFile != nil {
		return nil, nil, errors.New("failed to open file")
	}
	return nil, nil, nil
}

func TestGetRecordsExceptionMessage(testHelper *testing.T) {

	// Arrange:
	mockRequestClient := &mockRequestContext{
		mockFormFile:  func(fileName string) (multipart.File, *multipart.FileHeader, error) {
			// this time we want to mock our a failure scenario
			return nil, nil, errors.New("Failed to open file")
		},
	}

	// Act: invole getMatrixFromRequest
	_, err := getMatrixFromRequest(mockRequestClient)

	// Assert:
	if err == nil {
		testHelper.Errorf("Expecting an error here")
	}
	if (err.Error() != "we are unable to process your request. can you try again with \nfile=@matrix.csv") {
		testHelper.Errorf("Unknown error!! %s", err.Error())
	}
}
