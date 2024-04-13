package lispy

import (
	"fmt"
	"strconv"
	"strings"
)

type AST struct {
	Tag      string
	Contents string
	Children []AST
}

type LispValue interface {
	Print()
}

type LVNumber struct {
	number int
}

func (number LVNumber) Print() {
	fmt.Printf("lisp value number", number.number)
}

func (ast AST) Print(depth int) {
	if ast.Contents != "" {
		fmt.Printf("%*s '%s'\n", depth*2+len(ast.Tag), ast.Tag, ast.Contents)
	} else {
		fmt.Printf("%*s\n", depth*2+len(ast.Tag), ast.Tag)
	}

	if len(ast.Children) > 0 {
		for _, child := range ast.Children {
			child.Print(depth + 1)
		}
	}
}

func eval(x int, op string, y int) int {
	switch op {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	default:
		return 0
	}
}

func (ast AST) Eval() int {
	if strings.Contains(ast.Tag, "number") {
		if i, err := strconv.Atoi(ast.Contents); err == nil {
			return i
		}
	}

	op := ast.Children[1].Contents
	x := ast.Children[2].Eval()

	for i := 3; strings.Contains(ast.Children[i].Tag, "expr"); i++ {
		x = eval(x, op, ast.Children[i].Eval())
	}
	return x

}
