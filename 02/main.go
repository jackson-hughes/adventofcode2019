package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculateFuel(moduleMass int) (fuelRequirement int, err error) {
	d := moduleMass / 3
	s := int(d) - 2
	return s, nil
}

func calculateFuelsFuel(fuel int) (totalFuel int) {
	d := fuel / 3
	s := d - 2
	if s <= 0 {
		return totalFuel
	}
	return s + calculateFuelsFuel(s)
}

func readModules() (modules []int, err error) {
	file, err := os.Open("modules.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	s := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if module, err := strconv.Atoi(scanner.Text()); err == nil {
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
