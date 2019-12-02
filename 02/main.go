package main

import (
	"fmt"
	"math"
)

func calculateFuelsFuel(fuel float64) (totalFuel float64) {
	d := fuel / float64(3)
	r := math.Floor(d)
	s := r - float64(2)
	if s <= 0 {
		return totalFuel
	}
	// fmt.Println(s)
	return s + calculateFuelsFuel(s)
}

func main() {
	moduleFuel := 3246455

	fuelFuel := calculateFuelsFuel(float64(moduleFuel))

	totalFuel := float64(moduleFuel) + fuelFuel
	fmt.Println(totalFuel)
}
