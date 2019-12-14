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

func main() {
	log.Print("starting...")
	b, err := ioutil.ReadFile("cmd/day14/input14-0.txt")
	f(err)
	rows := strings.Split(string(b), "\n")

	produceStrToProduce := make(map[string]AmountThing)
	produceToReq := make(map[AmountThing][]AmountThing)

	for _, r := range rows {
		splits := strings.Split(r, "=>")
		if len(splits) != 2 {
			log.Fatal("FDfd")
		}
		left := strings.TrimSpace(splits[0])
		right := strings.TrimSpace(splits[1])

		produces := StringToThing(right)

		var requires []AmountThing
		for _, s := range strings.Split(left, ",") {
			requires = append(requires, StringToThing(s))
		}

		produceToReq[produces] = requires
		produceStrToProduce[produces.thing] = produces
	}

	minVal := -1

	var round int64
	for true {
		round++
		have := make(map[string]int)
		excess := make(map[string]int)
		have["FUEL"] = 1
		//log.Printf("%v", produceToReq[produceStrToProduce["FUEL"]])
		alreadyMade := make(map[string]bool)
		for true {
			done := true
			for k := range have {
				if have[k] == 0 {
					continue
				}
				if k != "ORE" {
					done = false
					produce := produceStrToProduce[k]
					needToMake := max(produce.amount, have[k])
					roundsOfInputs := int(math.Ceil(float64(needToMake) / float64(produce.amount)))
					excess[k] += roundsOfInputs * produce.amount - have[k]
					reqs := produceToReq[produceStrToProduce[k]]
					//log.Printf("making %v of %v requires %v mult of %v", needToMake, k, roundsOfInputs, reqs)
					for _, r := range reqs {
						have[r.thing] += r.amount * roundsOfInputs
						for excess[r.thing] > 0 && have[r.thing] > 0 {
							have[r.thing] -= 1
							excess[r.thing] -= 1
						}
					}
					have[k] = 0
					alreadyMade[k] = true
					//log.Printf("after change, have %v", have)
				}
				//time.Sleep(500*time.Millisecond)
			}
			// not 169644
			if done {
				newV := have["ORE"]
				if minVal < 0 {
					minVal = newV
				} else {
					minVal = min(newV, minVal)
				}
				break
			}

		}
		// not 161753 either
		if round % 10000 == 0 {
			log.Printf("minVal = %v", minVal)




		}
	}

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
