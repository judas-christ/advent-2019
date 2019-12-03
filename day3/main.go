package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	inputs := readInputs()
	wire1 := parsePath(inputs[0])
	wire2 := parsePath(inputs[1])

	//part 1
	dist := calcDistance(wire1, wire2)
	fmt.Println("Part 1", dist)

	//part 2
	steps := calcSteps(wire1, wire2)
	fmt.Println("Part 2", steps)
}

func readInputs() []string {
	strInput, err := ioutil.ReadFile("day3/input.txt")
	check(err)
	return strings.Split(string(strInput), "\n")
}

// Coords is a list of Coord
type Coords []Coord

// Intersections gets all intersections
func (c Coords) Intersections(o Coords) Coords {
	r := make(Coords, 0)
	for ci := range c[1:] {
		lastC := c[ci]
		currC := c[ci+1]
		for oi := range o[1:] {
			lastO := o[oi]
			currO := o[oi+1]
			// go over lines
			minCx := min(lastC.x, currC.x)
			maxCx := max(lastC.x, currC.x)
			minCy := min(lastC.y, currC.y)
			maxCy := max(lastC.y, currC.y)
			minOx := min(lastO.x, currO.x)
			maxOx := max(lastO.x, currO.x)
			minOy := min(lastO.y, currO.y)
			maxOy := max(lastO.y, currO.y)
			for cx := minCx; cx <= maxCx; cx++ {
				for cy := minCy; cy <= maxCy; cy++ {
					for ox := minOx; ox <= maxOx; ox++ {
						for oy := minOy; oy <= maxOy; oy++ {
							if cx == ox && cy == oy && !(cx == 0 && cy == 0) {
								r = append(r, Coord{cx, cy})
							}
						}
					}
				}
			}
		}
	}
	return r
}

// MinDistance gets the minimum coord by manhattan distance
func (c Coords) MinDistance() Coord {
	minD := math.MaxInt32
	min := c[0]
	for i := range c {
		dist := c[i].Distance()
		if dist < minD {
			minD = dist
			min = c[i]
		}
	}
	return min
}

// StepsTo calculates the number of steps in a list to a specific coord
func (c Coords) StepsTo(coord Coord) int {
	steps := 0
	for ci := range c[1:] {
		last := c[ci]
		curr := c[ci+1]
		// case where this is the segment coord belongs to
		if last.x == coord.x && curr.x == coord.x && (last.y <= coord.y && curr.y >= coord.y || last.y >= coord.y && curr.y <= coord.y) ||
			last.y == coord.y && curr.y == coord.y && (last.x <= coord.x && curr.x >= coord.x || last.x >= coord.x && curr.x <= coord.x) {
			steps += last.DistanceTo(coord)
			return steps
		}
		// otherwise
		steps += last.DistanceTo(curr)
	}
	return steps
}

// Coord is a coordinate
type Coord struct {
	x, y int
}

// IsOrigo checks whether its 0,0
func (c Coord) IsOrigo() bool {
	return c.x == 0 && c.y == 0
}

// Add adds coords to Coord
func (c Coord) Add(x, y int) Coord {
	return Coord{c.x + x, c.y + y}
}

// Intersects checks whether two coords intersect
func (c Coord) Intersects(o Coord) bool {
	return c.x == o.x && c.y == o.y
}

// Distance calculates the manhattan distance
func (c Coord) Distance() int {
	return abs(c.x) + abs(c.y)
}

// DistanceTo calculates manhattan distance to other coord
func (c Coord) DistanceTo(other Coord) int {
	return abs(c.x-other.x) + abs(c.y-other.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func calcDistance(wire1, wire2 Coords) int {
	inters := wire1.Intersections(wire2)
	minInter := inters.MinDistance()
	return minInter.Distance()
}

func parsePath(path string) Coords {
	splitPath := strings.Split(path, ",")
	coords := make(Coords, len(splitPath)+1)
	coords[0] = Coord{0, 0}
	lastCoord := coords[0]
	for i := range splitPath {
		x, y := parseCoord(splitPath[i])
		newCoord := lastCoord.Add(x, y)
		coords[i+1] = newCoord
		lastCoord = newCoord
	}
	return coords
}

func parseCoord(str string) (int, int) {
	num, err := strconv.Atoi(str[1:])
	check(err)
	switch str[0] {
	case 'U':
		return 0, num
	case 'R':
		return num, 0
	case 'D':
		return 0, -num
	case 'L':
		return -num, 0
	default:
		panic("invalid direction")
	}
}

func calcSteps(wire1, wire2 Coords) int {
	inters := wire1.Intersections(wire2)
	minStep := math.MaxInt32
	// minInter := inters[0]
	for i := range inters {
		steps1 := wire1.StepsTo(inters[i])
		steps2 := wire2.StepsTo(inters[i])
		steps := steps1 + steps2
		if steps <= minStep {
			minStep = steps
			// minInter = inters[i]
		}
	}
	return minStep
}
