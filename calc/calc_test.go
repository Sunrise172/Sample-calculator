package calc

import "testing"

func TestSum(t *testing.T) {
	want := 4
	if got := Add(2, 2); got != want {
		t.Errorf("Add() = %q, want %q", got, want)
	}
	if got := Subtract(6, 2); got != want {
		t.Errorf("Subtract() = %q, want %q", got, want)
	}
	want1 := 4.0
	if got := Multiply(2.0, 2.0); got != want1 {
		t.Errorf("Multiply() = %f, want %f", got, want1)
	}
	if got := Divide(8.0, 2.0); got != want1 {
		t.Errorf("Divide() = %f, want %f", got, want1)
	}
}
