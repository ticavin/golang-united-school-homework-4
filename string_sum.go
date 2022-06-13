package string_sum

import (
	"errors"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	inputS := strings.ReplaceAll(input, " ", "")

	if inputS == "" {
		return "", errorEmptyInput
	}

	varOne := 0
	for _, i := range inputS[1:] {
		if i == '+' || i == '-' {
			varOne++
		}
	}

	if varOne != 1 {
		return "", errorNotTwoOperands
	}

	varTwo := 0

	for i := 0; i < len(inputS); i++ {
		if (inputS[i] == '-' || inputS[i] == '+') && i != 0 {
			varTwo = i
		}
	}

	firstNum, err := strconv.Atoi(inputS[:varTwo])
	if err != nil {
		return "", err
	}

	secondNum, err := strconv.Atoi(inputS[varTwo+1:])
	if err != nil {
		return "", err
	}

	if inputS[varTwo] == '+' {
		output = strconv.Itoa(firstNum + secondNum)
	} else {
		output = strconv.Itoa(firstNum - secondNum)
	}

	return output, nil
}
