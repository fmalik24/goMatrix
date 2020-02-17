package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

// Using an interfact for Http Request Client
type requestClient interface {
	FormFile(key string) (multipart.File, *multipart.FileHeader, error)
}

// getMatrixFromRequest takes in the request and provides with the matrix or an error
func getMatrixFromRequest(request requestClient) ([][]string, error) {
	file, _, err := request.FormFile("file")
	if err != nil {
		fmt.Println((fmt.Sprintf("info: User has provided incorrect form file name %s", err.Error())))
		return nil, errors.New("we are unable to process your request. can you try again with \nfile=@matrix.csv")
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, fmt.Errorf("we are having a hard time reading the file. can you make sure its a square and try again: \n%s", err.Error())
	}

	return records, nil
}

// getEcho takes in a 2d string array and returns it in string format
func getEcho(records [][]string) string {

	var accumulator string
	for _, row := range records {
		accumulator = fmt.Sprintf("%s%s\n", accumulator, strings.Join(row, ","))
	}
	return accumulator

}

// getInveretedMatrix takes in a 2d string array and inverts the matrix returns it as a string
func getInvertedMatrix(records [][]string) string {

	rowLength := len(records)
	inveretedMatrix := make([][]string, rowLength)

	var accumulator string
	for i := 0; i < len(records); i++ {
		coloumnLength := len(records[i])
		if rowLength != coloumnLength {
			return fmt.Sprintf("Invalid Entry: Row size is %d which is not equal to column of size %d\n", rowLength, coloumnLength)
		}
		inveretedMatrix[i] = make([]string, len(records[i]))
		copy(inveretedMatrix[i], records[i])
		for j := 0; j < len(records[0]); j++ {
			inveretedMatrix[i][j] = records[j][i]
		}
		accumulator = fmt.Sprintf("%s%s\n", accumulator, strings.Join(inveretedMatrix[i][:], ","))

	}
	return accumulator

}

// getFlattenedMatrix takes in a 2d string array and inverts the matrix returns it as a string
func getFlattenedMatrix(records [][]string) string {
	var strs []string
	for _, row := range records {
		value := strings.Join(row, ",")
		strs = append(strs, value)

	}
	flattened := strings.Join(strs, ",")
	return flattened
}

// getSumOfMatrixEnteries takes in a 2d string array and sums the matrix enteries returns it as a integer
func getSumOfMatrixEnteries(records [][]string) int {

	var sum = 0
	for i := 0; i < len(records); i++ {
		for j := 0; j < len(records[0]); j++ {
			entry, _ := strconv.Atoi(records[i][j])
			sum += entry
		}
	}

	return sum

}

//  getProductOfMatrixEnteries takes in a 2d string array and muliiplies the matrix enteries returns it as a number
func getProductOfMatrixEnteries(records [][]string) int {

	var product = 1
	for i := 0; i < len(records); i++ {
		for j := 0; j < len(records[0]); j++ {
			entry, _ := strconv.Atoi(records[i][j])
			if records[i][j] == "0" {
				return 0
			}
			product *= entry
		}
	}
	return product
}

func echo(responseWriter http.ResponseWriter, request *http.Request) {
	var matrix, error = getMatrixFromRequest(request)
	if error != nil {
		fmt.Fprint(responseWriter, error.Error())
		return
	}
	fmt.Fprint(responseWriter, getEcho(matrix))
}

func invert(responseWriter http.ResponseWriter, request *http.Request) {

	var matrix, error = getMatrixFromRequest(request)
	if error != nil {
		fmt.Fprint(responseWriter, error.Error())
		return
	}

	fmt.Fprint(responseWriter, getInvertedMatrix(matrix))
}

func flatten(responseWriter http.ResponseWriter, request *http.Request) {

	var matrix, error = getMatrixFromRequest(request)
	if error != nil {
		fmt.Fprint(responseWriter, error.Error())
		return
	}
	fmt.Fprint(responseWriter, getFlattenedMatrix(matrix))

}

func sumOfMatrixEnteries(responseWriter http.ResponseWriter, request *http.Request) {

	var matrix, error = getMatrixFromRequest(request)
	if error != nil {
		fmt.Fprint(responseWriter, error.Error())
		return
	}
	fmt.Fprint(responseWriter, getSumOfMatrixEnteries(matrix))

}

func productOfMatrixEnteries(responseWriter http.ResponseWriter, request *http.Request) {

	var matrix, error = getMatrixFromRequest(request)
	if error != nil {
		fmt.Fprint(responseWriter, error.Error())
		return
	}

	fmt.Fprint(responseWriter, getProductOfMatrixEnteries(matrix))
}

func main() {

	http.HandleFunc("/echo", echo)
	http.HandleFunc("/invert", invert)
	http.HandleFunc("/multiply", productOfMatrixEnteries)
	http.HandleFunc("/sum", sumOfMatrixEnteries)
	http.HandleFunc("/flatten", flatten)

	http.ListenAndServe(":8080", nil)
}
