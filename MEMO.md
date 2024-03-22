 # Build your own lisp/micro parser combinator

* https://www.buildyourownlisp.com/
* https://github.com/orangeduck/mpc

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

## Chapter 12 関数

* `\\`を関数のシンボルとしてあつかう
* 関数もlvalとしてあつかう

```
typedef lval *(lbuiltin)(lenv *,lval *)

struct lval {
  // ...
  
  builtin lbuiltin
  env  lenv*
  formals  lval*
  body lval*
  
  // ...
}
```

```
(* define *)
\ {x y} {+ x y}

(* apply *)
(\ {x y} {+ x y}) 10 20

(* bind *)
def {f} (\ {x y} {+ x y})
f 10 20
```

* lambda式の導入
* 環境の拡張

```
struct lenv {
    parent lenv*
    count  int
    syms   []string
    vals   []string
}
```

Function callingの手前まで。

## mpc/example/foobar.c

1. mpc_newでパーサを定義する
1. mpca_langで文法を定義する
1. mpc_parseで入力をパースする
1. mpc_ast_printで構文木を出力する
1. mpc_ast_deleteでメモリを解放する
1. mpc_cleanup


## Ch.6 Parsing

逆ポーランド記法の式をパースする

## Chapter.7 eval

* mpc_new,mpca_lang ... 文法
* mpc_parse ... ASTをつくる
* eval

mpca_lang -> mpca_lang_st 

mpc.cの1700行目からパーサになっている

### Parsing: ASTをつくる

mpcライブラリの仕事。

#### ASTの構造

* tag expr,number,regexなどのノードの型みたいなもの。ルールの定義につかう。
* contents ノードの中身。空のときもある。
* state このノードをパースしたときのパーサの状態。行数と列とか。
* children 子ノード。ポインタで木の構造にする

### Evaluation: ASTの評価する

この章では逆ポーランド記法の式を評価する。

* number そのまま数値をかえす
* 2つの数値をつかって四則演算をする

### Printing: ASTを印字する
