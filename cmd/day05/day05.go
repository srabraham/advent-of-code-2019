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

func RunPart2(filename string) int64 {
	b, err := ioutil.ReadFile(filename)
	f(err)
	nums := strings.Split(string(b), ",")
	cmds := make([]int64, 0)
	for _, num := range nums {
		n, err := strconv.ParseInt(num, 10, 64)
		f(err)
		cmds = append(cmds, n)
	}
	// 7616021
	return intcode.RunIntCodeMultiInput(cmds, []int64{5})
}

func main() {
	// 7616021
	log.Printf("output = %v", RunPart2("cmd/day05/input05-1.txt"))
}
