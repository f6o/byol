package lispy

import (
	"testing"
)

func TestPrint(t *testing.T) {
	// * 10 (+ 1 51)
	root := AST{">", "", make([]AST, 0)}

	regex := AST{"regex", "", make([]AST, 0)}
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

	root.Children = append(root.Children, regex)
	root.Children = append(root.Children, a)
	root.Children = append(root.Children, b)
	root.Children = append(root.Children, c)
	root.Children = append(root.Children, regex)

	root.Print(0)
}

func TestEval(t *testing.T) {
	// * 10 (+ 1 51)
	root := AST{">", "", make([]AST, 0)}

	regex := AST{"regex", "", make([]AST, 0)}
	a := AST{"operator|char", "*", make([]AST, 0)}
	b := AST{"expr|number|regex", "10", make([]AST, 0)}
	c := AST{"expr|>", "", make([]AST, 0)}

	h := AST{"char", "(", make([]AST, 0)}
	i := AST{"operator|char", "+", make([]AST, 0)}
	j := AST{"expr|number|regex", "1", make([]AST, 0)}
	k := AST{"expr|number|regex", "51", make([]AST, 0)}
	l := AST{"char", ")", make([]AST, 0)}

	c.Children = append(c.Children, h)
	c.Children = append(c.Children, i)
	c.Children = append(c.Children, j)
	c.Children = append(c.Children, k)
	c.Children = append(c.Children, l)

	root.Children = append(root.Children, regex)
	root.Children = append(root.Children, a)
	root.Children = append(root.Children, b)
	root.Children = append(root.Children, c)
	root.Children = append(root.Children, regex)

	x := root.Eval()
	ans := 520
	if x != ans {
		t.Errorf("Eval %d; expected %d", x, ans)
	}
}
