package main

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func f(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type AmountThing struct {
	amount int
	thing  string
}

type Formulas struct {
	produceToReq        map[AmountThing][]AmountThing
	produceStrToProduce map[string]AmountThing
}

func StringToThing(s string) AmountThing {
	s1 := strings.TrimSpace(s)
	splits := strings.Split(s1, " ")
	if len(splits) != 2 {
		log.Fatal("FDfdsfdsfd")
	}
	num, err := strconv.Atoi(splits[0])
	f(err)
	return AmountThing{amount: num, thing: splits[1]}
}

func CreateFormulas(rows []string) Formulas {
	produceStrToProduce := make(map[string]AmountThing)
	produceToReq := make(map[AmountThing][]AmountThing)

	for _, r := range rows {
		splits := strings.Split(r, "=>")
		if len(splits) != 2 {
			log.Fatalf("bad row in the file: %v", r)
		}
		left := strings.TrimSpace(splits[0])
		produces := StringToThing(strings.TrimSpace(splits[1]))
		var requires []AmountThing
		for _, s := range strings.Split(left, ",") {
			requires = append(requires, StringToThing(s))
		}
		produceToReq[produces] = requires
		produceStrToProduce[produces.thing] = produces
	}
	return Formulas{produceStrToProduce: produceStrToProduce, produceToReq: produceToReq}
}

// Solve finds the number of ore needed to produce the provided number of fuel.
func (f Formulas) Solve(requiredFuels int) int {
	requiredOutputs := map[string]int{"FUEL": requiredFuels}
	excess := make(map[string]int)

	continueLooping := true
	for continueLooping {
		continueLooping = false
		// We'll be removing things from the map during iteration.
		// That's fine in Go. https://golang.org/ref/spec#RangeClause
		for k := range requiredOutputs {
			if requiredOutputs[k] == 0 {
				delete(requiredOutputs, k)
				continue
			}
			if k != "ORE" {
				// If there's still a non-zero non-ore requirement, we must
				// keep looping.
				continueLooping = true
				produce := f.produceStrToProduce[k]
				outputUnits := max(produce.amount, requiredOutputs[k])
				roundsOfInputs := int(math.Ceil(float64(outputUnits) / float64(produce.amount)))
				excess[k] += roundsOfInputs*produce.amount - requiredOutputs[k]
				reqs := f.produceToReq[f.produceStrToProduce[k]]
				for _, r := range reqs {
					requiredOutputs[r.thing] += r.amount * roundsOfInputs
					for excess[r.thing] > 0 && requiredOutputs[r.thing] > 0 {
						requiredOutputs[r.thing] -= 1
						excess[r.thing] -= 1
					}
				}
				// Empty the requirements for k
				delete(requiredOutputs, k)
			}
		}
	}
	return requiredOutputs["ORE"]
}

func Part1(filename string) int {
	log.Print("starting...")
	b, err := ioutil.ReadFile(filename)
	f(err)
	formulas := CreateFormulas(strings.Split(string(b), "\n"))
	return formulas.Solve(1)
}

func Part2(filename string) int {
	log.Print("starting...")
	b, err := ioutil.ReadFile(filename)
	f(err)
	formulas := CreateFormulas(strings.Split(string(b), "\n"))

	log.Printf("Part 2: starting binary search")
	availableOre := 1000000000000
	var round int64
	minFuel := 0
	maxFuel := availableOre
	for minFuel != maxFuel {
		round++
		fuel := (maxFuel-minFuel+1)/2 + minFuel
		numOres := formulas.Solve(fuel)
		if round%10000 == 0 {
			log.Printf("requiredOutputs %v < %v ?", numOres, availableOre)
		}
		if numOres < availableOre {
			minFuel = fuel
			log.Printf("can do %v fuel", fuel)
		} else {
			maxFuel = fuel - 1
			log.Printf("can't do ore for %v fuels", fuel)
		}
	}
	log.Printf("Part 2: Done! can make a maximum of %v fuels", maxFuel)
	return maxFuel
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	log.Printf("Part 1: minimum ore/fuel is %v", Part1("cmd/day14/input14-1.txt"))
	log.Printf("Part 2: Done! can make a maximum of %v fuels", Part2("cmd/day14/input14-1.txt"))
}
