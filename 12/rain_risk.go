package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)


// facing direction: north is 0; west negative, east positive
type Boat struct {
	facing int
	x, y int
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

func (b *Boat) translate(dx, dy int) {
	b.x += dx
	b.y += dy
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

func (b *Boat) manhattanDistance() int {
	dist := 0

	if b.x < 0 {
		dist += -b.x
	} else {
		dist += b.x
	}

	if b.y < 0 {
		dist += -b.y
	} else {
		dist += b.y
	}

	return dist
}


func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var boat Boat

	// start east
	boat.facing = 90

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		action := line[0]
		v, _ := strconv.ParseInt(line[1:], 10, 32)
		value := int(v)

		switch action {
		case 'N':
			boat.translate(0, value)
		case 'E':
			boat.translate(value, 0)
		case 'W':
			boat.translate(-value, 0)
		case 'S':
			boat.translate(0, -value)
		case 'L':
			boat.turn(-value)
		case 'R':
			boat.turn(value)
		case 'F':
			boat.moveForward(value)
		}
		fmt.Printf("%+v\n", boat)
	}
	fmt.Println("manhattan distance", boat.manhattanDistance())
}