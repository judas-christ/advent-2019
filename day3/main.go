package main

import (
	"fmt"
	"io/ioutil"
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
	dist := calcDistance(inputs[0], inputs[1])
	fmt.Println(dist)
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
	for ci := 1; ci < len(c); ci++ {
		for oi := 1; oi < len(o); oi++ {
			currC := c[ci]
			lastC := c[ci-1]
			currO := o[oi]
			lastO := o[oi-1]
			// go over lines
			minC := Coords{lastC, currC}.Min()
			maxC := Coords{lastC, currC}.Max()
			minO := Coords{lastO, currO}.Min()
			maxO := Coords{lastO, currO}.Max()
			for cx := minC.x; cx <= maxC.x; cx++ {
				for cy := minC.y; cy <= maxC.y; cy++ {
					for ox := minO.x; ox <= maxO.x; ox++ {
						for oy := minO.y; oy <= maxO.y; oy++ {
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

// Min gets the minimum coord by manhattan distance
func (c Coords) Min() Coord {
	minD := 100000
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

// Max gets the Maximum coord by manhattan distance
func (c Coords) Max() Coord {
	maxD := 0
	max := c[0]
	for i := range c {
		dist := c[i].Distance()
		if dist > maxD {
			maxD = dist
			max = c[i]
		}
	}
	return max
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calcDistance(wire1, wire2 string) int {
	wire1coords := parsePath(wire1)
	wire2coords := parsePath(wire2)
	inters := wire1coords.Intersections(wire2coords)
	minInter := inters.Min()
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
