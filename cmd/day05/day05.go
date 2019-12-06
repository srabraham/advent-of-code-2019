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

//func Result() int64 {
//	return 0
//}

// 223
func main() {
	// b, err := ioutil.ReadFile("cmd/day05/input05-0.txt")
	b, err := ioutil.ReadFile("cmd/day05/input05-1.txt")
	f(err)
	nums := strings.Split(string(b), ",")
	cmds := make([]int64, 0)
	for _, num := range nums {
		n, err := strconv.ParseInt(num, 10, 64)
		f(err)
		cmds = append(cmds, n)
	}
	// 7616021
	log.Printf("output = %v", intcode.RunIntCode(cmds, 5))
}
