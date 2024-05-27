package main

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

func TestPolynomials(t *testing.T) {
	t.Run("basic case", func(t *testing.T) {
		fn := "3x^2+5x^4"
		signs, coeffs := polynomialCoefficientExtraction(fn)
		expected1, expected2 := []Monomial{{3, 2}, {5, 4}}, []rune{'+', '+'}
		AssertPolynomials(t, signs, expected2, coeffs, expected1)
	})
	t.Run("degree one", func(t *testing.T) {
		fn := "3x^2+5x"
		signs, coeffs := polynomialCoefficientExtraction(fn)
		expected1, expected2 := []Monomial{{3, 2}, {5, 1}}, []rune{'+', '+'}
		AssertPolynomials(t, signs, expected2, coeffs, expected1)
	})
	t.Run("degree zero", func(t *testing.T) {
		fn := "3x^2+5"
		signs, coeffs := polynomialCoefficientExtraction(fn)
		expected1, expected2 := []Monomial{{3, 2}, {5, 0}}, []rune{'+', '+'}
		AssertPolynomials(t, signs, expected2, coeffs, expected1)
	})
	t.Run("minus signs", func(t *testing.T) {
		fn := "-4x^4-5x^2-1"
		signs, coeffs := polynomialCoefficientExtraction(fn)
		expected1, expected2 := []Monomial{{4, 4}, {5, 2}, {1, 0}}, []rune{'-', '-', '-'}
		for _, a := range signs {
			t.Log(a)
		}
		for _, a := range coeffs {
			t.Log(a)
		}
		AssertPolynomials(t, signs, expected2, coeffs, expected1)
	})
}

func TestFunctions(t *testing.T) {
	t.Run("basic polynomial", func(t *testing.T) {
		fn := "3x^2"
		signs, coeffs := polynomialCoefficientExtraction(fn)
		poly := createFunc(coeffs, signs)
		expected := 3.0
		got := poly(1)
		if expected != got {
			t.Errorf("got %f expected %f", got, expected)
		}
	})
	t.Run("minus signs polynomial", func(t *testing.T) {
		fn := "-3x^2-6x^3"
		signs, coeffs := polynomialCoefficientExtraction(fn)
		poly := createFunc(coeffs, signs)
		expected := -9.0
		got := poly(1)
		if expected != got {
			t.Errorf("got %f expected %f", got, expected)
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
		t.Error("wrong output")

	}
}

func AssertPolynomials(t *testing.T, signs, expected_signs []rune, coeffs, expected_coeffs []Monomial) {
	t.Helper()
	diff := false
	for i, sign := range signs {
		if expected_signs[i] != sign {
			diff = true
		}
	}
	for i, c := range coeffs {
		if c.coeff != expected_coeffs[i].coeff {
			diff = true
		}
		if c.exponent != expected_coeffs[i].exponent {
			diff = true
		}
	}
	if diff {
		t.Error("wrong output")
	}
}
