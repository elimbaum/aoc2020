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
	// fmt.Println("PLUS")
	a := ns.pop().(int)
	b := ns.pop().(int)
	ns.push(a + b)
}


func (o Times) action(ns, os *Stack) {
	// fmt.Println("TIMES")
	a := ns.pop().(int)
	b := ns.pop().(int)
	ns.push(a * b)
}

func (o LeftParen) action(ns, os *Stack) {
}

func (o RightParen) action(ns, os *Stack) {
}


type Token interface {

}

type Stack []interface{}

func (s * Stack) pop() interface{} {
	r := (*s)[len(*s) - 1]
	*s = (*s)[:len(*s) - 1]
	return r
}

func (s * Stack) peek() interface{} {
	return (*s)[len(*s) - 1]
}

func (s * Stack) push(e interface{}) {
	*s = append(*s, e)
}

func (s * Stack) empty() bool {
	return len(*s) == 0
}

func (s * Stack) printT() {
	for _, e := range *s {
		k := ""
		switch e.(type) {
		case LeftParen:
			k = "("
		case RightParen:
			k = ")"
		case Plus:
			k = "+"
		case Times:
			k = "*"
		}
		fmt.Printf("%s ", k)
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
		// fmt.Printf("\nTOK %#v\n", t)
		switch v := t.(type) {
		// if we see a number...
		case Number:
			n_stack.push(v.v)

			// execute an operation
			if !op_stack.empty() {
				op := op_stack.pop().(Operator)
				_, isLeftP := op.(LeftParen)
				_, isTimes := op.(Times)
				if isLeftP || isTimes {
					// unless the last operation was a left bracket
					op_stack.push(op)
				} else {
					op.action(&n_stack, &op_stack)	
				}
			}
		case RightParen:
			// pop the left paren
			// op := op_stack.peek().(Operator)
			// op_stack.pop()

			// evaluate what remains inside parens
			for !op_stack.empty() {
				op := op_stack.pop().(Operator)
				_, isLeftP := op.(LeftParen)
				if isLeftP {
					break
				} else {
					// op = op_stack.pop().(Operator)
					op.action(&n_stack, &op_stack)
				}
			}
		default:
			op_stack.push(v)
		}
		fmt.Println("n:", n_stack)
		fmt.Print("   op: ")
		op_stack.printT()
	}

	fmt.Println("--")

	for !op_stack.empty() {
		op := op_stack.pop().(Operator)
		op.action(&n_stack, &op_stack)

		fmt.Println("n:", n_stack)
		fmt.Print("   op: ")
		op_stack.printT()
	}

	if len(n_stack) > 1 {
		panic(n_stack)	
	}
	

	return n_stack[0].(int)
}

func main() {
	file, _ := os.Open("advanced.txt")
	scanner := bufio.NewScanner(file)

	wrong := 0

	sum := 0

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

		sum += r

		if a != "" {
			if correct == r {
				fmt.Print(" (OK)")
			} else {
				fmt.Printf(" (XXX expected %d) ", correct)
				wrong++
			}
		}
		fmt.Println()
	}

	fmt.Println(wrong, "wrong")
	fmt.Println("sum:", sum)

}