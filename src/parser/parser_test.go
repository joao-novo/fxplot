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
		fn := "2(x + 1)"
		category := categorizeInput(fn)
		expected := []Category{NUM, PARENTHESIS, VARIABLE, OPERATION, NUM, PARENTHESIS}
		AssertCategory(t, category, expected)
	})
}

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
		t.Errorf("wrong output")

	}
}
