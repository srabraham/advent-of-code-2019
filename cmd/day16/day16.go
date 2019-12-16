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

func main() {
	log.Print("starting...")
	//b, err := ioutil.ReadFile("cmd/day16/input16-0.txt")
	b, err := ioutil.ReadFile("cmd/day16/input16-1.txt")
	f(err)
	nums := strings.Split(string(b), ",")
	cmds := make([]int64, 0)
	for _, num := range nums {
		n, err := strconv.ParseInt(num, 10, 64)
		f(err)
		cmds = append(cmds, n)
	}
}
