package main

import (
	"testing"
)

func TestGetFlattenMatrix1(testHelper *testing.T) {

	falattened, _ := getFlattenedMatrix([][]string{{"1", "1", "2"}, {"3", "4", "5"}, {"6", "7", "8"}})
	if falattened != "1,1,2,3,4,5,6,7,8" {
		testHelper.Errorf("Not expecting this: %s", falattened)
	}
}


func TestGetFlattenMatrix2(testHelper *testing.T) {

	falattened, _ := getFlattenedMatrix([][]string{{"1", "1", "2"}, {"3", "4", "5"}, {"6", "7", "8"}})
	if falattened != "1,1,2,3,4,5,6,7,8" {
		testHelper.Errorf("Not expecting this: %s", falattened)
	}
}
