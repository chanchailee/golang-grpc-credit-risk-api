/*
generate-fico-records main program

Overview:

This main program would generate fico record by calling pkg/mockrecords

command:
 go run main.go



*/
package main

import (
	"log"

	"github.com/chanchailee/golang-grpc-credit-risk-api/pkg/mockrecords"
)

var (
	filename = "../mock-fico-records.csv"
	size     = 50
)

func main() {
	err := mockrecords.Generate(size, filename)
	if err != nil {
		log.Fatalf("error: can't generate csv file %+v", err)
	}
}
