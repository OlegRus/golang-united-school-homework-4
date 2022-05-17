package string_sum

import (
	"errors"
	"fmt"
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

const (
	plus  = rune('+')
	minus = rune('-')
)

func readNumber(input string) (int, int, error) {
	var lastPosition = 0
	var b strings.Builder

	for i, r := range []rune(input) {
		lastPosition = i
		if b.Len() == 0 {
			b.WriteRune(r)
		} else if r == plus || r == minus {
			break
		}
	}
	num, err := strconv.ParseInt(b.String(), 10, 32)
	if err != nil {
		err = fmt.Errorf("readNumber: cant parse '%s': %w", b.String(), err)
	}
	return int(num), lastPosition, err
}

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) == 0 {
		return "", fmt.Errorf("StringSum.emptyStringCheck: %w", errorEmptyInput)
	}

	operands := make([]int, 0, 2)

	lastPosition := 0
	for lastPosition+1 != len(input) {
		num, position, err := readNumber(input[lastPosition:])
		if err != nil {
			return "", fmt.Errorf("StringSum.parseIntCheck: %w", err)
		}
		lastPosition = position
		operands = append(operands, num)
	}

	if len(operands) < 2 || len(operands) > 2 {
		return "", fmt.Errorf("StringSum.operatorAmountCheck: %w", errorNotTwoOperands)
	}

	return strconv.Itoa(operands[0] + operands[1]), nil
}
