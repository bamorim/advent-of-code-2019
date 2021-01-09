package day2

import (
	"fmt"
	"strconv"
	"strings"
)

// Run runs day 2
func Run(bytes []byte) {
	parsed := parse(bytes)

	part1 := part1(parsed)
	fmt.Printf("Day 2 - Part 1: ")
	fmt.Println(part1[0])

	part2 := part2(parsed)
	fmt.Printf("Day 1 - Part 2: ")
	fmt.Println(part2[1]*100 + part2[2])
}

func parse(bytes []byte) (program []int64) {
	numbers := strings.Split(strings.Trim(string(bytes), "\n"), ",")
	for _, numberString := range numbers {
		number, err := strconv.Atoi(numberString)

		if err != nil {
			panic(err)
		}

		program = append(program, int64(number))
	}

	return
}

func runIntcodeProgram(program []int64) (memory []int64) {
	memory = make([]int64, len(program))
	copy(memory, program)
	pos := 0

	for memory[pos] != 99 {
		switch memory[pos] {
		case 1:
			aPos := memory[pos+1]
			bPos := memory[pos+2]
			resultPos := memory[pos+3]
			result := memory[aPos] + memory[bPos]
			memory[resultPos] = result
		case 2:
			aPos := memory[pos+1]
			bPos := memory[pos+2]
			resultPos := memory[pos+3]
			result := memory[aPos] * memory[bPos]
			memory[resultPos] = result
		default:
			panic("Invalid opcode")
		}
		pos = pos + 4
	}

	return
}

func part1(program []int64) (memory []int64) {
	newProgram := make([]int64, len(program))
	copy(newProgram, program)
	newProgram[1], newProgram[2] = 12, 2
	memory = runIntcodeProgram(newProgram)
	return
}

func part2(program []int64) (memory []int64) {
	for a := int64(0); a <= 99; a++ {
		for b := int64(0); b <= 99; b++ {
			newProgram := make([]int64, len(program))
			copy(newProgram, program)
			newProgram[1], newProgram[2] = a, b

			memory = runIntcodeProgram(newProgram)
			if memory[0] == 19690720 {
				return
			}
		}
	}
	return
}
