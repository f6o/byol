package parser

import "os"

// State Type

type mpc_state_t struct {
	Pos  int
	Row  int
	Col  int
	Term int
}

// Error Type

type mpc_error_t struct {
	state        mpc_state_t
	expected_num int
	filename     string
	failure      string
	expected     []string
	received     rune
}

func mpc_state_invalid() *mpc_state_t {
	return &mpc_state_t{-1, -1, -1, 0}
}

func Mpc_state_new() *mpc_state_t {
	return &mpc_state_t{0, 0, 0, 0}
}

type MPC_INPUT uint16

const (
	STRING MPC_INPUT = iota
	FILE
	PIPE
	MARKS_MIN     = 32
	INPUT_MEM_NUM = 512
)

type mpc_mem_t struct {
	mem [64]byte
}

type mpc_input_t struct {
	typenum  int    // "type" in mpc.c
	filename string // char * in mpc.c
	state    mpc_state_t

	str    string // "char *string" in mpc.c
	buffer string // char * in mpc.c
	file   *os.File

	suppress    int
	backtrack   int
	marks_slots int
	marks_num   int
	marks       *mpc_state_t

	lasts string // char * in mpc.c
	last  rune   // char in mpc.c

	mem_index int                 // size_t in mpc.c
	mem_full  [INPUT_MEM_NUM]byte // char mem_full[MPC_INPUT_MEM_NUM] in mpc.c
	mem       [INPUT_MEM_NUM]mpc_mem_t
}
