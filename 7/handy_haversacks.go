package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

const BAG_SPEC_SEPARATOR = " bags contain "
const CONTAINED_BAG_SEPARATOR = ", "
const NO_BAGS = "no other bags"

const TARGET = "shiny gold"

type BagQuantity struct {
	count int
	kind  string // this will be "adjective color"
}

// map from string "kind" to list of quantities
var bag_definitions = make(map[string][]BagQuantity)

func containsTarget(bag string, target string) bool {
	// myIndent := strings.Repeat(" ", indentLevel)

	for _, b := range bag_definitions[bag] {
		if b.count == 0 {
			// fmt.Printf("%sempty\n", myIndent);
			return false
		}

		if b.kind == target {
			// fmt.Printf("%s !!%s!!", myIndent, TARGET)
			return true
		}

		// fmt.Printf("%s %d %s:\n", myIndent, b.count, b.kind)
		if containsTarget(b.kind, target) {
			return true
		}
	}
	return false
}

func countWithin(bag string, indentLevel int) int {
	myIndent := strings.Repeat(" ", indentLevel)
	count := 0
	for _, b := range bag_definitions[bag] {
		if b.count > 0 {
			fmt.Printf("%s %d %s\n", myIndent, b.count, b.kind)
			count += b.count + b.count * countWithin(b.kind, indentLevel + 1)
		}
	}
	fmt.Printf("%s%s contains %d\n", myIndent, bag, count)
	return count
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)


	for scanner.Scan() {
		line := scanner.Text()

		// split each line
		x := strings.Split(line, BAG_SPEC_SEPARATOR)
		container_kind := x[0]

		// remove the final char ('.') and split on ", " to get the contained
		// bags
		contains := strings.Split(
						x[1][:len(x[1]) - 1],
						CONTAINED_BAG_SEPARATOR)

		// fmt.Println(line)
		for _, b := range contains {
			var quantity int
			var adjective string
			var color string

			var contained_bag BagQuantity;

			if b == NO_BAGS {
				contained_bag = BagQuantity{}
			} else {
				fmt.Sscanf(b, "%d %s %s", &quantity, &adjective, &color)
				kind := adjective + " " + color
				contained_bag = BagQuantity{quantity, kind}
			}

			bag_definitions[container_kind] = append(
						bag_definitions[container_kind],
						contained_bag)

		}

		// fmt.Printf(" %s: %q\n", container_kind, contains)
	}

	count_contains_target := 0

	for k, _ := range bag_definitions {
		// printBag(k, 0)
		if containsTarget(k, TARGET) {
			// fmt.Println(k, "contains", TARGET)
			count_contains_target += 1
		} else {
			// fmt.Println(k)
		}
		// fmt.Println("----")
	}

	fmt.Println(count_contains_target, "ways of containing", TARGET)
	fmt.Print("====\n\n")
	fmt.Println(countWithin(TARGET, 0), "bags inside of", TARGET)
}
