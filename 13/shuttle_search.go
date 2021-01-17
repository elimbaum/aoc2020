package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"eli"
)

// part 1: soonest we could leave after current time
// part 2: looking for minute-by-minute sequential bus departures;
//   taking `x` for any bus.

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	curr_t, _ := strconv.Atoi(scanner.Text())

	fmt.Println("It is", curr_t)
	
	scanner.Scan()
	schedule_str := scanner.Text()

	var schedule []int
	
	var nextbus_id int
	nextbus_wait := curr_t

	for _, ts := range strings.Split(schedule_str, ",") {
		if ts == "x" {
			continue
		}
		t, _ := strconv.Atoi(ts)
		wait := t - curr_t % t
		fmt.Println("bus", t, "leaves in", wait)

		if wait < nextbus_wait {
			nextbus_wait = wait
			nextbus_id = t
		}

		schedule = append(schedule, t)
	}
	fmt.Println("schedule", schedule)

	fmt.Println("bus", nextbus_id, "next; puzzle answer is", nextbus_id * nextbus_wait)


	// ideally would do this at the same time with the same file but i want to
	// try out the examples as well.
	fmt.Println("=== part 2")

	file, _ = os.Open("offset_input.txt")
	scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		schedule := scanner.Text()
		fmt.Println(schedule[:8] + "...")

		var ai []int64
		var mi []int64

		for i, ts := range strings.Split(schedule, ",") {
			if ts == "x" { continue }

			s, _ := strconv.Atoi(ts)

			// make RHS positive
			a := (s - i) % s
			for a < 0 {
				a += s
			}

			fmt.Printf("bus %d @ t+%d; t == %d (mod %d) \n", s, i, a, s)

			ai = append(ai, int64(a))
			mi = append(mi, int64(s))
		}

		// fmt.Println("a", ai)
		// fmt.Println("m", mi)
		fmt.Println(eli.ChineseRemainder(ai, mi))
		fmt.Println()
	}



}

/*

first timestamp, t
departure offset, i
schedule, s

(t + i) % s = 0, for multiple buses
t + i == 0 (mod s)
t == -i (mod s)
t == s - i (mod s)

q*s = t + i
qs - i = t

we get multiple [qs - i] which all must equal t. don't know q or t.

ok, this is just chinese remainder theorem.

1025845156820303 was too high.
possible I'm running into a overflow issue? moved to int64...

600691418730595 is correct. overflow.

*/