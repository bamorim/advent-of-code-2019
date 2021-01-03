package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/bamorim/advent-of-code-2019/day1"
	"github.com/bamorim/advent-of-code-2019/day2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// By default, runs last day
	day := 2

	if len(os.Args) > 1 {
		parsed, err := strconv.Atoi(os.Args[1])
		check(err)
		day = parsed
	}

	filename := fmt.Sprintf("./day%d/input.txt", day)
	bytes, err := ioutil.ReadFile(filename)
	check(err)
	fmt.Println(os.Args)
	switch day {
	case 1:
		day1.Run(bytes)
	case 2:
		day2.Run(bytes)
	default:
		fmt.Println("Invalid Day")
	}
}
