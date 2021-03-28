package main

import (
	"fmt"
	"bufio"
	"os"
	// "strings"
	// "strconv"
)

type Coord struct {
	x, y int
}

type Board map[Coord]bool

var current Board
var next Board

const ALIVE_RUNE = '#'
const DEAD_RUNE = '.'

const ALIVE = true
const DEAD = false

func countNeighbors(c Coord) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			neighborCoord := Coord{c.x + dx, c.y + dy}
			alive, exists := current[neighborCoord]
			if alive {
				count++
			} else if !exists {
				// either not alive, or not in table
				current[neighborCoord] = DEAD
			}
		}
	}

	// if current cell alive return 1 less
	if current[c] {
		return count - 1
	} else {
		return count
	}
	
}

func nextState(alive bool, n int) bool {
	if n == 3 || (alive && n == 2) {
		return ALIVE
	}
	return DEAD
}

func (B Board) toGrid() {
	var NWcorner, SEcorner Coord

	first := true

	// compute the bounds of the active set
	for pos, state := range B {
		_ = state
		// if state == DEAD {
		// 	continue
		// }

		if first {
			NWcorner = pos
			SEcorner = pos
			first = false
			continue
		}

		if pos.x < NWcorner.x {
			NWcorner.x = pos.x
		}

		if pos.y < NWcorner.y {
			NWcorner.y = pos.y
		}

		if pos.x > SEcorner.x {
			SEcorner.x = pos.x
		}

		if pos.y > SEcorner.y {
			SEcorner.y = pos.y
		}
	}

	fmt.Println(NWcorner, SEcorner)

	// # = alive, . = dead but in active set
	// space = dead and inactive
	for y := NWcorner.y; y <= SEcorner.y; y++ {
		for x := NWcorner.x; x <= SEcorner.x; x++ {
			alive, exists := B[Coord{x, y}]
			if exists {
				if alive {
					// fmt.Printf("%c", ALIVE_RUNE)
					fmt.Print("█")
				} else {
					fmt.Printf("%c", DEAD_RUNE)
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	file, _ := os.Open("r_pent.txt")
	scanner := bufio.NewScanner(file)

	current = make(Board)
	next = make(Board)

	y := 0

	loadCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		
		fmt.Println(line)

		for x, c := range line {
			if c == ALIVE_RUNE {
				// fmt.Printf("(%d, %d) ", x, y)
				loadCount++
				pos := Coord{x, y}
				current[pos] = true
				countNeighbors(pos)
			}
		}
		// fmt.Println()

		y++
	}

	fmt.Println("loaded", loadCount, "cells")
	current.toGrid()

	for i := 0; i < 100; i++ {
		// check alive first; add dead neighbors
		for cell, state := range current {
			// skip dead for now
			if state == DEAD {
				continue
			}
			curr := "DEAD"
			if state {
				curr = "ALIVE"
			}

			neighbors := countNeighbors(cell)

			fate := "DEAD"
			if nextState(state, neighbors) {
				fate = "ALIVE"
				next[cell] = ALIVE
			} else {
				// next[cell] = DEAD
			}
			_ = fate
			_ = curr
			// fmt.Println("ALIVE", cell, "has", neighbors, "=>", fate)
		}

		current.toGrid()

		// now do dead only
		for cell, state := range current {
			if state == ALIVE {
				continue
			}

			neighbors := countNeighbors(cell)
			fate := "DEAD"
			if nextState(state, neighbors) {
				next[cell] = ALIVE
				fate = "ALIVE"
			}
			_ = fate
			// fmt.Println("DEAD", cell, "has", neighbors, "=>", fate)

		}


		// fmt.Println("next:", next)
		// next.toGrid()
		current = next
		next = make(Board)
		fmt.Println("\n======\n")
	}
}