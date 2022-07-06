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
		var total, error = addAll(args[1:])
		if error == "error" {
			return
		} else {
			fmt.Println("Sum of all numbers: " + strconv.Itoa(total))
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

func addAll(numArray []string) (total int, error string) {
	var sum = 0
	for _, numStr := range numArray {
		num, numErr := strconv.Atoi(numStr)
		if numErr != nil {
			fmt.Println("Bad Argument: " + numStr)
			return 0, "error"
		}
		sum += num
	}
	return sum, ""
}

func parse(parseStr string) {
	println("Coming Soon...")
}
