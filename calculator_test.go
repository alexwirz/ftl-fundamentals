package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	var want float64 = 4
	got := calculator.Add(2, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(4, 2)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

func TestSubstract2(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 5, b: 1, want: 4},
		{a: 5, b: 0, want: 5},
		{a: 3, b: -1, want: 4},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("Substract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 6
	got := calculator.Multiply(2, 3)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}

type testCase struct {
	a, b          float64
	want          float64
	errorExpected bool
	message       string
}

func TestMultiply2(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4, message: "Multiply two positive numbers"},
		{a: 3, b: 1, want: 3, message: "Multiply by one"},
		{a: 5, b: 0, want: 0, message: "Mutliply by zero"},
		{a: 3, b: -1, want: -3, message: "Multiply by minus one"},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("%s: Multiply(%f, %f): want %f, got %f", tc.message, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestAdd2(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 6, b: 3, want: 2, message: "Simple division with two positive integers"},
		{a: 10, b: -2, want: -5, message: "Divide by a negative integer"},
		{a: 6, b: 0, errorExpected: true, message: "Simple division with two positive integers"},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if tc.errorExpected != (err != nil) {
			t.Fatalf("Divide(%f, %f) unexpected error status: %s", tc.a, tc.b, err)
		}

		if !tc.errorExpected && got != tc.want {
			t.Errorf("%s: Divide(%f, %f): want %f, got %f", tc.message, tc.a, tc.b, tc.want, got)
		}
	}
}
