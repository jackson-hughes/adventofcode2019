package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func calculateFuel(moduleMass float64) (fuelRequirement int64, err error) {
	d := moduleMass / float64(3)
	r := int64(math.Floor(d))
	s := r - int64(2)
	return s, nil
}

func readModules() (modules []float64, err error) {
	file, err := os.Open("modules.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := []float64{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if module, err := strconv.ParseFloat(scanner.Text(), 64); err == nil {
			s = append(s, module)
		} else {
			return nil, err
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return s, nil

}

func main() {

	var totalFuelRequirement int64

	modules, err := readModules()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := []int64{}

	for _, x := range modules {
		if y, err := calculateFuel(x); err == nil {
			s = append(s, y)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	for _, x := range s {
		totalFuelRequirement += x
	}
	fmt.Println(totalFuelRequirement)
}
