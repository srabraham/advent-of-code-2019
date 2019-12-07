package main

import (
	"github.com/srabraham/advent-of-code-2019/internal/seanmath"
	"log"
)

func f(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func matchesPart1Criteria(n int64) bool {
	s := seanmath.ToString(n)
	counts := make(map[int64]int)
	for i := 0; i < len(s)-1; i++ {
		counts[seanmath.ParseString(s[i:i+1])]++
		if s[i:i+1] > s[i+1:i+2] {
			return false
		}
	}
	counts[seanmath.ParseString(s[len(s)-1:])]++
	for _, v := range counts {
		if v >= 2 {
			return true
		}
	}
	return false
}

func matchesPart2Criteria(n int64) bool {
	s := seanmath.ToString(n)
	counts := make(map[int64]int)
	for i := 0; i < len(s)-1; i++ {
		counts[seanmath.ParseString(s[i:i+1])]++
		if s[i:i+1] > s[i+1:i+2] {
			return false
		}
	}
	counts[seanmath.ParseString(s[len(s)-1:])]++
	for _, v := range counts {
		if v == 2 {
			return true
		}
	}
	return false
}

func main() {
	numsPt1 := make([]int64, 0)
	for i := int64(171309); i <= 643603; i++ {
		if matchesPart1Criteria(i) {
			numsPt1 = append(numsPt1, i)
		}
	}
	log.Printf("part 1 len = %v, numsPt1 = %v", len(numsPt1), numsPt1)

	numsPt2 := make([]int64, 0)
	for i := int64(171309); i <= 643603; i++ {
		if matchesPart2Criteria(i) {
			numsPt2 = append(numsPt2, i)
		}
	}
	log.Printf("part 2 len = %v, numsPt2 = %v", len(numsPt2), numsPt2)
}
