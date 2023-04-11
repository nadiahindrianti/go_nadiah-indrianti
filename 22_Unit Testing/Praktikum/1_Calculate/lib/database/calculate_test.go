package calculate

import "testing"

func TestAddition(t *testing.T) {
	if Addition(6, 3) != 9 {
		t.Error("Expected 6 (+) 3 to equal 9")
	}

	if Addition(-6, -3) != -9 {
		t.Error("Expected -6 (+) -3 to equal -9")
	}
}

func TestSubtraction(t *testing.T) {
	if Subtraction(6, 3) != 3 {
		t.Error("Expected 6 (-) 3 to equal 3")
	}

	if Subtraction(-6, -3) != -3 {
		t.Error("Expected -6 (-) -3 to equal -3")
	}
}

func TestDivison(t *testing.T) {
	if Divison(6, 3) != 2 {
		t.Error("Expected 6 (/) 3 to equal 2")
	}

	if Divison(-6, -3) != 2 {
		t.Error("Expected -6 (/) -3 to equal 2")
	}
}

func TestMultiplication(t *testing.T) {
	if Multiplication(6, 3) != 18 {
		t.Error("Expected 6 (*) 3 to equal 18")
	}

	if Multiplication(-6, -3) != 18 {
		t.Error("Expected -6 (*) -3 to equal -18")
	}
}
