// Package that parses mathematical functions and converts them into a function
// Currently supports polynomials
package main

import (
	"C"
	"reflect"
	"strconv"
	"strings"
)
import (
	"encoding/json"
	"net/http"
)

type Category int

type Monomial struct {
	coeff    int64 `json:"coeff"`
	exponent int64 `json:"exponent"`
}

// Each of the characters in the input can be of any of these types
const (
	NUM = iota
	OPERATION
	VARIABLE
	PARENTHESIS
)

// Reports whether a value is in an array of that type
// Works for any type
func inArray[T any](target T, arr []T) bool {
	for _, item := range arr {
		if reflect.DeepEqual(item, target) {
			return true
		}
	}
	return false
}

// Takes the user's input function and creates an array with the category of each of the characters
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
//
//export polynomialCoefficientExtraction
func polynomialCoefficientExtraction(fn string) ([]rune, []Monomial) {
	splitFn := []string{}
	var coeffs []Monomial
	signs := []rune{}
	var exponent int64
	for _, char := range fn {
		if char == '+' || char == '-' {
			signs = append(signs, char)
		}
	}
	if fn[0] != '-' {
		signs = append(signs, '+')
	}
	fn = strings.ReplaceAll(fn, "-", "+")
	splitFn = strings.Split(fn, "+")

	for _, monomial := range splitFn {
		category := categorizeInput(monomial)
		if monomial == "" {
			continue
		}
		if !inArray(VARIABLE, category) {
			coeff, _ := strconv.ParseInt(monomial, 10, 64)
			exponent = 0
			var mono Monomial
			mono.coeff = coeff
			mono.exponent = exponent
			coeffs = append(coeffs, mono)
		} else if !inArray(OPERATION, category) {
			i := strings.Index(monomial, "x")
			coeff, _ := strconv.ParseInt(monomial[:i], 10, 64)
			exponent = 1
			var mono Monomial
			mono.coeff = coeff
			mono.exponent = exponent
			coeffs = append(coeffs, mono)
		} else {
			tempSplit := strings.Split(monomial, "x^")
			coeff, _ := strconv.ParseInt(tempSplit[0], 10, 64)
			exponent, _ := strconv.ParseInt(tempSplit[1], 10, 64)
			var mono Monomial
			mono.coeff = coeff
			mono.exponent = exponent
			coeffs = append(coeffs, mono)
		}
	}
	return signs, coeffs
}

// Handles json requests to send the values to a web API, to be used for plotting
func handler(w http.ResponseWriter, r *http.Request) {
	// Grabs input
	r.ParseForm()
	expression := r.Form.Get("function")

	// Obtains the coefficients and signs, converts it into json
	signs, coeffs := polynomialCoefficientExtraction(expression)
	coeff_list := [][]int64{}
	for _, monomial := range coeffs {
		coeff_list = append(coeff_list, []int64{monomial.coeff, monomial.exponent})
	}
	responseData := map[string]interface{}{
		"coeffs": coeff_list,
		"signs":  signs,
	}

	// Sends the response
	json.NewEncoder(w).Encode(responseData)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
