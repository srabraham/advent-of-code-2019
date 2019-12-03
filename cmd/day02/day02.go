package main

import (
	"github.com/srabraham/advent-of-code-2019/internal/intcode"
	"io/ioutil"
	"log"
	"runtime"
	"strconv"
	"strings"
)

func f(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Printf("num CPU = %v", runtime.NumCPU())
	b, err := ioutil.ReadFile("cmd/day02/input02-1a.txt")
	f(err)
	nums := strings.Split(string(b), ",")
	cmds := make([]int64, 0)
	for _, num := range nums {
		n, err := strconv.ParseInt(num, 10, 64)
		f(err)
		cmds = append(cmds, n)
	}
	log.Printf("cmds = %v", cmds)
	//log.Printf("result = %v", run(cmds))

	for noun := int64(0); noun < 100; noun++ {
		for verb := int64(0); verb < 100; verb++ {
			cmdsCopy := make([]int64, len(cmds))
			copy(cmdsCopy, cmds)
			cmdsCopy[1] = noun
			cmdsCopy[2] = verb
			res := intcode.RunIntCode(cmdsCopy)[0]
			log.Printf("for noun %v, verb %v got res = %v", noun, verb, res)
			if res == 19690720 {
				return
			}
		}
	}
}
