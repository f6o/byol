# Step 5: Chapters 8 and 9

## Chapter 8: Error handling

Most of this is already in place: `LVError` and the `ERR_*` constants exist, and `eval` returns an error on division by zero.

What is left:

- Print readable messages (`Error: division by zero`) instead of `lisp value error #0`.
- Report an unknown operator. `eval` currently returns 0 for an operator it doesn't recognise. Which error should that be?
- Report a number that doesn't fit in an `int`. The book has `LERR_BAD_NUM` for this; you have no equivalent yet. Where does `strconv.Atoi` fail today, and what happens to that failure?

Then try the chapter's bonus exercises, such as adding `%` and `^`.

## Chapter 9: S-expressions

This is the larger change. Read the chapter fully before starting, then plan around these points.

### `LispValue` becomes a tree

The book's `lval` gains two variants: a symbol and an S-expression holding a list of child values. In Go, think about:

- one struct type per variant behind the `LispValue` interface (what you have now), or
- one struct with a type tag, as the C code does.

Which makes `Eval` easier to read? Which makes adding a variant easier?

### Reading replaces the AST

The book adds a "read" step that converts the mpc AST into an `lval` tree. Since you own the parser, it can produce `LispValue` directly, and the `AST` type can go away. This is why step 3 recommends not investing in the mpc-shaped tree.

### The grammar changes

```
number : /-?[0-9]+/
symbol : '+' | '-' | '*' | '/'
sexpr  : '(' expr* ')'
expr   : number | symbol | sexpr
lispy  : expr*
```

Operators become ordinary symbols. Look at how the tokenizer separates `TOKEN_OPERATOR` from `TOKEN_SYMBOL` and decide whether that split still makes sense.

### Evaluation changes

Evaluating an S-expression means: evaluate the children, check the first is a symbol, then apply it to the rest. Errors in any child propagate upward. The fold you fixed in step 2 moves into a `builtin_op` equivalent.

## Suggested order

1. Finish the chapter 8 leftovers.
2. Add the new `LispValue` variants with printing, and test printing alone.
3. Change the parser to produce `LispValue`.
4. Rewrite evaluation over `LispValue`.
5. Remove `AST` and the old helpers once nothing uses them.
