package main

import (
	"fmt"
	"github.com/srabraham/advent-of-code-2019/internal/intcode"
	"go/constant"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Cell int

const (
	Blank Cell = iota
	Wall
	Block
	Paddle
	Ball
)

func f(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Cell) String() string {
	switch c {
	case Blank:
		return " "
	case Wall:
		return "█"
	case Block:
		return "□"
	case Paddle:
		return "―"
	case Ball:
		return "⬤"
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
			s += fmt.Sprintf("%v", v)
		}
		s += "\n"
	}
	return s
}

func (g GameState) PosOf(c Cell) (GridPos, bool) {
	for k, v := range g.vals {
		if v == c {
			return k, true
		}
	}
	return GridPos{}, false
}

func RunGame(filename string, watchGame bool) *GameState {
	log.Print("starting...")
	b, err := ioutil.ReadFile(filename)
	f(err)
	nums := strings.Split(string(b), ",")
	cmds := make([]int64, 0)
	for _, num := range nums {
		n, err := strconv.ParseInt(num, 10, 64)
		f(err)
		cmds = append(cmds, n)
	}
	inCh := make(chan int64)
	outCh := make(chan int64)
	intcodeDone := make(chan bool)
	go func() {
		intcode.RunIntCodeWithChannels(cmds, inCh, outCh)
		intcodeDone <- true
	}()
	g := GameState{vals: make(map[GridPos]Cell)}
	doneInitialDrawing := false
loop:
	for true {
		ballPos, isBall := g.PosOf(Ball)
		paddlePos, isPaddle := g.PosOf(Paddle)
		if isBall && isPaddle {
			g.joystickDir = constant.Sign(constant.Make(ballPos.x - paddlePos.x))
		}
		select {
		case <-intcodeDone:
			break loop
		case x := <-outCh:
			y := <-outCh
			if x == -1 && y == 0 {
				g.score = <-outCh
				//log.Printf("score = %v", score)
			} else {
				tile := <-outCh
				g.vals[GridPos{x: x, y: y}] = Cell(int(tile))
				if doneInitialDrawing && watchGame {
					time.Sleep(10 * time.Millisecond)
					cmd := exec.Command("clear")
					cmd.Stdout = os.Stdout
					cmd.Run()
					log.Printf("score = %v, grid =\n%v", g.score, g)
				}
			}
		case inCh <- int64(g.joystickDir):
			doneInitialDrawing = true
		}
	}
	log.Printf("Done. Score = %v", g.score)
	return &g
}

func main() {
	// Part 1
	g := RunGame("cmd/day13/input13-0.txt", false)
	log.Printf("found %v blocks", g.count(Block))

	// Part 2
	RunGame("cmd/day13/input13-1.txt", true)
}
