package main

import (
	"fmt"
)

func gen_num(n int) func(...func(int) int) int {
	return func(funcs ...func(int) int) int {
		if len(funcs) == 0 {
			return n
		}
		if len(funcs) > 1 {
			panic("IDK what to do with more than one arg")
		}
		return funcs[0](n)
	}
}

func plus(right_op int) func(int) int {
	return func(left_op int) int { return left_op + right_op }
}

func minus(right_op int) func(int) int {
	return func(left_op int) int { return left_op - right_op }
}

func mul(right_op int) func(int) int {
	return func(left_op int) int { return left_op * right_op }
}

func divide_by(right_op int) func(int) int {
	return func(left_op int) int { return left_op / right_op }
}

var (
	zero  = gen_num(0)
	one   = gen_num(1)
	two   = gen_num(2)
	three = gen_num(3)
	four  = gen_num(4)
	five  = gen_num(5)
	six   = gen_num(6)
	seven = gen_num(7)
	eight = gen_num(8)
	nine  = gen_num(9)
)

func main() {
	fmt.Println(
		five(divide_by(two())),
		seven(plus(six())),
	)
}

