package main

// #cgo LDFLAGS: -Lmpc/build -ledit -lmpc
// #include <stdlib.h>
// #include <editline/readline.h>
// #include <editline/history.h>
// #include "mpc/mpc.h"
//
// static mpc_parser_t *mpc_or4(
//	 mpc_parser_t *p1,
//   mpc_parser_t *p2,
//   mpc_parser_t *p3,
//   mpc_parser_t *p4) {
//     return mpc_or(4, p1, p2, p3, p4);
// }
//
// static mpc_parser_t *mpc_or5(
//	 mpc_parser_t *p1,
//   mpc_parser_t *p2,
//   mpc_parser_t *p3,
//   mpc_parser_t *p4,
//   mpc_parser_t *p5) {
//     return mpc_or(5, p1, p2, p3, p4, p5);
// }
//
// static mpc_parser_t *mpc_and2(
//	 mpc_parser_t *p1,
//   mpc_parser_t *p2) {
//     return mpc_or(2, mpcf_strfold, p1, p2, free);
// }
//
// static mpc_parser_t *mpc_many_single(
//	 mpc_parser_t *p1) {
//     return mpc_many(mpcf_strfold, p1);
// }
//
import "C"
import (
	"fmt"
	"regexp"
	"unsafe"

	"github.com/f6o/byol/parser"
)

func symbol(name string) *C.mpc_parser_t {
	cs := C.CString(name)
	sym := C.mpc_sym(cs)
	fmt.Printf("%T\n", sym)
	return sym
}

func main() {
	// https://pkg.go.dev/cmd/cgo
	//    "Calling variadic C functions is not supported.
	//     It is possible to circumvent this by using a C function wrapper."
	adjective := C.mpc_or4(
		symbol("wow"),
		symbol("many"),
		symbol("so"),
		symbol("such"))

	noun := C.mpc_or5(
		symbol("lisp"),
		symbol("language"),
		symbol("book"),
		symbol("build"),
		symbol("c"),
	)

	phrase := C.mpc_and2(adjective, noun)

	doge := C.mpc_many_single(phrase)

	fmt.Printf("%T\n", doge)

	// TODO: parse string

	C.mpc_delete(doge)

	cs := C.CString("lispy> ")
	for {
		input := C.readline(cs)
		line := C.GoString(input)
		matched, err := regexp.MatchString(`[a-zA-Z0-9]+`, line)
		if err != nil {
			return
		}
		if matched {
			fmt.Println("add history:", line)
			C.add_history(input)
		}
		fmt.Printf("No you're a %s\n", C.GoString(input))
		C.free(unsafe.Pointer(input))
		state := parser.Mpc_state_new()
		fmt.Printf("%d %d\n", state.Pos, state.Term)
	}
}
