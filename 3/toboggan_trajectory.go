package main

import (
	"fmt"
	"bufio"
	"os"
	// "strings"
)

const MOVE_PER_ROW = 3

var tree_map []string
var line_length int

// returns number of trees the slope (right, down) would hit
func check_trees(right, down int) int {
	x := 0
	trees_hit := 0
	for y, L := range tree_map {
		if y % down > 0 {
			continue
		}

		if L[x] == '#' {
			trees_hit++
		}

		x += right
		x %= line_length


	}

	return trees_hit
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

    for scanner.Scan() {
		tree_map = append(tree_map, scanner.Text())
	}
	line_length = len(tree_map[0])
	fmt.Println("Line", line_length)

	slope_trees := []int{
			check_trees(1, 1),
			check_trees(3, 1),
			check_trees(5, 1),
			check_trees(7, 1),
			check_trees(1, 2),
	}

	fmt.Println(slope_trees)

	prod := 1
	for _, v := range slope_trees {
		prod *= v
	}
	fmt.Println(prod)

}
