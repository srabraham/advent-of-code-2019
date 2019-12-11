package main

import (
	"github.com/srabraham/advent-of-code-2019/internal/seanmath"
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

type GridPos struct {
	x int
	y int
}

type Grid struct {
	vals map[GridPos]string
	maxX int
	maxY int
}

func (g Grid) ValAt(x, y int) (string, bool) {
	val, inBounds := g.vals[GridPos{x: x, y: y}]
	return val, inBounds
}

func (g Grid) String() string {
	var s string
	for y := 0; y < g.maxY; y++ {
		for x := 0; x < g.maxX; x++ {
			s += g.vals[GridPos{x: x, y: y}]
			s += " "
		}
		s += "\n"
	}
	return s
}

func ReadGrid(filename string) Grid {
	b, err := ioutil.ReadFile(filename)
	f(err)
	rows := strings.Split(string(b), "\n")
	grid := Grid{
		vals: make(map[GridPos]string),
		maxX: len(rows[0]),
		maxY: len(rows),
	}
	for y, row := range rows {
		for x, col := range row {
			grid.vals[GridPos{x: x, y: y}] = string(col)
			if x > grid.maxX {
				grid.maxX = x
			}
		}
		if y > grid.maxY {
			grid.maxY = y
		}
	}
	//log.Printf("grid =\n%v", grid)
	return grid
}

func countCollisions(grid Grid, x int, y int) int {
	collisions := 0
	for checkX := 0; checkX < grid.maxX; checkX++ {
		for checkY := 0; checkY < grid.maxY; checkY++ {
			if checkX == x && checkY == y {
				continue
			}
			checkVal, _ := grid.ValAt(checkX, checkY)
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
					rayVal, inBounds := grid.ValAt(xRay, yRay)
					if !inBounds {
						break loop
					}
					if rayVal == "#" {
						//log.Printf("rise %v run %v intersect (%v,%v) <> (%v,%v) <> (%v,%v)", rise, run, x, y, xRay, yRay, checkX, checkY)
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

func Part1(filename string) int {
	g := ReadGrid(filename)
	collisions := make(map[GridPos]int)
	for x := 0; x < g.maxX; x++ {
		for y := 0; y < g.maxY; y++ {
			val, _ := g.ValAt(x, y)
			if val == "#" {
				collisions[GridPos{x: x, y: y}] = countCollisions(g, x, y)
			}
		}
	}
	var max int
	var maxX, maxY int
	for x := 0; x < g.maxX; x++ {
		//log.Print(row)
		for y := 0; y < g.maxY; y++ {
			v := collisions[GridPos{x: x, y: y}]
			if v > max {
				max = v
				maxX = x
				maxY = y
			}
		}
	}
	log.Printf("max = %v, maxX = %v, maxY = %v", max, maxX, maxY)
	return max
}

// RelAst is the relative position of an asteroid with respect to some other x,y point
type RelAst struct {
	pos GridPos
	distX int
	distY int
}

func (ra RelAst) angle() float64 {
	return math.Atan2(-float64(ra.distX), float64(ra.distY))
}

func (ra *RelAst) length() float64 {
	return math.Sqrt(float64(ra.distX*ra.distX+ra.distY*ra.distY))
}

func Part2(filename string, x, y int) {
	g := ReadGrid(filename)

	var ra []RelAst
	// find all the other asteroids on the board and calculate their positions
	// relative to the x,y input pair
	for checkX := 0; checkX < g.maxX; checkX++ {
		for checkY := 0; checkY < g.maxY; checkY++ {
			if checkX == x && checkY == y {
				continue
			}
			checkVal, _ := g.ValAt(checkX, checkY)
			if checkVal == "#" {
				r := RelAst{
					pos: GridPos{checkX, checkY},
					distX: checkX - x,
					distY: checkY - y,
				}
				ra = append(ra, r)
				//log.Printf("created ra %v", r)
			}
		}
	}

	// map of angles to other asteroids on the ray extending at that angle from x,y
	anglesToAsteroids := make(map[int64][]RelAst)

	for _, r := range ra {
		angle := int64(math.Round(r.angle()*1_000_000.0))
		anglesToAsteroids[angle] = append(anglesToAsteroids[angle], r)
	}

	// all of the known angles
	angles := make([]int64, 0)
	for k := range anglesToAsteroids {
		angles = append(angles, k)
	}
	seanmath.Sort(angles)
	for _, s := range anglesToAsteroids {
		sort.Slice(s, func(i, j int) bool {
			log.Printf("compare %v to %v", s[i].length(), s[j].length())
			return s[i].length() < s[j].length()
		})
	}



	blast := 0
	blastsThisPass := 1
	// each pass is a full 2pi rotation of the blaster ray.
	// keep doing passes until no asteroids remain.
	g.vals[GridPos{x:x,y:y}] = "S"
	for pass := 0; blastsThisPass > 0; pass++ {
		blastsThisPass = 0
		for _, a := range angles {
			ma, found := anglesToAsteroids[a]
			if !found {
				continue
			}
			if len(ma) <= pass {
				continue
			}
			blast++
			blastsThisPass++
			val := ma[pass]
			g.vals[val.pos] = "X"
			log.Printf("%v entry is %v, grid =\n%v", blast, val, g)
		}
	}
}

func main() {
	Part1("cmd/day10/input10-1.txt")
	Part2("cmd/day10/input10-1.txt", 11, 11)
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
	//log.Printf("gcd = %v", a)
	return a
}
