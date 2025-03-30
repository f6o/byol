package lispy

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []Token
		wantErr  bool
	}{
		{
			name:     "Simple Addition",
			input:    "(+ 10 20)",
			expected: []Token{{TOKEN_LPAREN, "("}, {TOKEN_OPERATOR, "+"}, {TOKEN_NUMBER, "10"}, {TOKEN_NUMBER, "20"}, {TOKEN_RPAREN, ")"}, {TOKEN_EOF, ""}},
			wantErr:  false,
		},
		{
			name:     "Subtraction",
			input:    "(- 50 30)",
			expected: []Token{{TOKEN_LPAREN, "("}, {TOKEN_OPERATOR, "-"}, {TOKEN_NUMBER, "50"}, {TOKEN_NUMBER, "30"}, {TOKEN_RPAREN, ")"}, {TOKEN_EOF, ""}},
			wantErr:  false,
		},
		{
			name:     "Multiplication",
			input:    "(* 5 2)",
			expected: []Token{{TOKEN_LPAREN, "("}, {TOKEN_OPERATOR, "*"}, {TOKEN_NUMBER, "5"}, {TOKEN_NUMBER, "2"}, {TOKEN_RPAREN, ")"}, {TOKEN_EOF, ""}},
			wantErr:  false,
		},
		{
			name:     "Division",
			input:    "(/ 10 2)",
			expected: []Token{{TOKEN_LPAREN, "("}, {TOKEN_OPERATOR, "/"}, {TOKEN_NUMBER, "10"}, {TOKEN_NUMBER, "2"}, {TOKEN_RPAREN, ")"}, {TOKEN_EOF, ""}},
			wantErr:  false,
		},
		{
			name:     "Nested Expression",
			input:    "(+ 10 (* 2 3))",
			expected: []Token{{TOKEN_LPAREN, "("}, {TOKEN_OPERATOR, "+"}, {TOKEN_NUMBER, "10"}, {TOKEN_LPAREN, "("}, {TOKEN_OPERATOR, "*"}, {TOKEN_NUMBER, "2"}, {TOKEN_NUMBER, "3"}, {TOKEN_RPAREN, ")"}, {TOKEN_RPAREN, ")"}, {TOKEN_EOF, ""}},
			wantErr:  false,
		},
		{
			name:     "Symbol",
			input:    "(define x 10)",
			expected: []Token{{TOKEN_LPAREN, "("}, {TOKEN_SYMBOL, "define"}, {TOKEN_SYMBOL, "x"}, {TOKEN_NUMBER, "10"}, {TOKEN_RPAREN, ")"}, {TOKEN_EOF, ""}},
			wantErr:  false,
		},
		{
			name:     "Invalid Token",
			input:    "(+ 10 $)",
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "Empty Input",
			input:    "",
			expected: []Token{{TOKEN_EOF, ""}},
			wantErr:  false,
		},
		{
			name:     "Multiple Spaces",
			input:    "(+    10   20)",
			expected: []Token{{TOKEN_LPAREN, "("}, {TOKEN_OPERATOR, "+"}, {TOKEN_NUMBER, "10"}, {TOKEN_NUMBER, "20"}, {TOKEN_RPAREN, ")"}, {TOKEN_EOF, ""}},
			wantErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Tokenize(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Errorf("Tokenize(%q) expected an error, but got none", tc.input)
				}
				return
			}
			if err != nil {
				t.Errorf("Tokenize(%q) returned an unexpected error: %v", tc.input, err)
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Tokenize(%q) = %v, expected %v", tc.input, actual, tc.expected)
			}
		})
	}
}
