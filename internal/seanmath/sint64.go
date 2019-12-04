package seanmath

import (
	"log"
	"sort"
	"strconv"
)

func Min(a []int64) int64 {
	if len(a) == 0 {
		log.Fatal("empty slice for min call")
	}
	min := a[0]
	for _, n := range a {
		if n < min {
			min = n
		}
	}
	return min
}

func Max(a []int64) int64 {
	if len(a) == 0 {
		log.Fatal("empty slice for max call")
	}
	max := a[0]
	for _, n := range a {
		if n > max {
			max = n
		}
	}
	return max
}

func Cp(a []int64) []int64 {
	cp := make([]int64, len(a))
	copy(cp, a)
	return cp
}

func Sort(a []int64) {
	sort.Slice(a, func(i, j int) bool {return a[i] < a[j]})
}

func ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

func ParseString(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("failed to parse %v: %v", s, err)
	}
	return n
}