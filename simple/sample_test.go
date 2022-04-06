package main

import (
	"testing"
)

func TestMulti(t *testing.T) {
	var x, y, result = 2, 2, 4
	realResult := Multi(x, y)

	if realResult != result {
		t.Errorf("%d ! = %d", result, realResult)
	}
}
func TestAdd(t *testing.T) {
	var num, result = 2, 4
	realResult := Add(num)

	if realResult != result {
		t.Errorf("%d ! = %d", result, realResult)
	}
}
func TestDescribeNumber(t *testing.T) {
	var f = 12.1
	var result = "This is the number 12.1"
	realResult := DescribeNumber(f)

	if realResult != result {
		t.Errorf("%s ! = %s", result, realResult)
	}

}
func TestDescribeNumber1(t *testing.T) {
	test := struct {
		description string
		input       float64
		expected    string
	}{

		description: "Describe 4.1",
		input:       4.1,
		expected:    "This is the number 4.1",
	}
	//Two options:
	//  If realResult IS the result of the DescribeNumber function, no action is taken
	//  If realResult is not equal to what is expected, we get an error message.
	if realResult := DescribeNumber(test.input); realResult != test.expected {
		// In error message we are comparing  the NUMBER from the struct (test.input)
		// to NUMBER used in the function
		// want expected value
		t.Errorf("DescribeNumber(%v) = %v; want %v", test.input, realResult, test.expected)
	}

}

var numTests = []struct {
	input    float64
	expected string
}{
	{4.1, "This is the number 4.1"},
	{3.2, "This is the number 3.2"},
}

func TestDescribeNumber2(t *testing.T) {
	for _, n := range numTests {
		got := DescribeNumber(n.input)
		if got != n.expected {
			t.Errorf("got: %v expected: %v\n", got, n.expected)
		}
	}
}

func TestDescribeNumber3(t *testing.T) {
	test := struct {
		description string
		input       float64
		want        string
	}{

		description: "Describe 4.1",
		input:       4.1,
		want:        "This is the number 4.1",
	}

	if got := DescribeNumber(test.input); got != test.want {
		t.Errorf("DescribeNumber(%v) = %v; want %v", test.input, got, test.want)
	}

}

// -json flag

// Run test as a suite or separately
// go test  -v -run  TestMultipleSuite/easy
func TestMultipleSuite(t *testing.T) {

	//setup
	//t.Parallel()
	t.Logf("easy")
	t.Run("easy", func(t *testing.T) {
		var x, y, result = 2, 2, 4
		realResult := Multi(x, y)

		if realResult != result {
			t.Errorf("%d ! = %d", result, realResult)
		}
	})

	t.Run("medium", func(t *testing.T) {
		//t.Parallel()
		t.Logf("medium")
		var x, y, result = 2, 2, 4
		realResult := Multi(x, y)

		if realResult != result {
			t.Errorf("%d ! = %d", result, realResult)
		}
	})

	t.Run("hard", func(t *testing.T) {
		//t.Parallel()
		// t.Logf("hard")
		var x, y, result = 2, 2, 4
		realResult := Multi(x, y)

		if realResult != result {
			t.Errorf("%d ! = %d", result, realResult)
		}
	})
}
