package main

import (
	"fmt"
	"os"
	"strconv"

	fizzbuzz "github.com/carrpet/backend-developer-tests/unit-testing/pkg"
)

func main() {
	fmt.Println("SP// Backend Developer Test - FizzBuzz")
	fmt.Println()

	// Set default argument values
	totalArgument := "20"
	fizzAtArgument := "3"
	buzzAtArgument := "5"

	// Pull custom argument values from the command line
	if len(os.Args) > 1 {
		totalArgument = os.Args[1]
	}

	if len(os.Args) > 2 {
		fizzAtArgument = os.Args[2]
	}

	if len(os.Args) > 3 {
		buzzAtArgument = os.Args[3]
	}

	// Convert argument values to numbers
	total, err := strconv.ParseInt(totalArgument, 10, 32)
	if err != nil {
		panic("The number of items to FizzBuzz should be an integer")
	}

	fizzAt, err := strconv.ParseInt(fizzAtArgument, 10, 32)
	if err != nil {
		panic("The number to Fizz at should be an integer")
	}

	buzzAt, err := strconv.ParseInt(buzzAtArgument, 10, 32)
	if err != nil {
		panic("The number to Buzz at should be an integer")
	}

	// FizzBuzz the input and print the results
	fmt.Println(fmt.Sprintf("FizzBuzzing %d number(s), fizzing at %d and buzzing at %d:", total, fizzAt, buzzAt))
	for _, result := range fizzbuzz.FizzBuzz(total, fizzAt, buzzAt) {
		fmt.Println(result)
	}

	fmt.Println()
	fmt.Println("Done")
	fmt.Println()
}
