package parser

import (
	"slices"
)

type Category int

const (
	NUM = iota
	OPERATION
	VARIABLE
	PARENTHESIS
)

type OperationTree struct {
	op   byte
	arg1 string
	arg2 string
}

type OperationNode struct {
	item OperationTree
	next *OperationNode
}

func inArray(target string, arr []string) bool {
	for _, item := range arr {
		if target == item {
			return true
		}
	}
	return false
}

func categorizeInput(fn string) []Category {
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

func convertToTree(operation string) OperationTree {
	var tree OperationTree
	category := categorizeInput(operation)
	opIdx := slices.Index(category, OPERATION)
	tree.op = operation[opIdx]
	tree.arg1 = operation[:opIdx]
	tree.arg2 = operation[opIdx+1:]
	return tree
}
