package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		//to use the %v and all other we need to use printf cause it only support in C
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		//to use the %v and all other we need to use printf cause it only support in C
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)

	}
	//Here we made it a condional switch statement by passing the remainder as condition
	switch remainder {
	case 0:
		fmt.Printf("The division was exact")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")

	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("Cannot Divide by Zero")
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, nil
}
