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

func fuelFor(mass int64) (fuel int64) {
	fuel = mass/3 - 2
	return
}

func fuelForModule(mass int64) (totalFuel int64) {
	fuel := fuelFor(mass)
	totalFuel = fuel

	for fuel != 0 {
		fuel = fuelFor(fuel)
		if fuel < 0 {
			fuel = 0
		}
		totalFuel = totalFuel + fuel
	}

	return
}

func main() {
	// Read File
	bytes, err := ioutil.ReadFile("./inputs/1.txt")
	check(err)

	// Parse input
	input := string(bytes)
	lines := strings.Split(input, "\n")
	masses := []int64{}
	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		// Last line might be empty so let's just ignore it
		if err != nil {
			continue
		}
		masses = append(masses, int64(mass))
	}

	// Get Solution for Part 1
	var part1 int64 = 0

	for _, mass := range masses {
		part1 = part1 + fuelFor(mass)
	}

	fmt.Println(part1)

	// Get Solution for Part 2
	var part2 int64 = 0

	for _, mass := range masses {
		part2 = part2 + fuelForModule(mass)
	}

	fmt.Println(part2)
}
