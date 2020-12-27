package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

// Preamble: 25 chars
const PREAMBLE_LEN = 25


// naive solution just using lists; generates a new list every time, i guess.
// hopefully go is smart enough to reuse slices, but they may have to get
// copied.
// perhaps a smarter data structure would somehow store, in a map, all sums that
// are possible with the current list? but that may not even be a good idea.
// updates would be very slow, and you couldn't exit early as you could when you
// are only looking for answers on the fly.

func slice_contains(s *[]int64, t int64) bool {
	for _, i := range *s {
		if t == i {
			return true
		}
	}
	return false
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	// don't love 1000 here (file is that long), but easiest for now
	var input_numbers = make([]int64, 1000)
	pos := 0;

	// read the preamble in first
	for pos = 0; pos < PREAMBLE_LEN; pos++ {
		scanner.Scan()
		// candidate_list[pos], _ = strconv.ParseInt(scanner.Text(), 10, 64)
		input_numbers[pos], _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}

	// fmt.Println("preamble:", candidate_list)

	// part 1 answer
	var first_invalid int64

	// read rest of numbers
	for ; scanner.Scan(); pos++ {
		n, _ := strconv.ParseInt(scanner.Text(), 10, 64);

		// fmt.Println(n)

		// find if sum in list
		sum_possible := false
		current_slice := input_numbers[pos - PREAMBLE_LEN:pos]
		// fmt.Println(current_slice)
		for _, a := range current_slice {
			b := n - a  
			if slice_contains(&current_slice, b) {
				fmt.Println(n, "=", a, "+", b)
				sum_possible = true
				break
			}
		}

		if ! sum_possible {
			fmt.Println(n, "invalid")
			first_invalid = n
			break
		}

		input_numbers[pos] = n

	}

	fmt.Println("looking for window sum", first_invalid, "starting at", pos)


	// now, need to find contiguous.
	// is there a better way than just scanning through?
	// reverse at least
	//
	// approach: some kind of window. grow the window down until greater than;
	// then remove from top until less than.

	window_start := pos - 1 // bottom
	window_end := pos - 1	// top
	window_sum := input_numbers[window_end]

	for {
		if window_sum == first_invalid {
			fmt.Println("found!")
			break
		} else if window_sum > first_invalid {
			window_sum -= input_numbers[window_end]
			window_end -= 1
		} else if window_sum < first_invalid {
			window_start -= 1
			window_sum += input_numbers[window_start]
		}

		fmt.Printf("NEW WINDOW %d:%d (len %d) = %d\n", window_start, window_end, window_end - window_start, window_sum)
	}

	window := input_numbers[window_start:window_end+1]
	fmt.Println("window is", window) 

	min := first_invalid
	var max int64 = 0
	for _, i := range window {
		if i < min {
			min = i
		} else if i > max {
			max = i
		}
	}
	fmt.Println("min", min, "max", max, "=", min + max)
}