package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// lol is a func
func lol() int {
	return 4
}

func f(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fuel(mass int64) int64 {
	return (mass / 3) - 2
}

func fuel2(mass int64) int64 {
	if mass <= 0 {
		return 0
	}
	newFuel := (mass / 3) - 2
	if newFuel < 0 {
		newFuel = 0
	}
	return newFuel + fuel2(newFuel)
}

func main() {
	b, err := ioutil.ReadFile("cmd/day01/input01-1.txt")
	f(err)
	s := strings.Split(string(b), "\n")
	var totalFuelPart1, totalFuelPart2 int64
	for _, fuelStr := range s {
		fuelInt, err := strconv.ParseInt(fuelStr, 10, 64)
		f(err)
		totalFuelPart1 += fuel(fuelInt)
		totalFuelPart2 += fuel2(fuelInt)
	}
	log.Printf("total fuel part 1 is %v", totalFuelPart1)
	log.Printf("total fuel part 2 is %v", totalFuelPart2)

	// 9769030
}
