package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		parse(args[1])
		return
	} else if len(args) < 3 {
		fmt.Println("One or more arguments missing")
		return
	} else if len(args) > 2 {
		var result float64
		var error string
		if args[1] == "add" {
			result, error = addAll(args[2:])
		} else {
			result, error = subtractAll(args[2:])
		}
		if error == "error" {
			return
		} else {
			fmt.Println("Sum of all numbers: ", result)
			return
		}
	}
	var num1, num2 int
	var num1Str = args[1]
	var num2Str = args[2]
	num1, num1Err := strconv.Atoi(num1Str)
	num2, num2Err := strconv.Atoi(num2Str)
	if num1Err != nil {
		fmt.Println("Bad argument: " + num1Str)
		return
	}
	if num2Err != nil {
		fmt.Println("Bad argument: " + num2Str)
		return
	}
	var sum = num1 + num2
	fmt.Println("Addition of two numbers: ", sum)
}

func addAll(numArray []string) (total float64, error string) {
	var result = 0.0
	for _, numStr := range numArray {
		num, numErr := strconv.ParseFloat(numStr, 64)
		if numErr != nil {
			fmt.Println("Bad Argument: " + numStr)
			return 0, "error"
		}
		result += num
	}
	return result, ""
}

func subtractAll(numArray []string) (total float64, error string) {
	var result float64
	for idx, numStr := range numArray {
		num, numErr := strconv.ParseFloat(numStr, 64)
		if numErr != nil {
			fmt.Println("Bad Argument: " + numStr)
			return 0.0, "error"
		}
		if idx == 0 {
			result = num
		}
		result = result - num
	}
	return result, ""
}

func parse(parseStr string) {
	println("Coming Soon...")
}
