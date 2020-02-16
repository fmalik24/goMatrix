package main

import (
	"testing"
)

func TestGetProductOfMatrixEnteries1(testHelper *testing.T) {

	product, _ := getProductOfMatrixEnteries([][]string{{"1", "1", "2"}, {"3", "4", "5"}, {"6", "7", "8"}})
	if product != "40320" {
		testHelper.Errorf("Not expecting this: %s", product)
	}
}
