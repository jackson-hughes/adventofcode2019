package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
	Opcode 1 adds together numbers read from two positions and stores the result in a third position.

	The three integers immediately after the opcode tell you these three positions.
	The first two indicate the positions from which you should read the input values, and the third indicates
	the position at which the output should be stored.
*/

func getOpCode(program string) (opcode string) {
	opSlice := []rune(program)
	code := string(opSlice[0])
	return code
}

func processOpCode(program string) (code string, err error) {
	opcode := getOpCode(program)
	if opcode == strconv.Itoa(1) {
		fmt.Printf("Opcode was %v\n", opcode)
	} else if opcode == strconv.Itoa(2) {
		fmt.Printf("Opcode was %v\n", opcode)
	} else if opcode == strconv.Itoa(99) {
		return "", nil
	} else {
		errors.New("unexpected opcode")
		return "", err
	}
	return "", errors.New("something went wonky")
}

func opCodeOne() {
	fmt.Println("Op Code 1")
}

func opCodeTwo() {
	fmt.Println("Op Code 2")
}

func main() {
	input := "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,6,19,1,19,6,23,2,23,6,27,2,6,27,31,2,13,31" +
		",35,1,9,35,39,2,10,39,43,1,6,43,47,1,13,47,51,2,6,51,55,2,55,6,59,1,59,5,63,2,9,63," +
		"67,1,5,67,71,2,10,71,75,1,6,75,79,1,79,5,83,2,83,10,87,1,9,87,91,1,5,91,95,1,95,6,99" +
		",2,10,99,103,1,5,103,107,1,107,6,111,1,5,111,115,2,115,6,119,1,119,6,123,1,123,10,127" +
		",1,127,13,131,1,131,2,135,1,135,5,0,99,2,14,0,0"
	fmt.Println(getOpCode(input))

	p, err := processOpCode(input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
