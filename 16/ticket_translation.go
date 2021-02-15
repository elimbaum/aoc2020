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

var fields map[string][]Range

func parseTicket(ts string) []int {
	split := strings.Split(ts, ",")
	n := len(split)

	ret := make([]int, n)
	for i, s := range split {
		ret[i], _ = strconv.Atoi(s)
	}
	return ret
}

// check if this ticket can _ever_ be valid.
func checkTicket(ticket []int) (bool, int) {
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

// check if value v could be valid for field f.
func could_be_valid(v int, f string) bool {
	for _, r := range fields[f] {
		if v >= r.min && v <= r.max {
			return true
		}
	}
	return false
}

/*
func init_field_map(possible * []map[string]bool) {
	for i := 0; i < len(*possible); i++ {
		(*possible)[i] = make(map[string]bool)
		for f, _ := range fields {
			(*possible)[i][f] = true
		}
	}
}*/

func init_field_map(possible * map[string]map[int]bool) {
	for F, _ := range fields {
		// build set of possible indexes for this field
		all := make(map[int]bool)
		for i := 0; i < len(fields); i++ {
			all[i] = true
		}

		(*possible)[F] = all
	}
}

func map_pprint(M *map[string]map[int]bool) {
	for k, v := range *M {
		fmt.Println(k)
		fmt.Println("  ", v)
	}
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	// Read fields
	fields = make(map[string][]Range)
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

	// possible_fields := make([]map[string]bool, len(fields))
	// init_field_map(&possible_fields)
	possible_idx := make(map[string]map[int]bool)
	init_field_map(&possible_idx)

	// Read nearby tickets
	for scanner.Scan() {
		t := parseTicket(scanner.Text())
		// fmt.Println(t)
	 	valid, n := checkTicket(t)
	 	if ! valid {
	 		// ticket cannot be valid. increment error rate
			ticket_scanning_error += n
	 	} else {
	 		// ticket is valid. adjust ranges.
	 		// each index should be a map of which field it could be.

	 		// loop through each field and check ticket position
	 		for F, _ := range fields {
	 			for i, v := range t {
	 				if ! could_be_valid(v, F) {
	 					delete(possible_idx[F], i)
	 				}
	 			}
	 		}
	 		// for i, fm := range possible_fields {
	 		// 	// loop over each field
	 		// 	for f, b := range fm {
				//  	if b && !could_be_valid(t[i], f) {
				//  		// fmt.Println("pos", i, f, "invalid:", t[i])
				//  		fm[f] = false
				//  	}
	 		// 	}
	 		// }
	 	}
	}

	// 18227
	fmt.Println("TSER", ticket_scanning_error)


	// build the inverse map
	// field name -> ticket position

	// map of field name to actual ticket position
	definite_fields := make(map[string]int)
	departure_prod := 1

	for {
		// map_pprint(&possible_idx)
		for F, idx_map := range possible_idx {
			if len(idx_map) == 1 {
				// got it!
				var pos int

				for k, _ := range idx_map {
					pos = k
					break
				}

				fmt.Println("== pos", pos, "must be", F)
				definite_fields[F] = pos 

				if strings.HasPrefix(F, "departure ") {
					departure_prod *= myticket[pos]
				}

				// remove pos from all others
				for _, idx_map2 := range possible_idx {
					delete(idx_map2, pos)
				}

				// don't look at this guy again
				delete(possible_idx, F)
			}

			continue
		}

		if len(possible_idx) == 0 {
			break
		}
	}

	// 2355350878831
	fmt.Println("DP", departure_prod)
	
}

/*

Want to store, for each ticket index, which fields it could still be.
init with all fields in each index; remove them as it happens

for each ticket keep track if any valid

*/