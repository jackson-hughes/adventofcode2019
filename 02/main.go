package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculateFuel(moduleMass float64) (fuelRequirement int, err error) {
	d := moduleMass / float64(3)
	s := int(d) - 2
	return s, nil
}

func calculateFuelsFuel(fuel int) (totalFuel int) {
	d := fuel / 3
	s := d - 2
	// fmt.Println(s)
	if s <= 0 {
		return totalFuel
	}
	return s + calculateFuelsFuel(s)
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
	var totalFuelRequirement int

	modules, err := readModules()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := []int{}
	for _, x := range modules {
		if y, err := calculateFuel(x); err == nil {
			z := calculateFuelsFuel(y)
			s = append(s, y, z)
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
