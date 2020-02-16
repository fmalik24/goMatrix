package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
	"strconv"
	"mime/multipart"
	"errors"

)

type requestClient interface {
	FormFile(key string) (multipart.File, *multipart.FileHeader, error)
}


func getMatrixFromRequest(request requestClient) ([][]string, error) {
	file, _, err := request.FormFile("file")
		if (err != nil) {
			fmt.Println((fmt.Sprintf("\nInfo: User has provided incorrect form file name %s", err.Error())))
			return nil, errors.New("File not opened: We are unable to process your request. Can you try again with \nfile=@matrix.csv\n\n")
		}
		defer file.Close()
		records, err := csv.NewReader(file).ReadAll()
		if (err != nil) {
			return nil, errors.New("\nFileOpened: We are unable to process your request. Can you try again with \nfile=@matrix.csv\n\n")
		}

		return records, nil;
}



func getEcho(records [][]string) (string, error) {

		var accumulator string
		for _, row := range records {
			accumulator = fmt.Sprintf("%s%s\n", accumulator, strings.Join(row, ","))
		}
		return accumulator, nil;

}

func getTransposedMatrix(records [][]string) (string, error) {
	 var transposeMatrix [3][3]string
	 var accumulator string
		for i := 0;  i<len(records); i++ {
           for j :=0; j<len(records[0]); j++ {
           		transposeMatrix[i][j] = records[j][i]
  			}
  			accumulator = fmt.Sprintf("%s%s\n", accumulator, strings.Join(transposeMatrix[i][:], ","))
 		}
	 	return accumulator, nil;

}

func getFlattenedMatrix(records[][]string) (string, error) {
	 var strs []string
    	for _, v1 := range records {
        	s := strings.Join(v1, ",")
        	strs = append(strs, s)

    	}
    	s := strings.Join(strs, ",")
    	return s, nil;
}

func getSumOfMatrixEnteries(records[][] string) (string, error) {

			var sum = 0;
   			for i := 0;  i<len(records); i++ {
	           for j :=0; j<len(records[0]); j++ {
    		       i1, err := strconv.Atoi(records[j][i])
					if err == nil {
						 sum += i1
						}
					}
 			}

		return strconv.Itoa(sum), nil;

}

func getProductOfMatrixEnteries(records [][]string) (string, error) {

	 		var product = 1;
   			for i := 0;  i<len(records); i++ {
	           for j :=0; j<len(records[0]); j++ {
    		       i1, err := strconv.Atoi(records[j][i])
					if err == nil {
						 product *= i1
						}
					}
 			}
 			return strconv.Itoa(product), nil
}


func echo(responseWriter http.ResponseWriter, request *http.Request) {
	var matrix, error = getMatrixFromRequest(request)
	if (error != nil) {
		fmt.Fprint(responseWriter, error.Error())
		return
	}
	var echo, err = getEcho(matrix) 
	if (err != nil) {
		fmt.Fprint(responseWriter, error.Error())
		return
	}

	fmt.Fprint(responseWriter, echo)	 	
}

func invert(responseWriter http.ResponseWriter, request *http.Request) {
	
	var matrix, error = getMatrixFromRequest(request)
	if (error != nil) {
		fmt.Fprint(responseWriter, error.Error())
		return
	}
	var invert, err = getTransposedMatrix(matrix) 
	if (err != nil) {
		fmt.Fprint(responseWriter, error.Error())
		return
	}

	fmt.Fprint(responseWriter, invert)	 	
}


func flatten(responseWriter http.ResponseWriter, request *http.Request) {
	
	var matrix, error = getMatrixFromRequest(request)
	if (error != nil) {
		fmt.Fprint(responseWriter, error.Error())
		return
	}

	
	var flattened, err = getFlattenedMatrix(matrix) 
	if (err != nil) {
		fmt.Fprint(responseWriter, error.Error())
		return
	}

	fmt.Fprint(responseWriter, flattened)	 	
	 	
}


func sumOfMatrixEnteries(responseWriter http.ResponseWriter, request *http.Request) {
	
	var matrix, error = getMatrixFromRequest(request)
	if (error != nil) {
		fmt.Fprint(responseWriter, error.Error())
		return
	}
	
	var sum, err = getSumOfMatrixEnteries(matrix) 
	if (err != nil) {
		fmt.Fprint(responseWriter, error.Error())
		return
	}

	fmt.Fprint(responseWriter, sum)	 	
	 	
}

func productOfMatrixEnteries(responseWriter http.ResponseWriter, request *http.Request) {

	var matrix, error = getMatrixFromRequest(request)
	if (error != nil) {
		fmt.Fprint(responseWriter, error.Error())
		return
	}

	var product, err = getProductOfMatrixEnteries(matrix) 
	if (err != nil) {
		fmt.Fprint(responseWriter, error.Error())
		return
	}

	fmt.Fprint(responseWriter, product)	 	
}

func main() {

	http.HandleFunc("/echo", echo)
	http.HandleFunc("/invert", invert)
	http.HandleFunc("/multiply", productOfMatrixEnteries)
	http.HandleFunc("/sum", sumOfMatrixEnteries)		
	http.HandleFunc("/flatten", flatten)
	
	http.ListenAndServe(":8080", nil)
}