package main

// #cgo LDFLAGS: -ledit
// #include <stdlib.h>
// #include <editline/readline.h>
// #include <editline/history.h>
// #include "mpc/mpc.h"
import "C"
import (
	"fmt"
	"regexp"
	"unsafe"

	"github.com/f6o/byol/parser"
)

func symbol(name string) interface{}  {
	cs := C.CString(name)
	return C.mpc_sym(cs)
}

func main() {

	adjective := C.mpc_or(4,
		symbol("wow"),
		symbol("many"),
		symbol("so"),
		symbol("such"))
	
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
