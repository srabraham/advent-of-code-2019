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

type GridPos struct {
	x int
	y int
}

type Grid struct {
	vals map[GridPos]string
}

func (g Grid) countVisited() int {
	return len(g.vals)
}

func (g Grid) String() string {
	var s string

	var minX, maxX, minY, maxY int
	for gp := range g.vals {
		if gp.x < minX {
			minX = gp.x
		}
		if gp.x > maxX {
			maxX = gp.x
		}
		if gp.y < minY {
			minY = gp.y
		}
		if gp.y > maxY {
			maxY = gp.y
		}
	}

	for y := minY; y < maxY+1; y++ {
		for x := minX; x < maxX+1; x++ {
			v := g.vals[GridPos{x: x, y: y}]
			if v == "." {
				v = "."
			}
			if v == "#" {
				v = "#"
			}
			if v == "" {
				v = "."
			}
			s += v
			s += " "
		}
		s += "\n"
	}
	return s
}

func Part1(filename string) int {
	b, err := ioutil.ReadFile(filename)
	f(err)
	nums := strings.Split(string(b), ",")
	cmds := make([]int64, 0)
	for _, num := range nums {
		n, err := strconv.ParseInt(num, 10, 64)
		f(err)
		cmds = append(cmds, n)
	}
	inCh := make(chan int64, 1)
	outCh := make(chan int64, 1)
	intcodeDone := make(chan bool, 1)
	go func() {
		intcode.RunIntCodeWithChannels(cmds, inCh, outCh)
		intcodeDone <- true
	}()

	g := Grid{vals: make(map[GridPos]string)}
	pos := GridPos{}
	dir := "u"
bigLoop:
	for true {
		// log.Printf("at pos %v", pos)
		val := g.vals[pos]
		var newVal int
		if val == "" || val == "." {
			newVal = 0
		} else if val == "#" {
			newVal = 1
		}
		select {
		case inCh <- int64(newVal):
		case <-intcodeDone:
			log.Printf("intcode program is done")
			break bigLoop
		}
		switch <-outCh {
		case 0:
			g.vals[pos] = "."
		case 1:
			g.vals[pos] = "#"
		default:
			log.Fatal("badddd")
		}
		pos, dir = getNewDirAndPos(<-outCh, dir, pos)
		//log.Printf("Grid = \n%v", g)
	}
	count := g.countVisited()
	log.Printf("visited %v spots", count)
	return count
}

func getNewDirAndPos(newDir int64, dir string, pos GridPos) (GridPos, string) {
	switch newDir {
	case 0:
		switch dir {
		case "u":
			dir = "l"
			pos.x = pos.x - 1
		case "l":
			dir = "d"
			pos.y = pos.y + 1
		case "r":
			dir = "u"
			pos.y = pos.y - 1
		case "d":
			dir = "r"
			pos.x = pos.x + 1
		}
	case 1:
		switch dir {
		case "d":
			dir = "l"
			pos.x = pos.x - 1
		case "r":
			dir = "d"
			pos.y = pos.y + 1
		case "l":
			dir = "u"
			pos.y = pos.y - 1
		case "u":
			dir = "r"
			pos.x = pos.x + 1
		}
	default:
		log.Fatal("bdsjbdf")
	}
	return pos, dir
}

func Part2(filename string) string {
	b, err := ioutil.ReadFile(filename)
	f(err)
	nums := strings.Split(string(b), ",")
	cmds := make([]int64, 0)
	for _, num := range nums {
		n, err := strconv.ParseInt(num, 10, 64)
		f(err)
		cmds = append(cmds, n)
	}
	inCh := make(chan int64, 1)
	outCh := make(chan int64, 1)
	intcodeDone := make(chan bool, 1)
	go func() {
		intcode.RunIntCodeWithChannels(cmds, inCh, outCh)
		intcodeDone <- true
	}()

	g := Grid{vals: make(map[GridPos]string)}
	pos := GridPos{}
	g.vals[pos] = "#"
	dir := "u"
bigLoop:
	for true {
		// log.Printf("at pos %v", pos)
		val := g.vals[pos]
		var newVal int
		if val == "" || val == "." {
			newVal = 0
		} else if val == "#" {
			newVal = 1
		}
		select {
		case inCh <- int64(newVal):
		case <-intcodeDone:
			log.Printf("intcode program is done")
			break bigLoop
		}
		// color painted
		switch <-outCh {
		case 0:
			g.vals[pos] = "."
		case 1:
			g.vals[pos] = "#"
		default:
			log.Fatal("badddd")
		}
		// read new direction
		pos, dir = getNewDirAndPos(<-outCh, dir, pos)
	}
	log.Printf("got grid:\n%v", g)
	return g.String()
}

func main() {
	Part1("cmd/day11/input11-1.txt")
	Part2("cmd/day11/input11-1.txt")
}
