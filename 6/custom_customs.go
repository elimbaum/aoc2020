package main

import (
	"os"
	"fmt"
	"bufio"
)

func count_all(m map[rune]int, n int) int {
	pass := 0
	for _, cnt := range m {
		if cnt == n {
			pass += 1
		}
	}
	fmt.Println("all count", pass)
	return pass
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	group_answers_any := make(map[rune]bool)
	group_answers_all := make(map[rune]int)

	group_count_any := 0

	total_count_any := 0
	total_count_all := 0

	n := 0

	for scanner.Scan() {
		line := scanner.Text()

		// TODO: can end early if count == 26
		if line == "" {
			fmt.Println("any count", group_count_any)

			total_count_any += group_count_any
			total_count_all += count_all(group_answers_all, n)

			fmt.Println()

			// next group
			group_answers_any = make(map[rune]bool)
			group_answers_all = make(map[rune]int)

			group_count_any = 0

			n = 0

			continue
		}

		fmt.Println(line)
		n += 1

		for _, ch := range line {
			_, ok := group_answers_any[ch]
			if ! ok {
				// haven't seen yet
				group_count_any += 1
				group_answers_any[ch] = true
			} // else, duplicate; ignore
			group_answers_all[ch] += 1
		}
	}

	// get the last one
	total_count_any += group_count_any
	total_count_all += count_all(group_answers_all, n)

	fmt.Println("total any", total_count_any)
	fmt.Println("total all", total_count_all)
}
