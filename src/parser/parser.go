// Package that parses mathematical functions and converts them into a function
// Currently working on supporting polynomials
package parser

import (
	"reflect"
	"strconv"
	"strings"
)

type Category int

// Each of the characters in the input can be of any of these types
const (
	NUM = iota
	OPERATION
	VARIABLE
	PARENTHESIS
)

// Reports whether a string is in an array
// May be extended to support other types in the future
func inArray[T any](target T, arr []T) bool {
	for _, item := range arr {
		if reflect.DeepEqual(item, target) {
			return true
		}
	}
	return false
}

// Takes the user's input function and creates an array with the category of each of the characters
// Currently not being used but may be useful in the future
func categorizeInput(fn string) []Category {
	ops := []string{"+", "-", "*", "/", "^"}
	length := len(fn)
	res := make([]Category, length)
	for i, char := range fn {
		// iterates over the array to check for each kind of category and fills the result accordingly
		if inArray(string(char), ops) {
			res[i] = OPERATION
		} else if 'a' <= char && char <= 'z' {
			res[i] = VARIABLE
		} else if '0' <= char && char >= '9' {
			res[i] = NUM
		}
		if char == '(' || char == ')' {
			res[i] = PARENTHESIS
		}
	}
	return res
}

// Extracts the coefficients and the signs of a polynomial function into a slice with signs and a map with the coefficients
func polynomialCoefficientExtraction(fn string) ([]rune, map[int64]int64) {
	splitFn := []string{}
	coeffs := make(map[int64]int64)
	signs := []rune{}
	i := 0
	var exponent int64
	for _, char := range fn {
		if char == '+' || char == '-' {
			signs[i] = char
			i++
		}
	}
	fn = strings.ReplaceAll(fn, "-", "+")
	splitFn = strings.Split(fn, "+")

	for _, monomial := range splitFn {
		category := categorizeInput(monomial)
		if !inArray(VARIABLE, category) {
			coeff, _ := strconv.ParseInt(monomial, 10, 64)
			exponent = 0
			coeffs[coeff] = exponent
		} else if !inArray(OPERATION, category) {
			i := strings.Index(monomial, "x")
			coeff, _ := strconv.ParseInt(monomial[:i], 10, 64)
			exponent = 1
			coeffs[coeff] = exponent
		} else {
			tempSplit := strings.Split(monomial, "x^")
			coeff, _ := strconv.ParseInt(tempSplit[0], 10, 64)
			exponent, _ := strconv.ParseInt(tempSplit[1], 10, 64)
			coeffs[coeff] = exponent
		}
	}
	return signs, coeffs
}
