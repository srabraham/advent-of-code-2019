package main

import (
	"fmt"
	"github.com/srabraham/advent-of-code-2019/internal/intcode"
	"io/ioutil"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

type Cell int

const (
	Unknown Cell = iota
	Empty
	Wall
	Oxygen
)

func f(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Cell) String() string {
	switch c {
	case Unknown:
		return " "
	case Empty:
		return "."
	case Wall:
		return "█"
	case Oxygen:
		return "O"
	}
	return strconv.Itoa(int(c))
}

type GridPos struct {
	x int64
	y int64
}

type GameState struct {
	vals              map[GridPos]Cell
	necessaryUnknowns map[GridPos]bool
	droidPos          GridPos
	startPos          GridPos
	finishPos         GridPos
}

func (g GameState) count(c Cell) int {
	var count int
	for _, v := range g.vals {
		if v == c {
			count++
		}
	}
	return count
}

func (g GameState) discoveredWholeBoard() bool {
	for k := range g.necessaryUnknowns {
		if g.vals[k] == Unknown {
			return false
		}
		delete(g.necessaryUnknowns, k)
	}
	for k := range g.vals {
		if g.vals[k] != Wall {
			if g.vals[GridPos{x: k.x - 1, y: k.y}] == Unknown {
				g.necessaryUnknowns[GridPos{x: k.x - 1, y: k.y}] = true
				return false
			}
			if g.vals[GridPos{x: k.x + 1, y: k.y}] == Unknown {
				g.necessaryUnknowns[GridPos{x: k.x + 1, y: k.y}] = true
				return false
			}
			if g.vals[GridPos{x: k.x, y: k.y - 1}] == Unknown {
				g.necessaryUnknowns[GridPos{x: k.x, y: k.y - 1}] = true
				return false
			}
			if g.vals[GridPos{x: k.x, y: k.y + 1}] == Unknown {
				g.necessaryUnknowns[GridPos{x: k.x, y: k.y + 1}] = true
				return false
			}
		}
	}
	return true
}

func (g GameState) String() string {
	var s string

	var minX, maxX, minY, maxY int64
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
			pos := GridPos{x: x, y: y}
			var cellStr string
			if pos == g.droidPos {
				cellStr = "D"
			} else if pos == g.startPos {
				cellStr = "S"
			} else if pos == g.finishPos {
				cellStr = "F"
			} else {
				cellStr = g.vals[pos].String()
			}
			s += fmt.Sprintf("%v ", cellStr)
		}
		s += "\n"
	}
	return s
}

func main() {
	log.Print("starting...")
	b, err := ioutil.ReadFile("cmd/day15/input15-1.txt")
	f(err)
	nums := strings.Split(string(b), ",")
	cmds := make([]int64, 0)
	for _, num := range nums {
		n, err := strconv.ParseInt(num, 10, 64)
		f(err)
		cmds = append(cmds, n)
	}
	g := GameState{vals: make(map[GridPos]Cell), necessaryUnknowns: make(map[GridPos]bool)}
	g.vals[GridPos{x: 0, y: 0}] = Empty
	g.startPos = GridPos{x: 0, y: 0}
	//droidPos := GridPos{x: 0, y: 0}
	inCh := make(chan int64)
	outCh := make(chan int64)
	intcodeDone := make(chan bool)
	go func() {
		intcode.RunIntCodeWithChannels(cmds, inCh, outCh)
		intcodeDone <- true
	}()
	var oxSys GridPos
	for round := 1; ; round++ {
		movementCommand := rand.Intn(4) + 1
		var targetPos GridPos
		switch movementCommand {
		case 1:
			targetPos = GridPos{x: g.droidPos.x, y: g.droidPos.y - 1}
		case 2:
			targetPos = GridPos{x: g.droidPos.x, y: g.droidPos.y + 1}
		case 3:
			targetPos = GridPos{x: g.droidPos.x - 1, y: g.droidPos.y}
		case 4:
			targetPos = GridPos{x: g.droidPos.x + 1, y: g.droidPos.y}
		default:
			log.Fatal("bad move")
		}
		inCh <- int64(movementCommand)
		result := <-outCh
		switch result {
		case 0:
			g.vals[targetPos] = Wall
		case 1:
			g.droidPos = targetPos
			g.vals[targetPos] = Empty
		case 2:
			oxSys = targetPos
			g.vals[targetPos] = Empty
			g.finishPos = targetPos
			g.droidPos = targetPos
		}
		if round%20000 == 0 {
			log.Printf("grid after %v rounds =\n%v", round, g)
		}
		if g.discoveredWholeBoard() {
			log.Printf("discovered whole board after %v rounds", round)
			log.Printf("grid after %v rounds =\n%v", round, g)
			break
		}
		// 270 by manual count :/
	}

	g.vals[oxSys] = Oxygen
	for min := 1; ; min++ {
		var makeOxy []GridPos
		//var minX, maxX, minY, maxY int64
		for gp, val := range g.vals {
			if val == Oxygen {
				up := GridPos{x: gp.x, y: gp.y - 1}
				if EmptyPosLol(g.vals[up]) {
					makeOxy = append(makeOxy, up)
				}
				down := GridPos{x: gp.x, y: gp.y + 1}
				if EmptyPosLol(g.vals[down]) {
					makeOxy = append(makeOxy, down)
				}
				left := GridPos{x: gp.x - 1, y: gp.y}
				if EmptyPosLol(g.vals[left]) {
					makeOxy = append(makeOxy, left)
				}
				right := GridPos{x: gp.x + 1, y: gp.y}
				if EmptyPosLol(g.vals[right]) {
					makeOxy = append(makeOxy, right)
				}
			}
		}
		for _, mo := range makeOxy {
			g.vals[mo] = Oxygen
		}
		log.Printf("grid after %v mins =\n%v", min, g)
		if g.count(Empty) == 0 {
			log.Fatalf("done after %v minutes", min)
		}
	}
}

func EmptyPosLol(c Cell) bool {
	return c == Empty
}
