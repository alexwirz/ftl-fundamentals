package calculator_test

import (
	"calculator"
	"math"
	"math/rand"
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
		checkTestCase(tc, calculator.Subtract, t)
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
		checkTestCase(tc, calculator.Multiply, t)
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
		checkTestCase(tc, calculator.Add, t)
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

func TestAddRandom(t *testing.T) {
	t.Parallel()
	testCases := []testCase{}

	for i := 0; i < 1000; i++ {
		sum := rand.Float64() * 100
		part := rand.Float64() * 100
		testCases = append(testCases, testCase{a: part, b: sum - part, want: sum})
	}

	for _, tc := range testCases {
		checkTestCase(tc, calculator.Add, t)
	}
}

func checkTestCase(tc testCase, act func(a, b float64) float64, t *testing.T) {
	got := act(tc.a, tc.b)

	if !closeEnough(got, tc.want) {
		t.Errorf("xxx(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
	}
}

func roundAtFour(x float64) float64 {
	return round(x, 0.0005)
}

func round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func closeEnough(a, b float64) bool {
	return roundAtFour(a) == roundAtFour(b)
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 4, want: 2},
		{a: 0, want: 0},
		{a: -1, errorExpected: true},
	}

	for i := 0; i < 100; i++ {
		factor := rand.Float64() * 100
		testCases = append(testCases, testCase{a: factor * factor, want: factor})
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if tc.errorExpected != (err != nil) {
			t.Fatalf("Sqrt(%f, %f) unexpected error status: %s", tc.a, tc.b, err)
		}

		if !closeEnough(got, tc.want) {
			t.Errorf("Sqrt(%f): want %f, got %f", tc.a, tc.want, got)
		}
	}
}
