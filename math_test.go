package main

import "testing"

func TestSum(t *testing.T) {
	if sum(1, 2) != 3 {
		t.Errorf("Sum(1, 2) = %d; want 3", sum(1, 2))
	}
}

func TestSub(t *testing.T) {
	if sub(5, 3) != 2 {
		t.Errorf("Sub(5, 3) = %d; want 2", sub(5, 3))
	}
}

func TestMul(t *testing.T) {
	if mul(3, 4) != 12 {
		t.Errorf("Mul(3, 4) = %d; want 12", mul(3, 4))
	}
}

func TestDiv(t *testing.T) {
	if div(10, 2) != 5 {
		t.Errorf("Div(10, 2) = %d; want 5", div(10, 2))
	}
}

func TestMod(t *testing.T) {
	if mod(10, 3) != 1 {
		t.Errorf("Mod(10, 3) = %d; want 1", mod(10, 3))
	}
}

func TestSqrt(t *testing.T) {
	if sqrt(16) != 4 {
		t.Errorf("Sqrt(16) = %d; want 4", sqrt(16))
	}
}

func TestCbrt(t *testing.T) {
	if cbrt(27) != 3 {
		t.Errorf("Cbrt(27) = %d; want 3", cbrt(27))
	}
}

func TestLog(t *testing.T) {
	if log(2) != 0 {
		t.Errorf("Log(2) = %d; want 0", log(2))
	}
}