package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"testing"
)

func createMultipartFormData(testHelper *testing.T, csvFileContents []byte) (bytes.Buffer, *multipart.Writer) {
	var fileBuffer bytes.Buffer
	var err error
	multipartWriter := multipart.NewWriter(&fileBuffer)
	var fw io.Writer
	if fw, err = multipartWriter.CreateFormFile("file", "iShallNotExist.txt"); err != nil {
		testHelper.Errorf("Error creating writer: %v", err)
	}
	fw.Write(csvFileContents)

	multipartWriter.Close()
	return fileBuffer, multipartWriter
}

func createMultipartFormDataWithWrongFileName(testHelper *testing.T, csvFileContents []byte) (bytes.Buffer, *multipart.Writer) {
	var fileBuffer bytes.Buffer
	var err error
	multipartWriter := multipart.NewWriter(&fileBuffer)
	var fw io.Writer
	if fw, err = multipartWriter.CreateFormFile("incorrect", "iShallNotExist.txt"); err != nil {
		testHelper.Errorf("Error creating writer: %v", err)
	}
	fw.Write(csvFileContents)

	
	multipartWriter.Close()

	return fileBuffer, multipartWriter
}
