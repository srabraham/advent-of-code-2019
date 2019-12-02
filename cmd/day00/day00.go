package main

import (
	"io/ioutil"
	"log"
	"strings"
)

// lol is a func
func lol() int {
	return 4
}

func f(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Print("hello world")
	b, err := ioutil.ReadFile("cmd/day00/input.txt")
	f(err)
	s := strings.Split(string(b), "\n")
	log.Printf("arr = %v", s)

}
