package main

// #cgo LDFLAGS: -ledit
// #include <stdlib.h>
// #include <editline/readline.h>
// #include <editline/history.h>
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/f6o/byol/parser"
)

func main() {
	cs := C.CString("lispy> ")
	for {
		input := C.readline(cs)
		C.add_history(input)
		fmt.Printf("No you're a %s\n", C.GoString(input))
		C.free(unsafe.Pointer(input))
		state := parser.Mpc_state_new()
		fmt.Printf("%d %d\n", state.Pos, state.Term)
	}
}
