
package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

const ADDR_LEN_BITS int = 36 

func get_addr_helper(addr, x_mask uint64, bit int) (ret []uint64) {
	// fmt.Println("helper called with", addr, x_mask, bit)
	if bit >= ADDR_LEN_BITS {
		return append(ret, addr)
	}

	if x_mask & (1 << bit) != 0 {
		// floating at this bit
		ret = append(ret,
				get_addr_helper(addr & ^(1 << bit),
							    x_mask, bit + 1)...)
		ret = append(ret,
				get_addr_helper(addr |  (1 << bit),
								x_mask, bit + 1)...)
	} else {
		ret = get_addr_helper(addr, x_mask, bit + 1)
	}
	return ret
}

func get_addr(addr, x_mask uint64) []uint64 {
	return get_addr_helper(addr, x_mask, 0)
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	memory1 := make(map[uint64]uint64)
	memory2 := make(map[uint64]uint64)
	
	one_mask := uint64(0)
	zero_mask := uint64(0)
	x_mask := uint64(0)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "mask = ") {
			new_mask := line[7:]
			fmt.Printf("\nmask = %s\n", new_mask)

			one_mask = 0
			zero_mask = 0
			x_mask = 0

			n_x := 0

			for _, m := range new_mask {
				zero_mask <<= 1
				one_mask <<= 1
				x_mask <<= 1

				if m == '0' {
					zero_mask |= 1
				} else if m == '1' {
					one_mask |= 1
				} else if m == 'X' {
					x_mask |= 1
					n_x += 1
				}
	
			}
			zero_mask = ^zero_mask
			fmt.Println("   ", n_x, "floating bits")
			// fmt.Printf("1s %36b\n", one_mask)
			// fmt.Printf("0s %36b\n", zero_mask)

		} else if strings.HasPrefix(line, "mem[") {
			close_bracket_idx := strings.Index(line, "]")
			addr_str := line[4 : close_bracket_idx]
			addr, _ := strconv.ParseInt(addr_str, 10, 64)

			val_str := line[close_bracket_idx + 4:]
			val, _ := strconv.ParseInt(val_str, 10, 64)

			fmt.Printf("mem[%d] = %d\n", addr, val)

			n_addr := 0
			for _, a := range get_addr(uint64(addr) | one_mask, x_mask) {
				// fmt.Printf("  %36b\n", a)
				memory2[a] = uint64(val)
				n_addr += 1
			}
			fmt.Println("   ", n_addr, "floating addresses")

			memory1[uint64(addr)] = (uint64(val) | one_mask) & zero_mask
		}
	}

	sum := uint64(0)
	for _, v := range memory1 {
		sum += v
	}
	fmt.Printf("part 1: sum %d\n", sum)

	sum = 0
	for _, v := range memory2 {
		sum += v
	}
	fmt.Printf("part 2: sum %d\n", sum)
}