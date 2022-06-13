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
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return "", errorEmptyInput
	}
	if len(input) > 4 && len(input) < 3 {
		return "", errorNotTwoOperands
	}
	if input[0] == '-' {

		firstNum, err := strconv.ParseInt(string(input[1]), 10, 32)
		if err != nil {
			return "", err
		}

		secondNum, err := strconv.ParseInt(string(input[3]), 10, 32)
		if err != nil {
			return "", err
		}
		if firstNum >= secondNum {
			output = strconv.FormatInt(firstNum-secondNum, 10)
		} else {
			output = strconv.FormatInt(secondNum-firstNum, 10)
		}
	} else {

		firstNum, err := strconv.ParseInt(string(input[0]), 10, 32)
		if err != nil {
			return "", err
		}
		secondNum, err := strconv.ParseInt(string(input[2]), 10, 32)
		if err != nil {
			return "", err
		}

		output = strconv.FormatInt(firstNum+secondNum, 10)
	}
	return output, nil
}
