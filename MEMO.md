# BYOL memo

## mpc.hでのASTの定義

```
mpc_ast_t {
    tag string
    contents string
    state mpc_state_t
    children_num int
    children mpc_ast_t**
}
```

## ASTの操作

* new
* build
* add_root
* add_child
* add_tag
* add_route_tag
* tag
* state
* delete
* print
* get_index
* get_child

## ASTの走査順

* pre
* post

`mpc_parse`の結果としてASTがある

## Lisp Value

* lisp value, lval としてエラー、数字、シンボル、S式
* read: AST -> lval
* eval: lval -> lval

## quote

Common Lispだと`'(1 2 3 4)`だけど、ここでは
`{1 2 3 4}`という形式

Chapter9まででS式のREPLができる
Chapter11では環境を導入して、変数をつかえるようにした


