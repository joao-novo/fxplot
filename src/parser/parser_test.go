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
		fn := "2(x + 1)"
		category := categorizeInput(fn)
		expected := []Category{NUM, PARENTHESIS, VARIABLE, OPERATION, NUM, PARENTHESIS}
		AssertCategory(t, category, expected)
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
