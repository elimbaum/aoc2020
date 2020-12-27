package main

import (
	"fmt"
	"bufio"
	"os"
	"sort"
	"strconv"
)

/* rules:
 *   - wall is 0 jolts
 *	 - adapters can take up to 3 lower than output
 * 	 - phone takes (max adapter + 3) jolts
 *   - must use all adapters
 */

// plan: sort
func can_make(adapters map[int]bool, target int) bool {
	_, ok := adapters[target]
	return ok
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var adapters []int
	adapter_map := make(map[int]bool)

	adapter_map[0] = true

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		joltage, _ := strconv.ParseInt(scanner.Text(), 10, 64);
		adapters = append(adapters, int(joltage))
		adapter_map[int(joltage)] = true
	}

	sort.Ints(adapters)

	fmt.Println("adapter joltages:", adapters)

	diff_histogram := make([]int, 4)

	last := 0
	for _, j := range adapters {
		// fmt.Println(last, "to", j, "=", j - last)
		diff_histogram[j - last] += 1
		last = j
	}

	// to device
	diff_histogram[3] += 1

	fmt.Println(diff_histogram)

	// part 1
	fmt.Println(diff_histogram[1] * diff_histogram[3])


	// part 2: dynamic programming!
	// we have only jumps of 1 and 3.
	// WaysToMake(x) = WaysToMake(x - 1) + WaysToMake(x - 3)

	// plus 3 for device
	max_joltage := adapters[len(adapters) - 1] + 3 + 1

	config_counts := make([]uint64, max_joltage)

	// base cases
	config_counts[0] = 1

	for i, _ := range config_counts {
			fmt.Println("output joltage", i, "requires >=", i - 3)

			if can_make(adapter_map, i - 1) {
				p := config_counts[i - 1]
				config_counts[i] = p
				fmt.Println("  can make", i - 1, "   ", p)
			}

			if can_make(adapter_map, i - 2) {
				p := config_counts[i - 2]
				config_counts[i] += p
				fmt.Println("  can make", i - 2, "   ", p)
			}

			if can_make(adapter_map, i - 3) {
				p := config_counts[i - 3]
				config_counts[i] += p
				fmt.Println("  can make", i - 3, "   ", p)
			}

		fmt.Println("    output", i, "has", config_counts[i])
	}
}