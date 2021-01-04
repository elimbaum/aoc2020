package main

// part 1: compass directions are translations
// part 2: are waypoint commands

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)


type Point struct {
	x, y int
}

func (p *Point) translateXY(dx, dy int) {
	p.x += dx
	p.y += dy
}

func (p *Point) translatePoint(q Point, multiplier int) {
	p.x += multiplier * q.x
	p.y += multiplier * q.y
}

func (p *Point) rotate(angle int) {
	// canonicalize angle
	if angle > 180 {
		angle -= 360
	} else if angle <= -180 {
		angle += 360
	}

	// rotate the point `angle` degrees about the origin
	switch angle {
	case 0:
		break
	case -90:
		p.x, p.y = -p.y, p.x
	case 180:
		p.x, p.y = -p.x, -p.y
	case 90:
		p.x, p.y = p.y, -p.x
	default:
		fmt.Println("bad angle", angle)
	}
}

// facing direction: north is 0; west negative, east positive
type Boat struct {
	facing int
	location Point
	waypoint Point
}

// enforce angle -179 to +180
func (b *Boat) turn(angle int) {
	b.facing += angle
	
	if b.facing > 180 {
		b.facing -= 360
	} else if b.facing <= -180 {
		b.facing += 360
	}
}

// translate the boat itself (part 1)
func (b *Boat) translate(dx, dy int) {
	b.location.translateXY(dx, dy)
}

// translate the waypoint
func (b *Boat) translateWaypoint(dx, dy int) {
	b.waypoint.translateXY(dx, dy)
}

func (b *Boat) moveForward(amt int) {
	switch b.facing {
	case 0:
		b.translate(0, amt)
	case 90:
		b.translate(amt, 0)
	case 180:
		b.translate(0, -amt)
	case -90:
		b.translate(-amt, 0)
	default:
		fmt.Println("bad angle", b.facing)
	}
}

func (b *Boat) moveTowardsWaypoint(amt int) {
	b.location.translatePoint(b.waypoint, amt)
}

func (b *Boat) manhattanDistance() int {
	dist := 0

	if b.location.x < 0 {
		dist += -b.location.x
	} else {
		dist += b.location.x
	}

	if b.location.y < 0 {
		dist += -b.location.y
	} else {
		dist += b.location.y
	}

	return dist
}


func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	// part 1 & part 2
	var boat1, boat2 Boat

	// start east
	boat1.facing = 90
	boat2.facing = 90

	// default 10 E,  1 N
	boat2.waypoint = Point{10, 1}

	for scanner.Scan() {
		line := scanner.Text()

		action := line[0]
		v, _ := strconv.ParseInt(line[1:], 10, 32)
		value := int(v)

		switch action {
		case 'N':
			boat1.translate(0, value)
			boat2.waypoint.translateXY(0, value)
		case 'E':
			boat1.translate(value, 0)
			boat2.waypoint.translateXY(value, 0)
		case 'W':
			boat1.translate(-value, 0)
			boat2.waypoint.translateXY(-value, 0)
		case 'S':
			boat1.translate(0, -value)
			boat2.waypoint.translateXY(0, -value)
		case 'L':
			boat1.turn(-value)
			boat2.waypoint.rotate(-value)
		case 'R':
			boat1.turn(value)
			boat2.waypoint.rotate(value)
		case 'F':
			boat1.moveForward(value)
			boat2.moveTowardsWaypoint(value)
		}
		fmt.Printf("%-4v:  1 %v\n", line, boat1)
		fmt.Printf("       2 %v\n", boat2)
	}
	fmt.Println("1 manhattan distance:", boat1.manhattanDistance())
	fmt.Println("2 manhattan distance:", boat2.manhattanDistance())
}