package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var ErrInvalidNumber = errors.New("the value is not a number")

func main() {

	if len(os.Args) < 2 {
		log.Fatal("One arg is required for this program")
	}

	//checkNumber(getValues())
	msg, err := checkNumber(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(msg)
}

func getValues() string {
	if len(os.Args)-1 < 1 {
		fmt.Println("One arg is requerid for this program")
		os.Exit(1)
	}
	value := os.Args[1]
	return value
}

func checkNumber(value string) (string, error) {
	fValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return "", ErrInvalidNumber
		//fmt.Printf("The value <%s> is not number\n", value)
		//os.Exit(1)
	}
	return fmt.Sprintf("The value <%f> is a valid number\n", fValue), nil
}
