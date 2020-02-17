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
		return nil, nil, errors.New("failed to open file")
	}
	return nil, nil, nil
}

func TestGetRecordsExceptionMessage(testHelper *testing.T) {
	mockRequestClient := &mockRequestContext{
		mockFormFile: func(fileName string) (multipart.File, *multipart.FileHeader, error) {
			// this time we want to mock our a failure scenario
			return nil, nil, errors.New("Failed to open file")
		},
	}

	// perform the function call we want to test
	_, err := getMatrixFromRequest(mockRequestClient)
	if err == nil {
		testHelper.Errorf("Not expecting this")
	}
}
