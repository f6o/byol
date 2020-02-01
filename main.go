package main

// #cgo LDFLAGS: -ledit
// #include <stdlib.h>
// #include <editline/readline.h>
// #include <editline/history.h>
import "C"
import (
	"unsafe"
	"fmt"
)

func main() {
	cs := C.CString("lispy> ")
	for {
		input := C.readline(cs)
		C.add_history(input)
		fmt.Printf("No you're a %s\n", C.GoString(input))
		C.free(unsafe.Pointer(input))
	}
}
