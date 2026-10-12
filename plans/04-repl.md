# Step 4: Wire a pure-Go REPL together

`main.go` needs cgo, `libedit` and a built `mpc` library, which is awkward on Windows. It also never parses its input. This step gets you to the end state of chapter 7 without the C toolchain.

## Goal

A loop that does: read a line → tokenize → parse → eval → print.

```
lispy> + 1 (* 2 3)
7
lispy> / 1 0
Error: division by zero
```

## Things to decide

- **Reading input.** `bufio.Scanner` on `os.Stdin` is enough. What should happen at end of input (Ctrl+Z on Windows, Ctrl+D elsewhere)?
- **Where errors go.** There are three kinds: tokenizer errors, parser errors, and `LVError` values from `Eval`. None of them should end the loop.
- **Printing.** `LispValue.Print()` writes straight to stdout with a debug-style message. Consider whether a `String()` method would serve both the REPL and the tests better.
- **The old `main.go`.** Keep it as a reference under a build tag or in a separate directory, or delete it. Note that its `mpc_and2` wrapper at line 29 calls `mpc_or` where `mpc_and` looks intended.

## Testability

Put the read-eval-print logic in a function that takes an `io.Reader` and an `io.Writer`, and have `main` pass stdin and stdout. Then a test can feed it a string and check the output.

## Done when

- `go run .` works without cgo.
- A bad line prints an error and the prompt comes back.
