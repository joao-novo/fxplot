package parser

import "strings"

type Category int

const (
	NUM = iota
	OPERATION
	VARIABLE
	PARENTHESIS
)

func inArray(target string, arr []string) bool {
	for _, item := range arr {
		if target == item {
			return true
		}
	}
	return false
}

func categorizeInput(fn string) []Category {
	fn = strings.ReplaceAll(fn, " ", "")
	ops := []string{"+", "-", "*", "/", "^"}
	length := len(fn)
	res := make([]Category, length)
	for i, char := range fn {
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
