package main

import (
	"fmt"
	"strings"

	"github.com/uberx/advent-of-code-2019/mathy"
	"github.com/uberx/advent-of-code-2019/util"
)

func main() {
	input := util.ReadFile("day1.txt")
	answer1 := sumOfFuelRequirements1(input)
	fmt.Println("Part 1:", answer1)

	answer2 := sumOfFuelRequirements2(input)
	fmt.Println("Part 2:", answer2)
}

func sumOfFuelRequirements1(input string) int {
	moduleMasses := parseInput(input)
	totalFuelRequired := 0
	for _, moduleMass := range moduleMasses {
		totalFuelRequired += fuelRequired(moduleMass)
	}
	return totalFuelRequired
}

func sumOfFuelRequirements2(input string) int {
	moduleMasses := parseInput(input)
	totalFuelRequired := 0
	for _, moduleMass := range moduleMasses {
		currFuelRequired := fuelRequired(moduleMass)
		for currFuelRequired > 0 {
			totalFuelRequired += currFuelRequired
			currFuelRequired = fuelRequired(currFuelRequired)
		}
	}
	return totalFuelRequired
}

func fuelRequired(mass int) int {
	fuelRequired := mass/3 - 2
	return mathy.Max(0, fuelRequired)
}

func parseInput(input string) (moduleMasses []int) {
	for _, moduleMass := range strings.Split(input, "\n") {
		moduleMasses = append(moduleMasses, util.ToInt(moduleMass))
	}
	return moduleMasses
}
