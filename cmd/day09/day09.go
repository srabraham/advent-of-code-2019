package main

import (
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

func RunBoost(filename string) {
	b, err := ioutil.ReadFile(filename)
	f(err)
	nums := strings.Split(string(b), ",")
	cmds := make([]int64, 0)
	for _, num := range nums {
		n, err := strconv.ParseInt(num, 10, 64)
		f(err)
		cmds = append(cmds, n)
	}
	inCh := make(chan int64, 10000)
	outCh := make(chan int64, 10000)
	inCh <- 2
	go intcode.RunIntCodeWithChannels(cmds, inCh, outCh)

	var outs []int64
	for a := range outCh {
		log.Printf("GOT %v", a)
		outs = append(outs, a)
	}
	log.Printf("outs = %v", outs)
}

func main() {
	//b, err := ioutil.ReadFile("cmd/day09/input09-0.txt")
	b, err := ioutil.ReadFile("cmd/day09/input09-1.txt")
	f(err)
	nums := strings.Split(string(b), ",")
	cmds := make([]int64, 0)
	for _, num := range nums {
		n, err := strconv.ParseInt(num, 10, 64)
		f(err)
		cmds = append(cmds, n)
	}
	inCh := make(chan int64, 10000)
	outCh := make(chan int64, 10000)
	inCh <- 2
	go intcode.RunIntCodeWithChannels(cmds, inCh, outCh)

	var outs []int64
	for a := range outCh {
		log.Printf("GOT %v", a)
		outs = append(outs, a)
	}
	log.Printf("outs = %v", outs)
	// 3906448201 correct
	// 59785 correct

}

// 203, 0
