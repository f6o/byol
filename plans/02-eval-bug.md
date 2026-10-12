# Step 2: Find and fix the fold bug in `Eval`

`AST.Eval` in `lispy/lispy.go` (line 78) gives wrong answers for some inputs. The existing test only uses two operands, which hides it.

## Write the tests first

Build these ASTs by hand, the same way `ast01` does, and predict each result before running:

| Input | Expected |
|---|---|
| `+ 1 2 3` | 6 |
| `- 10 2 3` | 5 |
| `+ 5` | 5 |
| `+ 1 (/ 1 0)` | a divide-by-zero error |

## Then read the loop at line 90

Answer these by tracing the code on paper:

1. What is the value of `x` on the second iteration of the loop?
2. What is `r` if the loop body never runs?
3. What stops the loop if the last child of the node is an `expr`? (The test ASTs end with a `regex` or `char` node. What if one didn't?)
4. If a nested expression returns `LVError{ERR_DIV_ZERO}`, which error does the outer call report?

## Fix

Change `Eval` so all four tests pass. Keep the change small: the shape of the function is right, the accumulation is not.

## Extra

The book treats `- 5` (one operand) as negation. Decide whether you want that, and add a test either way.

## Done when

- The four new tests pass.
- The tests from step 1 still pass.
