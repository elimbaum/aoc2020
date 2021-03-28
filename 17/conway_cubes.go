package main

import (
	"fmt"
	"bufio"
	"os"
	// "time"
)

type Coord struct {
	x, y, z, w int
}

type Board map[Coord]bool

var current Board
var next Board

const ALIVE_RUNE = '#'
const DEAD_RUNE = '.'

const ALIVE = true
const DEAD = false

const NUM_GEN = 6

// 4d versions
func countNeighbors(c Coord) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					neighborCoord := Coord{c.x + dx, c.y + dy, c.z + dz, c.w + dw}
					alive, exists := current[neighborCoord]
					if alive {
						count++
					} else if !exists {
						// either not alive, or not in table
						current[neighborCoord] = DEAD
					}
				}
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


//// 2D versions (original game of life)
// func countNeighbors(c Coord) int {
// 	count := 0
// 	for dx := -1; dx <= 1; dx++ {
// 		for dy := -1; dy <= 1; dy++ {
// 			neighborCoord := Coord{c.x + dx, c.y + dy}
// 			alive, exists := current[neighborCoord]
// 			if alive {
// 				count++
// 			} else if !exists {
// 				// either not alive, or not in table
// 				current[neighborCoord] = DEAD
// 			}
// 		}
// 	}

// 	// if current cell alive return 1 less
// 	if current[c] {
// 		return count - 1
// 	} else {
// 		return count
// 	}
	
// }

// func (B Board) toGrid() {
// 	var NWcorner, SEcorner Coord

// 	first := true

// 	// compute the bounds of the active set
// 	for pos, state := range B {
// 		_ = state
// 		// if state == DEAD {
// 		// 	continue
// 		// }

// 		if first {
// 			NWcorner = pos
// 			SEcorner = pos
// 			first = false
// 			continue
// 		}

// 		if pos.x < NWcorner.x {
// 			NWcorner.x = pos.x
// 		}

// 		if pos.y < NWcorner.y {
// 			NWcorner.y = pos.y
// 		}

// 		if pos.x > SEcorner.x {
// 			SEcorner.x = pos.x
// 		}

// 		if pos.y > SEcorner.y {
// 			SEcorner.y = pos.y
// 		}
// 	}

// 	fmt.Println("====")
// 	fmt.Println(NWcorner, SEcorner)

// 	// # = alive, . = dead but in active set
// 	// space = dead and inactive
// 	for y := NWcorner.y; y <= SEcorner.y; y++ {
// 		for x := NWcorner.x; x <= SEcorner.x; x++ {
// 			alive, exists := B[Coord{x, y}]
// 			if exists {
// 				if alive {
// 					// fmt.Printf("%c", ALIVE_RUNE)
// 					fmt.Print("\033[107m \033[49m")
// 				} else {
// 					fmt.Print("\033[48;5;236m \033[49m")
// 					// fmt.Printf("%c", DEAD_RUNE)
// 				}
// 			} else {
// 				fmt.Print(" ")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

func (B * Board) countAlive() int {
	n := 0
	for _, state := range *B {
		if state == ALIVE {
			n++
		}
	}
	return n
}

func main() {
	file, _ := os.Open("input.txt")
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
				pos := Coord{x, y, 0, 0}
				current[pos] = true
				countNeighbors(pos)
			}
		}
		// fmt.Println()

		y++
	}

	fmt.Println("loaded", loadCount, "cells")
	// current.toGrid()

	for gen := 0; gen < NUM_GEN; gen++ {
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

		// current.toGrid()

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

		fmt.Printf("After gen %d: %d alive\n", gen, next.countAlive())


		// fmt.Println("next:", next)
		// next.toGrid()
		current = next
		next = make(Board)
		// fmt.Println("\n======\n")

		// time.Sleep(100 * time.Millisecond)
	}
}