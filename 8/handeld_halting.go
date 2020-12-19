package main

import (
    "os"
    "bufio"
    "fmt"
    "strings"
    "strconv"
)

type Instruction struct {
	op string
	arg int
	executed int // run number
}

// determines if a program loops (lol)
// returns
//   - bool: did it loop or terminate
//   - int:  accumulator value when loop detected, or terminated
func loops(program *[]*Instruction, runno int) (bool, int) {
	pc := 0
	acc := 0

	for pc < len(*program) {
		curr := (*program)[pc]
		// fmt.Printf("pc:%4d  acc:%5d  exec:%t\n", pc, acc, curr.executed)
		if curr.executed == runno {
			// fmt.Println("Loop!")
			// fmt.Println("pc =", pc, "acc =", acc)
			return true, acc
		}

		curr.executed = runno
		if curr.op == "jmp" {
			pc += curr.arg - 1
		} else if curr.op == "acc" {
			acc += curr.arg
		}

		pc += 1
	}
	return false, acc
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var program []*Instruction

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, " ")
		cmd := split[0]
		amt, _ := strconv.ParseInt(split[1], 10, 32)

		program = append(program, &Instruction{cmd, int(amt), -1})
	}

	for i, inst := range program {
		// inst.executed = false

		if inst.op == "acc" {
			continue
		}

		old := inst.op
		new := ""
		if inst.op == "jmp" {
			new = "nop"
		} else {
			new = "jmp"
		}

		inst.op = new
		does_loop, acc := loops(&program, i)
		inst.op = old

		if does_loop {
			fmt.Println("LOOP Instruction", i, *inst, "acc:", acc)
		} else {
			fmt.Println("TERM Instruction", i, *inst, "acc:", acc)
			break
		} 
	}
}