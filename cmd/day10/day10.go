package main

import (
	"io/ioutil"
	"log"
	"math"
	"sort"
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
	var maxX, maxY int
	for x, row := range collisions {
		log.Print(row)
		for y, v := range row {
			if v > max {
				max = v
				maxX = x
				maxY = y
			}
		}
	}
	log.Printf("max = %v, maxX = %v, maxY = %v", max, maxX, maxY)

	getRelAsts(grid, 11,11)
}


type RelAst struct {
	x int
	y int
	distX int
	distY int
}

func (ra RelAst) dist() int {
	return ra.distX * ra.distX + ra.distY * ra.distY
}

func getRelAsts(grid [][]string, x int, y int) {
	var ra []RelAst
	for checkX := range grid {
		for checkY, checkVal := range grid[checkX] {
			if checkX == x && checkY == y {
				continue
			}
			if checkVal == "#" {
				r := RelAst{
					x: checkX,
					y: checkY,
					distX: checkX - x,
					distY: checkY-y,
				}
				ra = append(ra,r)
				log.Printf("created ra %v", r)
			}
		}
	}


	mappedAngles := make(map[float64][]RelAst)

	for _, r := range ra {
		angle := math.Atan2(-float64(r.distY), float64(r.distX))
		mappedAngles[angle] = append(mappedAngles[angle], r)
	}

	angles := make([]float64, 0)
	for k := range mappedAngles {
		angles = append(angles, k)
	}

	sort.Float64s(angles)
	for _, s := range mappedAngles {
		sort.Slice(s, func(i, j int) bool {
			return math.Sqrt(float64(ra[i].distX)*float64(ra[i].distX)+float64(ra[i].distY)*float64(ra[i].distY)) < math.Sqrt(float64(ra[j].distX)*float64(ra[j].distX)+float64(ra[j].distY)*float64(ra[j].distY))
		})
	}

	blast := 0
	blastsThisPass := 1
	for pass := 0; blastsThisPass > 0; pass++ {
		blastsThisPass = 0
		for _, a := range angles {
			ma, found := mappedAngles[a]
			if !found {
				continue
			}
			if len(ma) <= pass {
				continue
			}
			blast++
			blastsThisPass++
			val := ma[pass]
			log.Printf("%v entry is %v", blast, val)
		}
	}
}


func countCollisions(grid [][]string, x int, y int) int {
	collisions := 0
	for checkX := range grid {
		for checkY, checkVal := range grid[checkX] {
			if checkX == x && checkY == y {
				continue
			}
			if checkVal == "#" {
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
