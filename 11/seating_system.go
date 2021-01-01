package main

import (
	"fmt"
	"bufio"
	"os"
	"errors"
	// "strconv"
)

/* Rules
 * 
 * empty (L) and no occupied seats adjacent => occupied
 * occupied (#) and 4+ adjacent also occupied => empty
 *
 * include diagonals
 */

const (
	Floor 	 = '.'
	Empty 	 = 'L'
	Occupied = '#'
)

type Seats struct {
	state [][]rune
	width, height int
}

func (s *Seats) print() {
	for _, r := range(s.state) {
		fmt.Println(string(r))
	}
}

func (s *Seats) get(x, y int) rune {
	return s.state[y][x]
}

func (s *Seats) runRound() {
	newstate := make([][]rune, s.height)

	countF := s.countNeighborsLOS
	occupiedThresh := 5

	for y, r := range s.state {
		newstate[y] = make([]rune, s.width)
		for x := range r {
			if s.get(x, y) == Empty && countF(x, y) == 0 {
				newstate[y][x] = Occupied
			} else if s.get(x, y) == Occupied && countF(x, y) >= occupiedThresh {
				newstate[y][x] = Empty
			} else {
				newstate[y][x] = s.get(x, y)
			}
		}
	}

	// copy back
	s.state = newstate
}

// count neighbors of the seat at x, y
// returns -1 if x, y is floor
// PART 1
func (s *Seats) countNeighborsAdjacent(x, y int) int {
	if s.get(x, y) == Floor {
		return -1
	}

	c := 0
	for i := x - 1; i <= x + 1; i++ {
		for j := y - 1; j <= y + 1; j++ {
			// out of bounds
			if s.outOfBounds(i, j) {
				continue
			}
			// skip self
			if i == x && j == y {
				continue
			}

			if s.get(i, j) == Occupied {
				c++
			}
		}
	}

	return c
}

func (s *Seats) outOfBounds(x, y int) bool {
	return x < 0 || x >= s.width || y < 0 || y >= s.height
}

// check if there is an occupied seat in direction dx, dy starting at
// x, y
// true: occupied
// false: empty or floor
func (s *Seats) checkDirection(x, y, dx, dy int) bool {
	for {
		x += dx
		y += dy

		if s.outOfBounds(x, y) {
			return false
		} else if s.get(x, y) == Occupied {
			return true
		} else if s.get(x, y) == Empty {
			return false
		}


	}
}

// PART 2: just need to rewrite countNeighbors
// count neighbors not adjacent, but line of sight
func (s *Seats) countNeighborsLOS(x, y int) int {
	if s.get(x, y) == Floor {
		return -1
	}

	c := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			if s.checkDirection(x, y, dx, dy) {
				c += 1
			}
		}
	}
	return c
}

func (s *Seats) addRow(row []rune) error {
	if len(s.state) > 0 {
		if len(s.state[0]) != len(row) {
			return errors.New(
				fmt.Sprint("length mismatch: ", len(row),
						   " (expect ", len(s.state[0]), ")"))
		}
	} else {
		s.width = len(row)
	}
	s.state = append(s.state, row)
	s.height = len(s.state)

	return nil
}

func (s *Seats) countOccupied() int {
	occ := 0
	for _, row := range s.state {
		for _, seat := range row {
			if seat == Occupied {
				occ++
			}
		}
	}
	return occ
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var seats Seats

	for scanner.Scan() {
		line := scanner.Text()
		err := seats.addRow([]rune(line))

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// fmt.Printf("%q\n", seats.state)
	fmt.Println(seats.width, "x", seats.height)

	// for y, r := range seats.state {
	// 	var c string
	// 	for x, seat := range r {
	// 		fmt.Printf("%c", seat)
	// 		count := seats.countNeighbors(x, y)
	// 		if count < 0 {
	// 			c += " "
	// 		} else {
	// 			c += strconv.Itoa(count)
	// 		}
	// 	}
	// 	fmt.Println()
	// 	fmt.Println(c)
	// 	fmt.Println()
	// }

	lastCount := -1
	for {
		seats.runRound()
		// seats.print()
		count := seats.countOccupied()
		fmt.Println(count)

		if lastCount == count {
			break
		}
		lastCount = count
	}

	// seats.print()
	// seats.runRound()
	// fmt.Println()
	// seats.print()
	// seats.runRound()
	// fmt.Println()
	// seats.print()
}