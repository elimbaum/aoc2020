package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type Range struct {
	min, max int
}

func parseTicket(ts string) []int {
	split := strings.Split(ts, ",")
	n := len(split)

	ret := make([]int, n)
	for i, s := range split {
		ret[i], _ = strconv.Atoi(s)
	}
	return ret
}

func checkTicket(ticket []int,
				 fields map[string][]Range) (bool, int) {
	for _, v := range ticket {
		valid := false
		for _, valid_ranges := range fields {
			for _, r := range valid_ranges {
				if v >= r.min && v <= r.max {
					valid = true
					break
				}
			}
		}
		if ! valid {
			return false, v
		}
	}
	return true, 0
}

func init_field_map(possible * []map[string]bool) {
	for i := 0; i < len(*possible); i++ {
		(*possible)[i] = make(map[string]bool)
	}
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	// Read fields
	fields := make(map[string][]Range)
	for {
		scanner.Scan()
		line := scanner.Text()

		if line == "" {
			// done
			break
		}

		split := strings.Split(line, ": ")
		f, vals := split[0], split[1]

		ranges := strings.Split(vals, " or ")

		for _, r := range ranges {
			n := strings.Split(r, "-")
			min, _ := strconv.Atoi(n[0])
			max, _ := strconv.Atoi(n[1])
			fields[f] = append(fields[f], Range{min, max})
		}

		fmt.Println(f, fields[f])
	}

	// Read my ticket
	scanner.Scan() // header
	scanner.Scan() // my ticket
	myticket := parseTicket(scanner.Text())

	_ = myticket
	// fmt.Println(myticket)

	// valid, n := checkTicket(myticket, fields)

	scanner.Scan() // blank line
	scanner.Scan() // header

	ticket_scanning_error := 0

	possible_fields := make([]map[string]bool, len(fields))
	init_field_map(&possible_fields)

	// Read nearby tickets
	for scanner.Scan() {
		t := parseTicket(scanner.Text())
		// fmt.Println(t)
	 	valid, n := checkTicket(t, fields)
	 	if ! valid {
	 		// ticket is not valid. increment error rate
			ticket_scanning_error += n
	 	} else {
	 		// ticket is valid. adjust ranges.
	 		// each index should be a map of which field it could be.

	 		// loop through each ticket position and check
	 		for i, fm := range possible_fields {
	 			// loop over each field
	 			for f, _ := range fm {
				 	if could_be_valid(t[i], f, &fields) {

				 	} else {
				 		fm[f] = false
				 	}
	 			}
	 		}
	 	}
	}

	fmt.Println("TSER", ticket_scanning_error)

	departure_prod := 1
}

/*

Want to store, for each ticket index, which fields it could still be.
init with all fields in each index; remove them as it happens

for each ticket keep track if any valid

*/