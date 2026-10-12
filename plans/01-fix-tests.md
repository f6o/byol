# Step 1: Make the existing tests able to fail

Both tests in `lispy/lispy_test.go` pass regardless of what `Eval` returns. Fix them before relying on them.

## The problems

### `TestEval` (line 52)

The condition is `ok && x.number != ans`.

- If `Eval` returns something that is not an `LVNumber`, `ok` is false.
- The whole condition is then false, and the test passes silently.

### `TestDivByZeroError` (line 71)

The condition is `!ok && x.errno != ERR_DIV_ZERO`.

- When `ok` is false, `x` is the zero value of `LVError`.
- Its `errno` is 0, and `ERR_DIV_ZERO` is also 0 (the first `iota`).
- The condition can never be true, so the test can never fail.

## Exercise

1. Rewrite each test so that a wrong type fails and a wrong value fails, as two separate checks.
2. Fix the `t.Errorf` format in `TestEval`: it prints an `LVNumber` struct with `%d`.
3. Prove the tests work: temporarily break `eval` (for example, make `/` never return an error), run the tests, and confirm they go red. Then revert.

## Questions to think about

- Why is it risky that `ERR_DIV_ZERO` is the zero value of its type? What would change if the first constant were a "no error" placeholder?
- `go vet` reports some suspicious test conditions. Does it catch either of these?

## Done when

- Both tests fail when `Eval` is deliberately broken.
- Both tests pass on the unmodified evaluator.
