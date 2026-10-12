# Step 3: Implement `Parser.Parse()`

`lispy/parser.go` is a stub. This step turns a token stream into an `AST`, which connects the tokenizer to the evaluator.

## Grammar

From chapter 6, in Polish notation:

```
lispy : operator expr+ EOF
expr  : NUMBER | '(' operator expr+ ')'
```

Recursive descent fits this well: one function per rule, each consuming tokens and returning a node.

## Decisions to make before writing code

### 1. AST shape

`Eval` currently expects the layout mpc produces:

- tags like `expr|number|regex` and `operator|char`
- `char` nodes for the parentheses
- `regex` sentinel nodes at both ends of the root

You have two options:

- **Keep the mpc shape.** `Eval` works unchanged, and parsing `* 10 (+ 1 51)` should produce a tree equal to `ast01`.
- **Design a cleaner tree** and adapt `Eval` and the test helpers.

Recommendation: the cleaner tree. Chapter 9 replaces this AST with S-expressions, so the mpc mimicry is throwaway work.

### 2. Token source

- Take the `[]Token` slice from `Tokenize` and index into it, or
- pull from `Tokenizer.NextToken()` with one token of lookahead.

The slice is easier to start with. The streaming version is a good follow-up exercise.

### 3. The `Parser` struct

It currently holds `src string` and `ast *AST`. Decide what it really needs: tokens and a position? A constructor? Should `Parse` return the AST instead of storing it?

### 4. Errors

Decide what `Parse` returns for each of these, and write them as tests first:

| Input | Problem |
|---|---|
| `(+ 1` | missing `)` |
| `+ 1 )` | unexpected `)` |
| empty line | nothing to parse |
| `1 2` | no operator |
| `+` | operator with no operands |
| `+ 1 foo` | a symbol where an expression is expected |

A good error says what was expected and what was found.

## Suggested order of work

1. Parse a single number expression.
2. Parse `operator expr+` at the top level, with numbers only.
3. Add the parenthesised case, which makes `expr` recursive.
4. Add the error cases.
5. Add end-to-end tests: source string → tokens → AST → `Eval` → number.

## Tokenizer gap to handle

`-` is always tokenized as an operator, so `-5` cannot be a number. The book's number rule is `-?[0-9]+`. Decide where to fix this: in the tokenizer (needs lookahead, and `peek` already exists) or in the parser.

## Done when

- `* 10 (+ 1 51)` parses and evaluates to 520.
- Every error case in the table returns an error and does not panic.
