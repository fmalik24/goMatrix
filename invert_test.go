package main

import (
	"testing"
)

func TestGetTransposedMatrix1(testHelper *testing.T) {

	trasnposedMatrix, _ := getTransposedMatrix([][]string{{"0", "1", "2"}, {"3", "4", "5"}, {"6", "7", "8"}})
	if trasnposedMatrix != "0,3,6\n1,4,7\n2,5,8\n" {
		testHelper.Errorf("Not expecting this: %s", trasnposedMatrix)
	}
}
