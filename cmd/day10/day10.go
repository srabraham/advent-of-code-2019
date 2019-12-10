package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func f(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	b, err := ioutil.ReadFile("cmd/day10/input10-1.txt")
	f(err)
	rows := strings.Split(string(b), "\n")
	grid := make([][]string, 0)
	var collisions [][]int
	for _, r := range rows {
		var g []string
		for _, c := range r {
			g = append(g, string(c))
		}
		grid = append(grid, g)
		collisions = append(collisions, make([]int, len(g)))
	}
	log.Print(grid)

	for x := range grid {
		for y, val := range grid[x] {
			if val == "#" {
				collisions[x][y] = countCollisions(grid, x, y)
			}
		}
	}
	var max int
	for _, row := range collisions {
		log.Print(row)
		for _, v := range row {
			if v > max {
				max = v
			}
		}
	}
	log.Printf("max = %v", max)
}

func countCollisions(grid [][]string, x int, y int) int {
	collisions := 0
	for checkX := range grid {
		for checkY, checkVal := range grid[checkX] {
			if checkX == x && checkY == y {
				continue
			}
			if checkVal == "#" {

				if x == 9 && y == 9 && checkX == 0 && checkY == 6 {
					log.Print("here")
				}

				rise := checkY - y
				run := checkX - x
				gcd := GCD(rise, run)
				rise = rise / gcd
				run = run / gcd
				xRay := x
				yRay := y
				intersect := false
				loop:
				for true {
					xRay += run
					yRay += rise
					if xRay == checkX && yRay == checkY {
						break loop
					}
					if xRay < 0 || xRay >= len(grid) || yRay < 0 || yRay >= len(grid[xRay]) {
						break loop
					}
					if grid[xRay][yRay] == "#" {
						log.Printf("rise %v run %v intersect (%v,%v) <> (%v,%v) <> (%v,%v)", rise, run, x,y,xRay,yRay,checkX,checkY)
						intersect = true
						break loop
					}
				}
				if !intersect {
					collisions++
				}
			}
		}
	}
	return collisions
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	if a < 0 {
		a = -a
	}
	log.Printf("gcd = %v", a)
	return a
}
