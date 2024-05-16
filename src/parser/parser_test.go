package parser

import (
	"reflect"
	"testing"
)

func TestCategory(t *testing.T) {
	t.Run("number", func(t *testing.T) {
		fn := "2"
		category := categorizeInput(fn)
		expected := []Category{NUM}
		AssertCategory(t, category, expected)
	})
	t.Run("linear function", func(t *testing.T) {
		fn := "2x"
		category := categorizeInput(fn)
		expected := []Category{NUM, VARIABLE}
		AssertCategory(t, category, expected)
	})
	t.Run("addition", func(t *testing.T) {
		fn := "2+x"
		category := categorizeInput(fn)
		expected := []Category{NUM, OPERATION, VARIABLE}
		AssertCategory(t, category, expected)
	})
	t.Run("multiplication", func(t *testing.T) {
		fn := "2*x"
		category := categorizeInput(fn)
		expected := []Category{NUM, OPERATION, VARIABLE}
		AssertCategory(t, category, expected)
	})
	t.Run("subtraction", func(t *testing.T) {
		fn := "2-x"
		category := categorizeInput(fn)
		expected := []Category{NUM, OPERATION, VARIABLE}
		AssertCategory(t, category, expected)
	})
	t.Run("division", func(t *testing.T) {
		fn := "2/x"
		category := categorizeInput(fn)
		expected := []Category{NUM, OPERATION, VARIABLE}
		AssertCategory(t, category, expected)
	})
	t.Run("exponentiation", func(t *testing.T) {
		fn := "2^x"
		category := categorizeInput(fn)
		expected := []Category{NUM, OPERATION, VARIABLE}
		AssertCategory(t, category, expected)
	})
	t.Run("parenthesis", func(t *testing.T) {
		fn := "2(x+1)"
		category := categorizeInput(fn)
		expected := []Category{NUM, PARENTHESIS, VARIABLE, OPERATION, NUM, PARENTHESIS}
		AssertCategory(t, category, expected)
	})
}

func TestTreeConversion(t *testing.T) {
	t.Run("args length 1", func(t *testing.T) {
		tree := convertToTree("2+3")
		expected := OperationTree{
			'+',
			"2",
			"3",
		}
		AssertTreeConversion(t, tree, expected)
	})
	t.Run("variable length arguments with variables", func(t *testing.T) {
		tree := convertToTree("52+x")
		expected := OperationTree{
			'+',
			"52",
			"x",
		}
		AssertTreeConversion(t, tree, expected)
	})
}

func AssertCategory(t *testing.T, category, expected []Category) {
	t.Helper()
	if !reflect.DeepEqual(category, expected) {
		t.Log("Output: ")
		for _, item := range category {
			t.Log(item)
		}
		t.Log("Expected: ")
		for _, item := range category {
			t.Log(item)
		}
		t.Error("wrong output")

	}
}

func AssertTreeConversion(t *testing.T, tree, expected OperationTree) {
	t.Helper()
	if !reflect.DeepEqual(tree, expected) {
		t.Logf("op: got %b expected %b\narg1: got %s expected %s\narg2: got %s expected %s",
			tree.op, expected.op, tree.arg1, expected.arg1, tree.arg2, expected.arg2)
		t.Error("wrong output")
	}

}
