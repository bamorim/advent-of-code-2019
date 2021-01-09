package day3

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	up = iota
	right
	down
	left
)

type instruction struct {
	direction int
	distance  int
}

type wire []instruction

type input struct {
	wire1 wire
	wire2 wire
}

type pos [2]int

type f func() f

// Run runs day 3
func Run(bytes []byte) {
	parsed := parse(bytes)

	fmt.Printf("Day 3 - Part 1: ")
	fmt.Println(part1(parsed))

	fmt.Printf("Day 3 - Part 2: ")
	fmt.Println(part2(parsed))
}

func parse(bytes []byte) (result input) {
	lines := strings.Split(string(bytes), "\n")

	result = input{parseWire(lines[0]), parseWire(lines[1])}
	return
}

func parseWire(line string) (result wire) {
	for _, inst := range strings.Split(line, ",") {
		c := inst[0]
		distance, err := strconv.Atoi(inst[1:])

		if err != nil {
			panic(err)
		}

		switch c {
		case 'U':
			result = append(result, instruction{up, distance})
		case 'R':
			result = append(result, instruction{right, distance})
		case 'D':
			result = append(result, instruction{down, distance})
		case 'L':
			result = append(result, instruction{left, distance})
		default:
			panic("What are you doing?")
		}
	}
	return
}

func part1(parsed input) (minDistance int) {
	visited := map[pos]bool{}
	intersections := []pos{}
	center := [2]int{0, 0}

	traverse(parsed.wire1, func(p pos, _ int) {
		visited[p] = true
	})

	traverse(parsed.wire2, func(p pos, _ int) {
		value, _ := visited[p]
		if value {
			intersections = append(intersections, p)
		}
	})

	for _, intersection := range intersections {
		distance := manhattanDistance(intersection, center)
		if minDistance == 0 || minDistance > distance {
			minDistance = distance
		}
	}
	return
}

func part2(parsed input) (minSteps int) {
	stepsForWire1 := map[pos]int{}
	visitedByWire2 := map[pos]bool{}

	traverse(parsed.wire1, func(p pos, steps int) {
		_, found := stepsForWire1[p]
		if !found {
			stepsForWire1[p] = steps
		}
	})

	traverse(parsed.wire2, func(p pos, steps2 int) {
		steps1, alreadyVisitedByWire1 := stepsForWire1[p]
		_, alreadyVisitedByWire2 := visitedByWire2[p]
		if alreadyVisitedByWire1 && !alreadyVisitedByWire2 {
			combinedSteps := steps1 + steps2
			if minSteps == 0 || minSteps > combinedSteps {
				minSteps = combinedSteps
			}
		}
		visitedByWire2[p] = true
	})

	return
}

func traverse(w wire, onPoint func(pos, int)) {
	pos := [2]int{0, 0}
	steps := 0

	for _, instruction := range w {
		for i := 1; i <= instruction.distance; i++ {
			switch instruction.direction {
			case up:
				pos[1] = pos[1] + 1
			case right:
				pos[0] = pos[0] + 1
			case down:
				pos[1] = pos[1] - 1
			case left:
				pos[0] = pos[0] - 1
			}
			steps++
			onPoint(pos, steps)
		}
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func manhattanDistance(p1 pos, p2 pos) (distance int) {
	distance = abs(p1[0]-p2[0]) + abs(p1[1]-p2[1])
	return
}
