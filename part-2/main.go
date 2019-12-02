package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("part-1/input.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fuelRequired := 0
	for scanner.Scan() {
		moduleWeight, err := strconv.Atoi(scanner.Text())
		check(err)
		fuelRequired += getFuelRequired(moduleWeight)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fuelRequired)
}

func getFuelRequired(moduleWeight int) int {
	return recurseFuelRequired(moduleWeight, 0)
}

func recurseFuelRequired(weight, fuel int) int {
	req := weight/3 - 2
	if req <= 0 {
		return fuel
	}
	return recurseFuelRequired(req, req+fuel)
}
