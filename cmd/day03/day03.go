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
	b, err := ioutil.ReadFile("cmd/day03/input03-1.txt")
	f(err)
	lines := strings.Split(string(b), "\n")
	w1 := strings.Split(lines[0], ",")
	w2 := strings.Split(lines[1], ",")

	grid1 := getGrid(w1)
	grid2 := getGrid(w2)

	log.Print(grid1)
	log.Print(grid2)

	for x, yMap := range grid1 {
		for y := range yMap {
			if grid2[x] != nil && grid2[x][y] > 0 {
				log.Printf("found intersection at %v, %v. === %v", x, y, grid1[x][y]+grid2[x][y])
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}


// 16592
func getGrid(w1 []string) map[int]map[int]int {
	grid1 := make(map[int]map[int]int)
	var currX, currY int
	stepCount := 0
	for _, w := range w1 {
		dir := w[:1]
		length, err := strconv.Atoi(w[1:])
		f(err)
		switch dir {
		case "U":
			for y := 0; y < length; y++ {
				currY++
				stepCount++
				initG(grid1, currX)
				if grid1[currX][currY] == 0 {
					//stepCount = grid1[currX][currY]
					grid1[currX][currY] = stepCount
				}
				//grid1[currX][currY] = stepCount
			}
		case "D":
			for y := 0; y < length; y++ {
				currY--
				stepCount++
				initG(grid1, currX)
				if grid1[currX][currY] == 0 {
					//stepCount = grid1[currX][currY]
					grid1[currX][currY] = stepCount
				}
				//grid1[currX][currY] = stepCount
			}
		case "L":
			for x := 0; x < length; x++ {
				currX--
				stepCount++
				initG(grid1, currX)
				if grid1[currX][currY] == 0 {
					//stepCount = grid1[currX][currY]
					grid1[currX][currY] = stepCount
				}
				//grid1[currX][currY] = stepCount
			}
		case "R":
			for x := 0; x < length; x++ {
				currX++
				stepCount++
				initG(grid1, currX)
				if grid1[currX][currY] == 0 {
					//stepCount = grid1[currX][currY]
					grid1[currX][currY] = stepCount
				}
				//grid1[currX][currY] = stepCount
			}
		default:
			log.Fatalf("bad %v", dir)
		}
	}
	return grid1
}

func initG(g map[int]map[int]int, x int) {
	if g[x] == nil {
		g[x] = make(map[int]int)
	}
}