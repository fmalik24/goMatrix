package main

import (
	"testing"
)

func TestGetSumOfMatrixEnteries1(testHelper *testing.T) {

	sum, _ := getSumOfMatrixEnteries([][]string{{"0", "1", "2"}, {"3", "4", "5"}, {"6", "7", "8"}})
	if sum != "36" {
		testHelper.Errorf("Not expecting this: %s", sum)
	}
}
