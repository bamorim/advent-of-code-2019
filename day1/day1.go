package day1

import (
	"fmt"
	"strconv"
	"strings"
)

// Run runs day 1
func Run(bytes []byte) {
	parsed := parse(bytes)

	part1 := part1(parsed)
	fmt.Printf("Day 1 - Part 1: ")
	fmt.Println(part1)

	part2 := part2(parsed)
	fmt.Printf("Day 1 - Part 2: ")
	fmt.Println(part2)
}

func parse(bytes []byte) (masses []int64) {
	input := string(bytes)
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		// Last line might be empty so let's just ignore it
		if err != nil {
			continue
		}
		masses = append(masses, int64(mass))
	}

	return
}

func part1(masses []int64) (fuel int64) {
	for _, mass := range masses {
		fuel = fuel + fuelFor(mass)
	}
	return
}

func part2(masses []int64) (totalFuel int64) {
	for _, mass := range masses {
		totalFuel = totalFuel + fuelForModule(mass)
	}
	return
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
