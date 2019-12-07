package main

import (
	"github.com/gitchander/permutation"
	"github.com/srabraham/advent-of-code-2019/internal/intcode"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func f(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Part1(filename string) int64 {
	b, err := ioutil.ReadFile(filename)
	f(err)
	nums := strings.Split(string(b), ",")
	cmds := make([]int64, 0)
	for _, num := range nums {
		n, err := strconv.ParseInt(num, 10, 64)
		f(err)
		cmds = append(cmds, n)
	}

	seq := []int{4, 3, 2, 1, 0}
	p := permutation.New(permutation.IntSlice(seq))
	var max int64
	for p.Next() {
		aOut := intcode.RunIntCodeMultiInput(cmds, []int64{int64(seq[0]), 0})
		bOut := intcode.RunIntCodeMultiInput(cmds, []int64{int64(seq[1]), aOut})
		cOut := intcode.RunIntCodeMultiInput(cmds, []int64{int64(seq[2]), bOut})
		dOut := intcode.RunIntCodeMultiInput(cmds, []int64{int64(seq[3]), cOut})
		eOut := intcode.RunIntCodeMultiInput(cmds, []int64{int64(seq[4]), dOut})
		log.Printf("got result %v", eOut)
		if eOut > max {
			max = eOut
		}
	}
	log.Printf("max = %v", max)
	return max
}

func Part2(filename string) int64 {
	b, err := ioutil.ReadFile(filename)
	f(err)
	nums := strings.Split(string(b), ",")
	cmds := make([]int64, 0)
	for _, num := range nums {
		n, err := strconv.ParseInt(num, 10, 64)
		f(err)
		cmds = append(cmds, n)
	}

	seq := []int{9, 8, 7, 6, 5}
	p := permutation.New(permutation.IntSlice(seq))
	var max int64
	for p.Next() {
		aIn := make(chan int64, 1)
		bIn := make(chan int64, 1)
		cIn := make(chan int64, 1)
		dIn := make(chan int64, 1)
		eIn := make(chan int64, 1)

		aIn <- int64(seq[0])
		bIn <- int64(seq[1])
		cIn <- int64(seq[2])
		dIn <- int64(seq[3])
		eIn <- int64(seq[4])

		go intcode.RunIntCodeWithChannels(cmds, aIn, bIn)
		aIn <- 0
		go intcode.RunIntCodeWithChannels(cmds, bIn, cIn)
		go intcode.RunIntCodeWithChannels(cmds, cIn, dIn)
		go intcode.RunIntCodeWithChannels(cmds, dIn, eIn)
		result := intcode.RunIntCodeWithChannels(cmds, eIn, aIn)
		log.Printf("result = %v", result)
		if result > max {
			max = result
		}
	}
	log.Printf("max = %v", max)
	return max
}

func main() {
	Part1("cmd/day07/input07-1.txt")
	Part2("cmd/day07/input07-1.txt")
}
