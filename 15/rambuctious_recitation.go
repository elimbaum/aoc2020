package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

const MAX int = 30000000

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println("---")
		line := scanner.Text()

		split := strings.Split(line, " ")
		answer := -1
		if len(split) == 2 {
			answer, _ = strconv.Atoi(split[1])
		}

		// map from n to turn last spoken
		seen := make(map[int]int)

		turn := 0
		last := -1
		for t, ns := range strings.Split(split[0], ",") {
			n, _ := strconv.Atoi(ns)
			turn = t + 1
			if last >= 0 {
				seen[last] = t
			}
			// fmt.Printf("[initial] turn %d said %d (last %d)\n",
			// 			turn, n, last)
			last = n
		}

		turn += 1

		for ; turn <= MAX; turn++ {
			next := 0
			last_turn, already_seen := seen[last]
			if already_seen {
				next = turn - last_turn - 1
				// fmt.Printf("turn %d said %d (last %d, saw on t%d)\n", turn, next, last, last_turn)
			} else {
				next = 0
				// fmt.Printf("turn %d said %d (last %d, was new)\n", turn, next, last)
			}

			seen[last] = turn - 1
			last = next
		}
		fmt.Println(last)
		fmt.Println("correct", answer)

		// fmt.Printf("%v %v\n", seen, answer)
	}
}