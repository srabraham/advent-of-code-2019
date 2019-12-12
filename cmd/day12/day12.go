package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input string, numSteps int64) int {
	bodies, _ := ReadBodies(input)

	log.Printf("starting simulation")

	var finalEnergy int
	for step := int64(1); step <= numSteps; step++ {
		updateVelocitiesFromPositions(bodies)
		updatePositionsFromVelocities(bodies)
		var energy int
		for _, b := range bodies {
			energy += b.Energy()
		}
		finalEnergy = energy
	}
	return finalEnergy
}

func Part2(input string) int64 {
	bodies, startPos := ReadBodies(input)

	log.Printf("starting simulation")
	var xPeriod, yPeriod, zPeriod int64
	for step := int64(1); true; step++ {
		updateVelocitiesFromPositions(bodies)
		updatePositionsFromVelocities(bodies)

		// for each of x,y,z, figure out if the all of the bodies have the same
		// position as the start position. These are independently periodic, and
		// due to conservation of position and momentum, when all three x's are
		// back at the same position as the start, they will have velocity 0 again.
		sameXAsStart, sameYAsStart, sameZAsStart := samePositionsAsStart(bodies, startPos)
		if xPeriod == 0 && sameXAsStart {
			// The bodies will have the same position for two steps at this point.
			// It's the second step we want, when velocity is 0.
			xPeriod = step + 1
			log.Printf("found x period = %v", step)
		}
		if yPeriod == 0 && sameYAsStart {
			yPeriod = step + 1
			log.Printf("found y period = %v", step)
		}
		if zPeriod == 0 && sameZAsStart {
			zPeriod = step + 1
			log.Printf("found x period = %v", step)
		}
		if xPeriod > 0 && yPeriod > 0 && zPeriod > 0 {
			return LCM(xPeriod, yPeriod, zPeriod)
		}
	}
	log.Fatal("somehow escaped infinite loop!")
	return -1
}

func samePositionsAsStart(bodies []*Body, startPos []*Body) (bool, bool, bool) {
	startXPos := true
	startYPos := true
	startZPos := true
	for b := range bodies {
		if bodies[b].pos.x != startPos[b].pos.x {
			startXPos = false
		}
		if bodies[b].pos.y != startPos[b].pos.y {
			startYPos = false
		}
		if bodies[b].pos.z != startPos[b].pos.z {
			startZPos = false
		}
	}
	return startXPos, startYPos, startZPos
}

func main() {
	input := `<x=5, y=4, z=4>
<x=-11, y=-11, z=-3>
<x=0, y=7, z=0>
<x=-13, y=2, z=10>`

	log.Printf("part 1 answer is %v energy units", Part1(input, 1000))
	log.Printf("part 2 answer is %v", Part2(input))
}

func updateVelocitiesFromPositions(bodies []*Body) {
	for _, bod1 := range bodies {
		for _, bod2 := range bodies {
			if bod1.pos.x < bod2.pos.x {
				bod1.vel.x++
			}
			if bod1.pos.x > bod2.pos.x {
				bod1.vel.x--
			}
			if bod1.pos.y < bod2.pos.y {
				bod1.vel.y++
			}
			if bod1.pos.y > bod2.pos.y {
				bod1.vel.y--
			}
			if bod1.pos.z < bod2.pos.z {
				bod1.vel.z++
			}
			if bod1.pos.z > bod2.pos.z {
				bod1.vel.z--
			}
		}
	}
}

func updatePositionsFromVelocities(bodies []*Body) {
	for _, bod := range bodies {
		bod.pos.x += bod.vel.x
		bod.pos.y += bod.vel.y
		bod.pos.z += bod.vel.z
	}
}

type Triple struct {
	x int
	y int
	z int
}

type Body struct {
	pos Triple
	vel Triple
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (b Body) Energy() int {
	return (abs(b.pos.x) + abs(b.pos.y) + abs(b.pos.z)) * (abs(b.vel.x) + abs(b.vel.y) + abs(b.vel.z))
}

func (t Triple) String() string {
	return fmt.Sprintf("<x= %v, y= %v, z= %v>", t.x, t.y, t.z)
}

func (b Body) String() string {
	return fmt.Sprintf("pos=%v, vel=%v", b.pos, b.vel)
}

func MustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	f(err)
	return i
}

func f(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadBodies(input string) ([]*Body, []*Body) {
	var bodiesCopy1 []*Body
	var bodiesCopy2 []*Body
	for _, s := range strings.Split(input, "\n") {
		bodiesCopy1 = append(bodiesCopy1, ReadBodyFromString(s))
		bodiesCopy2 = append(bodiesCopy2, ReadBodyFromString(s))
	}
	return bodiesCopy1, bodiesCopy2
}

func ReadBodyFromString(s string) *Body {
	re := regexp.MustCompile("<x=([0-9-]+), y=([0-9-]+), z=([0-9-]+)>")
	matches := re.FindStringSubmatch(s)
	if len(matches) != 4 {
		log.Fatal("bad string")
	}
	return &Body{pos: Triple{x: MustParseInt(matches[1]), y: MustParseInt(matches[2]), z: MustParseInt(matches[3])}}
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int64) int64 {
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

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
