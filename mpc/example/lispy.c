#include "../mpc.h"

void print_tag(mpc_ast_t* a, int depth) {
    printf("[%d]%s\n", depth, a->tag);
    for ( int i = 0; i < a->children_num; i++ ) {
        mpc_ast_t* child = a->children[i];
        print_tag(child, depth+1);
    }
}

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

        printf("PRINT TAG ONLY\n");
        print_tag(r.output, 0);

        mpc_ast_delete(r.output);
    } else {
        mpc_err_print(r.error);
        mpc_err_delete(r.error);
    }
    
    mpc_cleanup(1, Lispy);
    
    return 0;
}