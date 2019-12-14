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
	requiredOutputs := make(map[string]int)
	excess := make(map[string]int)
	requiredOutputs["FUEL"] = requiredFuels
	//log.Printf("%v", produceToReq[produceStrToProduce["FUEL"]])
	alreadyMade := make(map[string]bool)
	for true {
		done := true
		for k := range requiredOutputs {
			if requiredOutputs[k] == 0 {
				continue
			}
			if k != "ORE" {
				done = false
				produce := f.produceStrToProduce[k]
				outputUnits := max(produce.amount, requiredOutputs[k])
				roundsOfInputs := int(math.Ceil(float64(outputUnits) / float64(produce.amount)))
				excess[k] += roundsOfInputs*produce.amount - requiredOutputs[k]
				reqs := f.produceToReq[f.produceStrToProduce[k]]
				//log.Printf("making %v of %v requires %v mult of %v", outputUnits, k, roundsOfInputs, reqs)
				for _, r := range reqs {
					requiredOutputs[r.thing] += r.amount * roundsOfInputs
					for excess[r.thing] > 0 && requiredOutputs[r.thing] > 0 {
						requiredOutputs[r.thing] -= 1
						excess[r.thing] -= 1
					}
				}
				requiredOutputs[k] = 0
				alreadyMade[k] = true
				//log.Printf("after change, requiredOutputs %v", requiredOutputs)
			}
		}
		if done {
			break
		}
	}
	return requiredOutputs["ORE"]
}

func main() {
	log.Print("starting...")
	b, err := ioutil.ReadFile("cmd/day14/input14-1.txt")
	f(err)
	formulas := CreateFormulas(strings.Split(string(b), "\n"))

	log.Printf("Part 1: minimum ore/fuel is %v", formulas.Solve(1))

	log.Printf("Part 2: starting binary search")
	availableOre := 1000000000000
	var round int64
	minFuel := 0
	maxFuel := availableOre
	for minFuel != maxFuel {
		round++
		fuel := (maxFuel - minFuel + 1) / 2 + minFuel
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
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	if a < 0 {
		a = -a
	}
	//log.Printf("gcd = %v", a)
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
