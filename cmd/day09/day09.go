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

func RunBoost(filename string, inputs []int64) []int64 {
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
	for _, in := range inputs {
		inCh <- in
	}
	go intcode.RunIntCodeWithChannels(cmds, inCh, outCh)

	var outputs []int64
	for a := range outCh {
		log.Printf("GOT %v", a)
		outputs = append(outputs, a)
	}
	log.Printf("outs = %v", outputs)
	return outputs
}

func main() {
	RunBoost("cmd/day09/input09-0.txt", []int64{})
	RunBoost("cmd/day09/input09-1.txt", []int64{1})
}
