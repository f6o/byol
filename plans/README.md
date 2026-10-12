# Plans

Next steps for the Go port of *Build Your Own Lisp*. These are instructions and questions only; the implementation is yours.

## Where the project stands

- The tokenizer (`lispy/tokenizer.go`) works and is tested.
- The evaluator (`lispy/lispy.go`) works on hand-built ASTs.
- The parser (`lispy/parser.go`) is a stub, so nothing connects the two.
- `main.go` is a cgo REPL that builds a grammar with mpc but never parses input.

## Order

| Step | File | Goal |
|---|---|---|
| 1 | [01-fix-tests.md](01-fix-tests.md) | Make the existing tests able to fail |
| 2 | [02-eval-bug.md](02-eval-bug.md) | Find and fix the fold bug in `Eval` |
| 3 | [03-parser.md](03-parser.md) | Implement `Parser.Parse()` |
| 4 | [04-repl.md](04-repl.md) | Wire a pure-Go REPL together |
| 5 | [05-chapters-8-9.md](05-chapters-8-9.md) | Error messages, then S-expressions |

Steps 1 and 2 come first because the parser tests in step 3 rely on `Eval` and its tests being trustworthy.

## Postponed

Porting mpc to Go (`parser/mpc.go`, the README's other TODO) is a far larger job than all five steps combined: `mpc/mpc.c` is several thousand lines. It doesn't block the book, because the hand-written parser in step 3 replaces what mpc does. Treat it as a separate project.
