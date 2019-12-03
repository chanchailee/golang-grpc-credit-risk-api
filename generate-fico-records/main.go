/*
generate-fico-records main program

Overview:

This main program would generate fico record by calling pkg/mockrecords

command:
 go run main.go



*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/chanchailee/golang-grpc-credit-risk-api/pkg/mockrecords"
)

var (
	size           = 50
	errMissingArgs = fmt.Errorf("error : missing argument")
	errInvalidArgs = fmt.Errorf("error : invalid argument")
)

func readArgs(args []string) (string, string, error) {
	length := len(args)

	switch {
	case length < 3:
		return "", "", errMissingArgs
	case length == 3:
		return strings.TrimSpace(args[1]), strings.TrimSpace(args[2]), nil
	default:
		return "", "", errInvalidArgs
	}
}

func main() {
	filename, s, err := readArgs(os.Args)
	if err != nil {
		log.Fatalf("read argument error: %s", err)
	}

	size, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("convert string to int error: %s", err)
	}

	err = mockrecords.Generate(size, filename)
	if err != nil {
		log.Fatalf("error: can't generate csv file %+v", err)
	}
}
