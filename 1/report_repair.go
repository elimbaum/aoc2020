package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

const TARGET = 2020

// return one of the pair
// (really this will be the second one)

// part one
func findDoubleTarget(input []int, target int) int {
	seen_numbers := make(map[int]bool)

	for _, v := range input {
		// fmt.Print("Checking ", curr, "... ")

		// have we seen TARGET - curr?
		_, ok := seen_numbers[TARGET - v]
		if ok {
			return v
		}
		// fmt.Println("haven't seen", TARGET - curr)
		seen_numbers[v] = true
	}
	return 0
}

// part two
func findTripleTarget(input []int, target int) (int, int, int) {
	// this maps: [a + b] to a, for all sums a + b
	seen_double_sums := make(map[int]int)

	for i := 1; i < len(input); i++ {
		x := input[i]
		for _, y := range input {

			v, ok := seen_double_sums[2020 - y]
			if ok {
				return v, y, 2020 - y - v
			}

			seen_double_sums[x + y] = x
		}
	}

	return 0, 0, 0
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	input := []int{}

    for scanner.Scan() {
		i, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		input = append(input, int(i))
    }

	// fmt.Println(input)

	a := findDoubleTarget(input, TARGET)
	b := TARGET - a
	fmt.Printf("%v * %v = %v\n", a, b, a*b)

	a, b, c := findTripleTarget(input, TARGET)
	fmt.Printf("%v * %v * %v = %v \n", a, b, c, a * b * c)
}
