package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

// cid not required
var required_fields = []string { "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid" }


// make an empty map
func build_map() map[string]string {
	m := make(map[string]string)
	for _, f := range required_fields {
		m[f] = ""
	}
	return m
}

// we remove each field we find; if there are none left, this is valid
func checkPassport(fields map[string]string) bool {
	if len(fields) == 0 {
		return true
	} else {
		return false
	}
}

var valid_eyes = map[string]bool {
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

func checkFieldValidity(f, v string) bool {
	switch f {

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	case "byr":
		yr, _ := strconv.Atoi(v)
		return yr >= 1920 && yr <= 2002
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	case "iyr":
		yr, _ := strconv.Atoi(v)
		return yr >= 2010 && yr <= 2020
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	case "eyr":
		yr, _ := strconv.Atoi(v)
		return yr >= 2020 && yr <= 2030

	// hgt (Height) - a number followed by either cm or in:
	// 	 If cm, the number must be at least 150 and at most 193.
	// 	 If in, the number must be at least 59 and at most 76.
	case "hgt":
		unit := v[len(v)-2:]
		height, _ := strconv.Atoi(v[:len(v)-2])
		if unit == "cm" {
			return height >= 150 && height <= 193
		} else if unit == "in" {
			return height >= 59 && height <= 76
		} else {
			return false // bad unit
		}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	case "hcl":
		if v[0] != '#' {
			return false
		}

		for _, r := range v[1:] {
			if ! ((r >= '0' && r <= '9') ||
				  (r >= 'a' && r <= 'f') ||
				  (r >= 'A' && r <= 'F')) {
				  return false
			  }
		}
		return true

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	case "ecl":
		_, ok := valid_eyes[v]
		return ok

	// ignored
	case "cid":
		return true

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	case "pid":
		if len(v) != 9 {
			return false
		}

		_, err := strconv.Atoi(v)
		return err == nil

	default:
		fmt.Println(f, "not a valid field?")
		return false
	}
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	// fields only
	part1_passport_fields := build_map()
	part1_all_present := 0

	// fields and values
	part2_passport_fields := build_map()
	part2_all_valid := 0

    for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// next passport!
			fmt.Println("---")

			if checkPassport(part1_passport_fields) {
				fmt.Println("[1] OK")
				part1_all_present += 1

				if checkPassport(part2_passport_fields) {
					fmt.Println("[2] OK")
					part2_all_valid += 1
				} else {
					fmt.Println("[2] Bad:", part2_passport_fields)
				}
			} else {
				fmt.Println("[1] Missing:", part1_passport_fields)
			}

			part1_passport_fields = build_map()
			part2_passport_fields = build_map()
			fmt.Println()
			continue
		} else {
			fmt.Println(">", line)
			for _, x := range strings.Split(line, " ") {
				f := x[:3]
				delete(part1_passport_fields, f)

				v := x[4:]
				if checkFieldValidity(f, v) {
					delete(part2_passport_fields, f)
				} else {
					part2_passport_fields[f] = v
				}
				// fmt.Println(f)
			}
		}
	}
	fmt.Println("---")

	// gotta check the last passport
	if checkPassport(part1_passport_fields) {
		fmt.Println("[1] OK")
		part1_all_present += 1

		if checkPassport(part2_passport_fields) {
			fmt.Println("[2] OK")
			part2_all_valid += 1
		} else {
			fmt.Println("[2] Bad:", part2_passport_fields)
		}
	} else {
		fmt.Println("[1] Missing:", part1_passport_fields)
	}



	fmt.Println(part1_all_present, "passports have all fields")
	fmt.Println(part2_all_valid, "passports have all fields")
}
