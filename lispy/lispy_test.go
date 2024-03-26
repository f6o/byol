package lispy

import (
	"testing"
)

func TestPrint(t *testing.T) {
	// * 10 (+ 1 51)
	a := AST{"operator", "*", make([]AST, 0)}
	b := AST{"number", "10", make([]AST, 0)}
	c := AST{"expr", "", make([]AST, 0)}

	h := AST{"char", "(", make([]AST, 0)}
	i := AST{"operator", "+", make([]AST, 0)}
	j := AST{"number", "1", make([]AST, 0)}
	k := AST{"number", "51", make([]AST, 0)}
	l := AST{"char", ")", make([]AST, 0)}
	c.Children = append(c.Children, h)
	c.Children = append(c.Children, i)
	c.Children = append(c.Children, j)
	c.Children = append(c.Children, k)
	c.Children = append(c.Children, l)

	a.Children = append(a.Children, b)
	a.Children = append(a.Children, c)

	a.Print(0)
}
