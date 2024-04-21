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

func (n LVNumber) Print() {
	fmt.Printf("lisp value number %d", n.number)
}

type LispValueErrorNumber int

const (
	ERR_DIV_ZERO LispValueErrorNumber = iota
	ERR_ARGS_COUNT
	ERR_TYPE_ASSERT
)

type LVError struct {
	errno LispValueErrorNumber
}

func (err LVError) Print() {
	fmt.Printf("lisp value error #%d", err.errno)
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

func eval(x LVNumber, op string, y LVNumber) LispValue {
	r := LVNumber{number: 0}

	switch op {
	case "+":
		r.number = x.number + y.number
	case "-":
		r.number = x.number - y.number
	case "*":
		r.number = x.number * y.number
	case "/":
		if y.number == 0 {
			return LVError{ERR_DIV_ZERO}
		} else {
			r.number = x.number / y.number
		}
	}

	return r
}

func (ast AST) Eval() LispValue {
	if strings.Contains(ast.Tag, "number") {
		if i, err := strconv.Atoi(ast.Contents); err == nil {
			return LVNumber{number: i}
		}
	}

	op := ast.Children[1].Contents
	var r LispValue

	switch x := ast.Children[2].Eval().(type) {
	case LVNumber:
		for i := 3; strings.Contains(ast.Children[i].Tag, "expr"); i++ {
			if y, ok := ast.Children[i].Eval().(LVNumber); ok {
				r = eval(x, op, y)
			} else {
				return LVError{ERR_TYPE_ASSERT}
			}
		}
	default:
		return LVError{ERR_TYPE_ASSERT}
	}

	return r

}
