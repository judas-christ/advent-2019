package main

import "testing"

func TestRunProgram(t *testing.T) {
	inputs := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	runProgram(inputs)
	if inputs[0] != 3500 {
		t.Errorf("expected %d but got %d", 3500, inputs[0])
	}
}
