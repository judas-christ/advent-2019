package main

import (
	"testing"
)

func Test_parsePath(t *testing.T) {
	path := "R8,U5,L5,D3"
	got := parsePath(path)
	if got[0].x != 0 && got[0].y != 0 {
		t.Error("Expected starting point to be 0,0")
	}
	if got[1].x != 8 && got[1].y != 0 {
		t.Error("Expected [1] to be 8,0")
	}
	if got[2].x != 8 && got[2].y != 5 {
		t.Error("Expected [2] to be 8,5")
	}
}

func Test_Intersections1(t *testing.T) {
	wire1 := parsePath("R8,U5,L5,D3")
	wire2 := parsePath("U7,R6,D4,L4")
	got := wire1.Intersections(wire2)
	if len(got) != 2 {
		t.Errorf("got %d expected %d", got, 2)
	}
}

func Test_CalcDistance1(t *testing.T) {
	wire1 := "R8,U5,L5,D3"
	wire2 := "U7,R6,D4,L4"
	got := calcDistance(wire1, wire2)
	if got != 6 {
		t.Errorf("got %d expected %d", got, 6)
	}
}

func Test_CalcDistance2(t *testing.T) {
	wire1 := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	wire2 := "U62,R66,U55,R34,D71,R55,D58,R83"
	got := calcDistance(wire1, wire2)
	if got != 159 {
		t.Errorf("got %d expected %d", got, 159)
	}
}
func Test_CalcDistance3(t *testing.T) {
	wire1 := "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	wire2 := "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
	got := calcDistance(wire1, wire2)
	if got != 135 {
		t.Errorf("got %d expected %d", got, 135)
	}
}
