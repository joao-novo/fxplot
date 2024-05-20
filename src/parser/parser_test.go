package parser

import (
	"fmt"
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

<<<<<<< HEAD
func TestPolynomials(t *testing.T) {
	t.Run("basic case", func(t *testing.T) {
		fn := "3x^2+5x^4"
		signs, coeffs := polynomialCoefficientExtraction(fn)
		expected1, expected2 := map[int]int{5: 4, 3: 2}, map[rune]int{'+': 1, '-': 0}
		if fmt.Sprint(coeffs) != fmt.Sprint(expected1) || fmt.Sprint(signs) != fmt.Sprint(expected2) {
			t.Errorf("wrong output")
		}
	})
	t.Run("degree one", func(t *testing.T) {
		fn := "3x^2+5x"
		signs, coeffs := polynomialCoefficientExtraction(fn)
		expected1, expected2 := map[int]int{3: 2, 5: 1}, map[rune]int{'+': 1, '-': 0}
		if fmt.Sprint(coeffs) != fmt.Sprint(expected1) || fmt.Sprint(signs) != fmt.Sprint(expected2) {
			t.Errorf("wrong output")
		}
	})
	t.Run("degree zero", func(t *testing.T) {
		fn := "3x^2+5"
		signs, coeffs := polynomialCoefficientExtraction(fn)
		expected1, expected2 := map[int]int{3: 2, 5: 0}, map[rune]int{'+': 1, '-': 0}
		if fmt.Sprint(coeffs) != fmt.Sprint(expected1) || fmt.Sprint(signs) != fmt.Sprint(expected2) {
			t.Errorf("wrong output")
		}
	})
}

||||||| 09cd4b6
=======
func TestTreeConversion(t *testing.T) {
	t.Run("args length 1", func(t *testing.T) {
		tree := convertToTree("2+3")
		expected := OperationTree{
			'+',
			"2",
			"3",
			0,
		}
		AssertTreeConversion(t, tree, expected)
	})
	t.Run("variable length arguments with variables", func(t *testing.T) {
		tree := convertToTree("52+x")
		expected := OperationTree{
			'+',
			"52",
			"x",
			0,
		}
		AssertTreeConversion(t, tree, expected)
	})
	t.Run("operation with parentheses", func(t *testing.T) {
		tree := convertToTree("(42*x)")
		expected := OperationTree{
			'*',
			"42",
			"x",
			3,
		}
		AssertTreeConversion(t, tree, expected)
	})
}

>>>>>>> refs/remotes/origin/master
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
		t.Logf("op: got %b expected %b\narg1: got %s expected %s\narg2: got %s expected %s\npriority: got %d expected %d",
			tree.op, expected.op, tree.arg1, expected.arg1, tree.arg2, expected.arg2, tree.priority, expected.priority)
		t.Error("wrong output")
	}

}
