# CLAUDE.md - AI Assistant Guide for byol

## Project Overview

**byol** (Build Your Own Lisp) is a Lisp interpreter implemented in Go, following the tutorial at http://buildyourownlisp.com/. The project is actively converting the C-based implementation into idiomatic Go, with an embedded C parser combinator library ([mpc](https://github.com/orangeduck/mpc)) used for early prototyping.

## Repository Structure

```
byol/
├── main.go              # CGo entry point (REPL using editline + mpc)
├── go.mod               # Go module: github.com/f6o/byol (Go 1.21.3)
├── lispy/               # Core Lisp interpreter package (pure Go)
│   ├── lispy.go         # AST, LispValue interface, eval logic
│   ├── tokenizer.go     # Stream-based tokenizer (bufio.Reader)
│   ├── parser.go        # Parser stub (TODO - not yet implemented)
│   ├── lispy_test.go    # AST evaluation tests
│   ├── tokenizer_test.go           # Tokenizer unit tests
│   └── tokenizer_iterator_test.go  # Iterator/streaming tokenizer tests
├── parser/              # Go port of mpc types (in progress)
│   └── mpc.go           # Go wrapper types for mpc structures
├── mpc/                 # Vendored C library (orangeduck/mpc)
│   ├── mpc.c            # MPC implementation (~4100 lines)
│   ├── mpc.h            # MPC header
│   ├── Makefile          # Builds libmpc.a static library
│   └── example/         # C examples (lispy.c, foobar.c)
├── README.md            # Brief project description
└── MEMO.md              # Detailed development notes (Japanese)
```

## Key Concepts

### Architecture

The project has two parallel tracks:

1. **CGo path** (`main.go`): Uses the C mpc library via CGo for parsing, with editline for the REPL. This was the initial approach following the book.
2. **Pure Go path** (`lispy/`): A native Go reimplementation of the tokenizer, parser, and evaluator. This is where active development is focused.

### Core Types (`lispy/lispy.go`)

- **`AST`**: Tree node with `Tag`, `Contents`, and `Children []AST`
- **`LispValue`** interface: Implemented by `LVNumber` (integers) and `LVError` (error codes)
- **`LispValueErrorNumber`** enum: `ERR_DIV_ZERO`, `ERR_ARGS_COUNT`, `ERR_TYPE_ASSERT`
- **`eval()`**: Evaluates binary operations (`+`, `-`, `*`, `/`) on `LVNumber` values
- **`AST.Eval()`**: Recursively evaluates the AST

### Tokenizer (`lispy/tokenizer.go`)

- Stream-based using `bufio.Reader` (supports `io.Reader` input)
- Token types: `TOKEN_NUMBER`, `TOKEN_OPERATOR`, `TOKEN_LPAREN`, `TOKEN_RPAREN`, `TOKEN_SYMBOL`, `TOKEN_EOF`
- Iterator pattern via `NewTokenizer(r)` + `NextToken()` loop
- Convenience functions: `Tokenize(string)` and `TokenizeReader(io.Reader)`

### Parser (`lispy/parser.go`)

- Stub only - not yet implemented (marked TODO)

## Build & Test Commands

```bash
# Run all Go tests
go test ./...

# Run tests with verbose output
go test -v ./lispy/...

# Build the CGo REPL (requires libedit and libmpc.a)
cd mpc && make shared && cd ..
go build -o byol .

# Build mpc C library only
cd mpc && make shared

# Clean mpc build artifacts
cd mpc && make clean
```

### Build Prerequisites (for main.go CGo path)

- GCC (C compiler)
- `libedit-dev` (editline library)
- The mpc static library must be built first: `cd mpc && make shared`

### Testing (lispy/ package - no external dependencies)

The `lispy/` package is pure Go and can be tested without any C toolchain:

```bash
go test ./lispy/...
```

## Code Conventions

- **Language**: Go (standard library only, no external Go dependencies)
- **Formatting**: Standard `gofmt`
- **Naming**: Go conventions - `CamelCase` for exported, `camelCase` for unexported
- **Testing**: Table-driven tests with `t.Run()` subtests
- **Error handling**: Custom types implementing the `LispValue` interface for Lisp-level errors; Go `error` for system-level errors
- **Package structure**: `lispy` is the main package for the interpreter; `parser` is the Go port of mpc types

## Development Notes

- `MEMO.md` contains detailed chapter-by-chapter development notes in Japanese, covering chapters 6-14 of the book
- The project is progressively replacing CGo/mpc usage with pure Go implementations
- The tokenizer is complete; the parser is the next component to implement
- Active development is on the `lispy/` package; `main.go` serves as a reference/prototype

## Git Workflow

- Primary branch: `master`
- No CI/CD configured
- Commit messages are short and descriptive (e.g., "add tokenizer", "use bufio", "fix")
