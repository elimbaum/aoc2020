package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)

const SEAT_ID_LEN = 10

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var max_id int = 0

	occupied := make(map[int]bool)

	// load up occupied list
	for i := 0; i < (1 << SEAT_ID_LEN); i++ {
		occupied[i] = true
	}

	for scanner.Scan() {
		line := scanner.Text()
		vertical := line[:7]
		horizontal := line[7:]

		// -1 means replace all
		vb := strings.Replace(vertical, "F", "0", -1)
		vb = strings.Replace(vb, "B", "1", -1)

		hb := strings.Replace(horizontal, "L", "0", -1)
		hb = strings.Replace(hb, "R", "1", -1)

		row, _ := strconv.ParseInt(vb, 2, 32)
		col, _ := strconv.ParseInt(hb, 2, 32)

		id := int(row * 8 + col)

		// fmt.Printf("%v%v = row %v, col %v (id %v)\n",
		// 			vertical, horizontal, row, col, id)

		if id > max_id {
			max_id = id
		}

		delete(occupied, id)
	}

	fmt.Println("max id:", max_id)

	empty_seats := make([]int, 1 << SEAT_ID_LEN)
	for k, _ := range occupied {
		empty_seats = append(empty_seats, k)
	}

	sort.Ints(empty_seats)

	// check backwards, since all the empty zeros are @ beginning
	for i := len(empty_seats) - 1; i > 1; i-- {
		if empty_seats[i] - empty_seats[i - 1] > 1 {
			fmt.Println("my seat:", empty_seats[i - 1])
			break
		}
	}
	// fmt.Println("unoccupied:", empty_seats)
}
