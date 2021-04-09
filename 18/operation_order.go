package main

import (
	"fmt"
	"bufio"
	"os"
	"unicode"
	"strconv"
	"strings"
)

var op_map = map[rune]Operator {
	'+': Plus{},
	'*': Times{},
	'(': LeftParen{},
	')': RightParen{},
}

type Number struct {
	v int
}

type Operator interface {
	action(ns, os *Stack)
}

type Plus struct {}
type Times struct {}
type LeftParen struct {}
type RightParen struct {}

func (o Plus) action(ns, os *Stack) {
	a := ns.pop().(int)
	b := ns.pop().(int)
	ns.push(a + b)
}


func (o Times) action(ns, os *Stack) {
	a := ns.pop().(int)
	b := ns.pop().(int)
	ns.push(a * b)
}

func (o LeftParen) action(ns, os *Stack) {
	fmt.Println("LEFT", *ns)
	// os.push(o)
}

func (o RightParen) action(ns, os *Stack) {
	fmt.Println("RIGHT", *ns)
	// os.push(o)
}


type Token interface {

}

type Stack []interface{}

func (s * Stack) pop() interface{} {
	r := (*s)[len(*s) - 1]
	*s = (*s)[:len(*s) - 1]
	return r
}

func (s * Stack) push(e interface{}) {
	*s = append(*s, e)
}

func (s * Stack) empty() bool {
	return len(*s) == 0
}

func (s * Stack) printT() {
	for _, e := range *s {
		fmt.Printf("%T ", e)
	}

	fmt.Println()
}

func tokenize(eq string) []Token {
	var toks []Token

	var curr_tok string

	for _, r := range eq {
		if unicode.IsDigit(r) {
			curr_tok += string(r)
		} else {
			if curr_tok != "" {
				n, _ := strconv.Atoi(curr_tok)
				toks = append(toks, Number{n})
				curr_tok = ""
			}

			if unicode.IsSpace(r) {
				continue
			}

			toks = append(toks, op_map[r])
		}
	}

	// may still have int to process
	if curr_tok != "" {
		n, _ := strconv.Atoi(curr_tok)
		toks = append(toks, Number{n})
	}

	return toks
}

func compute(tok []Token) int {
	var n_stack Stack
	var op_stack Stack

	for _, t := range tok {
		switch v := t.(type) {
		case Number:
			// fmt.Print(v.v, " ")
			n_stack.push(v.v)

			if !op_stack.empty() {
				op := op_stack.pop().(Operator)
				op.action(&n_stack, &op_stack)
			}
		default:
			_, isRightP := v.(RightParen)
			if isRightP {
				var q Operator
				for ! op_stack.empty(){
					q = op_stack.pop().(Operator)
					q.action(&n_stack, &op_stack)
				}
			} else {
				op_stack.push(v)
			}
		}
		fmt.Println("n:", n_stack)
		fmt.Print("   op: ")
		op_stack.printT()
	}

	for !op_stack.empty() {
		op := op_stack.pop().(Operator)
		op.action(&n_stack, &op_stack)
	}

	return n_stack[0].(int)
}

func main() {
	file, _ := os.Open("more.txt")
	scanner := bufio.NewScanner(file)

	wrong := 0

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)

		line_split := strings.Split(line, "=")

		var tok []Token
		var q, a string
		var correct int

		if len(line_split) > 1 {
			q, a = line_split[0], line_split[1]
			correct, _ = strconv.Atoi(strings.TrimSpace(a))
			tok = tokenize(q)
		} else {
			tok = tokenize(line)
		}

		r := compute(tok)

		fmt.Print("= ", r)

		if a != "" {
			if correct == r {
				fmt.Print(" (OK)")
			} else {
				fmt.Printf(" (XXX expected %d) ", correct)
				wrong++
			}
		}

		fmt.Println("\n\n")

		// for _, k := range tok {
		// 	fmt.Printf("%T %+v\n", k, k)
		// }
	}

	fmt.Println(wrong, "wrong")

}