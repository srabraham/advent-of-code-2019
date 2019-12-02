package main

import (
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

func run(in []int64) []int64 {
	currPos := 0
	for {
		switch in[currPos] {
		case 1:
			log.Printf("doing an add: %v, %v, %v, %v", in[currPos], in[currPos+1], in[currPos+2], in[currPos+3])
			aVal := in[currPos+1]
			bVal := in[currPos+2]
			cVal := in[currPos+3]
			in[cVal] = in[aVal] + in[bVal]
			currPos += 4
		case 2:
			log.Printf("doing an mult: %v, %v, %v, %v", in[currPos], in[currPos+1], in[currPos+2], in[currPos+3])
			aVal := in[currPos+1]
			bVal := in[currPos+2]
			cVal := in[currPos+3]
			in[cVal] = in[aVal] * in[bVal]
			currPos += 4
		case 99:
			return in
		default:
			log.Fatalf("bad opcode %v", in[currPos])
		}
	}
	log.Fatal("shouldn't happen")
	return nil
}

func main() {
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
			res := run(cmdsCopy)[0]
			log.Printf("for noun %v, verb %v got res = %v", noun, verb, res)
			if res == 19690720 {
				return
			}
		}
	}
}
