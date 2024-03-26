#include "../mpc.h"

int main(int argc, char** argv) {
    mpc_parser_t* Number = mpc_new("number");
    mpc_parser_t* Operator = mpc_new("operator");
    mpc_parser_t* Expr = mpc_new("expr");
    mpc_parser_t* Lispy = mpc_new("lispy");

    if (argc != 2) {
      printf("Usage: ./lispy <lispy>\n");
      exit(0);
    }

    mpca_lang(MPCA_LANG_DEFAULT,
    "\
     number : /-?[0-9]+/ ; \
     operator : '+' | '-' | '*' | '/' ; \
     expr : <number> | '(' <operator> <expr>+ ')' ; \
     lispy : /^/ <operator> <expr>+ /$/ ; \
     ", Number, Operator, Expr, Lispy);

    mpc_result_t r;

    if (mpc_parse("<stdin>", argv[1], Lispy, &r)) {
        mpc_ast_print(r.output);
        mpc_ast_delete(r.output);
    } else {
        mpc_err_print(r.error);
        mpc_err_delete(r.error);
    }
    
    mpc_cleanup(1, Lispy);
    
    return 0;
}