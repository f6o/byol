package lispy

import (
	"testing"
)

func NewRoot(children []AST) AST {
	return AST{">", "", children}
}

func NewNode(children []AST) AST {
	return AST{"expr|>", "", children}
}

func NewOperator(op string) AST {
	return AST{"operator|char", op, make([]AST, 0)}
}

func NewNumber(num string) AST {
	return AST{"expr|number|regex", num, make([]AST, 0)}
}

func NewChar(ch string) AST {
	return AST{"char", ch, make([]AST, 0)}
}

func ast01() (AST, int) {
	// * 10 (+ 1 51)
	ans := 520
	regex := AST{"regex", "", make([]AST, 0)}

	root := NewRoot([]AST{
		regex,
		NewOperator("*"),
		NewNumber("10"),
		NewNode([]AST{
			NewChar("("),
			NewOperator("+"),
			NewNumber("1"),
			NewNumber("51"),
			NewChar(")"),
		}),
		regex,
	})

	return root, ans
}

func TestEval(t *testing.T) {
	root, ans := ast01()
	root.Print(0)
	if x, ok := root.Eval().(LVNumber); ok && x.number != ans {
		t.Errorf("Eval %d; expected %d", x, ans)
	}
}

func ast02() AST {
	regex := AST{"regex", "", make([]AST, 0)}
	return NewRoot([]AST{
		regex,
		NewOperator("/"),
		NewNumber("1"),
		NewNumber("0"),
		regex,
	})
}

func TestDivByZeroError(t *testing.T) {
	root := ast02()
	root.Print(0)
	if x, ok := root.Eval().(LVError); !ok && x.errno != ERR_DIV_ZERO {
		x.Print()
		t.Errorf("did not got LVError")
	}
}
