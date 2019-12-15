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
	Droid
	OxygenSys
	Start
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
	case Droid:
		return "D"
	case Start:
		return "S"
	case OxygenSys:
		return "F"
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
	vals        map[GridPos]Cell
	score       int64
	joystickDir int
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
			v := g.vals[GridPos{x: x, y: y}]
			s += fmt.Sprintf("%v ", v)
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
	g := GameState{vals: make(map[GridPos]Cell)}
	droidPos := GridPos{x: 0, y: 0}
	inCh := make(chan int64)
	outCh := make(chan int64)
	intcodeDone := make(chan bool)
	go func() {
		intcode.RunIntCodeWithChannels(cmds, inCh, outCh)
		intcodeDone <- true
	}()
	var round int
	var oxSys GridPos
	for true {
		round++
		//if g.vals[GridPos{x:}]
		movementCommand := rand.Intn(4) + 1
		var targetPos GridPos
		switch movementCommand {
		case 1:
			targetPos = GridPos{x: droidPos.x, y: droidPos.y - 1}
		case 2:
			targetPos = GridPos{x: droidPos.x, y: droidPos.y + 1}
		case 3:
			targetPos = GridPos{x: droidPos.x - 1, y: droidPos.y}
		case 4:
			targetPos = GridPos{x: droidPos.x + 1, y: droidPos.y}
		default:
			log.Fatal("bad move")
		}
		inCh <- int64(movementCommand)
		result := <-outCh
		switch result {
		case 0:
			g.vals[targetPos] = Wall
		case 1:
			g.vals[targetPos] = Droid
			g.vals[droidPos] = Empty
			droidPos = targetPos
		case 2:
			oxSys = targetPos
			g.vals[targetPos] = OxygenSys
			g.vals[droidPos] = Empty
			droidPos = targetPos
			//
			//log.Printf("grid =\n%v", g)
			//log.Fatal("done")
		}
		g.vals[GridPos{x: 0, y: 0}] = Start
		g.vals[oxSys] = OxygenSys
		if round%1000 == 0 {
			log.Printf("grid =\n%v", g)
		}
		if round > 1000000 {
			break
		}
		// 270
		// not 270 for part 2
	}
	g.vals[GridPos{x: 0, y: 0}] = Start
	g.vals[oxSys] = OxygenSys

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
	return c == Empty || c == Droid || c == Start
}
