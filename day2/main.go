package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	inputs := readInputs()
	// initial value of [0]
	fmt.Println(inputs[0])

	// restore to last good state
	inputs[1] = 12
	inputs[2] = 2

	// run program
	runProgram(inputs)

	fmt.Println("Answer")
	fmt.Println(inputs[0])

	// part 2
	maxVal := len(inputs)
	for noun := 0; noun < maxVal; noun++ {
		for verb := 0; verb < maxVal; verb++ {
			inputs = readInputs()
			inputs[1] = noun
			inputs[2] = verb
			runProgram(inputs)
			result := inputs[0]
			if result == 19690720 {
				fmt.Printf("noun(%d) * 100 + verb(%d) = %d", noun, verb, noun*100+verb)
				return
			}
		}
	}
}

func readInputs() []int {
	strInput, err := ioutil.ReadFile("day2/input.txt")
	check(err)
	return parseInput(strInput)
}

func parseInput(strInput []byte) []int {
	strInputs := strings.Split(strings.TrimRight(string(strInput), " \n"), ",")
	intOutput := make([]int, len(strInputs))
	for i := 0; i < len(strInputs); i++ {
		val, err := strconv.Atoi(strInputs[i])
		check(err)
		intOutput[i] = val
	}
	return intOutput
}

func runProgram(inputs []int) {
	pos := 0
	for opcode := inputs[pos]; opcode != 99; opcode = inputs[pos] {
		switch opcode {
		case 1:
			posA, posB, posResult := parseAddresses(pos, inputs)
			inputs[posResult] = inputs[posA] + inputs[posB]
			pos += 4
		case 2:
			posA, posB, posResult := parseAddresses(pos, inputs)
			inputs[posResult] = inputs[posA] * inputs[posB]
			pos += 4
		case 99:
			return
		default:
			panic("UNEXPECTED OPCODE! " + string(opcode))
		}
	}
}

func parseAddresses(pos int, inputs []int) (posA, posB, posResult int) {
	posA = inputs[pos+1]
	posB = inputs[pos+2]
	posResult = inputs[pos+3]
	return
}
