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
	x int64
	y int64
}

type Grid struct {
	vals map[GridPos]string
}

func (g Grid) countVisited() int {
	return len(g.vals)
}

func (g Grid) String() string {
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

	for y := minY - 1; y < maxY+2; y++ {
		for x := minX - 1; x < maxX+2; x++ {
			v := g.vals[GridPos{x: x, y: y}]
			if v == "0" {
				v = " "
			}
			s += v
			s += " "
		}
		s += "\n"
	}
	return s
}

func (g Grid) PosOfBall() (GridPos, bool) {
	for k, v := range g.vals {
		if v == "4" {
			return k, true
		}
	}
	return GridPos{}, false
}

func (g Grid) PosOfPaddle() (GridPos, bool) {
	for k, v := range g.vals {
		if v == "3" {
			return k, true
		}
	}
	return GridPos{}, false
}

func main() {
	b, err := ioutil.ReadFile("cmd/day13/input13-0.txt")
	// b, err := ioutil.ReadFile("cmd/day13/input13-1.txt")
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
	intcodeDone := make(chan bool, 1)
	go func() {
		intcode.RunIntCodeWithChannels(cmds, inCh, outCh)
		log.Fatal("done")
		intcodeDone <- true
	}()

	g := Grid{vals: make(map[GridPos]string)}
	var score int64
	var joystickTilt int64
	doneDrawing := false
	for true {
	innerLoop:
		for true {
			select {
			case x := <-outCh:
				y := <-outCh
				if x == -1 && y == 0 {
					score = <-outCh
					log.Printf("score = %v", score)
				} else {
					tile := <-outCh
					g.vals[GridPos{x: x, y: y}] = strconv.Itoa(int(tile))
					if doneDrawing {
						log.Printf("got %v,%v,%v", x, y, tile)
						//log.Printf("score = %v, grid =\n%v", score, g)
					}
				}
			default:
				break innerLoop
			}
		}
		//x := <- outCh

		ballPos, isBall := g.PosOfBall()
		paddlePos, isPaddle := g.PosOfPaddle()
		if isBall && isPaddle {
			ballX := ballPos.x
			paddleX := paddlePos.x
			if ballX < paddleX {
				joystickTilt = -1
			}
			if ballX > paddleX {
				joystickTilt = 1
			}
			if ballX == paddleX {
				joystickTilt = 0
			}
			log.Printf("try to set tilt = %v", joystickTilt)
		}
		select {
		case inCh <- joystickTilt:
			log.Printf("did set tilt = %v", joystickTilt)
			doneDrawing = true
		default:
			//break innerLoop2
		}

	}
}
