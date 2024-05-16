// calc_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", result)
	}
}

func TestSub(t *testing.T) {
	result := Sub(3, 2)
	if result != 1 {
		t.Errorf("Sub(3, 2) = %d; want -1", result)
	}
}

func TestMul(t *testing.T) {
	result := Mul(2, 4)
	if result != 8 {
		t.Errorf("Mul(2, 4) = %d; want 8", result)
	}
}

func TestDiv(t *testing.T) {
	result := Div(9, 3)
	if result != 3 {
		t.Errorf("Div(9, 3) = %d; want 3", result)
	}
}
