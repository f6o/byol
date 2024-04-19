package lispy

import (
	"testing"
)

func ast01() (AST, int) {
	// * 10 (+ 1 51)
	ans := 520
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

	return root, ans
}

func TestEval(t *testing.T) {
	root, ans := ast01()
	if x, ok := root.Eval().(LVNumber); ok && x.number != ans {
		t.Errorf("Eval %d; expected %d", x, ans)
	}
}
