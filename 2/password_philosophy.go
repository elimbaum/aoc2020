package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	// by count
	valid_part1 := 0

	// by index
	valid_part2 := 0

    for scanner.Scan() {
		var min int
		var max int
		var ch rune
		var password string

		line := scanner.Text()
		// i, _ := strconv.ParseInt(scanner.Text(), 10, 32)
		// input = append(input, int(i))
		fmt.Sscanf(line, "%d-%d %c: %s", &min, &max, &ch, &password)
		fmt.Println(line)
		// fmt.Printf("  %d %d %c %s\n", min, max, ch, password)

		// PART 1 - counts must between min/max
		if n := strings.Count(password, string(ch)); n >= min && n <= max {
			// fmt.Println("  OK")
			valid_part1++
		} else {
			// fmt.Printf("  BAD (%d %c)\n", n, ch)
		}

		// PART 2 - the numbers are actually index requirements
		ch_b := byte(ch)
		char_in_pos := 0
		if password[min - 1] == ch_b {
			char_in_pos++
		}
		if password[max - 1] == ch_b {
			char_in_pos++
		}

		if char_in_pos == 1 {
			valid_part2++
		}
    }

	fmt.Println("Part 1:", valid_part1, "valid passwords.")
	fmt.Println("Part 12", valid_part2, "valid passwords.")

}
